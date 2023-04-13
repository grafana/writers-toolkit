---
title: Reference topic
menuTitle: Reference
description: Learn how to write a reference topic.
weight: 300
keywords:
  - topic types
  - template
  - reference
---

# Reference topic

A reference topic provides users with the information they might need to refer to when performing a task.
An effective reference provides a comprehensive list of data,
such as functions and parameters, error messages, and return codes.
A reference is usually presented as a table, a bulleted list, or a sample script.

API information is also included in reference topics.

Because reference topics contain information the user needs to accomplish a task, reference topics are often linked to task topics.

> **Note:** Do not include steps or conceptual information in reference topics.

## Reference structure

- **Topic title:** Reference topic titles contain a qualifier and noun, for example, _Grafana CLI_. This helps the reader distinguish between reference topics and tasks.
- **Introduction:** Provide an introduction that explains what to expect from this topic.
- **Body:** Use tables or lists to provide information within reference topics.

<figure>
<img src="reference.png" alt="Reference structure" width="600">
</figure>

## Write a reference topic

To write a reference, complete these steps:

1. Determine where you want to add reference documentation for a Grafana Labs product.
1. Create a child directory within the parent directory that follows this naming convention:

   - Begin the directory name with a qualifier followed by an noun.
   - Use lowercase letters.
   - Add a hyphen between words.

     For example:

     - `calculation-types`
     - `standard-field-definitions`
     <p>

1. Create an `index.md` file within the reference directory.
1. Add front matter to the `index` file.

   For more information about front matter, refer to [Front matter]({{< relref "../../front-matter" >}}).

1. Make a copy of the [Reference template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/reference-template.md), and add your content to it.

## Reference topic examples

Refer to the following topics for a reference topic examples:

- [Calculation types](/docs/grafana/latest/panels/calculation-types/)
- [Standard field definitions](/docs/grafana/latest/panels/standard-field-definitions/)
- [Grafana CLI](/docs/grafana/latest/administration/cli/)

## Reference template

When you are ready to write, make a copy of the [Reference template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/task-template.md) and add your content to it.
