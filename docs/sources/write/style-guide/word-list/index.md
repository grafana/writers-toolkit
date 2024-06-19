---
title: Word list
description: Guidelines for words to use in writing Grafana documentation.
weight: 600
keywords:
  - Grafana
  - word list
---

# Word list

In most cases, you can refer to the [word list in the Google developer documentation style guide](https://developers.google.com/style/word-list) to determine if you should use a word or not. The following guidelines cover cases in which:

- Grafana guidelines differ from Google guidelines.
- The word isn't included in Google guidelines.
- It's still easy to use an incorrect word because it's widely used, generally or in other Grafana media.

{{< admonition type="note" >}}
This page is a work in progress.
{{< /admonition >}}

<!-- vale Grafana.Headings = NO -->
<!-- vale Grafana.Spelling = NO -->
<!-- vale Grafana.WordList = NO -->

## A

<!-- vale Grafana.Agentless = NO -->

### agentless

Don't use.
Grafana Agent has been replaced by Grafana Alloy, so you shouldn't use agent-based terminology.
Refer to [no-collector](#no-collector).

<!-- vale Grafana.Agentless = YES -->

### alert rule

Grafana Alerting uses the term _alert rule_ to describe the Grafana feature that includes both [Grafana-managed alert rules](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/alert-rule-types/#grafana-managed-alert-rules) and [Data source-managed alert rules](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/alert-rule-types/#data-source-managed-alert-rules).
For more information, refer to [Alert rules](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/).

{{< admonition type="caution" >}}
Don't confuse this with an [_alerting rule_ ](#alerting-rule).
{{< /admonition >}}

### alerting rule

An _alerting rule_ is a Prometheus concept reused in Grafana Mimir.
For more information refer to [Alerting Rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/).

{{< admonition type="caution" >}}
Don't confuse this with an [_alert rule_ ](#alert-rule).
{{< /admonition >}}

<!--
## B -->

## C

### CHANGELOG

When naming a file or making a general reference to CHANGELOGs, spell using all caps. When referencing a specific CHANGELOG file, match the same capitalization of that file.

## D

### data source

Use this rather than _datasource_ for the noun form.

Also, use _data source plugin_ rather than _data-source plugin_.

While most compound adjectives require a hyphen, it's left out in this case to maintain consistency with the naming of data sources in the application and reduce confusion.

{{< admonition type="note" >}}
For other compound adjectives, use a hyphen unless otherwise specified.
{{< /admonition >}}

### dialog box

<!-- vale Grafana.DialogBox = NO -->

Use this rather than _modal_ or _dialog_.

This guidance intentionally differs from Google style guide advice, which prefers just [_dialog_](https://developers.google.com/style/word-list#dialog) because _dialog box_ is a user-friendly term that's easy to understand.

<!-- vale Grafana.YES = NO -->

### drop-down

Use this rather than _dropdown_ or _drop down_.

<!--
## E
## F
## G -->

## H

### hover over

Use this rather than _hold the pointer over_ or _point to_.

<!--
## I
## J
## K
## L
-->

## M

### menu icon

Use this rather than _hamburger menu_ or _kebab menu_.

## N

### no-collector

<!-- vale Grafana.Agentless = NO -->

Use this to refer to deployments that don't have a collector.
Use this instead of _agentless_.

<!-- vale Grafana.Agentless = YES -->

### Node Exporter

<!-- vale Grafana.PrometheusExporters = NO -->
<!-- This is demonstrating improper usage. -->

When referring to the product, Prometheus Node Exporter, capitalize both words in the term _Node Exporter_.
Don't use _Node exporter_ or _node exporter_.

<!-- vale Grafana.PrometheusExporters = YES -->

When referring to the tool, use `node_exporter`. The text should always be pre-formatted (in "code" font).

<!--
## O
## P -->

## Q

### quickstart

Use the compound adjective without a hyphen whether the noun is implied or explicit. For example, you can use _quickstart guide_ or just _quickstart_. If you are using the noun form, write as two words.

## R

### React

<!-- vale Grafana.React = NO -->
<!-- This is demonstrating improper usage. -->

Use this rather than _React.js_ or _ReactJS_.

<!-- vale Grafana.React = YES -->

### README

When naming a file or making a general reference to READMEs, spell using all caps. When referencing a specific README file, match the same capitalization of that file.

## S

### SQL (Structured Query Language)

The article&mdash;_a_ or _an_&mdash;that you use before the acronym _SQL_ depends on how the word is pronounced.

When referring to the product Microsoft SQL Server, _SQL_ should be pronounced "sequel". In this case, use the article _a_, as in _a SQL Server analysis_.

When referring to the term in any other context, such as SQL databases, errors, or servers, _SQL_ should be pronounced "ess-cue-el". In this case, use the article _an_, as in _an SQL error_.

## T

### time series

Use this rather than _timeseries_ for the noun form.

When you need to use the adjective form, use _time-series_ rather than _timeseries_.

<!--

## U

## V

## W

## X

-->

<!-- vale on -->
