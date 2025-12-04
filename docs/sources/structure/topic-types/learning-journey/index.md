---
date: "2025-10-30T16:43:50-04:00"
description: Learn how to write a learning journey topic.
menuTitle: Learning journey
review_date: "2024-05-30"
title: Learning journey topic
---

<!-- vale Grafana.Gerunds = NO -->

# Learning journey topic

<!-- vale Grafana.Gerunds = YES -->

A learning journey provides an opinionated guide to performing achieving a user goal across in Grafana Cloud.

## Interactivity

Use shortcodes to make learning journey topics interactive so that the interactive learning plugin can show users how to perform actions or even perform the action for them.

### `interactive/sequence`

Groups multiple interactive steps with coordinated execution.
Creates a container with progress tracking and a single **Do section** button that executes all child steps in sequence.

**Usage:**

```markdown
{{</* interactive/sequence
    id="unique-sequence-id"
    requirements="has-datasource:prometheus"
    objectives="section-completed:previous-section" */>}}
<CHILD INTERACTIVE STEPS>
{{</* /interactive/sequence */>}}
```

For example:

The following plain Markdown:

```markdown
1. Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`.
1. On the Grafana Cloud home page, open the navigation menu on the left side of the screen and click **Drilldown > Metrics**.
1. Select the data source you want to query.

   The Metrics Drilldown interface containing a visualization for each metric appears.

1. Use the **Filter by label values** field to narrow metrics by specific labels (optional).

   Use **label filters** to narrow down metrics by specific label values like `job="my-service"` or `instance="localhost"`.
   This helps focus on metrics from particular services or systems.

1. Enter a search term like `cpu` in the **Quick search metrics** field to find specific metrics.
1. Review the filtered metrics list showing visualizations for each matching metric.

   The **metrics list** shows all matching metrics with preview charts.
   Each metric panel displays the metric name, a time series preview, and aggregation information to help you quickly identify interesting patterns.
```

Could become the following interactive sequence:

{{< admonition type="note" >}}
Because the steps are part of an interactive sequence, you no longer need to number them.
Each interactive step is automatically numbered.
{{< /admonition >}}

```markdown
{{</* interactive/sequence
    id="search-filter-metrics"
    requirements="has-datasource:prometheus" */>}}

{{</* interactive/ignore */>}}
Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`.
{{</* /interactive/ignore */>}}

{{</* interactive/highlight
    reftarget="a[data-testid='data-testid Nav menu item'][href='/a/grafana-metricsdrilldown-app/drilldown']"
    requirements="navmenu-open,exists-reftarget"
    verify="on-page:/a/grafana-metricsdrilldown-app/drilldown" */>}}
On the Grafana Cloud home page, open the navigation menu on the left side of the screen and click **Drilldown > Metrics**.
{{</* /interactive/highlight */>}}

{{</* interactive/highlight
    reftarget="div[id='ds']"
    requirements="exists-reftarget"
    doit="false" */>}}
Select the data source you want to query.

The Metrics Drilldown interface containing a visualization for each metric appears.
{{</* /interactive/highlight */>}}

{{</* interactive/highlight
    reftarget="input[placeholder='Filter by label values']"
    requirements="exists-reftarget"
    doit="false" */>}}
Use the **Filter by label values** field to narrow metrics by specific labels (optional).

Use **label filters** to narrow down metrics by specific label values like `job="my-service"` or `instance="localhost"`.
This helps focus on metrics from particular services or systems.
{{</* /interactive/highlight */>}}

{{</* interactive/formfill
    reftarget="input[placeholder='Quick search metrics']"
    targetvalue="cpu"
    requirements="exists-reftarget" */>}}
Enter a search term like `cpu` in the **Quick search metrics** field to find specific metrics.
{{</* /interactive/formfill */>}}

{{</* interactive/highlight
    reftarget="div[data-testid='metrics-list']"
    requirements="exists-reftarget"
    doit="false" */>}}
Review the filtered metrics list showing visualizations for each matching metric.

The **metrics list** shows all matching metrics with preview charts.
Each metric panel displays the metric name, a time series preview, and aggregation information to help you quickly identify interesting patterns.
{{</* /interactive/highlight */>}}

{{</* /interactive/sequence */>}}
```

**Parameters:**

- `id` (required) - Unique identifier for the sequence
- `requirements` - Comma-separated conditions that must be met before execution
- `objectives` - Auto-completion conditions (skips entire section if already met)
- `hint` - User-friendly description displayed in tooltips

### `interactive/highlight`

Highlights and optionally clicks DOM elements using CSS selectors. Supports both educational (show-only) and interactive modes.

**Usage:**

