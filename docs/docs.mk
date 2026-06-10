# The source of this file is https://raw.githubusercontent.com/grafana/writers-toolkit/main/docs/docs.mk.
# A changelog is included in the head of the `make-docs` script.
include variables.mk
-include variables.mk.local

.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

.DEFAULT_GOAL: help

# Adapted from https://www.thapaliya.com/en/writings/well-documented-makefiles/
.PHONY: help
help: ## Display this help.
help:
	@awk 'BEGIN { \
		FS = ": ##"; \
		printf "Usage:\n  make <target>\n\nTargets:\n" \
	} \
	/^[a-zA-Z0-9_\.\-\/%]+: ##/ { printf "  %-15s %s\n", $$1, $$2 }' \
	$(MAKEFILE_LIST)

GIT_ROOT := $(shell git rev-parse --show-toplevel)

PODMAN := $(shell if command -v podman >/dev/null 2>&1; then echo podman; else echo docker; fi)

ifeq ($(PROJECTS),)
$(error "PROJECTS variable must be defined in variables.mk")
endif

# First project slug from PROJECTS, without any optional :version/:repo/:path suffixes.
BROKEN_LINKS_PROJECT := $(firstword $(subst :, ,$(word 1,$(PROJECTS))))

# Host port to publish container port to.
ifeq ($(origin DOCS_HOST_PORT), undefined)
export DOCS_HOST_PORT := 3002
endif

# Host port used by local nginx during broken-link checks.
ifeq ($(origin BROKEN_LINKS_NGINX_PORT), undefined)
export BROKEN_LINKS_NGINX_PORT := 3002
endif

# Git ref used to determine changed files for the local broken-links comment.
ifeq ($(origin BROKEN_LINKS_BASE_REF), undefined)
export BROKEN_LINKS_BASE_REF := origin/main
endif

# Sources mapping used by local broken-links checks; mirrors CI deploy-pr-preview sources.
ifeq ($(origin BROKEN_LINKS_SOURCES), undefined)
export BROKEN_LINKS_SOURCES := [{"index_file":null,"relative_prefix":"/docs/$(BROKEN_LINKS_PROJECT)/","repo":"$(BROKEN_LINKS_PROJECT)","source_directory":"docs/sources","website_directory":"content/docs/$(BROKEN_LINKS_PROJECT)"}]
endif

# Container image used to perform Hugo build.
ifeq ($(origin DOCS_IMAGE), undefined)
export DOCS_IMAGE := grafana/docs-base:latest
endif

# Container ancestor used to copy dist from a running make docs server.
ifeq ($(origin BROKEN_LINKS_CONTAINER_ANCESTOR), undefined)
export BROKEN_LINKS_CONTAINER_ANCESTOR := $(DOCS_IMAGE)
endif

# Container image used for Vale linting.
ifeq ($(origin VALE_IMAGE), undefined)
export VALE_IMAGE := grafana/vale:latest
endif

# PATH-like list of directories within which to find projects.
# If all projects are checked out into the same directory, ~/repos/ for example, then the default should work.
ifeq ($(origin REPOS_PATH), undefined)
export REPOS_PATH := $(realpath $(GIT_ROOT)/..)
endif

# How to treat Hugo relref errors.
ifeq ($(origin HUGO_REFLINKSERRORLEVEL), undefined)
export HUGO_REFLINKSERRORLEVEL := WARNING
endif

# Whether to pull the latest container image before running the container.
ifeq ($(origin PULL), undefined)
export PULL := true
endif

.PHONY: docs-rm
docs-rm: ## Remove the docs container.
	$(PODMAN) rm -f $(DOCS_CONTAINER)

.PHONY: docs-pull
docs-pull: ## Pull documentation base image.
	$(PODMAN) pull -q $(DOCS_IMAGE)

make-docs: ## Fetch the latest make-docs script.
make-docs:
	if [[ ! -f "$(CURDIR)/make-docs" ]]; then
		echo 'WARN: No make-docs script found in the working directory. Run `make update` to download it.' >&2
		exit 1
	fi

.PHONY: docs
docs: ## Serve documentation locally, which includes pulling the latest `DOCS_IMAGE` (default: `grafana/docs-base:latest`) container image. To not pull the image, set `PULL=false`.
ifeq ($(PULL), true)
docs: docs-pull make-docs
else
docs: make-docs
endif
	$(CURDIR)/make-docs $(PROJECTS)

.PHONY: docs-debug
docs-debug: ## Run Hugo web server with debugging enabled. TODO: support all SERVER_FLAGS defined in website Makefile.
docs-debug: make-docs
	WEBSITE_EXEC='hugo server --bind 0.0.0.0 --port 3002 --logLevel debug' $(CURDIR)/make-docs $(PROJECTS)

