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
	@awk 'BEGIN {FS = ": ##"; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_\.\-\/%]+: ##/ { printf "  %-45s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

PODMAN := $(shell if command -v podman >/dev/null 2>&1; then echo podman; else echo docker; fi)

dictionaries/en_US-grafana.%: dictionaries/wordlist
	$(MAKE) -C $(@D) $(@F)

.PHONY: grafana/vale
grafana/vale: ## Builds a Docker image for Vale with the technical-documentation styles loaded.
grafana/vale: dictionaries/en_US-grafana.aff dictionaries/en_US-grafana.dic
	$(PODMAN) build . -t grafana/vale:latest