```markdown
{{</* interactive/highlight
    reftarget="button[data-testid='save-dashboard']"
    requirements="exists-reftarget,on-page:/dashboard"
    doit="false"
    verify="has-dashboard-named:My Dashboard" */>}}
Click the **Save** button to save your dashboard
{{</* /interactive/highlight */>}}
```

**Parameters:**

- `reftarget` (required) - CSS selector for target element
- `requirements` - Default: "exists-reftarget"
- `doit` - Set to "false" for show-only mode (educational highlighting)
- `verify` - Post-action verification condition
- `skippable` - Set to "true" to allow skipping if requirements fail
- `hint` - User-friendly description

**Common requirements:**

- `exists-reftarget` - Target element exists in DOM
- `navmenu-open` - Navigation menu is open
- `on-page:/path` - User is on specific page

### `interactive/button`

Clicks buttons by their visible text content. More stable than CSS selectors for buttons with consistent labels.

**Usage:**

```markdown
{{</* interactive/button
    reftarget="Save & test"
    requirements="exists-reftarget"
    verify="has-datasource:prometheus" */>}}
Save and test the data source configuration
{{</* /interactive/button */>}}
```

**Parameters:**

- `reftarget` (required) - Button text (exact or partial match)
- `requirements` - Default: `exists-reftarget`
- `verify` - Post-action verification condition
- `skippable` - Allow skipping if requirements fail
- `hint` - User-friendly description

**Advantages:**

- More stable than CSS selectors for UI changes
- Works with internationalized button text
- Handles dynamic button states

### `interactive/formfill`

Fills form and text inputs values.

**Usage:**

```markdown
{{</* interactive/formfill
    reftarget="input[id='connection-url']"
    targetvalue="http://prometheus:9090"
    requirements="exists-reftarget"
    verify="form-field-value:connection-url:http://prometheus:9090" */>}}
Enter the Prometheus server URL
{{</* /interactive/formfill */>}}
```

**Parameters:**

- `reftarget` (required) - CSS selector for form element
- `targetvalue` (required) - Value to enter in the form field
- `requirements` - Default: `exists-reftarget`
- `verify` - Post-action verification condition
- `hint` - User-friendly description

### `interactive/navigate`

Navigates to internal routes or external URLs.

**Usage:**

```markdown
{{</* interactive/navigate
    reftarget="/dashboard/new"
    requirements="navmenu-open"
    verify="on-page:/dashboard/new" */>}}
Navigate to create a new dashboard
{{</* /interactive/navigate */>}}
```

**Parameters:**

- `reftarget` (required) - URL path
- `requirements` - Default: `navmenu-open`
- `verify` - Post-navigation verification condition
- `hint` - User-friendly description

### `interactive/multistep`

Executes multiple actions as a single atomic step.
Contains child steps that are executed in sequence.

**Usage:**

```markdown
{{</* interactive/multistep
    requirements="on-page:/dashboard/new"
    hint="Add visualization and select data source" */>}}

Add a new visualization with Prometheus data source

{{</* interactive/ignore */>}}
Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`.
{{</* /interactive/ignore */>}}

{{</* interactive/highlight
    reftarget="a[data-testid='data-testid Nav menu item'][href='/a/grafana-metricsdrilldown-app/drilldown']"
    requirements="navmenu-open,exists-reftarget"
    verify="on-page:/a/grafana-metricsdrilldown-app/drilldown" */>}}
On the Grafana Cloud home page, open the navigation menu on the left side of the screen and click **Drilldown > Metrics**.
{{</* /interactive/highlight */>}}

{{</* /interactive/multistep */>}}
```

**Parameters:**

- `requirements` - Overall requirements for the `multistep` action
- `objectives` - Auto-completion conditions
- `hint` - Description of what the `multistep` accomplishes

### `interactive/guided`

Provides guided instruction where users must manually perform each step.
Similar to `interactive/multistep` but it requires user interaction rather than automated execution.
Includes timeout and skip functionality for complex interactions.

**Usage:**

```markdown
{{</* interactive/guided
    step-timeout="45000"
    skippable="true"
    hint="Complete both hover and click actions" */>}}

{{</* interactive/ignore */>}}
Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`.
{{</* /interactive/ignore */>}}

{{</* interactive/highlight
    reftarget="a[data-testid='data-testid Nav menu item'][href='/a/grafana-metricsdrilldown-app/drilldown']"
    requirements="navmenu-open,exists-reftarget"
    verify="on-page:/a/grafana-metricsdrilldown-app/drilldown" */>}}
On the Grafana Cloud home page, open the navigation menu on the left side of the screen and click **Drilldown > Metrics**.
{{</* /interactive/highlight */>}}

{{</* /interactive/guided */>}}
```

**Parameters:**

