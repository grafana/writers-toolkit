---
title: Reference title
menuTitle: Reference
description: Use this template when you write a reference topic.
aliases:
  - /docs/writers-toolkit/latest/templates/reference-template
weight: 100
keywords:
  - keyword 1
  - keyword 2
  - keyword 3
---

<!-- Refer to [Topic front matter]({{< relref "../../front-matter/" >}}) for more information about how to populate front matter. -->

# Reference title

<!-- The reference title is required. Use a noun-based title. -->

Add an introduction to the reference.

<!-- The introduction is required. Include reference information, such as extensive tables, lists, or other information that is used as support for a task. Reference topics are also designed for API information.

For guidelines about writing a reference topic, see [Reference topic](https://grafana.com/docs/writers-toolkit/writing-guide/topic-types/reference/).

Often reference topics are linked from _task_ topics, because they contain information the user needs when performing a task.
-->

## Lists

Lists of commands or parameters are often organized in reference topics. The information you need to present will dictate the format.

- They might
- be in
- unordered lists.

[Configuration](https://grafana.com/docs/grafana/latest/installation/configuration/) is an example of lists.

## Tables

If you have a large list of things to store in a table, then you are probably dealing with reference information. Hugo accepts either tables in Markdown or in HTML format, so use whichever is easier for you.

The [Glossary](https://grafana.com/docs/grafana/latest/guides/glossary/) provides an example of reference data in a table.

### Empty Markdown table

Although you might not need a heading for each table, headings are a way to chunk information if you have several tables.
Headings also make content easy to skim. Use headings (or introductory paragraphs such as this one) to provide the reader with context about the information in the table and its use.

|     |     |     |     |     |     |
| :-- | :-- | :-: | :-: | --: | --: |
|     |     |     |     |     |     |
|     |     |     |     |     |     |
|     |     |     |     |     |     |
|     |     |     |     |     |     |

## API documentation

Although API documentation is a reference topic, rather than a task topic, it has its own guidelines.
