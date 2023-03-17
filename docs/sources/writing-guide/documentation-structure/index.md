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

The Grafana Labs Docs team makes intentional decisions about how to organize and structure product documentation. The topic levels discussed on this page reflect common user goals. For example, a first-time Grafana user needs to understand the basic concepts before installing the product.

Before you begin contributing to documentation, it is important to understand the structure of the content.

## User goals and documentation structure

Any documentation you create should be focused on what your user's goals are. Refer to [Focus on user goals]({{< relref "../../style-guide/style-conventions" >}}) for more information.

Your documentation structure should reflect these user goals. Think about what your users want to do, what they need to know, and how they can accomplish the tasks.

This approach applies not only to content on a page but also to how you organize a set of documentation, whether it's for a product or a feature.

With well-structured content, you can find what you need quickly and easily. Topics flow in a logical progression.

<!-- Commenting this paragraph out until we have an actual section on information architecture.

According to STC’s [Information Design and Information Architecture: Why Technical Communicators Should Care About These Fields](https://www.stc.org/intercom/2022/05/information-design-and-information-architecture-why-technical-communicators-should-care-about-these-fields/),
information architecture is "...the practice of organizing, structuring, and labeling information to it’s easy to find, use, and understand..."
-->

## Structure of published content

Generally, documentation structure determines how content is:

- Titled
- Grouped
- Combined (or not combined) with other, related content

A standardized content structure that spans sets of documentation provides a consistent user experience.
Information flows from a higher (less specific) level to a lower (more specific) level.
For example, a new Grafana user wants to learn conceptual information first, so _Introduction_ comes before _Installation_.

### Use the topics you need

Depending on your product design and maturity, you may not need every topic listed below.
If a topic does not apply to your project, you don't need to use it.
For Grafana OSS for example, you might use all of the headings.
For Grafana Enterprise Traces for example, you might only use a subset of topics.

Some topics are optional and are usually found in specific contexts.
For example, the _Create_, _Manage_, and _Monitor_ topics are used in Grafana OnCall but are not used in Grafana Tempo.

### Topic list

This table provides a list of the high-level topics used for grouping content. When writing new content, consider where it should appear given this content structure. For example, a conceptual page explaining metrics would go under the Introduction topic.


| Topic | Example link | Contents |
| --- | --- | --- |
| _Introduction_ | [Introduction to Grafana](/docs/grafana/latest/introduction/) | Conceptual, fundamental, or architectural information.  |
| _Get started_ | [Get started with Tempo](/docs/tempo/latest/getting-started/) | Opinionated walk-throughs and tutorials. |
| _Set up_ | [Set up Grafana](/docs/grafana/latest/setup-grafana/) | System requirements, and subsections titled _Set up_, _Configure_, _Upgrade_, or _Migrate_. |
| _Configure_ | [Configure Grafana Agent](/docs/agent/latest/configuration/) | _Configure_ can be its own section directory if the number of pages warrants it. Making this determination is not an exact science; use your best judgement. |
| _Create alerts_ | [Create a Grafana managed alert rule](/docs/grafana-cloud/alerting/alerting-rules/create-grafana-managed-rule/) | Specific to Grafana Ops products (Alerting, Oncall, Incident, SLOs). The word _alert_ may be changed, depending upon the product. If used with Grafana SLO, this then topic would be _Create SLO_. Do not use with backend database products, such as Tempo and Loki. Use _Alerts_ instead, and refer to an operational product for details. |
| _Manage alerts_ | | Specific to Grafana Ops products (Alerting, Oncall, Incident, SLOs). Do not use with backend database products, such as Tempo and Loki. Use _Alert_ instead, and refer to an operational product for details. |
| _Monitor alerts_ | | Specific to Grafana Ops products (Alerting, Oncall, Incident, SLOs). Do not use with backend database products, such as Tempo and Loki. Use _Alert_ instead, and refer to an operational product for details.  |
| _Integrate (with) [product]_ or _Send data_ | [Connect your datat to Grafana Cloud](/docs/grafana-cloud/data-configuration/get-started-data/) | How to set up data integrations, product integrations, data sources, clients, plugins, and more. |
| _Query data_ | [TraceQL](/docs/tempo/latest/traceql/) | Query languages, query tools, and examples. |
| _Visualize data_ | [Dashboards](/docs/grafana/latest/dashboards/) | Dashboard concepts and procedures. Link to the definitive source of dashboard documentation, rather than duplicating the information here. |
| _Alert_ | [Alerting rules](/docs/loki/latest/rules/) | This topic level is used for pages that discuss alerting features, like Grafana Loki's alerting rules. It provides a place for alerting content that is not specific to the Grafana Operations products. |
| _Monitor [product]_ | [Monitor Mimir](/docs/mimir/latest/operators-guide/monitor-grafana-mimir/) | Information about using tools to monitor a Grafana Labs product. |
| _References_ | [HTTP API reference](/docs/grafana-cloud/api-reference/http-api/) | APIs, configuration references, SDKs, and more. Material that is usually not procedural. |



## Table of contents levels

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
