---
title: Visualization title
menuTitle: Visualization
description: Use this template when you write a visualization topic.
aliases:
  - /docs/writers-toolkit/latest/templates/concept-template
weight: 100
keywords:
  - keyword 1
  - keyword 2
  - keyword 3
---

<!-- Refer to [Front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/) for more information about how to populate front matter. -->

# Visualization title

<!-- The visualization title is required. This is the name of the visualization as it appears in the UI. For example: Time series.

A visualization topic provides an overview to help end users understand how to best use the visualization type and the options available to them in the user interface (UI). Visualization topics always include conceptual and reference content and may include task content.

Refer to the [Visualization topic documentation](https://grafana.com/docs/writers-toolkit/writing-guide/topic-types/visualization/) for guidelines on writing a visualization topic.
-->

Introduce the concept.

<!-- The introduction is required. Add an introduction to the visualization that explains what visualization type can help the user do. -->

Main visual representing the visualization.

<!-- A screenshot is generally preferred because it's easiest to maintain, but a five-to-ten second video is acceptable. Don't use longer videos as the main visual aid here. They can be added elsewhere in the page, if needed. -->

## Configure a visualization

<!-- Optional {{< youtube id="videoidhere" >}} -->

Grafana Play shortcode block. Place this either after the main image or after a configuration video.

<!-- Optional  {{< docs/play title="Time Series Visualizations in Grafana" url="playurlhere" >}} -->

## Supported data formats

<!-- Provide guidance about which data formats are supported by the visualization type with example use cases. Does not apply to a widget topic. -->

## Special configuration or task

<!-- Optional special configurations or tasks to achieve relevant tasks with the visualization. Each of these should be their own heading unless it makes sense to group some of them under a heading. -->

## Special configuration or tak

<!-- Optional special configurations or tasks to achieve relevant tasks with the visualization. Each of these should be their own heading unless it makes sense to group some of them under a heading. -->

### Special configuration or task

<!-- Grouped optional special configuration or task -->

## Configuration options

{{< docs/shared lookup="visualizations/config-options-intro.md" source="grafana" version="<GRAFANA_VERSION>" >}}

<!-- Arrange following sections in the order in which they appear in the UI -->

### [Section name] options

<!-- For each expandable section of options in the panel edit pane, add a section in the topic with the name of the UI section, followed by the word "options." -->

### [Section name] options

<!-- For each expandable section of options in the panel edit pane, add a section in the topic with the name of the UI section, followed by the word "options." -->

### Data links

{{< docs/shared lookup="visualizations/datalink-options.md" source="grafana" version="<GRAFANA_VERSION>" >}}

### Field overrides

{{< docs/shared lookup="visualizations/overrides-options.md" source="grafana" version="<GRAFANA_VERSION>" >}}

### Legend

{{< docs/shared lookup="visualizations/legend-options-1.md" source="grafana" version="<GRAFANA_VERSION>" >}}

<!-- OR -->

{{< docs/shared lookup="visualizations/legend-options-2.md" source="grafana" version="<GRAFANA_VERSION>" >}}

### Panel options

{{< docs/shared lookup="visualizations/panel-options.md" source="grafana" version="<GRAFANA_VERSION>" >}}

### Standard options

{{< docs/shared lookup="visualizations/standard-options.md" source="grafana" version="<GRAFANA_VERSION>" >}}

### Thresholds

{{< docs/shared lookup="visualizations/thresholds-options-1.md" source="grafana" version="<GRAFANA_VERSION>" >}}

<!-- OR -->

{{< docs/shared lookup="visualizations/thresholds-options-2.md" source="grafana" version="<GRAFANA_VERSION>" >}}

### Tooltips

{{< docs/shared lookup="visualizations/tooltip-options-1.md" source="grafana" version="<GRAFANA_VERSION>" >}}

<!-- OR -->

{{< docs/shared lookup="visualizations/tooltip-options-2.md" source="grafana" version="<GRAFANA_VERSION>" >}}

### Value mappings

{{< docs/shared lookup="visualizations/value-mappings-options.md" source="grafana" version="<GRAFANA_VERSION>" >}}
