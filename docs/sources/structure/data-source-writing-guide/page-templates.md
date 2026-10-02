---
date: "2025-09-10T00:00:00+01:00"
description: Templates and required sections for each page in a data source's documentation.
keywords:
  - data source
  - plugin
  - templates
menuTitle: Page templates
review_date: "2027-09-04"
title: Data source page templates
weight: 200
---

# Data source page templates

Use these templates as a starting point for each page in a data source's documentation.
Copy a template, then adapt it to your data source.
Replace placeholders such as `<Data Source Name>`, `<plugin-id>`, and `<GRAFANA_VERSION>` with real values.

Every example link uses the fully qualified form described in [File structure and front matter](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/file-structure-and-front-matter/).

## Overview page (`_index.md`)

The overview page is the landing page users encounter first.
Keep it concise and focused on helping users get started.
Go directly from supported features to "Get started" and "Additional features".
Don't include marketing-style sections that list benefits.

For an Enterprise data source, add a note after the intro paragraph.
Enterprise plugins aren't available on the Grafana Cloud free tier, so always specify "Pro or Advanced":

```markdown
{{< admonition type="note" >}}
The <Data Source Name> data source is an Enterprise plugin. It's available with a Grafana Cloud Pro or Advanced plan and Grafana Enterprise. For installation instructions, refer to [Install Grafana Enterprise plugins](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/administration/plugin-management/#install-grafana-enterprise-plugins).
{{< /admonition >}}
```

Use the following template for the overview page:

```markdown
# <Data Source Name> data source

<One paragraph explaining what this data source does and its primary use case.>

## Supported features

| Feature     | Supported |
| ----------- | --------- |
| Metrics     | Yes/No    |
| Logs        | Yes/No    |
| Traces      | Yes/No    |
| Alerting    | Yes/No    |
| Annotations | Yes/No    |

## Get started

The following documents help you get started:

- [Install the <Data Source Name> data source](https://grafana.com/docs/plugins/<plugin-id>/latest/install/) (Enterprise plugins)
- [Configure the <Data Source Name> data source](https://grafana.com/docs/plugins/<plugin-id>/latest/configure/)
- [<Data Source Name> query editor](https://grafana.com/docs/plugins/<plugin-id>/latest/query-editor/)
- [Template variables](https://grafana.com/docs/plugins/<plugin-id>/latest/template-variables/)
- [Troubleshooting](https://grafana.com/docs/plugins/<plugin-id>/latest/troubleshooting/)

## Additional features

After you configure the data source, you can:

- Use [Explore](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/explore/) to query data without building a dashboard.
- Add [transformations](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/panels-visualizations/query-transform-data/transform-data/) to manipulate query results.
- Set up [alerting](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/alerting/) rules.

## Pre-built dashboards

<If the data source includes pre-built dashboards, list them here with import instructions.>

## Plugin updates

Always ensure that your plugin version is up-to-date so you have access to all current features and improvements. Navigate to **Plugins and data** > **Plugins** to check for updates. Grafana recommends upgrading to the latest Grafana version, and this applies to plugins as well.

{{< admonition type="note" >}}
On Grafana Cloud, the <Data Source Name> plugin is managed by Grafana and updates automatically. On self-managed Grafana, you must update Enterprise plugins manually. Refer to [Version and upgrade guidance](https://grafana.com/docs/plugins/<plugin-id>/latest/troubleshooting/#version-and-upgrade-guidance).
{{< /admonition >}}

## Related resources

- [Official <Data Source Name> documentation](external-link)
- [Grafana community forum](https://community.grafana.com/)
```

## Install and upgrade page (`install.md`)

Include a dedicated install and upgrade page for Enterprise data sources.
It covers activation, installation across deployment environments, verification, upgrade, rollback, and uninstall.
Set `menuTitle: Installation` and a low `weight` so it sorts near the top.

Update behavior differs by environment, so be precise:

- On Grafana Cloud, Grafana Labs-managed Enterprise plugins update automatically.
- On self-managed Grafana, you update plugins manually.
- In other managed environments, such as Azure Managed Grafana, the platform provider controls the version.

Don't write a blanket statement that plugins always, or never, update automatically.

