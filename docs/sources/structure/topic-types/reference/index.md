---
aliases:
  - /docs/writers-toolkit/structure/topic-types/reference/
  - /docs/writers-toolkit/writing-guide/documentation-structure/topic-types/reference
date: "2022-10-27T16:43:50-04:00"
description: Learn how to write a reference topic.
keywords:
  - topic types
  - template
  - reference
menuTitle: Reference
review_date: "2024-05-30"
title: Reference topic
---

# Reference topic

A reference topic provides users with the information they might need to refer to when performing a task.
An effective reference provides a comprehensive list of data,
such as functions and parameters, error messages, and return codes.
A reference is usually presented as a table, a bulleted list, or a sample script.

API information is also included in reference topics.

Because reference topics contain information the user needs to accomplish a task, reference topics are often linked to task topics.

{{< admonition type="note" >}}
Don't include steps or conceptual information in reference topics.
{{< /admonition >}}

## Reference structure

- **Topic title:** Reference topic titles contain a qualifier and noun, for example, _Grafana CLI_.
  This helps the reader distinguish between reference topics and tasks.
- **Introduction:** Provide an introduction that explains what to expect from this topic.
- **Body:** Use tables or lists to provide information within reference topics.

{{< figure src="/media/docs/writers-toolkit/reference.png" alt="Annotated example of a reference page's structure" >}}

## Write a reference topic

To write a reference, complete these steps:

1. Determine where you want to add reference documentation for a Grafana Labs product.
1. Create a child directory within the parent directory that follows this naming convention:

   - Begin the directory name with a qualifier followed by an noun.
   - Use lowercase letters.
   - Add a hyphen between words.

1. Create an `index.md` file within the reference directory.
1. Add front matter to the `index.md` file.

   For more information about front matter, refer to [Front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/).

1. Make a copy of the [Reference template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/reference-template.md), and add your content to it.

## Reference topic examples

Refer to the following topics for a reference topic examples:

- [Calculation types](https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/calculation-types/)
- [Grafana CLI](https://grafana.com/docs/grafana/latest/cli/)

## Reference template

When you are ready to write, make a copy of the [Reference template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/reference-template.md) and add your content to it.
