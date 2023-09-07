---
title: Visualization topic
menuTitle: Visualization
description: Learn how to write a visualization topic.
weight: 400
aliases:
  - /docs/writers-toolkit/structure/topic-types/visualization
keywords:
  - topic types
  - template
  - visualization
---

# Visualization topic

A _visualization_ topic provides an overview of a visualization type, its use cases, and the display options available in the user interface (UI). Visualization topics always include conceptual and reference content and may include task content.

The following types of content can be included in visualization topics:

- A high-level overview of a visualization's features, with use cases
- Descriptions of the UI options unique to that visualization (common options are documented separately)
- Screenshots that help users understand UI interactions or hard-to-describe UI
- As needed, guidance on how to use the visualization in special use cases
- Best practice guidelines

A visualization topic does not include:

- Tutorial content

Most visualizations require a data source; the ones that don't are called _widgets_. As a result, a widget requires less documentation than a typical visualization. The structures for both of these are explained in the following sections:

- [Visualization topic structure](#visualization-topic-structure)
- [Widget topic structure](#widget-topic-structure)

## Visualization topic structure

The following applies to visualizations that require a data source. For visualizations that don't require one, refer to [Widget topic structure](#widget-topic-structure).

A visualization topic includes the following elements:

- **Topic title:** A visualization topic title is the same as the name of the visualization as it's written in the UI (for example, "Time series"). Don't add the word "panel" or "visualization" to the title.
- **Introduction:** Include an introduction that explains what the visualization type can help the user do.
- **Main image:** Add a clear screenshot depicting the visualization displaying a typical use case.
- **Body:** The body may include task/how-to information as needed. The body must include the following sections:

  - **Supported data formats:** Provide guidance about which data formats are supported by the visualization type, with example use cases.
  - **[Section] options:** For each drop-down section of options in the edit panel, add a section in the topic with the name of the UI section, followed by the word "options." For example, if a drop-down section of options is called "Axis" in the UI, include a section in the topic called "Axis options". This doesn't include the sections linked from the **Other options** section.
  - **Other visualization options:** Include a section for links to the documentation for common UI options:

    - Panel options
    - Standard options
    - Thresholds
    - Legend
    - Data links
    - Field overrides
    - Value mappings

![Visualization structure](/media/docs/writers-toolkit/visualization-topic-example-ann-2.png)

## Widget topic structure

_Widgets_ are visualizations that don't require a data source. These visualizations have far fewer options and don't need as much explanation. For visualizations that require a data source, refer to [Visualization topic structure](#visualization-topic-structure). 

A widget topic includes the following elements:

- **Topic title:** A widget topic title is the same as the name of the widget as it's written in the UI (for example, "Time series"). Don't add the word "panel", "visualization", or "widget" to the title.
- **Introduction:** Include an introduction that explains what the visualization type can help the user do.
- **Main image:** Add a clear screenshot depicting the visualization displaying a typical use case.
- **Body:** The body may include task/how-to information as needed. The body must include the following sections:

  - **[Section] options:** For each drop-down section of options in the edit panel, add a section in the topic with the name of the UI section, followed by the word "options." For example, if a drop-down section of options is called "Axis" in the UI, include a section in the topic called "Axis options".

## Visualization naming conventions

When writing about visualizations (or widgets), use:

- Noun form
- Lower case, unless at the beginning of a sentence

For example:

```markdown
Use bar charts to graph categorical data.
A bar chart lets you graph categorical data.
```

If the name of a visualization is _already in plural form_, or _is highly confusing on its own_, use the adjective form. Do this by adding the word "visualization". For example:

```markdown
Use time series visualizations to display time series data as a graph.
A traces visualization turns traces data into a diagram.
Text visualizations allow you to directly include text or HTML in your dashboards.
```

## Write a visualization topic

To write a visualization (or widget) topic, follow these steps.

1. In the `grafana/grafana` repository, go to `/docs/sources/panels-visualizations/visualizations/`.
1. In the `/visualizations/` directory, create a parent directory with the following naming convention:

   - Use the name of the visualization type
   - Use lowercase letters
   - Add a hyphen between words
     <br>
     <br>
     For example: `time-series`, `node-graph`
     <br>
     <br>

1. Within the parent directory, create an `index.md` file.
1. Add front matter to the `index` file.

   For more information about front matter, refer to [Front matter]({{< relref "../../../write/front-matter" >}}).

1. Copy the content of the [Visualization template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/visualization-template.md) into your new index file.

   For more information about the kinds of content you can add to a concept topic, refer to [Visualization topic](#visualization-topic).

<!-- Add examples when some of these follow the template
## Visualization topic examples

Refer to the following topics for visualization topic examples:
-->

## Visualization template

When you're ready to write, make a copy of the [Visualization template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/concept-template.md) and add your content.