Cover these sections on the install page:

- **Before you begin:** License, minimum Grafana version, required role, and network access.
- **Activate the Enterprise plugin:** Steps for Grafana Cloud and self-managed Grafana Enterprise.
- **Install the plugin:** Steps for Grafana Cloud, the Grafana CLI, Docker, Kubernetes, and air-gapped installs.
- **Verify the installation:** How to confirm the plugin shows a status of **Installed**.
- **Upgrade the plugin:** Per-environment upgrade steps, how to install a specific version, and how to roll back.
- **Uninstall the plugin:** Removal steps, and a note that existing configurations are preserved but become non-functional.
- **Troubleshoot installation issues:** Activation, role, license, and unsigned-plugin errors.

## Configure page (`configure.md`)

The Configure page is a comprehensive setup guide.
Use the following template:

````markdown
# Configure the <Data Source Name> data source

This document explains how to configure the <Data Source Name> data source.

## Before you begin

Before you configure the data source, ensure you have:

- **Grafana permissions:** Organization administrator role.
- **<Data source> prerequisites:** <List specific requirements.>

## Add the data source

To add the data source:

1. Click **Connections** in the left-side menu.
1. Click **Add new connection**.
1. Type `<Data Source Name>` in the search bar.
1. Select **<Data Source Name>**.
1. Click **Add new data source**.

## Configure settings

| Setting     | Description                                                      |
| ----------- | ---------------------------------------------------------------- |
| **Name**    | The name used to refer to the data source in panels and queries. |
| **Default** | Toggle to make this the default data source for new panels.      |
| **URL**     | The URL of your <data source> instance.                          |

## Authentication

<Document each authentication method in its own subsection.>

## Verify the connection

Click **Save & test** to verify the connection. Document the exact success message users see when the connection test passes. Check the plugin's health check implementation for the exact message.

## Provision the data source

You can define the data source in YAML files as part of Grafana's provisioning system. For more information, refer to [Provision Grafana](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/administration/provisioning/#data-sources).

```yaml
apiVersion: 1

datasources:
  - name: <Data Source Name>
    type: <plugin-id>
    access: proxy
    url: <URL>
    jsonData:
      <key>: <value>
    secureJsonData:
      <secret-key>: <secret-value>
```
````

