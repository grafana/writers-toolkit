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

If you use an AI assistant to draft or review documentation, refer to the [AI quick reference](https://grafana.com/docs/writers-toolkit/write/style-guide/ai-quick-reference/) for a concise set of style rules written for AI agents and documentation skills.

## Headings

- Use sentence case. Capitalize only the first word and proper nouns. Write "Configure authentication settings", not "Configure Authentication Settings".
- Don't use gerunds (-ing verbs). Use imperative verbs instead: "Configure authentication", not "Configuring authentication".
- Don't place a heading directly after another heading. Always include at least one sentence of introductory content between a section heading and its first subheading.
- Don't skip heading levels. After a `#` heading, use `##`, not `###`.
- Don't use hyphens in headings.
- Don't add parenthesized qualifiers such as (Important) to headings. The exception is (Optional).
- Don't duplicate headings on a page. If you must reuse a heading, keep its meaning consistent.

For more guidance, refer to [Heading don'ts](https://grafana.com/docs/writers-toolkit/write/markdown-guide/#heading-donts).

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

## Admonitions

Use the `admonition` shortcode to call out exceptional information, such as a prerequisite, a side effect, or a risk of data loss.
Use admonitions sparingly.
If every paragraph is an admonition, none of them stand out.

Place the type in quotes and put the content between the opening and closing tags:

```markdown
{{</* admonition type="note" */>}}
Not every data source supports every feature.
Confirm which features a data source supports by checking `plugin.json`.
{{</* /admonition */>}}
```

Choose the type that matches the information:

| Type      | Use for                                                                                   |
| --------- | ----------------------------------------------------------------------------------------- |
| `note`    | Supplementary information the reader shouldn't miss, such as a prerequisite or a default. |
| `tip`     | Optional, helpful advice that isn't essential to the task.                                |
| `caution` | An action that can have unintended consequences, such as changing a shared setting.       |
| `warning` | An action that can cause data loss, downtime, or a security risk.                         |

For the full syntax and parameters, refer to [Admonition](https://grafana.com/docs/writers-toolkit/write/shortcodes/#admonition).

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

## Related resources

For more detailed guidance from the Grafana [Style guide](https://grafana.com/docs/writers-toolkit/write/style-guide/), refer to the following pages:

- [Write for developers](https://grafana.com/docs/writers-toolkit/write/style-guide/write-for-developers/): How to write for software developers and engineers, the primary audience for data source documentation.
- [UI elements list](https://grafana.com/docs/writers-toolkit/write/style-guide/ui-elements/): How to refer to UI elements, which configure and query editor pages use often.
- [UX writing](https://grafana.com/docs/writers-toolkit/write/style-guide/ux-writing/): How to write UI text, such as tooltips and other microcopy.
- [Security](https://grafana.com/docs/writers-toolkit/write/style-guide/security/): How to write about credentials, secrets, and authentication.
- [Voice and tone guidelines](https://grafana.com/docs/writers-toolkit/write/style-guide/voice-tone-guidelines/): How to apply a consistent voice and tone.
- [Capitalization and punctuation](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/): Capitalization and punctuation rules.
- [Inclusive writing](https://grafana.com/docs/writers-toolkit/write/style-guide/inclusive-writing/): How to write inclusively.
- [Word list](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/): Preferred terms and spellings.
- [AI quick reference](https://grafana.com/docs/writers-toolkit/write/style-guide/ai-quick-reference/): Concise style rules for AI agents and documentation skills.
