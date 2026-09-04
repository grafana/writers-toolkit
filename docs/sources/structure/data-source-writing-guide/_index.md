---
date: "2025-09-10T00:00:00+01:00"
description: Learn how to create, structure, and maintain documentation for Grafana data sources.
keywords:
  - data source
  - plugin
  - documentation structure
menuTitle: Data sources writing guide
review_date: "2027-09-04"
title: Data sources writing guide
weight: 500
---

# Data sources writing guide

This guide helps you create and maintain documentation for Grafana data sources.
It's written for the data source development team and anyone who documents a data source, whether you're publishing a brand-new plugin or updating an existing one.

Grafana offers hundreds of data sources.
Some are maintained by Grafana Labs, and others are maintained by Grafana Champions, partners, or community contributors.
This guide gives you a consistent structure and set of conventions so that every data source's documentation feels cohesive, no matter who writes it.

## Document every data source as a standalone plugin

Grafana data sources are moving to standalone plugin repositories, and core data sources are following the same path.
Write documentation for every data source as if it lives in a standalone plugin repository, regardless of where the docs currently live.

This means you always use:

- A flat file structure: one Markdown file per guide at the top level of `docs/sources/`, with no subfolders.
- Fully qualified links: link to sibling pages and Grafana product docs with complete `https://grafana.com/docs/...` URLs.

Following one model keeps every data source consistent and means the docs work correctly when a plugin moves to its own repository.

For the details, refer to [File structure and front matter](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/file-structure-and-front-matter/).

## Pages every data source needs

Structure each data source's documentation as a set of focused pages, each covering one task or topic.
This makes information easier to find than one long page that requires a lot of scrolling.

The following table lists the standard pages.
Create only the pages your data source needs: a small data source might need only an `_index.md`, while a complex one uses the full set.

| Page                | File                    | Purpose                                                             |
| ------------------- | ----------------------- | ------------------------------------------------------------------- |
| Overview            | `_index.md`             | Introduces the data source and links to the other pages.            |
| Install and upgrade | `install.md`            | Installation and upgrade steps. Enterprise plugins only.            |
| Configure           | `configure.md`          | How to add, configure, authenticate, and provision the data source. |
| Query editor        | `query-editor.md`       | How to build and run queries.                                       |
| Template variables  | `template-variables.md` | How to use template variables with the data source.                 |
| Annotations         | `annotations.md`        | How to add annotations from the data source.                        |
| Alerting            | `alerting.md`           | How to use the data source with Grafana Alerting.                   |
| Troubleshooting     | `troubleshooting.md`    | Common issues and how to resolve them.                              |

For example, the [MySQL data source documentation](https://grafana.com/docs/grafana/latest/datasources/mysql/) uses this structure, with separate Configure, Query editor, Template variables, Annotations, Alerting, and Troubleshooting pages.

## How to use this guide

Work through the following pages in order when you create documentation for a new data source, or jump to the one you need when you update an existing data source:

- [File structure and front matter](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/file-structure-and-front-matter/): Where files live, how they publish, the front matter template, and linking rules.
- [Page templates](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/page-templates/): What to include on each of the standard pages, with copy-and-adapt templates.
- [Style conventions](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/style-conventions/): Headings, voice, formatting, spelling, and other shared style rules.
- [Maintain and update documentation](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/maintain-and-update/): How to verify accuracy against source code and keep docs current.
- [Checklists](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/checklists/): Quick checklists for creating and reviewing data source documentation.
