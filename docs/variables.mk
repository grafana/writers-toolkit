# List of projects to provide to the make-docs script.
PROJECTS = writers-toolkit

# Test the nightly image.
export DOCS_IMAGE := grafana/docs-base:nightly

# Set the DOC_VALIDATOR_IMAGE to match the one defined in CI.
export DOC_VALIDATOR_IMAGE := $(shell sed -En 's, *image: (grafana/doc-validator.*),\1,p' "$(shell git rev-parse --show-toplevel)/.github/workflows/validate-documentation.yml")

# Set the VALE_IMAGE to match the one defined in CI.
export VALE_IMAGE := $(shell sed -En 's, *image: (grafana/vale.*),\1,p' "$(shell git rev-parse --show-toplevel)/.github/workflows/validate-documentation.yml")

export VALE_MINALERTLEVEL := warning

# Skip some doc-validator checks.
export DOC_VALIDATOR_SKIP_CHECKS := $(shell sed -En "s, *'--skip-checks=(.+)',\1,p" "$(shell git rev-parse --show-toplevel)/.github/workflows/validate-documentation.yml")
