---
description: Get started with Writers' Toolkit
menuTitle: Get started
title: Get started with Writers' Toolkit
weight: 50
aliases:
  - /docs/writers-toolkit/get-started/
---

# Get started with Writers' Toolkit

This guide includes everything you need to complete your documentation project from start to finish.
In short, it quickly points you in the right direction.

## Goals

By reading this guide, you will:

- Learn how to successfully document new products or features
- Have a step-by-step guide from project kick-off to release
- Be introduced to the craft of creating good documentation

## Audience

- Software developers who need to create docs alongside code
- Other technical roles that create and contribute to docs

## Before you begin

- Consider what you want someone reading your documentation to understand or accomplish

## Project checklist

If you have questions, you can ask them in the [Grafana Community Slack #docs channel](https://grafana.slack.com/archives/CNCRV74GP).

### 1. Plan the work

1. Decide what kind of documentation deliverables you need. What do you want a user to accomplish?

1. Know the release life cycle stage of the product or feature. Review the [Release cycle for Grafana Labs](/docs/release-life-cycle/).

1. Consider content reuse, if applicable. Read the [Reuse strategy documentation]({{< relref "../write/reuse-content/reuse-directories" >}}).

### 2. Create the structure

1. Add docs to your GitHub repository. Add a `docs/sources` folder to your repository, if it doesn’t already exist.

1. Create the folders and structure. Use the content framework for consistency across product docs.

### 3. Draft the docs

1. Learn how to best write and organize your docs | Read the [topic types]({{< relref "../structure/topic-types" >}}) topic and sub-topics in the Writers’ Toolkit.

1. Add front matter. Read the [Front matter]({{< relref "../write/front-matter" >}}) documentation, and other key information to Hugo (our publishing platform).

1. Add images and media. Read the [Images and media]({{< relref "../write/image-guidelines" >}}) documentation.

1. Use links. Refer to the [Links](https://grafana.com/docs/writers-toolkit/write/links/) documentation for guidance.

1. Add code samples. Refer to the [Code samples]({{< relref "../write/markdown-guide#code-blocks" >}}) documentation for guidance.

1. Refer to our [Style guide]({{< relref "../write/style-guide/style-conventions" >}}).

### 4. Review the docs

1. Build the docs locally. Refer to the [build locally guidelines]({{< relref "../review/run-a-local-webserver" >}}) and check for errors, particularly in linking.

1. Open a PR and tag the docs team. Use the label type/docs.

1. Use `doc-validator` and Vale linter.
   These are automatic checks from within the PR itself.
   Learn more about the [`doc-validator`]({{< relref "../review/doc-validator" >}}) and the [Vale linter]({{< relref "../review/lint-prose" >}}).

### 5. Publish the docs

1. Review publishing options. Engage #docs to discuss the publishing options for different stages in the release cycle.

1. Backport, if required. Refer to the [backporting guidelines]({{< relref "../review/backporting" >}}).

1. Add What’s New or release notes, if required. Refer to the [Contribute to release notes guidelines]({{< relref "../contribute-documentation/contribute-release-notes" >}}). In Grafana, use the label `add-to-whats-new` in your PR.
