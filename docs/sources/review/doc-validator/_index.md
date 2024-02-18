---
description: How to validate technical documentation with the doc-validator tool.
menuTitle: Automated validation
title: Automated validation with doc-validator
weight: 300
aliases:
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/validate-technical-documentation/
  - /docs/writers-toolkit/review/doc-validator/
---

# Automated validation with doc-validator

{{< docs/shared source="writers-toolkit" lookup="make-help.md" >}}

To validate technical documentation with `doc-validator`, run `make doc-validator` from the `docs/` directory.

## Error codes

All errors include an error code.
You can find documentation for each error code in [Errata for `doc-validator`](https://grafana.com/docs/writers-toolkit/review/doc-validator/errata/).

## Run on specific files

The script that invokes `doc-validator` mounts projects using the Hugo website structure.
All projects are subdirectories of `/hugo/content/docs/`.

To run `doc-validator` on specific files, provide the _`DOC_VALIDATOR_INCLUDE`_ argument to your `make` command.
It's value is a regular expression that the tool matches against file paths.
`doc-validator` only lints the paths that the regular expression matches.

### Writers' Toolkit repository, `/docs/sources/write/` directory

When in the writers-toolkit repository, to only validate content in the `/docs/sources/write/` directory, run the following command:

```console
make doc-validator DOC_VALIDATOR_INCLUDE='^/hugo/content/docs/writers-toolkit/write/.*$'
```

#### Explanation of the regular expression

- `^` matches the empty string at the beginning of a line, anchoring the regular expression to the start of the input.
- Writers' Toolkit isn't a versioned project and its documentation is synced directly into the `/hugo/content/docs/writers-toolkit/` directory in the website repository.

  `make doc-validator` puts content in the same directories as the website.

- `write/` matches literal string, only matching paths that contain `write/`.
- `.*` matches zero or more additional characters, matching all pages under the `write/` directory.
- `$` matches the empty string at the end of a line, effectively anchoring the regular expression to the end of the input.

### Grafana repository, `/docs/sources/developers/plugins/` directory

When in the Grafana repository, to only validate content in the `/docs/sources/developers/plugins/` directory, run the following command:

```console
make doc-validator DOC_VALIDATOR_INCLUDE='^/hugo/content/docs/grafana/latest/developers/plugins.*$'
```

#### Explanation of the regular expression

- `^` matches the empty string at the beginning of a line, anchoring the regular expression to the start of the input.
- Grafana is a versioned project.
  By default the script puts Grafana content into the `/hugo/content/docs/grafana/latest/` directory.
- `developers/plugins/` matches the literal string, only matching paths that contain `developers/plugins/`
- `.*` matches zero or more additional characters, matching all pages under the `developers/plugins/` directory.
- `$` matches the empty string at the end of a line, anchoring the regular expression to the end of the input.
