---
aliases:
date: 
description:
  Learn what is required to document data sources at Grafana. 
keywords:
  - data source
  - plugin
menuTitle: Data sources writing guide
review_date: "2024-09-03"
title: Data sources writing guide
weight: 500
---

# Data sources writing guide

Grafana offers almost 400 data sources to Grafana users, both open source and enterprise. While many are maintained by Grafana, others are maintained by Grafana Champions and other users or companies that partner with Grafana Labs. The data sources vary in type, from key data sources to less-used data source plugins.

The following data sources come bundled with Grafana and are considered core data sources:



Enterprise plugins reside in the [Plugins-Private](https://github.com/grafana/plugins-private) repo, which is a private repo. Other plugins have their own separate repos, so you will need to clone each repo to work on data source documentation updates. These repos are also private. 

## General writing at Grafana

Use active voice and present tense


## Docs by topic

In the past, a data source document tended to be one long document containing background, requirements, configuration options, query editor options and templates and variables. To make information easier to find, the technical writing team, with agreement from product management and engineering, decided to refactor the documentation into separate topics. The idea behind this refactor was to better serve our users. If a customer wants to know how to configure the data source, it’s easier to go directly to that topic page rather than wade through one long document where lots of scrolling is required. This refactor has gotten positive feedback from the user community.

## Introduction page

This document introduces the data source and serves as the "landing page" that users encounter first.

**File structure:**

- **File name:** _index.md
- **Location:** Main data source folder

This document should include the following information:

- What is this data source? 
  - Provide a clear, concise description of the data source.
  - Primary use cases and capabilities.
- What’s unique about this data source in general? Introduce these features or concepts. (Example: Prometheus uses exemplars). What’s unique about working with this data source in Grafana?
- Is there native support for this data source or do I need to install a plugin?
- What are the best practices when working with this data source?
- Are there any requirements or prerequisites for adding the data source.
  - Example of a prerequisite:  Salesforce requires a Salesforce Connected App.
  - Include minimum version requirements.
- What are any known limitations the user needs to be aware of?
- Link to configure, query editor, etc,
  - configure the data source topic
  - query editor topic
  - templates and variable topic
  - other data source-related topics
- What can I do once I’ve configured this data source?
  - Add links to relevant docs (e.g., add variables, alerting, annotations).
- Any details regarding pre-made dashboards for the data source, and how to find these dashboards
- Add a "next steps" section at the bottom with link to the following:
  - Add annotations
  - Add transformations
  - Set up alerting
  - Additional things you can do with the specific data source
If the data source requires plugin installation, add this to the bottom of this introduction doc.

Standard text:

Always ensure your plugin version is up-to-date to access all current features. Navigate to **Plugins and data > Plugins** to check for updates.

Key points:

- Avoid technical jargon in opening paragraphs

## Configure document

**File structure:**

- **File name:** configure.md
- **Location:** Configure folder within the ain data source folder
- **menuTitle:** Configure
- **title:** Configure the <insert data source name> data source

The "Configure" document describes how to configure the data source and should serve as a comprehensive setup guide for users. This topic doc should include:


Start with a clear introduction that includes the following:

Standard text:
This document provides instructions for configuring the <data_source_name> data source and explains available configuration options. For general information on adding and managing data sources, refer to[Data source management](ref:data-source-management).

- "Before you begin" section. Include any prerequisites necessary before configuring the data source. 
  - Native support availability.
  - Plugin installation requirements with a link to installation instructions. This applies to any non-core data source. 
  - Note that the `Organization administrator` role is required to add a data source.
  - Is a separate account needed for the data source?
  - Any prerequisites that may be unique to the data source. (example: Salesforce requires that a specific account be created prior to connecting to Grafana.)
  - Know your authentication method (e.g., API key, OAuth, certificates). Have your security keys and certificates handy.

- How to provision the data source with YAML.

Standard text:
You can define and configure the data source in YAML files as part of the Grafana provisioning system. For more information about provisioning and available configuration options, refer to [Provision Grafana](ref:provisioning-data-sources).

- Add instructions for provisioning via YAML.
- Include a working example or examples.  

- Troubleshooting:
  - List common issues and how to resolve them.
  - What can users try before contacting support?
  - How to confirm a working connection in Explore.

## The query editor document

The query editor is part of the Explore page. It is the section where you create the query, either by a visual query builder or via a field where you add your query and some additional query options. Some query editors do not include a visual query builder, such as MongoDB or Jira. Datadog and Prometheus have very detailed query builders.

**File structure:**

- **File name:** query-editor.md
- **Location:** Configure folder within the ain data source folder
- **menuTitle:** Query editor
- **title:** Configure the <insert data source name> data source


This top document should include:

- Where to access the query editor (e.g., Explore, dashboard panels).
  Grafana provides a query editor for the CloudWatch data source, which allows you to query, visualize, and alert on logs and metrics stored in Amazon CloudWatch.  It's located on the [Explore](ref:explore) page. For general documentation on querying data sources in Grafana, refer to [Query and transform data](ref:query-transform-data).
- Available query modes 
  - Visual query builder or raw query
  - What query types are supported?
    - Table
    - Time series
- Explain how to build queries using the editor. 
- Describe each query editor option. This should include a description of each query type, an explanation of what each component or field is, and include examples where appropriate.
-Highlight data source-specific components or syntax. Provide examples for each component where applicable. 
- Include example queries, both simple and complex.
- Any other information useful when working with the query editor, such as information on aggregations or aggregating data.
- Are Grafana macros supported?
  - Provide a table of supported macros with examples.
  - Provide an example of a query with a macro.
- Recorded queries
  - Does this data source support recorded queries?
  - If yes, how to configure and use them.
- Validating and troubleshooting queries
  - How to check syntax and test queries.
  - What to do when queries fail.

## Template variables topic doc

template-variables.md

This topic explains templates and variables associated with the data source. You will need to work with the data source SME on this topic to ensure all information is accurate.

- What are variables? 
- What types of variables are supported with this data source?
  - Query variables, custom variables, interval variables, etc.
- Provide syntax and examples.
- Usage tips and options.
- Common issues with variables and how to fix them.

## Annotations topic doc

annotations.md

Add annotations, graphs, and dashboards to provide contextual information about specific points in time. This topic describes how annotations work with this data source.

- Overview - What are annotations in Grafana?
  - Add links to general Annotations documentation.
  - What types of annotations can you create for this data source?
  - Are there any limitations to using annotations with this data source?
- Querying for annotations
  - Field requirements: time, text, tags, optional timeEnd
- Provide one or two example annotation queries.
- Configuration
  - How to add annotation queries to a dashboard.
  - UI steps and important settings.
- Troubleshooting
  - Common problems and how to resolve them.

## Alerting topic doc

alerting.md

Can you use this data source for alerting? Create alerts to incoming metrics data or log entries and set up your Grafana Alerting system to watch for specific events or circumstances.

- Does this data source support alerting? 
  - Add links for general info on alerting and the Grafana alerting framework.
  - What kinds of alerts can you create?
- Setup
  - Describe how to create an alert rule using this data source.
  - List all query types compatible with alerts.
  - Describe UI steps for creating and configuring alerts.
  - List any notification types unique to the data source.
  - Special considerations for this data source.
- Troubleshooting
  - Known limitations or common alerting issues.
  - Tips for debugging alert rules.

## Troubleshooting topic doc

troubleshooting.md

This document provides troubleshooting steps users can try before opening a support case. If support is still needed, users can reference this list to show what they've already attempted.

- Common connection issues
  - List any common connection issues with this data source.
  - Provide a list of connection error messages and how to troubleshoot them.
  - Possible causes - Why might this happen?
- Logs and Diagnostics
  - Where can the user find relevant logs to help with troubleshooting
  - How to interpret common error messages in logs.
  - Add step-by-step resolutions. Provide the specific actions to take.
- Troubleshooting checklist
  - Include a troubleshooting checklist that users can work through systematically and attach to support cases when needed.
