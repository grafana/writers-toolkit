---
date: "2023-10-05T10:43:16+01:00"
description: Understand GNU Make and see and example of Make targets.
review_date: "2024-05-29"
title: GNU Make help output
---

Every project keeps technical documentation that's published to the website in the `docs/sources` directory.
That documentation is published to the `content/docs` directory in the website repository.

Additionally, every project uses [GNU Make](https://www.gnu.org/software/make/) to perform tasks related to technical documentation.
To learn more about GNU Make, refer to [GNU Make Manual](https://www.gnu.org/software/make/manual/).

To see a list of targets and their descriptions, run `make help` from the `docs/` directory.
If you are in the website repository, run `make help` from the root of the repository instead.

The output is similar to the following:

```console
Usage:
  make <target>

Targets:
  help             Display this help.
  docs-rm          Remove the docs container.
  docs-pull        Pull documentation base image.
  make-docs        Fetch the latest make-docs script.
  docs             Serve documentation locally, which includes pulling the latest `DOCS_IMAGE` (default: `grafana/docs-base:latest`) container image. To not pull the image, set `PULL=false`.
  docs-debug       Run Hugo web server with debugging enabled. TODO: support all SERVER_FLAGS defined in website Makefile.
  vale             Run vale on the entire docs folder which includes pulling the latest `VALE_IMAGE` (default: `grafana/vale:latest`) container image. To not pull the image, set `PULL=false`.
  update           Fetch the latest version of this Makefile and the `make-docs` script from Writers' Toolkit.
```