If the data source supports private data source connect (PDC), document it in the Configure page.
PDC lets you establish a private, secured connection between a Grafana Cloud instance and data sources secured within a private network.
Refer to [Private data source connect (PDC)](https://grafana.com/docs/grafana-cloud/connect-externally-hosted/private-data-source-connect/).

For a major data source, also include a Terraform example that references the Grafana Terraform Provider.

## Query editor page (`query-editor.md`)

The query editor page explains how to build and run queries.
Use the following template:

````markdown
# <Data Source Name> query editor

This document explains how to use the <Data Source Name> query editor.

## Before you begin

- Ensure you have [configured the <Data Source Name> data source](https://grafana.com/docs/plugins/<plugin-id>/latest/configure/).
- Verify your credentials have appropriate permissions.

## Key concepts

| Term       | Description |
| ---------- | ----------- |
| **Term 1** | Definition  |
| **Term 2** | Definition  |

## Query types

The query editor supports the following query types:

- **<Query Type 1>:** <Brief description>
- **<Query Type 2>:** <Brief description>

## Create a query

To create a query:

1. Select the **<Data Source Name>** data source.
1. Select a query type.
1. <Additional steps>

### Query examples

<Provide practical examples with code blocks.>

```<query-language>
<example query>
```

## Macros

| Macro             | Description                                  |
| ----------------- | -------------------------------------------- |
| `$__timeFilter()` | Filters results to the dashboard time range. |

## Next steps

- [Use template variables](https://grafana.com/docs/plugins/<plugin-id>/latest/template-variables/)
- [Set up alerting](https://grafana.com/docs/plugins/<plugin-id>/latest/alerting/)
````

Include both query examples and use cases.
Query examples show how to configure fields step by step.
Use cases show why and when to use a query, with real-world scenarios such as "Monitor DNS performance" that give users starting points for their own dashboards.

## Template variables page (`template-variables.md`)

The template variables page explains how to create dynamic, reusable dashboards.
Use the following template:

```markdown
# <Data Source Name> template variables

Use template variables to create dynamic, reusable dashboards.

## Before you begin

- [Configure the <Data Source Name> data source](https://grafana.com/docs/plugins/<plugin-id>/latest/configure/).
- Understand [Grafana template variables](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/dashboards/variables/).

## Supported variable types

| Variable type | Supported |
| ------------- | --------- |
| Query         | Yes/No    |
| Custom        | Yes/No    |
| Data source   | Yes       |

## Create a query variable

To create a query variable:

1. Navigate to **Dashboard settings** > **Variables**.
1. Click **Add variable**.
1. Select **Query** as the variable type.
1. Select the <Data Source Name> data source.
1. Enter your query.

## Query examples

<Provide examples of variable queries.>

## Use variables in queries

<Explain how to reference variables in data source queries.>
```

## Annotations page (`annotations.md`)

The annotations page describes how annotations work with the data source.
Annotations are visual markers that provide context about specific events or time periods.
They appear as vertical lines or regions on time-series charts.

Cover these sections:

- **Overview:** What annotations are in Grafana. Link to [Annotate visualizations](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/dashboards/build-dashboards/annotate-visualizations/).
- **Supported annotations:** What types of annotations you can create for this data source, and any limitations.
- **Query for annotations:** Field requirements, such as `time`, `text`, `tags`, and the optional `timeEnd`, with one or two example annotation queries.
- **Configuration:** How to add annotation queries to a dashboard, including UI steps and important settings.
- **Troubleshooting:** Common problems and how to resolve them.

## Grafana Alerting page (`alerting.md`)

The alerting page explains how to use the data source with Grafana Alerting.
Cover these sections:

- **Overview:** What Grafana Alerting is. Link to [Grafana Alerting](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/alerting/).
- **Supported alerts:** What kinds of alerts you can create with this data source.
- **Set up an alert rule:** How to create an alert rule, which query types are compatible with alerts, and the UI steps to configure alerts.
- **Considerations:** Notification types or special considerations unique to the data source.
- **Troubleshooting:** Known limitations, common alerting issues, and tips for debugging alert rules.

<!-- vale Grafana.Gerunds = NO -->

## Troubleshooting page (`troubleshooting.md`)

<!-- vale Grafana.Gerunds = YES -->

The troubleshooting page helps users resolve issues before opening a support case.
Organize issues into logical categories.
Common categories include:

| Category                         | When to use                                                                |
| -------------------------------- | -------------------------------------------------------------------------- |
| **License and setup errors**     | Enterprise licensing, plugin not installable, administrator role required. |
| **Version and upgrade guidance** | Outdated plugin versions, manual updates, symptoms of an old version.      |
| **Authentication errors**        | Credential issues, permissions, token expiration.                          |
| **Connection errors**            | Network, firewall, endpoint issues.                                        |
| **Query errors**                 | Syntax, no data, timeouts, invalid queries.                                |
| **Template variable errors**     | Variables not loading, incorrect values.                                   |
| **Performance issues**           | Throttling, slow queries, API limits.                                      |

For an Enterprise data source, include a "Version and upgrade guidance" section near the top of the page, after licensing and setup but before connection errors.
Running an outdated plugin version is a common root cause.
Use the heading exactly so it produces the `#version-and-upgrade-guidance` anchor that the "Plugin updates" note on the overview page links to.

When you document a specific error:

- **Use the exact error message as the heading.** Users search for the error text. For example, use `### "Invalid client secret"` rather than `### Secret problems`.
- **Start with a Symptoms section.** Describe what users see: the error message quoted exactly, the UI behavior, and when it occurs.
- **Use tables for multiple causes.** When an error has several possible causes, use a Cause and Solution table.
- **Use numbered lists for sequential solutions.** When steps must be followed in order.
- **Include code examples for complex fixes**, such as IAM policies or configuration file changes.

End the page with these sections:

- **Enable debug logging:** How to set the Grafana log level to `debug` and where to find the logs.
- **Get additional help:** Links to the [Grafana community forum](https://community.grafana.com/), the plugin's GitHub issues, and the external documentation, plus what to include when reporting an issue.
