---
aliases:
  - /docs/writers-toolkit/writing-guide/documentation-structure/
  - /docs/writers-toolkit/structure
date: "2022-06-27T11:51:13-05:00"
description: How to organize concepts and tasks in a Grafana Labs repository.
keywords:
  - information architecture
  - structure
menuTitle: Structure
review_date: "2024-05-29"
title: Documentation structure
weight: 300
---

# Documentation structure

The Grafana Labs documentation team makes decisions about how to organize and structure product documentation.
The topic levels discussed on this page reflect common user goals.
For example, a first-time Grafana user must understand the basic concepts before installing the product.

Before you begin contributing to documentation, it's important to understand the structure of the content.

## User goals and documentation structure

When writing documentation, [focus on what your user's goals](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#focus-on-user-goals) are.

Use the structure of your documentation to reflect the user's goals.
Think about what your users want to do, what they need to know, and how they can accomplish the tasks.

This approach applies not only to content on a page but also to how you organize a set of documentation, whether it's for a product or a feature.

With well-structured content, you can find what you need quickly and easily.
Topics flow in a logical progression.

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

### Give prominence to important topics

When you don't have a long top-level table of contents, you can include them at the top level as individual topics or modifications to the standard topic list.

For example, _Metrics-generator_ and _TraceQL_ are two of the most viewed topics in Tempo documentation.
_Metrics-generator_ is a top-level topic and _TraceQL_ is in the topic _Query with TraceQL_, a modification of the standard _Query data_ topic.

### Avoid extra hierarchy

When you only have a single topic, don't nest it inside a standard topic just to use the standard topic list.

For example, in the Tempo documentation, _API documentation_ would be the only entry under _References_ standard topic.
Instead of adding another layer in the table of contents, _API documentation_ is at the top level.

### Topic list

You can use the following high-level topics to group content.
When writing new content, consider where it should appear given this content structure.
For example, a conceptual page explaining metrics would go under the _Introduction_ topic.

<table>
  <thead>
    <tr>
      <th scope="col">Topic</th>
      <th scope="col">Example link</th>
      <th scope="col">Contents</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><em>Introduction</em></td>
      <td><a href=https://grafana.com/docs/grafana/latest/fundamentals/>Introduction to Grafana</a></td>
      <td>Conceptual, fundamental, or architectural information.</td>
    </tr>
    <tr>
      <td><em>Get started</em></td>
      <td><a href=https://grafana.com/docs/tempo/latest/getting-started/>Get started with Grafana Tempo</a></td>
      <td>Opinionated walk-throughs and tutorials.</td>
    </tr>
    <tr>
      <td><em>Set up</em></td>
      <td><a href=https://grafana.com/docs/grafana/latest/setup-grafana/>Set up Grafana</a></td>
      <td>System requirements, and subsections titled <em>Set up</em>, <em>Configure</em>, <em>Upgrade</em>, or <em>Migrate</em>.</td>
    </tr>
    <tr>
      <td><em>Configure</em></td>
      <td><a href=https://grafana.com/docs/tempo/latest/configuration/>Configure Tempo</a></td>
      <td>
        <em>Configure</em> can be its own section directory if the number of pages warrants it.
        Making this determination isn't an exact science; use your best judgement.
      </td>
    </tr>
    <tr>
      <td><em>Create alerts</em></td>
      <td>
        <a href=https://grafana.com/docs/grafana-cloud/alerting-and-irm/alerting/alerting-rules/create-grafana-managed-rule/>
          Create a Grafana managed alert rule
        </a>
      </td>
      <td>
        Specific to Grafana Ops products such as Alerting, OnCall, Incident, and SLO.
        The word <em>alert</em> may be changed, depending upon the product.
        If used with Grafana SLO, this then topic would be <em>Create SLO</em>.
        Don't use with backend database products, such as Tempo and Loki.
        Use <em>Alerts</em> instead, and refer to an operational product for details.
      </td>
    </tr>
    <tr>
      <td><em>Manage alerts</em></td>
      <td><a href=https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/manage/>Manage SLOs</a></td>
      <td>
        Specific to Grafana Ops products such as Alerting, OnCall, Incident, and SLO.
        The word <em>alert</em> may be changed, depending upon the product.
        If used with Grafana SLO, this then topic would be <em>Create SLO</em>.
        Don't use with backend database products, such as Tempo and Loki.
        Use <em>Alert</em> instead, and refer to an operational product for details.
      </td>
    <tr>
    <tr>
      <td><em>Integrate [with] &lt;PRODUCT&gt;</em> or <em>Send data</em></td>
      <td><a href=https://grafana.com/docs/grafana-cloud/monitor-infrastructure/get-started-data/>Instrument and send data to Grafana Cloud</a>
      <td>How to set up data integrations, product integrations, data sources, clients, plugins, and more.</td>
    </tr>
    <tr>
      <td><em>Query data</em></td>
      <td><a href=https://grafana.com/docs/tempo/latest/traceql/>Query with TraceQL</a>
      <td>Query languages, query tools, and examples.</td>
    </tr>
    <tr>
      <td><em>Visualize data</em></td>
      <td><a href=https://grafana.com/docs/grafana/latest/dashboards/>Dashboards</a>
      <td>Dashboard concepts and procedures. Link to the definitive source of dashboard documentation, rather than duplicating the information here.</td>
    </tr>
    <tr>
      <td><em>Alert</em></td>
      <td><a href=https://grafana.com/docs/loki/latest/alert/>Alerting and recording rules</a>
      <td>
        This topic level is used for pages that discuss alerting features, like alerting rules in Grafana Loki.
        It provides a place for alerting content that's not specific to the Grafana Operations products.
      </td>
    </tr>
    <tr>
      <td><em>Manage &lt;PRODUCT&gt;</em></td>
      <td><a href=https://grafana.com/docs/grafana-cloud/alerting-and-irm/oncall/manage/user-and-team-management/>Manage users and teams for Grafana OnCall</a>
      <td>
        Information about managing a Grafana Labs product.
        For example, content in this topic helps you view, edit, and iterate on the Grafana product you installed.
      </td>
    </tr>
    <tr>
      <td><em>Monitor &lt;PRODUCT&gt;</em></td>
      <td><a href=https://grafana.com/docs/mimir/latest/operators-guide/monitor-grafana-mimir/>Monitor Grafana Mimir</a>
      <td>Information about using tools to monitor a Grafana Labs product.</td>
    </tr>
    <tr>
      <td><em>References</em></td>
      <td><a href=https://grafana.com/docs/mimir/latest/references/>Grafana Mimir references</a>
      <td>
        APIs, configuration references, SDKs, and more.
        Material that's usually not procedural.
      </td>
    </tr>
  </tbody>
<table>

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

**Child:** This level of the information architecture includes concepts, tasks, or reference topics.

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
