.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

# Adapted from https://suva.sh/posts/well-documented-makefiles/
.PHONY: help
help: ## Display this help
help:
	@awk 'BEGIN {FS = ": ##"; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_\.\-\/% ]+: ##/ { printf "  %-45s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

PODMAN := $(shell if command -v podman >/dev/null 2>&1; then echo podman; else echo docker; fi)
GIT_ROOT := $(shell git rev-parse --show-toplevel)

dictionaries/en_US-grafana.%: dictionaries/wordlist
	$(MAKE) -C $(@D) $(@F)

.PHONY: grafana/vale
grafana/vale: ## Builds a container image for Vale with the Grafana style loaded.
grafana/vale: dictionaries/en_US-grafana.aff dictionaries/en_US-grafana.dic
	DOCKER_DEFAULT_PLATFORM=linux/amd64 DOCKER_BUILDKIT=1 $(PODMAN) build -t grafana/vale  .

.PHONY: sync
sync: ## Update the vendored Google style.
sync:
	cd $(GIT_ROOT) && vale sync

.vale.ini ../.vale.ini: ## Generate Vale INI configuration from the Jsonnet data template.
.vale.ini ../.vale.ini: .vale.jsonnet
	jsonnet -Se "(import '.vale.jsonnet').container" | sed '$${/^$$/d;}' > .vale.ini
	jsonnet -Se "(import '.vale.jsonnet').repository" | sed '$${/^$$/d;}' > ../.vale.ini
