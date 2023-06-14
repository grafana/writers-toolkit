---
description: How to validate technical documentation with the doc-validator tool.
menuTitle: Validate technical documentation
title: Validate technical documentation with doc-validator
---

# Validate technical documentation with doc-validator

Every project keeps technical documentation in the `docs/sources` directory.
Additionally, every project uses [GNU Make](https://www.gnu.org/software/make/) to perform tasks related to technical documentation.
To learn more about GNU Make, refer to [GNU Make Manual](https://www.gnu.org/software/make/manual/).

To see a list of targets and their descriptions, run `make` from the `docs/` directory.
The output is similar to the following:

```console
Usage:
  make <target>

Targets:
  help             Display this help.
  docs-rm          Remove the docs container.
  docs-pull        Pull documentation base image.
  make-docs        Fetch the latest make-docs script.
  docs             Serve documentation locally, which includes pulling the latest `DOCS_IMAGE` (default: `grafana/docs-base:latest`) container image. See also `docs-no-pull`.
  docs-no-pull     Serve documentation locally without pulling the `DOCS_IMAGE` (default: `grafana/docs-base:latest`) container image.
  docs-debug       Run Hugo web server with debugging enabled. TODO: support all SERVER_FLAGS defined in website Makefile.
  doc-validator    Run docs-validator on the entire docs folder.
  doc-validator/%  Run doc-validator on a specific path. To lint the path /docs/sources/administration, run 'make doc-validator/administration'.
  docs.mk          Fetch the latest version of this Makefile from Writers' Toolkit.
```

To validate technical documentation with `doc-validator`, run `make doc-validator` from the `docs/` directory.
Any linting errors are logged by the tool as JSON.

For more human-readable output, pipe the output to [`jq`](https://jqlang.github.io/jq/).
For example:

```console
make -s doc-validator | jq -r '"ERROR: \(.location.path):\(.location.range.start.line):\(.location.range.start.column): \(.message)"'
```

## Error codes

All errors include an error code.
Each error code is documented in [Errata]({{< relref "./errata" >}}).

## Run on specific files

The script that invokes `doc-validator` mounts projects using the Hugo website structure.
All projects are subdirectories of `/hugo/content/docs/`.

To run `doc-validator` on specific files, provide the _`DOC_VALIDATOR_INCLUDE`_ argument to your `make` command.
It's value is a regular expression to be matched against file paths.

### Writers' Toolkit repository, `/docs/sources/writing-guide/` directory

When in the Writer's Toolkit repository, to only validate content in the `/docs/sources/writing-guide/` directory, run the following command:

```console
make doc-validator DOC_VALIDATOR_INCLUDE='^/hugo/content/docs/writers-toolkit/writing-guide/.*$'
```

#### Explanation of the regular expression

- `^` matches the empty string at the beginning of a line, effectively anchoring the regular expression to the start of the input.
- Writers' Toolkit is an unversioned project so the content is mounted directly into the `/hugo/content/docs/writers-toolkit/` directory.
- `writing-guide/` is appended to the literal string so hat only the content in the `/docs/sources/writing-guide/` directory is validated.
- `.*` matches zero or more additional characters, effectively matching any file paths in the `/hugo/content/docs/writers-toolkit/` directory.
- `$` matches the empty string at the end of a line, effectively anchoring the regular expression to the end of the input.

### Grafana repository, `/docs/sources/developers/plugins/` directory

When in the Grafana repository, to only validate content in the `/docs/sources/developers/plugins/` directory, run the following command:

```console
make doc-validator DOC_VALIDATOR_INCLUDE='^/hugo/content/docs/grafana/latest/developers/plugins.*$'
```

#### Explanation of the regular expression

- `^` matches the empty string at the beginning of a line, effectively anchoring the regular expression to the start of the input.
- Grafana is a versioned project.
  By default the script mounts Grafana content into the `/hugo/content/docs/grafana/latest/` directory.
- `developers/plugins/` is appended to the literal string so hat only the content in the `/docs/sources/developers/plugins/` directory is validated.
- `.*` matches zero or more additional characters, effectively matching any file paths in the `/hugo/content/docs/developers/plugins/` directory.
- `$` matches the empty string at the end of a line, effectively anchoring the regular expression to the end of the input.
