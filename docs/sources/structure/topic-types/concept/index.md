---
aliases:
  - /docs/writers-toolkit/structure/topic-types/concept/
  - /docs/writers-toolkit/writing-guide/documentation-structure/topic-types/concept/
date: "2022-10-27T16:43:50-04:00"
description: Learn how to write a concept topic.
keywords:
  - topic types
  - template
  - concept
menuTitle: Concept
review_date: "2024-05-30"
title: Concept topic
---

# Concept topic

A concept provides an overview and background information to help end users understand a product, interface, or task.
Concepts answer the question "what is it?". Readers learn about features through concepts.

The following types of content can be included in concepts:

- Detailed overviews of features with benefits and clearly defined terms
- Diagrams that help users understand the components of a system
- Process flow diagrams
- Best practice guidelines
- An example of how a feature is used
  Examples might include screenshots or other supporting visuals

A concept topic doesn't include:

- Step-by-step instructions
- Reference information, such as lookup tables or lists of values

## Concept topic structure

A _concept_ topic includes the following elements:

- **Topic title:** Topic titles should be nouns, for example, _Grafana panels_.
  By using this naming convention, readers are able to distinguish between conceptual topics and tasks that begin with verbs.
  For best practice guidelines, use the title _Best practices_.
- **Introduction:** Include an introduction that explains what this topic is about.
- **Body:** Provide as much content as needed to explain the concept thoroughly.
  There can be sections, visuals, and text in the body of a concept.

{{< figure src="/media/docs/writers-toolkit/concept.png" alt="Annotated example of a concept page's structure" >}}

## Write a concept topic

To write a concept topic, follow these steps.

1. Determine where you want to add concept documentation to the Grafana Labs product documentation.
1. Within the top-level entity, create a parent directory with the following naming convention:

   - Use a noun
   - Use lowercase letters
   - Add a hyphen between words

1. Within the parent directory, create an `_index.md` file.
1. Add front matter to the `_index.md` file.

   For more information about front matter, refer to [Front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/).

1. Add the content to a copy of the [Concept template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/concept-template.md).

   For more information about the kinds of content you can add to a concept topic, refer to [Concept topic](#concept-topic).

## Concept topic examples

Refer to the following topics for concept topic examples:

- [Roles and permissions](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/)
- [Loki deployment modes](https://grafana.com/docs/loki/latest/get-started/deployment-modes/)
- [Grafana dashboard best practices](https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/best-practices/)

## Concept template

When you are ready to write, make a copy of the [Concept template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/concept-template.md) and add your content.
