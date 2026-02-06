# Learning journey step content template

This template gives patterns and examples for creating learning journey step content based on analysis of existing Grafana learning journey files.

## Required sections and order

Each learning journey step must include these sections in order:

1. **YAML front matter** (step number, weight, CTA configuration)
1. **H1 heading** (must match the `title` in front matter exactly)
1. **Introduction paragraph** (goal of this step)
1. **Numbered instructions** (step-by-step actions)
1. **Optional: Verification section** (how to confirm success)
1. **Optional: Call to action** (if using CTA front matter)

## Main heading pattern

- Include an H1 heading that matches the YAML front matter title exactly.
- Use action-oriented titles starting with verbs.

Always start with the same heading as the title in the front matter:

```markdown
# [Title from front matter]
```

Examples:

- `# The value of Logs Drilldown`
- `# Add Logs Drilldown visualization to a dashboard`
- `# Create a dashboard and add a visualization`
- `# Destination reached!{.text-center}`

## Common content patterns

### Business value and introduction pages

**Opening paragraph pattern**: Brief description of the feature or product value.

```markdown
[Product Name] is a [description of what it does]. [Key benefit statement].

With [Product name], you can:

- [Benefit 1 – starts with action verb]
- [Benefit 2 – starts with action verb]
- [Benefit 3 – starts with action verb]
- [Additional benefits as needed]
```

Examples:

- "Grafana Logs Drilldown is a queryless experience for browsing logs..."
- "Grafana Metrics Drilldown is a queryless experience for browsing **Prometheus-compatible** metrics..."
- "Grafana Alloy combines the strengths of the leading telemetry collectors..."

### Installation and setup pages

**Standard opening**:

```markdown
In this [milestone/step], you [action description].

[Optional context paragraph explaining why this step matters.]

To [perform the action], complete the following steps:
```

**Step format**:

```markdown
1. [Action step with specific UI navigation]

   [Optional additional context or screenshot reference]

1. [Next action step]

1. [Continue with numbered steps]
```

Common step patterns:

- "Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`."
- "On the Grafana Cloud home page, open the navigation menu on the left side of the screen and click **[Menu item]**."
- "Click **[Button name]**."
- "Enter [value] in the **[Field name]** field."

### Investigation and analysis pages

**Opening structure**:

```markdown
[Brief explanation of what this step accomplishes and why it's valuable]

[Description of insights you can gain from this action]

- [Insight 1]
- [Insight 2]
- [Insight 3]

## How to [perform the action]

To [action description], complete the following steps:

[Numbered steps]

## Insights from [the action/data]

The image below provides the following insights:

- [Specific insight with data examples]
- [Pattern or trend observation]
- [Performance indicator or issue identification]

![Image description](/path/to/image)
```

### Query writing pages

**Structure pattern**:

```markdown
# Write a [Query language] query using the Query Builder

Now that you've [previous context], it's time to write the query. A query is an expression that specifies the dataset used in the visualization.

[Query language] generates [types of queries]. [Explanation of query types].

## Elements of a [Query language] query

A [Query language] query consists of the following basic elements:

### [Element 1 name]

[Description with code example in backticks]

For example, [example explanation].

### [Element 2 name] (optional)

[Description]

For example, [example explanation].

## How to write a [Query language] query

To write a [Query language] query using the Query Builder, complete the following steps:

1. [Step with specific UI instructions]

   {{< admonition type="did you know?" >}}
   [Helpful tip or additional context]
   {{< /admonition >}}

1. [Continue with steps]
```

### Dashboard creation pages

**Standard opening**:

```markdown
{{< shared-snippet path="/docs/grafana/next/dashboards/_index.md" id="dashboard-overview" >}}

To [create/build] a dashboard [with specific purpose], complete the following steps:
```

Common step patterns:

- "Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`."
- "Click **New > Dashboard**."
- "Click **Add visualization**."
- "Search for and select a data source."
- "From the **Visualization** drop-down list in the upper right, select `[Visualization type]`."

### Conclusion pages

Always include a conclusion page as the last step.

Name the folder `end-journey`.

**Standard structure**:

```markdown
# Destination reached!{.text-center}

Congratulations on completing this journey! Job well done!
{.text-center}

We hope you've enjoyed traveling with us as you:

- [Achievement 1 – past tense]
- [Achievement 2 – past tense]
- [Achievement 3 – past tense]
- [Additional achievements]
```

## Verification section

- Use the heading "## Verify the installation" or "## Verify [action completed]" to match the step's goal
- Provide clear instructions to confirm the step was completed successfully
- Include specific commands or UI checks users can perform
- Use numbered lists for verification steps, starting with 1
- Include expected output or results users should see
- Add explanatory text after commands describing what successful output looks like

Example verification section format:

```markdown
## Verify the installation

To confirm Alloy is running correctly:

1. Check the service status:

    <CODE_BLOCK>

1. View the logs for any errors:

    <CODE_BLOCK>

The output should show Alloy is active and running without errors.
```

## Common phrases and patterns

### Navigation instructions

- "Sign in to your Grafana Cloud environment, for example `mystack.grafana.net`"
- "On the Grafana Cloud home page, open the navigation menu on the left side of the screen"
- "Click **[Menu item] > [Submenu item]**"
- "In the upper right of the page"
- "Located in the toolbar"

### Transition phrases

- "In the next milestone, you'll [action]"
- "In your next milestone, you'll [action]"
- "In your final milestone, you'll [action]"
- "Your next step in this journey is to [action]"
- "Having [completed previous action], [next action context]"

### Instruction starters

- "To [action], complete the following steps:"
- "To [action], perform the following steps:"
- "Complete the following steps:"

### Value propositions

- "This [action/feature] helps you:"
- "By [doing action], you can:"
- "With [feature name], you can:"
- "[Feature] provides the following advantages:"

### Conditional instructions

```markdown
| If you want to | Then |
| -------------- | ---- |
| [Scenario 1] | [Instructions] |
| [Scenario 2] | [Instructions] |
```

### Admonition patterns

```markdown
{{< admonition type="did you know?" >}}
[Helpful tip or additional context]
{{< /admonition >}}

{{< admonition type="tip" >}}
[Performance or best practice tip]
{{< /admonition >}}

{{< admonition type="caution" >}}
[Warning about potential issues]
{{< /admonition >}}

{{< admonition type="note" >}}
[Important information to remember]
{{< /admonition >}}
```

## Video integration pattern

```markdown
{{< docs/video id="[VIDEO_ID]" align="right" start="[START_TIME]" end="[END_TIME]" >}}

[Context about what the video shows]

{{< /docs/video >}}
```

## Shared snippets pattern

```markdown
{{< shared-snippet path="/docs/[path]/index.md" id="[snippet-id]" >}}
```

## Code examples pattern

```markdown
[Introduction sentence ending with a colon]:

<CODE_BLOCK>

[Explanation of the code and its purpose]
```

## Image integration pattern

```markdown
![Image description that conveys the essential information](/path/to/image)
```

## Common link patterns

- Internal docs: `/docs/[product]/[version]/[path]/`
- External links: Full URLs for external resources
- Video links: YouTube URLs for demonstrations

## Data examples pattern

When showing data structures, use realistic examples:

- JSON: Traffic density data, API responses
- CSV: Timestamped data with meaningful field names
- Metrics: CPU usage, memory consumption, error rates
- Logs: Application logs with realistic service names

## Prerequisites pattern (when applicable)

```markdown
Before you begin ensure you have the following:

- [Prerequisite 1]
- [Prerequisite 2]
- [Additional prerequisites as needed]
```
