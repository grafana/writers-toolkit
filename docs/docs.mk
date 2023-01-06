.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

DOCS_IMAGE   = docker.io/grafana/docs-base:latest
DOCS_PROJECT = writers-toolkit
DOCS_DIR     = sources

# Support podman over Docker if it is available.
PODMAN := $(shell if command -v podman &>/dev/null; then echo podman; else echo docker; fi)

# This allows ports and base URL to be overridden, so services like ngrok.io can
# be used to share a local running docs instances.
DOCS_HOST_PORT    = 3002
DOCS_LISTEN_PORT  = 3002
DOCS_BASE_URL    ?= "localhost:$(DOCS_HOST_PORT)"

HUGO_REFLINKSERRORLEVEL ?= WARNING
DOCS_CONTAINER = $(DOCS_PROJECT)-docs

.PHONY: docs-docker-rm
docs-rm:
	$(PODMAN) rm -f $(DOCS_CONTAINER)

.PHONY: docs-pull
docs-pull:
	$(PODMAN) pull $(DOCS_IMAGE)

.PHONY: docs
docs: ## Serve documentation locally.
docs: docs-pull
	@echo "Documentation will be served at:"
	@echo "http://$(DOCS_BASE_URL)/docs/$(DOCS_PROJECT)/"
	@echo ""
	@if [[ -z $${NON_INTERACTIVE} ]]; then \
		read -p "Press a key to continue"; \
	fi
	$(PODMAN) run -it --name $(DOCS_CONTAINER) \
		-v $(CURDIR)/$(DOCS_DIR):/hugo/content/docs/$(DOCS_PROJECT)/:rw,z \
		-e HUGO_REFLINKSERRORLEVEL=$(HUGO_REFLINKSERRORLEVEL) \
		-p $(DOCS_HOST_PORT):$(DOCS_LISTEN_PORT) \
		--rm $(DOCS_IMAGE) \
		make server BUILD_DRAFTS=true
