---
aliases:
  - /docs/writers-toolkit/structure/topic-types/visualization
review_date: 2024-02-29
description: Learn how to write a visualization topic.
keywords:
  - topic types
  - template
  - visualization
menuTitle: Visualization
title: Visualization topic
weight: 400
---

# Visualization topic

A _visualization_ topic provides an overview of a visualization type, its use cases, and the display options available in the user interface (UI).
Visualization topics always include conceptual and reference content and may include task content.

The following types of content can be included in visualization topics:

- A high-level overview of a visualization's features, with use cases
- Descriptions of the UI options unique to that visualization (common options are documented separately)
- Screenshots that help users understand UI interactions or hard-to-describe UI
- As needed, guidance on how to use the visualization in special use cases
- Best practice guidelines

A visualization topic doesn't include:

- Tutorial content

Most visualizations require a data source; the ones that don't are called _widgets_.
As a result, a widget requires less documentation than a typical visualization.
The structures for both of these are explained in the following sections:

- [Visualization topic structure](#visualization-topic-structure)
- [Widget topic structure](#widget-topic-structure)

## Visualization topic structure

The following applies to visualizations that require a data source.
For visualizations that don't require one, refer to [Widget topic structure](#widget-topic-structure).

A visualization topic includes the following elements:

- **Topic title:** A visualization topic title is the same as the name of the visualization as it's written in the UI (for example, "Time series").
  Don't add the word "panel" or "visualization" to the title.

- **Introduction:** Include an introduction that explains what the visualization type can help the user do.

- **Main visual:** Add a clear screenshot depicting the visualization displaying a typical use case.
  A short (5-10 seconds) embedded video is also acceptable.
  Don't use longer videos here.

- **Body:** The body may include task/how-to information as needed.
  The body must include the following sections:

  - **Supported data formats:** Provide guidance about which data formats are supported by the visualization type, with example use cases.
  - **_`<SECTION>`_ options:** For each drop-down section of options in the edit panel, add a section in the topic with the name of the UI section, followed by the word "options".
    For example, if a drop-down section of options is called "Axis" in the UI, include a section in the topic called "Axis options".
    This doesn't include the sections linked from the **Other options** section.
  - **Other visualization options:** Include a section for links to the documentation for common UI options:

    - Panel options
    - Standard options
    - Thresholds
    - Legend
    - Data links
    - Field overrides
    - Value mappings

{{< figure src="/media/docs/writers-toolkit/visualization-topic-example-ann-2.png" alt="Annotated example of a visualization page's structure" >}}

## Widget topic structure

_Widgets_ are visualizations that don't require a data source.
These visualizations have far fewer options and don't need as much explanation.
For visualizations that require a data source, refer to [Visualization topic structure](#visualization-topic-structure).

A widget topic includes the following elements:

- **Topic title:** A widget topic title is the same as the name of the widget as it's written in the UI (for example, "Time series").
  Don't add the word "panel", "visualization", or "widget" to the title.

- **Introduction:** Include an introduction that explains what the visualization type can help the user do.

- **Main visual:** Add a clear screenshot depicting the visualization displaying a typical use case.
  A short (5-10 seconds) embedded video is also acceptable.
  Don't use longer videos here.

- **Body:** The body may include task/how-to information as needed.
  The body must include the following sections:

  - **_`<SECTION>`_ options:** For each drop-down section of options in the edit panel, add a section in the topic with the name of the UI section, followed by the word "options".
    For example, if a drop-down section of options is called "Axis" in the UI, include a section in the topic called "Axis options".

## Visualization naming conventions

When writing about visualizations or widgets use:

- Noun form
- Lower case, unless at the beginning of a sentence

For example:

```markdown
Use bar charts to graph categorical data.
A bar chart lets you graph categorical data.
```

There are some cases where you should use the adjective form of a visualization.
You should do this if:

- The name of a visualization is already in plural form
- It's highly confusing because of other uses of the word
- You're referencing a visualization outside of the visualization documentation

Do this by adding the word "visualization". For example:

```markdown
Use time series visualizations to display time series data as a graph.
A traces visualization turns traces data into a diagram.
Text visualizations allow you to directly include text or HTML in your dashboards.
Click the **Enable node graph** switch to display a node graph visualization above the trace view.
```

## Write a visualization topic

To write a visualization or widget topic, follow these steps.

1. In the `grafana/grafana` repository, go to the `docs/sources/panels-visualizations/visualizations` directory.
1. In the `visualizations` directory, create a directory with the following naming convention:

   - Use the name of the visualization type
   - Use lowercase letters
   - Add a hyphen between words

1. Within that directory, create an `index.md` file.
1. Copy the content of the [Visualization template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/visualization-template.md) into your new index file.

   For more information about the kinds of content you can add to a concept topic, refer to [Visualization topic](#visualization-topic).

1. Add additional front matter to the `index.md` file.

   For more information about front matter, refer to [Front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/).

<!-- https://github.com/grafana/writers-toolkit/issues/560 -->

## Visualization template

When you're ready to write, make a copy of the [Visualization template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/visualization-template.md) and add your content.
