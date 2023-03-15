---
title: Documentation structure
description: How to organize concepts and tasks in the repository.
aliases:
  - /docs/writers-toolkit/latest/writing-guide/documentation-structure/
weight: 300
menuTitle: Documentation structure
keywords:
  - information architecture
  - structure
---

# Documentation structure

The Grafana Labs Docs team makes intentional decisions about how to organize and structure product documentation. With well-structured content, users can find what they need quickly and easily.

Before you begin contributing to documentation, it is important to understand the structure of the content.

According to STC’s [Information Design and Information Architecture: Why Technical Communicators Should Care About These Fields](https://www.stc.org/intercom/2022/05/information-design-and-information-architecture-why-technical-communicators-should-care-about-these-fields/),
information architecture is "...the practice of organizing, structuring, and labeling information to it’s easy to find, use, and understand..."

Generally, documentation structure determines how content is:

- Titled
- Grouped
- Combined (or not combined) with other, related content

The following examples are based on the Grafana OSS and Enterprise documentation.

## Structure of published content

Using the same content structure across documentation provides consistent experience. Topics are listed from high level to more specific. For example, a new Grafana user may wish to learn conceptual information first, so Introduction is listed before installation.

Nat all Topics are not used for every product. For example, Grafana OSS may use all of the headings, while Grafana Enterprise Traces only uses a subset.

This table provides a list of the high-level topics used for grouping content. For example, a conceptual page explaining metrics would go under the Introduction topic.

Italicized topics are optional and are usually found in specific contexts. For example, the Create, Manage, Monitor topics are used in Grafana OnCall but are not used in Grafana Tempo.

| Topic | Example link | Contains |
| --- | --- | --- |
| Introduction | Introduction to Grafana | Conceptual information, fundamentals, architecture, etc. |
| Get started | Get started with Tempo | Opinionated walk-throughs and tutorials |
| Set up | Set up Loki | System requirements, Set up, configure, upgrade, migrate, etc. |
| Configure |  | Configure may be it’s own directory if the number of pages |
| Create alerts | Create alerts for Grafana OnCall | Specific to operational products such as Grafana OnCall. Not used with backend database products like Tempo and Loki. |
| Manage alerts | Manage alerts for Grafana OnCall | Specific to operational products such as Grafana OnCall. Not used with backend database products like Tempo and Loki. |
| Monitor alerts | Monitor alerts for Grafana OnCall | Specific to operational products such as Grafana OnCall. Not used with backend database products like Tempo and Loki. |
| Integrate (with) product or Send data |  | How to set up data integrations, product integrations, data sources, clients, plugins, etc. |
| Query data | TraceQL query editor | Query languages, query tools, examples |
| Visualize data |  | Dashboard concepts and procedures |
| Alert |  |  |
| Monitor [product] |  | Information about using tools to monitor a Grafana product. |
| References | HTTP API Reference | API, configuration references, SDKs, etc. Material that is usually not procedural and infrequently used. |



### Table of contents levels

The table of contents consists of the following section levels.

![Grafana table of contents](grafana-toc.png)

**Top-level:** A table of contents top-level represents groups of features and functions of a product. The first step in contributing to the docs is to identify which top-level entity you will be contributing to.

> **Note:** You should not add a top-level entity to the table of contents. Reach out to the technical documentation team if it is not clear where your documentation belongs.

**Parent:** Each top-level entity has one or more parents, which are groups of related feature content. Parent topics assist users in navigating to child topics.

**Child:** This level of the information architecture includes includes concepts, tasks, or reference topics.

## Parent directory structure

Within the top-level directory, there is a parent directory.

The image below shows how the repository's `user-management` parent directory is structured.

- There is an `_index.md` file in the parent directory that serves as a landing page for the child topics. In most cases, `_index.md` contains conceptual content. For information about the types of conceptual content that you can add to the `_index.md` file, refer to [Concepts]({{< relref "../topic-types/concept/" >}}).
- There are also four task topics in the parent directory, each with a directory and `index.md` file.

For more information about how to write concepts, refer to [Concepts]({{< relref "../topic-types/concept/" >}}).
For more information about how to write tasks, refer to [Tasks]({{< relref "../topic-types/task/" >}}).

![Parent directory structure](parent-directory.png)

> **Note:** If a directory contains multiple pages or subdirectories, it is a branch bundle, and it must include an `_index.md` file. A directory containing only one page is a leaf bundle, and the content filename must be `index.md`.

## Pages and page bundles

Each web page generated by Hugo comes from one of three source files:
- `page/_index.md`: a Hugo branch bundle
- `page/index.md`: a Hugo leaf bundle
- `page.md`: a Hugo page

Although each of the preceding examples results in the same URL (`/page/`), Hugo works with each source file differently.

**Branch bundles** (`page/_index.md`) are required to produce page hierarchies.
For the `/page/subpage/` URL to generated, there must first be a `page/_index.md` branch bundle source file.

**Leaf bundles** (`page/index.md`) are required to conveniently and consistently bundle page assets.
To refer to a style sheet `page/style.css` with the link `./style.css`, the link must be in a `page/index.md` leaf bundle source file.

**Leaf bundles** (`page/index.md`) are also required to mount content from one part of the site to another using Hugo mounts.
Hugo mounts use a virtual filesystem before site generation, and only directories can be mounted.
To mount the page `/page/` to `/other/page/`, there must first be a `page/index.md` leaf bundle source file.

If you don't know whether you need to mount a leaf bundle, you probably don't and **default to using pages**.

**Pages** (`page.md`) are any source files that are not leaf or branch bundles.
It is convenient to use pages when none of the behaviors of leaf or branch bundles are required, as it can be easier to distinguish two source files in some text editors or IDEs.

For more information about branch bundles and leaf bundles, refer to [Page bundles](https://gohugo.io/content-management/page-bundles/).
