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

en_US-grafana.dic: ## Generate the Grafana dictionary from the wordlist.
en_US-grafana.dic: wordlist
	cat <(wc -l <$< | tr -d ' ') $< > $@

en_US-grafana.aff: ## Create the Grafana affix file.
en_US-grafana.aff:
	touch $@
