---
description: Get started with the Writer's Toolkit
menuTitle: Get started
title: Get started with the Writer's Toolkit
weight: 50
aliases:
  - /docs/writers-toolkit/get-started/
---

# Get started with the Writer's Toolkit

This guide provides includes everything you need to complete your documentation project from start to finish.
In short, it's a `tl;dr` on Writer’s Toolkit to quickly point you in the right direction.

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

### 1. Plan the work

| Action                                                           | Considerations                                                                                                                                                                                                                    |
| ---------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Decide what kind of documentation deliverables you need          | What do you want a user to accomplish?                                                                                                                                                                                            |
| Know which release life cycle stage the product or feature is at | Review the [Release cycle for Grafana Labs](/docs/release-life-cycle/) <br /> Check in at the [Grafana Community Slack #docs channel](https://grafana.slack.com/archives/CNCRV74GP) to learn about docs for each of these stages. |
| Consider content reuse, if applicable                            | Read the [Reuse strategy documentation]({{< relref "../write/reuse-content/reuse-directories" >}}) <br />Check in at #docs to find out how the docs team can support you in this.                                                 |

### 2. Create the structure

| Action                             | Considerations                                                         |
| ---------------------------------- | ---------------------------------------------------------------------- |
| Add docs to your GitHub repository | Add a `docs/sources` folder to your repo, if it doesn’t already exist. |
| Create the folders and structure   | Use the content framework for consistency across product docs.         |

### 3. Draft the docs

| Action                                         | Considerations                                                                                                                                                                                   |
| ---------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Learn how to best write and organize your docs | Read the [topic types]({{< relref "../structure/topic-types" >}}) topic and sub-topics in the Writers’ Toolkit.                                                                                  |
| Add front matter                               | [Front matter]({{< relref "../write/front-matter" >}})der, and other key information to Hugo (our publishing platform). Learn more about [Front matter]({{< relref "../write/front-matter" >}}). |
| Add images and media                           | [Images and media]({{< relref "../write/image-guidelines" >}})                                                                                                                                   |
| Use links and references                       | [Links and references]({{< relref "../write/references" >}})                                                                                                                                     |
| Add code samples                               | [Code samples]({{< relref "../write/markdown-guide#code-blocks" >}})                                                                                                                             |
| Refer to our style guide                       | [Style guide]({{< relref "../write/style-guide/style-conventions" >}})                                                                                                                           |

### 4. Review the docs

| Action                                 | Considerations                                                                                                                                                                                      |
| -------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Build the docs locally                 | See the [build locally guidelines]({{< relref "../review/run-a-local-webserver" >}}) and check for errors, particularly in linking.                                                                 |
| Open a PR and tag the docs team        | Use the label type/docs.                                                                                                                                                                            |
| Use the docs validator and vale linter | These are automatic checks from within the PR itself. Learn more about the [docs-validator]({{< relref "../review/doc-validator" >}}) and the [Vale linter]({{< relref "../review/lint-prose" >}}). |

### 5. Publish the docs

| Action                                       | Considerations                                                                                                                                                                    |
| -------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Review publishing options                    | Engage #docs to discuss the publishing options for different stages in the release cycle.                                                                                         |
| Backport, if required                        | See the [backporting guidelines]({{< relref "../review/backporting" >}}).                                                                                                         |
| Add What’s New or release notes, if required | See the [Contribute to release notes guidelines]({{< relref "../contribute-documentation/contribute-release-notes" >}}). \*In Grafana, use the label add-to-whats-new in your PR. |
