---
aliases:
  - /docs/writers-toolkit/writing-guide/documentation-structure/
  - /docs/writers-toolkit/structure
description: How to organize concepts and tasks in the repository.
keywords:
  - information architecture
  - structure
menuTitle: Structure
title: Documentation structure
weight: 300
---

# Documentation structure

The Grafana Labs Docs squad makes intentional decisions about how to organize and structure product documentation.
The topic levels discussed on this page reflect common user goals.
For example, a first-time Grafana user needs to understand the basic concepts before installing the product.

Before you begin contributing to documentation, it's important to understand the structure of the content.

## User goals and documentation structure

When writing documentation, focus on what your user's goals are.
For more information, refer to [Focus on user goals](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#focus-on-user-goals).

Use the structure of your documentation to reflect the user’s goals.
Think about what your users want to do, what they need to know, and how they can accomplish the tasks.

This approach applies not only to content on a page but also to how you organize a set of documentation, whether it's for a product or a feature.

With well-structured content, you can find what you need quickly and easily. Topics flow in a logical progression.

<!--
Commenting this paragraph out until we have an actual section on information architecture.

https://github.com/grafana/writers-toolkit/issues/532

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

Depending on your product design and maturity, you may not need every topic:

- If a topic doesn't apply to your project, you don't need to use it.
- For Grafana OSS for example, you might use all of the headings.
- For Grafana Enterprise Traces for example, you might only use a subset of topics.

Some topics are optional and are usually found in specific contexts.
For example, the _Create_ and _Monitor_ topics are used in Grafana OnCall, but aren't used in Grafana Tempo.

### Topic list

You can use the following high-level topics to group content.
When writing new content, consider where it should appear given this content structure.
For example, a conceptual page explaining metrics would go under the _Introduction_ topic.

| Topic                                       | Example link                                                                                                                                        | Contents                                                                                                                                                                                                                                                                                                                                     |
| ------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| _Introduction_                              | [Introduction to Grafana](https://grafana.com/docs/grafana/latest/introduction/)                                                                    | Conceptual, fundamental, or architectural information.                                                                                                                                                                                                                                                                                       |
| _Get started_                               | [Get started with Tempo](https://grafana.com/docs/tempo/latest/getting-started/)                                                                    | Opinionated walk-throughs and tutorials.                                                                                                                                                                                                                                                                                                     |
| _Set up_                                    | [Set up Grafana](https://grafana.com/docs/grafana/latest/setup-grafana/)                                                                            | System requirements, and subsections titled _Set up_, _Configure_, _Upgrade_, or _Migrate_.                                                                                                                                                                                                                                                  |
| _Configure_                                 | [Configure Grafana Agent](https://grafana.com/docs/agent/latest/static/configuration/)                                                              | _Configure_ can be its own section directory if the number of pages warrants it. Making this determination is not an exact science; use your best judgement.                                                                                                                                                                                 |
| _Create alerts_                             | [Create a Grafana managed alert rule](https://grafana.com/docs/grafana-cloud/alerting-and-irm/alerting/alerting-rules/create-grafana-managed-rule/) | Specific to Grafana Ops products (Alerting, OnCall, Incident, SLOs). The word _alert_ may be changed, depending upon the product. If used with Grafana SLO, this then topic would be _Create SLO_. Do not use with backend database products, such as Tempo and Loki. Use _Alerts_ instead, and refer to an operational product for details. |
| _Manage alerts_                             | [Manage and monitor SLOs](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/manage/)                                                      | Specific to Grafana Ops products (Alerting, OnCall, Incident, SLOs). Do not use with backend database products, such as Tempo and Loki. Use _Alert_ instead, and refer to an operational product for details.                                                                                                                                |
| _Integrate (with) [product]_ or _Send data_ | [Connect your data to Grafana Cloud](https://grafana.com/docs/grafana-cloud/monitor-infrastructure/get-started-data/)                               | How to set up data integrations, product integrations, data sources, clients, plugins, and more.                                                                                                                                                                                                                                             |
| _Query data_                                | [TraceQL](https://grafana.com/docs/tempo/latest/traceql/)                                                                                           | Query languages, query tools, and examples.                                                                                                                                                                                                                                                                                                  |
| _Visualize data_                            | [Dashboards](https://grafana.com/docs/grafana/latest/dashboards/)                                                                                   | Dashboard concepts and procedures. Link to the definitive source of dashboard documentation, rather than duplicating the information here.                                                                                                                                                                                                   |
| _Alert_                                     | [Alerting and recording rules](https://grafana.com/docs/loki/latest/alert/)                                                                         | This topic level is used for pages that discuss alerting features, like Grafana Loki's alerting rules. It provides a place for alerting content that's not specific to the Grafana Operations products.                                                                                                                                      |
| _Manage [product]_                          | [Manage users and teams with Grafana OnCall](https://grafana.com/docs/grafana-cloud/alerting-and-irm/oncall/manage/user-and-team-management/)       | Information about managing a Grafana Labs product. For example, content in this topic helps you view, edit, and iterate on the Grafana product you installed.                                                                                                                                                                                |
| _Monitor [product]_                         | [Monitor Mimir](https://grafana.com/docs/mimir/latest/operators-guide/monitor-grafana-mimir/)                                                       | Information about using tools to monitor a Grafana Labs product.                                                                                                                                                                                                                                                                             |
| _References_                                | [HTTP API reference](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/http-api/)                                            | APIs, configuration references, SDKs, and more. Material that is usually not procedural.                                                                                                                                                                                                                                                     |

## Table of contents levels

The table of contents consists of the following section levels.

![Grafana table of contents](/media/docs/writers-toolkit/grafana-toc.png)

<!-- vale Grafana.GoogleWill = NO -->
<!-- Valid use of future tense. -->

**Top-level:** A table of contents top-level represents groups of features and functions of a product.
The first step in contributing to the documentation is to identify which top-level entity you will be contributing to.

<!-- vale Grafana.GoogleWill = YES -->

{{< admonition type="note" >}}
Don't add a top-level entity to the table of contents.
If you're not sure where your documentation belongs, contact the technical documentation team.
{{< /admonition >}}

**Parent:** Each top-level entity has one or more parents, which are groups of related feature content.
Parent topics assist users in navigating to child topics.

**Child:** This level of the information architecture includes includes concepts, tasks, or reference topics.

## Parent directory structure

Within the top-level directory, there is a parent directory.

The image below shows how the repository's `user-management` parent directory is structured.

- There is an `_index.md` file in the parent directory that serves as a landing page for the child topics.
  In most cases, `_index.md` contains conceptual content.
  For information about the types of conceptual content that you can add to the `_index.md` file, refer to [Concepts](https://grafana.com/docs/writers-toolkit/structure/topic-types/concept/).
- There are also four task topics in the parent directory, each with a directory and `index.md` file.

For more information about how to write concepts, refer to [Concepts](https://grafana.com/docs/writers-toolkit/structure/topic-types/concept/).
For more information about how to write tasks, refer to [Tasks](https://grafana.com/docs/writers-toolkit/structure/topic-types/task/).

![Parent directory structure](/media/docs/writers-toolkit/parent-directory.png)

{{< admonition type="note" >}}
If a directory contains multiple pages or subdirectories, it's a branch bundle, and it must include an `_index.md` file.
{{< /admonition >}}

## Pages and page bundles

Each web page generated by Hugo comes from one of three source files:

- `page/_index.md`: a Hugo branch bundle
- `page/index.md`: a Hugo leaf bundle
- `page.md`: a Hugo page

Although each of the preceding examples results in the same URL (`/page/`), Hugo works with each source file differently.

**Branch bundles** (`page/_index.md`) produce page hierarchies.
For the `/page/subpage/` URL to generated, there must first be a `page/_index.md` branch bundle source file.

**Leaf bundles** (`page/index.md`) bundle page assets.
To refer to a style sheet `page/style.css` with the link `./style.css`, the link must be in a `page/index.md` leaf bundle source file.

You need to use leaf bundles if you intend to mount content from one part of the site to another using Hugo mounts.
Hugo mounts use a virtual filesystem before site generation, and you can only mount directories.
To mount the page `/page/` to `/other/page/`, there must first be a `page/index.md` leaf bundle source file.

If you don't know whether you need to mount a leaf bundle, you probably don't and can default to using pages.

**Pages** (`page.md`) are any source files that aren't leaf or branch bundles.
It's convenient to use pages when you don't require any of the behaviors of leaf or branch bundles, as it can be easier to distinguish two source files in some text editors or IDEs.

For more information about branch bundles and leaf bundles, refer to [Page bundles](https://gohugo.io/content-management/page-bundles/).