- `requirements` - Overall requirements for the guided steps
- `objectives` - Auto-completion conditions
- `step-timeout` - Timeout in milliseconds for each step
- `skippable` - Set to `"true"` to allow skipping if steps fail
- `hint` - Description of what the guided steps accomplish

### `interactive/assistant`

Marks content sections for AI assistant integration within interactive learning journeys.
Creates a container that the interactive learning plugin can use to provide contextual AI assistance or enhanced explanations.

**Usage:**

```markdown
{{</* interactive/assistant
    id="unique-assistant-id"
    type="query" */>}}
Use **label filters** to narrow down metrics by specific label values like `job="my-service"` or `instance="localhost"`.
This helps focus on metrics from particular services or systems.
{{</* /interactive/assistant */>}}
```

**Parameters:**

- `id` (required) - Unique identifier for the assistant section
- `type` (required) - Type of assistance to provide

**Types:**

| Type     | Use for                   | Example                         |
| -------- | ------------------------- | ------------------------------- |
| `query`  | PromQL, LogQL, SQL        | `rate(http_requests_total[5m])` |
| `config` | URLs, hostnames, settings | `http://prometheus:9090`        |
| `code`   | YAML, JSON, scripts       | Alert rules, recording rules    |

**Integration with other shortcodes:**

The `interactive/assistant` shortcode works well nested within other interactive elements to provide enhanced context:

```markdown
{{</* interactive/highlight
    reftarget="input[placeholder='Filter by label values']"
    requirements="exists-reftarget"
    doit="false" */>}}
Use the **Filter by label values** field to narrow metrics by specific labels (optional).

{{</* interactive/assistant id="label-filter-help" type="query" */>}}
Use **label filters** to narrow down metrics by specific label values like `job="my-service"` or `instance="localhost"`.
This helps focus on metrics from particular services or systems.
{{</* /interactive/assistant */>}}
{{</* /interactive/highlight */>}}
```

### `interactive/ignore`

Ignore the inner content in the interactive learning plugin.
Useful for steps that only make sense outside of the Grafana UI.

**Usage:**

```markdown
{{</* interactive/ignore */>}}
Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`.
{{</* /interactive/ignore */>}}
```

### `interactive/comment`

Displays information in the box the interactive learning plugin uses to highlight a UI element.

**Usage:**

```markdown
{{</* interactive/highlight
    reftarget="div[data-testid='query-editor']"
    doit="false" */>}}
Learn about the query editor interface

{{</* interactive/comment */>}}
The **query editor** is where you write queries in your data source's native language.
For Prometheus, use **PromQL**. The editor provides syntax highlighting and auto-completion.
{{</* /interactive/comment */>}}
{{</* /interactive/highlight */>}}
```

## Requirements system

### Common requirements

- `exists-reftarget` - Target element exists in UI
- `navmenu-open` - Navigation menu is open
- `on-page:/path` - User is on specific page
- `is-admin` - User has administrator privileges
- `has-datasource:name` - Specific data source exists
- `has-datasource:type:prometheus` - Prometheus-type data source exists
- `section-completed:id` - Previous section must be complete

### Objectives versus requirements

- **Requirements**: Must be met before action can execute
- **Objectives**: Auto-complete the step/section if already met
- **Priority**: Objectives always override requirements

### Verification conditions

Post-action verification ensures steps completed successfully:

- `on-page:/path` - Verify navigation succeeded
- `has-datasource:name` - Verify data source was created
- `form-field-value:field:value` - Verify form field contains value

## Best practices

### Selector guidelines

1. **Prefer stable selectors**: Use `data-testid` attributes when available
1. **Use button action for text**: Prefer `interactive/button` for buttons with stable text
1. **Include element type**: `button[data-testid='save']` vs `[data-testid='save']`
1. **Avoid CSS classes**: Unless they're semantic (like `interactive`)

### Requirement patterns

1. **Always include `exists-reftarget`** for UI interactions
1. **Add `navmenu-open`** for navigation menu interactions
1. **Include page requirements** for page-specific actions
1. **Use objectives** for expensive state checks that can auto-complete

### Content organization

1. **Use sequences** for multi-step workflows with checkpoints
1. **Use multisteps** for atomic multi-actions
1. **Use show-only mode** (`doit="false"`) for educational content
1. **Include interactive comments** for complex UI explanations

## Error handling

### Auto-fixable requirements

- `navmenu-open` - Shows **Fix this** button to open navigation
- Navigation expansion - Can auto-expand parent sections

### Skip actions

- Use `skippable="true"` for administrator-only or plugin-dependent features
- Shows **Skip** button when requirements aren't met

### Graceful degradation

- Provide alternative paths for different user states
- Include helpful error messages in `hint` parameters
- Design learning journeys to work across different Grafana configurations
