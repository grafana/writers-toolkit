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

DICTIONARIES := Grafana/styles/config/dictionaries/en_US-grafana.dic
TESTS := fixtures/Grafana
STYLES := Grafana/styles/Grafana/GrafanaGoogle Grafana/styles/Grafana/Readability
RULES := Grafana/styles/Grafana/Acronyms.yml Grafana/styles/Grafana/AmazonProductNames.yml Grafana/styles/Grafana/ApacheProjectNames.yml Grafana/styles/Grafana/GoogleProductNames.yml Grafana/styles/Grafana/Headings.yml Grafana/styles/Grafana/ProductPossessives.yml Grafana/styles/Grafana/WordList.yml
DOCUMENTATION := ../docs/sources/review/lint-prose/rules.md

all: $(DICTIONARIES) $(DOCUMENTATION) $(RULES) $(STYLES) $(TESTS) ## Build all the Grafana Vale rules and styles.
$(DOCUMENTATION): $(wildcard Grafana/styles/Grafana/*) $(wildcard ../tools/cmd/generate-documentation/*)
	cd ../tools/ && go run ./cmd/generate-documentation ../ tools/cmd/generate-documentation vale/Grafana docs/sources/review/lint-prose/rules.md
	prettier -w $(DOCUMENTATION)

Grafana/styles/config/dictionaries/en_US-grafana.%: dictionary.libsonnet
	$(MAKE) -C $(@D) $(@F)

.PHONY: grafana/vale
grafana/vale: ## Builds a container image for Vale with the Grafana style loaded.
grafana/vale: Grafana/styles/config/dictionaries/en_US-grafana.aff Grafana/styles/config/dictionaries/en_US-grafana.dic Grafana/styles/Grafana/GrafanaGoogle .vale.ini
	$(PODMAN) build \
		--platform linux/x86_64,linux/arm64 \
		--progress plain \
		--tag grafana/vale:$(shell git rev-parse --short HEAD) \
		--tag grafana/vale:latest \
			.

.PHONY: grafana/vale/push
grafana/vale/push: ## Builds and pushes container image for Vale with the Grafana style loaded.
grafana/vale/push: dictionaries/en_US-grafana.aff dictionaries/en_US-grafana.dic Grafana/styles/Grafana/Google .vale.ini
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

.PHONY: Grafana/styles/Grafana/GrafanaGoogle
Grafana/styles/Grafana/GrafanaGoogle: ## Generate the Grafana style from the Google style.
Grafana/styles/Grafana/GrafanaGoogle: google.jsonnet
	rm -f "$$(jsonnet -Se 'std.join(" ", std.objectFields(import "google.jsonnet"))')"
	jsonnet -S -m . google.jsonnet
	prettier -w $@*.yml

Grafana/styles/Grafana/%.yml: ## Generate Grafana style rules from the corresponding Jsonnet file.
Grafana/styles/Grafana/%.yml: %.jsonnet
	jsonnet -S $< > $@
	prettier -w $@

.PHONY: Grafana/styles/Grafana/Readability
Grafana/styles/Grafana/Readability: ## Generate the Grafana style from the Readability style.
Grafana/styles/Grafana/Readability: readability.jsonnet
	rm -f $@*.yml
	jsonnet -S -m . readability.jsonnet
	prettier -w $@*.yml

fixtures/Grafana/%: ## Generate the scaffolding to test a new Grafana Vale rule.
fixtures/Grafana/%:
	mkdir -p $@
	printf 'MinAlertLevel = suggestion\n[*.md]\nGrafana.%s = YES\n' $(@F) > $@/.vale.ini
	touch $@/test{,in}valid.{golden,md}

.PHONY: fixtures/Grafana
fixtures/Grafana: ## Recreate the golden files for all the Grafana Vale rule fixtures.
fixtures/Grafana:
	find $@ -maxdepth 1 -mindepth 1 -print -exec /bin/bash -c 'cd {}; vale --output=line testvalid.md > testvalid.golden; vale --output=line testinvalid.md > testinvalid.golden; cd -' \;

.PHONY: Grafana.zip
Grafana.zip: ## Create an archive of the Grafana Vale style.
Grafana.zip:
	rm -f $@
	zip -r Grafana.zip Grafana
