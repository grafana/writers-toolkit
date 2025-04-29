# List of projects to provide to the make-docs script.
PROJECTS = writers-toolkit

# Test the nightly image.
export DOCS_IMAGE := grafana/docs-base:nightly

# Set the VALE_IMAGE to match the one defined in CI.
export VALE_IMAGE := $(shell sed -En 's, *image: (grafana/vale.*),\1,p' "$(shell git rev-parse --show-toplevel)/.github/workflows/validate-documentation.yml" | head -n 1)

export VALE_MINALERTLEVEL := warning

export WEBSITE_MOUNTS := true
