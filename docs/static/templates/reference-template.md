---
title: "Reference title"
menuTitle: "Reference"
description: “Use this template when you write a reference topic.”
aliases: ["/docs/writers-toolkit/latest/templates/reference-template"]
weight: 100
keywords:
    - keyword 1
    - keyword 2
    - keyword 3
---
<!-- Refer to [Topic front matter]({{< relref "../../front-matter/" >}}) for more information about how to populate front matter. -->

# Reference title
<!-- vale Grafana.Quotes = NO -->
<!-- The reference title is required. Use a noun-based title. -->
<!-- vale Grafana.Quotes = YES -->

Add an introduction to the reference.

<!-- The introduction is required. Include reference information, such as extensive tables, lists, or other information that is used as support for a task. Reference topics are also designed for API information.

Often reference topics are linked from _task_ topics, because they contain information the user needs in order to perform a task. -->

## Lists

Lists of commands or parameters are often organized in reference topics. The information you need to present will dictate the format.

- They might
- be in
- unordered lists.

[Configuration](https://grafana.com/docs/grafana/latest/installation/configuration/) is an example of lists.

## Tables

If you have a large list of things to store in a table, then you are probably dealing with reference information. Hugo accepts either tables in Markdown or in HTML format, so use whichever is easier for you.

The [Glossary](https://grafana.com/docs/grafana/latest/guides/glossary/) provides an example of reference data in a table.

### Empty markdown table

While you might not need a heading for each table, headings are a good way to chunk information if you have several tables. They also make the content easy to skim. Use headings or intro paragraphs like this one to explain to the reader what the information in the table is used for.

|     |     |     |     |     |     |
| :-- | :-- | :-: | :-: | --: | --: |
|     |     |     |     |     |     |
|     |     |     |     |     |     |
|     |     |     |     |     |     |
|     |     |     |     |     |     |

## API documentation

API documentation is always a reference topic rather than a task topic, but it has its own rules.
