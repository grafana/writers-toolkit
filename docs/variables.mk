# List of projects to provide to the make-docs script.
PROJECTS = writers-toolkit

# Set the DOC_VALIDATOR_IMAGE to match the one defined in CI.
export DOC_VALIDATOR_IMAGE := $(shell sed -En 's, *image: "(grafana/doc-validator.*)",\1,p' "$(shell git rev-parse --show-toplevel)/.github/workflows/test.yml")

# Use alternative image until make-docs 3.0.0 is rolled out.
export DOCS_IMAGE := grafana/docs-base:dbd975af06
