---
aliases:
  - /docs/writers-toolkit/get-started/
date: "2023-07-11T10:24:20-04:00"
description:
  Everything you need to complete your documentation project from start
  to finish.
menuTitle: Get started
review_date: "2024-09-04"
title: Get started with writing documentation
weight: 150
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

If you have questions:

- and you are a Grafana Labs employee, reach out in the `#docs` channel on the internal Slack workspace.
- and you aren't a Grafana Labs employee, reach out in the `#docs` channel on the [community Slack workspace](https://grafana.slack.com/archives/CNCRV74GP).

To get started:

1. [Plan the work](#plan-the-work)
1. [Create the structure](#create-the-structure)
1. [Draft the documentation](#draft-the-documentation)
1. [Review the documentation](#review-the-documentation)
1. [Publish the documentation](#publish-the-documentation)

### Plan the work

1. Decide what kind of documentation deliverables you need.
   What do you want a user to accomplish?

1. Know the release life cycle stage of the product or feature.
   Review the [Release cycle for Grafana Labs](https://grafana.com/docs/release-life-cycle/).

### Create the structure

1. Add documentation to your GitHub repository.
   Add a `docs/sources` directory to your repository, if it doesn't already exist.

1. Create the folders and structure.
   Use the content framework for consistency across product documentation.
   For more information about the content framework, refer to [Documentation structure](https://grafana.com/docs/writers-toolkit/structure/).

### Draft the documentation

1. Learn how to best write and organize your documentation.
   Read the [topic types](https://grafana.com/docs/writers-toolkit/structure/topic-types/) topic and sub-topics in the Writers' Toolkit.

1. Add front matter.
   Read the [front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/) documentation.
1. Add images and media.
   Read the [image, diagram, screenshot, and video guidelines](https://grafana.com/docs/writers-toolkit/write/image-guidelines/) documentation.
1. Use links.
   Refer to the [links](https://grafana.com/docs/writers-toolkit/write/links/) documentation.
1. Add code samples.
   To add samples with Markdown, refer to [Code blocks](https://grafana.com/docs/writers-toolkit/write/markdown-guide/#code-blocks).
   For style guidance, refer to [Code examples](https://grafana.com/docs/writers-toolkit/write/style-guide/write-for-developers/#code-examples).
   For samples in multiple languages, refer to [Code](https://grafana.com/docs/writers-toolkit/write/shortcodes/#code).
1. Refer to the [Style conventions](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/).

### Review the documentation

1. Build the documentation locally.
   For instructions, refer to [Test documentation changes](https://grafana.com/docs/writers-toolkit/review/test-documentation-changes/).
1. Open a pull request and use the label `type/docs`.
1. Use the Vale prose linter.
   These are often automatic checks from within the PR itself.
   To learn more about the Vale prose linter, refer to [Lint prose with the Vale linter](https://grafana.com/docs/writers-toolkit/review/lint-prose/).

### Publish the documentation

<!-- vale Grafana.Timeless = NO -->

1. Review publishing options.
   If you are a Grafana Labs employee, reach out in the `#docs` channel on the internal Slack workspace.
   If you aren't a Grafana Labs employee, reach out in the `#docs` channel on the [community Slack workspace](https://grafana.slack.com/archives/CNCRV74GP).
1. Backport, if required.
   For guidance, refer to [Backport changes](https://grafana.com/docs/writers-toolkit/review/backport-changes/).
1. Add What's new or release notes, if required.
   For guidance, refer to [Contribute to What's new or release notes](https://grafana.com/docs/writers-toolkit/contribute/release-notes/).

<!-- vale Grafana.Timeless = YES -->
