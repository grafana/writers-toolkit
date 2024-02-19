---
title: GNU Make help output
description: Understand GNU Make and see and example of Make targets.
---

All projects keep website documentation in the `docs/sources` directory.
Additionally, all projects use [GNU Make](https://www.gnu.org/software/make/) to perform tasks related to that documentation.
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
  doc-validator    Run doc-validator on the entire docs folder.
  vale             Run vale on the entire docs folder.
  update           Fetch the latest version of this Makefile and the `make-docs` script from Writers' Toolkit.
```
