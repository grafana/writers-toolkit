---
date: "2025-09-10T00:00:00+01:00"
description: Style conventions to follow when writing data source documentation.
keywords:
  - data source
  - plugin
  - style
menuTitle: Style conventions
review_date: "2027-09-04"
title: Data source documentation style
weight: 300
---

# Data source documentation style

Data source documentation follows the same conventions as the rest of Grafana documentation.
This page summarizes the conventions you use most often and adds notes specific to data sources.
For complete guidance, refer to the [Style guide](https://grafana.com/docs/writers-toolkit/write/style-guide/).

## Headings

- Use sentence case. Capitalize only the first word and proper nouns. Write "Configure authentication settings", not "Configure Authentication Settings".
- Don't use gerunds (-ing verbs). Use imperative verbs instead: "Configure authentication", not "Configuring authentication".
- Don't place a heading directly after another heading. Always include at least one sentence of introductory content between a section heading and its first subheading.

## Voice and word choice

- Write in active voice. Write "Click **Save** to save the configuration", not "The configuration is saved by clicking Save".
- Address users as "you", and use present tense.
- Use contractions for a conversational tone, such as "it's", "don't", and "you're".
- Choose plain words: "use" instead of "utilize", "help" instead of "assist", "start" instead of "commence".

## Text formatting

- Format UI elements in bold, using sentence case as they appear: Click **Save & test**.
- Use single backticks for paths, configuration options, values, variables, and status codes.
- Use triple backticks with a language tag for code blocks.
- Use uppercase with angle brackets for placeholders, such as `<YOUR_ENDPOINT_URL>`, and explain them after the code block.
- Use dashes (`-`) for unordered lists and `1.` for every item in an ordered list.
- Use tables for settings and options. Bold the setting name in the first column and include defaults when applicable.
- Use admonitions sparingly, only for exceptional information. The available types are `note`, `caution`, and `warning`.

## Preferred spellings

Use the following spellings consistently.
For the full list, refer to the [Word list](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/).

<!-- vale off -->

| Term                    | Correct     | Incorrect     |
| ----------------------- | ----------- | ------------- |
| drop-down               | drop-down   | dropdown      |
| time series (noun)      | time series | timeseries    |
| time series (adjective) | time-series | timeseries    |
| data source (noun)      | data source | datasource    |
| dialog box              | dialog box  | modal, dialog |

<!-- vale on -->

## Query language code blocks

Use the appropriate language tag for query examples:

| Data source         | Language tag |
| ------------------- | ------------ |
| Azure Monitor       | `kusto`      |
| Prometheus          | `promql`     |
| Loki                | `logql`      |
| SQL databases       | `sql`        |
| Elasticsearch       | `json`       |
| InfluxDB (Flux)     | `flux`       |
| InfluxDB (InfluxQL) | `sql`        |

## Screenshots

Store data source screenshots in `/media/docs/<data-source-name>/` and reference them with the Hugo figure shortcode:

```markdown
{{< figure src="/media/docs/azure-monitor/screenshot-query-editor.png" max-width="800px" class="docs-image--no-shadow" caption="Azure Monitor query editor showing a Metrics query" >}}
```

Add a descriptive caption and use `max-width` to control the display size.

## Key concepts tables

For data sources with platform-specific terminology, such as Azure, AWS, GCP, or Cloudflare, add a **Key concepts** table where a reader first encounters unfamiliar vocabulary.
This is most commonly:

- The Configure page, for authentication and account-model terminology such as IAM, service principal, or workspace.
- The Query editor page, for query-language and data-model terminology such as KQL, metric namespace, or filter expression.

Place the table near the top of the page, after "Before you begin" and before the first task.
Keep it to four to eight terms: this is a primer, not a glossary.
The configuration and query-editor tables can, and usually should, be different.

## Avoid positional language

Don't use "above", "below", "following", or "previous" to refer to content.
Content moves during editing and restructuring, which breaks these references.

| Don't use              | Use instead                                   |
| ---------------------- | --------------------------------------------- |
| "the table above"      | "this table", or link to the specific section |
| "as shown below"       | Remove it, or use "as shown in this example"  |
| "the previous section" | Link to the section by name                   |

## Match external terminology

When you document permissions, authentication, or platform-specific concepts:

- Use the exact terminology from the external platform's UI. Users look at both docs at the same time.
- Include the navigation path users follow in the external platform, such as "In the Cloudflare dashboard, navigate to **My Profile** > **API Tokens**".
- Format permissions as users see them in the external platform.
