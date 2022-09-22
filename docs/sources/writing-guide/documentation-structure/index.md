---
title: Documentation structure
description: How to organize concepts and tasks in the repository.
aliases:
  - /docs/writers-toolkit/latest/writing-guidelines/documentation-structure/
weight: 300
menuTitle: Documentation structure
keywords:
  - information architecture
  - structure
---

# Documentation structure

The technical writing team makes very intentional decisions about how we organize and structure product documentation. 
Well-structured content enables users to find what they need, quickly and reliably.

The way in which we organize content is referred to as the information architecture (IA). 
The IA of each set of product documentation varies, so it is important that you understand the structure of the content before you start contributing.

In general, the IA determines how content is:
- Titled
- Grouped
- Combined (or not combined) with other, related content

The following examples are based on the Grafana OSS and Enterprise documentation.

## Information architecture

The information architecture consists of the following levels.

![Grafana table of contents](grafana-toc.png)

**Top-level:** The top-most level of the table of contents represents groups of product features and functionality. 
When you contribute to the docs, first identify which top-level entity you are contributing to.

> **Note:** Do not add a top-level entity to the table of contents. 
If it is not obvious where your documentation belongs, reach out to the technical documentation team.

**Parent:** Each root item contains one or more parents, which are groups of related feature content.
The parent level helps users navigate to child topics.

**Child:** This level of the information architecture includes includes concepts, tasks, or reference topics.

## Parent directory structure

A parent directory exists within the root directory.

The following image shows the `user-management` parent directory as it is structured in the repository.

- The parent directory contains an `_index.md` file that functions as a landing page for the child topics. 
The `_index.md` file typically contains conceptual content.
For more information about the kinds of conceptual content that you can add to the `_index.md` file, refer to [Concepts]({{< relref "../topic-types/#concepts" >}}).
- The parent directory also includes four task topics, each with its own directory and `index.md` file.

For more information about how to write concepts, refer to [Concepts]({{< relref "../topic-types/#concepts" >}}).
For more information about how to write tasks, refer to [Tasks]({{< relref "../topic-types/#tasks" >}}).

![Parent directory structure](parent-directory.png)

> **Note:** If a directory contains multiple pages or subdirectories, then it is considered a branch bundle and the index filename must be `_index.md`. 
If a directory contains only one page, then it is a leaf bundle and the content filename must be `index.md`.

## Useful links

- For more information about branch bundles and leaf bundles, refer to [Page bundles](https://gohugo.io/content-management/page-bundles/).
