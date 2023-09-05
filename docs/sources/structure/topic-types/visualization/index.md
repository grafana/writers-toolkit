---
title: Visualization topic
menuTitle: Visualization
description: Learn how to write a visualization topic.
weight: 100
aliases:
  - /docs/writers-toolkit/structure/topic-types/visualization
keywords:
  - topic types
  - template
  - visualization
---

# Visualization topic

A visualization topic provides an overview to help end users understand how to best use the visualization type and the options available to them in the user interface (UI). Visualization topics always include conceptual and reference content and may include task content.

The following types of content can be included in visualization topics:

- A high-level overview of a visualization's features with use cases
- Descriptions of the UI options unique to that visualization (common options are documented separately)
- Screenshots that help users understand hard-to-describe UI or UI interactions
- As needed, guidance on how to use the visualization in special use cases
- Best practice guidelines

A visualization topic does not include:

- Tutorial content

## Visualization topic structure

The following applies to visualizations that require a data source. For visualizations that don't require one, refer to [Widget topic structure](#widget-topic-structure).
A _visualization_ topic includes the following elements:

- **Topic title:** Topic titles should be the name of the visualization, for example, "Time series". Don't add the word "panel" or "visualization" to the title.
- **Introduction:** Include an introduction that explains what visualization type can help the user do.
- **Supported data formats:** Provide guidance about which data formats are supported by the visualization type with example use cases.
- **[Section] options:** For each drop-down section of options in the edit panel add a section in the topic with the name of the UI section, followed by the word "options." Generally, these don't require screenshots. Only add screenshots where the UI is potentially confusing or a UI interaction is unclear.This doesn't include the sections linked from the **Other options** section.
- **Other options:** Include a section for links to the docs for common UI options:

  - Panel options
  - Standard options
  - Thresholds
  - Legend
  - Data links
  - Field overrides
  - Value mappings

## Widget topic structure

_Widgets_ are visualizations that don't require a data source. These visualizations have far fewer options and don't need as much explanation. A widget topic includes the following elements:

- **Topic title:** Topic titles should be the name of the visualization, for example, "Time series". Don't add the word "panel" or "visualization" to the title.
- **Introduction:** Include an introduction that explains what visualization type can help the user do.
- **[Section] options:** For each drop-down section of options in the edit panel add a section in the topic with the name of the UI section, followed by the word "options".

### Visualizations naming conventions

When writing about visualizations (or widgets), use:

- Noun form
- Lower case, unless at the beginning of a sentence
- Plural form in most cases

For example: Use bar charts to graph categorical data.

Use the adjective form when the name of the visualization is:

- Already in plural form
- A word with a well-established meaning already

For example: Use time series visualizations to display time series data as a graph. Traces visualizations turns traces data into a diagram that helps you easily interpret that data.

<!-- add image here of good example visualization topic -->

## Write a visualization topic

To write a visualization topic, follow these steps.

1. In grafana/grafana, go to `/docs/sources/panels-visualizations/visualizations/`.
1. In the /visualizations/ directory, create a parent directory with the following naming convention:

   - Use the name of the visualization type
   - Use lowercase letters
   - Add a hyphen between words
     <br>
     <br>
     For example: - time-series node-graph
     <br>
     <br>

1. Within the parent directory, create an `index.md` file.
1. Add front matter to the `index` file.

   For more information about front matter, refer to [Front matter]({{< relref "../../../write/front-matter" >}}).

1. Copy the content of the [Visualization template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/visualization-template.md) into your new index file.

   For more information about the kinds of content you can add to a concept topic, refer to [Visualization topics](#visualization-topic).

<!-- Add examples when some of these follow the template 
## Visualization topic examples

Refer to the following topics for visualization topic examples:
-->

## Visualization template

When you are ready to write, make a copy of the [Visualization template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/concept-template.md) and add your content.
