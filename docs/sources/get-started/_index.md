---
aliases:
  - /docs/writers-toolkit/get-started/
description: Everything you need to complete your documentation project from start to finish.
date: 2024-02-14
menuTitle: Get started
title: Get started with writing documentation
weight: 50
---

# Get started with writing documentation

This guide includes everything you need to complete your documentation project from start to finish.

## Goals

This guide:

- Teaches you how to successfully document products and features
- Provides a step-by-step guide from project kick-off to release
- Introduces you to the craft of creating good documentation

## Audience

- Software developers who need to create documentation alongside code
- Other technical roles that create and contribute to documentation

## Before you begin

- Consider what you want someone reading your documentation to understand or accomplish

## Project checklist

If you have questions, you can ask them in the [Grafana Community Slack #docs channel](https://grafana.slack.com/archives/CNCRV74GP).

### 1. Plan the work

1. Decide what kind of documentation deliverables you need.
   What do you want a user to accomplish?

1. Know the release life cycle stage of the product or feature.
   Review the [Release cycle for Grafana Labs](https://grafana.com/docs/release-life-cycle/).

### 2. Create the structure

1. Add docs to your GitHub repository.
   Add a `docs/sources` directory to your repository, if it doesn’t already exist.

1. Create the folders and structure.
   Use the content framework for consistency across product documentation.

### 3. Draft the documentation

1. Learn how to best write and organize your documentation.
   Read the [topic types](https://grafana.com/docs/writers-toolkit/structure/topic-types/) topic and sub-topics in the Writers’ Toolkit.

1. Add front matter.
   Read the [front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/) documentation.
1. Add images and media.
   Read the [image, diagram, and screenshot guidelines](https://grafana.com/docs/writers-toolkit/write/image-guidelines/) documentation.
1. Use links.
   Refer to the [links](https://grafana.com/docs/writers-toolkit/write/links/) documentation.
1. Add code samples.
   To add samples with Markdown, refer to [Code blocks](https://grafana.com/docs/writers-toolkit/write/markdown-guide/#code-blocks).
   For style guidance, refer to [Code examples](https://grafana.com/docs/writers-toolkit/write/style-guide/write-for-developers/#code-examples).
   For samples in multiple languages, refer to [Code](https://grafana.com/docs/writers-toolkit/write/shortcodes/#code).
1. Refer to the [Style guide](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/).

### 4. Review the docs

1. Build the docs locally.
   For instructions, refer to [Test documentation changes](https://grafana.com/docs/writers-toolkit/review/run-a-local-webserver/).
1. Open a pull request and use the label `type/docs`.
1. Use `doc-validator` and the Vale prose linter.
   These are often automatic checks from within the PR itself.
   To learn more about [`doc-validator`], refer to [Automated validation with doc-validator](https://grafana.com/docs/writers-toolkit/review/doc-validator/).
   To learn more about the Vale prose linter, refer to [Lint prose with the Vale linter](https://grafana.com/docs/writers-toolkit/review/lint-prose/).

### 5. Publish the docs

<!-- vale Grafana.Timeless = NO -->

1. Review publishing options.
   If you are a Grafana Labs employee, reach out in the #docs channel on the internal Slack workspace.
   If you aren't a Grafana Labs employee, reach out in the #docs channel on the [community Slack workspace](https://grafana.slack.com/archives/CNCRV74GP).
1. Backport, if required.
   For guidance, refer to [Backport changes](https://grafana.com/docs/writers-toolkit/review/backporting/).
1. Add What’s new or release notes, if required.
   For guidance, refer to [Contribute to What’s new or release notes](https://grafana.com/docs/writers-toolkit/contribute-documentation/contribute-release-notes/).

<!-- vale Grafana.Timeless = YES -->
