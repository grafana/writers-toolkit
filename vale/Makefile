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

DICTIONARIES := dictionaries/en_US-grafana.dic
TESTS := fixtures/Grafana
STYLES := Grafana/Google Grafana/Readability
RULES := Grafana/Acronyms.yml Grafana/AmazonProductNames.yml Grafana/ApacheProjectNames.yml Grafana/_GoogleProductNames.yml Grafana/Headings.yml Grafana/ProductPossessives.yml
DOCUMENTATION := ../docs/sources/review/lint-prose/rules.md

all: $(DICTIONARIES) $(DOCUMENTATION) $(RULES) $(STYLES) $(TESTS) ## Build all the Grafana Vale rules and styles.
$(DOCUMENTATION): $(wildcard Grafana/*)
	cd ../tools/ && go run ./cmd/generate-documentation ../ vale/Grafana docs/sources/review/lint-prose/rules.md
	prettier -w $(DOCUMENTATION)

dictionaries/en_US-grafana.%: dictionary.jsonnet
	$(MAKE) -C $(@D) $(@F)

.PHONY: grafana/vale
grafana/vale: ## Builds a container image for Vale with the Grafana style loaded.
grafana/vale: dictionaries/en_US-grafana.aff dictionaries/en_US-grafana.dic Grafana/Google .vale.ini
	$(PODMAN) build \
		--platform linux/x86_64,linux/arm64 \
		--progress plain \
		--tag grafana/vale:$(shell git rev-parse --short HEAD) \
		--tag grafana/vale:latest \
			.

.PHONY: grafana/vale/push
grafana/vale/push: ## Builds and pushes container image for Vale with the Grafana style loaded.
grafana/vale/push: dictionaries/en_US-grafana.aff dictionaries/en_US-grafana.dic Grafana/Google .vale.ini
	$(PODMAN) build \
		--platform linux/x86_64,linux/arm64 \
		--progress plain \
		--push \
		--tag grafana/vale:$(shell git rev-parse --short HEAD) \
		--tag grafana/vale:latest \
			.

.PHONY: sync
sync: ## Update the vendored Google style.
sync:
	cd $(GIT_ROOT) && vale sync; cd -
	prettier -w .

.vale.ini ../.vale.ini: ## Generate Vale INI configuration from the Jsonnet data template.
.vale.ini ../.vale.ini: .vale.jsonnet
	jsonnet -Se "(import '.vale.jsonnet').container" | sed '$${/^$$/d;}' > .vale.ini
	jsonnet -Se "(import '.vale.jsonnet').repository" | sed '$${/^$$/d;}' > ../.vale.ini

.PHONY: Grafana/Google
Grafana/Google: ## Generate the Grafana style from the Google style.
Grafana/Google: google.jsonnet
	rm -f $@*.yml
	jsonnet -S -m . google.jsonnet
	prettier -w $@*.yml

Grafana/%.yml: ## Generate Grafana style rules from the corresponding Jsonnet file.
Grafana/%.yml: %.jsonnet
	jsonnet -S $< > $@
	prettier -w $@

.PHONY: Grafana/Readability
Grafana/Readability: ## Generate the Grafana style from the Readability style.
Grafana/Readability: readability.jsonnet
	rm -f $@*.yml
	jsonnet -S -m . readability.jsonnet
	prettier -w $@*.yml

fixtures/Grafana/%: ## Generate the scaffolding to test a new Grafana Vale rule.
fixtures/Grafana/%:
	mkdir -p $@
	printf 'StylesPath = ../../../\nMinAlertLevel = suggestion\n[*.md]\nGrafana.%s = YES\n' $(@F) > $@/.vale.ini
	touch $@/test{,in}valid.{golden,md}

.PHONY: fixtures/Grafana
fixtures/Grafana: ## Recreate the golden files for all the Grafana Vale rule fixtures.
fixtures/Grafana:
	find $@ -maxdepth 1 -mindepth 1 -print -exec /bin/bash -c 'cd {}; vale --output=line testvalid.md > testvalid.golden; vale --output=line testinvalid.md > testinvalid.golden; cd -' \;