.PHONY: broken-links-copy-dist
broken-links-copy-dist: ## Copy dist from running make docs container.
	@if ! command -v docker >/dev/null 2>&1; then
		echo 'ERRR: `docker` must be installed to run broken-links-copy-dist.' >&2
		exit 1
	fi
	cd "$(GIT_ROOT)"
	container_id="$$(docker ps -q --filter "ancestor=$(BROKEN_LINKS_CONTAINER_ANCESTOR)" | head -n 1)"
	if [[ -z "$$container_id" ]]; then
		echo "ERRR: no running docs container found for ancestor '$(BROKEN_LINKS_CONTAINER_ANCESTOR)'." >&2
		echo 'NOTE: run `make docs` first, or override BROKEN_LINKS_CONTAINER_ANCESTOR.' >&2
		exit 1
	fi

	rm -rf dist
	docker cp "$$container_id:/hugo/dist" ./dist

.PHONY: check-links
check-links: ## Run local broken-link checker and generate the broken-links comment.
	@if ! command -v go >/dev/null 2>&1; then
		echo 'ERRR: `go` must be installed to run broken-links-check.' >&2
		exit 1
	fi
	@if ! command -v nginx >/dev/null 2>&1; then
		echo 'ERRR: `nginx` must be installed to run broken-links-check.' >&2
		exit 1
	fi
	if [[ "$(BROKEN_LINKS_NGINX_PORT)" == "3002" ]]; then
		if command -v lsof >/dev/null 2>&1; then
			if lsof -nP -iTCP:3002 -sTCP:LISTEN >/dev/null 2>&1; then
				echo 'ERRR: port 3002 is already in use. Stop the running service before make broken-links-check.' >&2
				exit 1
			fi
		elif command -v ss >/dev/null 2>&1; then
			if ss -H -ltn 'sport = :3002' | grep -q .; then
				echo 'ERRR: port 3002 is already in use. Stop the running service before make broken-links-check.' >&2
				exit 1
			fi
		fi
	fi
	cd "$(GIT_ROOT)"
	SOURCES='$(BROKEN_LINKS_SOURCES)' go run ./link-checker/cmd/link-checker check -nginx-port "$(BROKEN_LINKS_NGINX_PORT)"

	SOURCES='$(BROKEN_LINKS_SOURCES)' go run ./link-checker/cmd/link-checker comment \
		-links links.json \
		-base-ref "$(BROKEN_LINKS_BASE_REF)" \
		-output broken-links-comment.md \
		-repo "$${BROKEN_LINKS_REPO:-$$(basename "$$(git rev-parse --show-toplevel)")}" \
		-title "$${BROKEN_LINKS_TITLE:-Local run}" \
		-artifact-url "$${BROKEN_LINKS_ARTIFACT_URL:-}"
	cat broken-links-comment.md

.PHONY: vale
vale: ## Run vale on the entire docs folder which includes pulling the latest `VALE_IMAGE` (default: `grafana/vale:latest`) container image. To not pull the image, set `PULL=false`.
vale: make-docs
ifeq ($(PULL), true)
	$(PODMAN) pull -q $(VALE_IMAGE)
endif
	DOCS_IMAGE=$(VALE_IMAGE) $(CURDIR)/make-docs $(PROJECTS)

.PHONY: update
update: ## Fetch the latest version of this Makefile and the `make-docs` script from Writers' Toolkit.
	curl -s -LO https://raw.githubusercontent.com/grafana/writers-toolkit/main/docs/docs.mk
	curl -s -LO https://raw.githubusercontent.com/grafana/writers-toolkit/main/docs/make-docs
	chmod +x make-docs

# ls static/templates/ | sed 's/-template\.md//' | xargs
TOPIC_TYPES := concept multiple-tasks reference section task tutorial visualization
.PHONY: $(patsubst %,topic/%,$(TOPIC_TYPES))
topic/%: ## Create a topic from the Writers' Toolkit template. Specify the topic type as the target, for example, `make topic/task TOPIC_PATH=sources/my-new-topic.md`.
$(patsubst %,topic/%,$(TOPIC_TYPES)):
	$(if $(TOPIC_PATH),,$(error "You must set the TOPIC_PATH variable to the path where the $(@F) topic will be created. For example: make $(@) TOPIC_PATH=sources/my-new-topic.md"))
	mkdir -p $(dir $(TOPIC_PATH))
	curl -s -o $(TOPIC_PATH) https://raw.githubusercontent.com/grafana/writers-toolkit/refs/heads/main/docs/static/templates/$(@F)-template.md
