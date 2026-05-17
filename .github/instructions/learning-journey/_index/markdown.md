# Learning journey content template

This template gives the standard content structure and language patterns for learning journey `_index.md` files. The template is based on analysis of existing journeys.

## Required sections and order

1. **YAML front matter** (refer to the front matter template)
1. **H1 heading** (must match the `title` in front matter)
1. **Introduction paragraph** (2-3 sentences about the goal and key concepts)
1. **"Here's what to expect" admonition**
1. **Before you begin** (prerequisites list)
1. **Troubleshooting** (common issues and solutions)
1. **More to explore** (related resources and next steps)

## Content structure and order

### 1. H1 heading

Use this pattern for the H1 heading:

```markdown
# [Action verb] [technology/goal] [with/using] [Grafana product]
```

Pattern examples:

- `# Explore logs using Logs Drilldown`
- `# Monitor Linux servers with Grafana Alloy`
- `# Visualize CSV data using the Infinity data source`
- `# Use Traces Drilldown to identify high latency service requests`
- `# Connect to a Prometheus data source in Grafana Cloud`

Guidelines:

- The H1 must match the YAML front matter `title` exactly
- Use action verbs: "Monitor", "Explore", "Visualize", "Send", "Create", "Connect", "Use", "Set up", "Migrate"
- Include the primary technology and Grafana product
- Keep under 60 characters for SEO
- Use verb phrases that describe the main goal
- Example headings: "Monitor Linux servers with Grafana Alloy", "Set up alerting with Grafana", "Migrate from Prometheus to Grafana Cloud"

### 2. Welcome statement and introduction paragraph

Start with a welcome statement and a short introduction:

```markdown
Welcome to the [specific description] learning journey [that shows/teaches/provides] [main goal].
```

Standard welcome phrases:

- `Welcome to the Grafana learning journey that shows you how to...`
- `Welcome to the Grafana learning journey that teaches you...`
- `Welcome to the Grafana learning journey that provides...`
- `Welcome to the [specific feature] learning journey.`

Introduction structure:

1. Welcome statement (required) – Start with "Welcome to the [journey name] learning journey"
1. Primary goal statement – State the main goal of the learning journey in the second sentence
1. Product context – Mention the main Grafana products used and their purpose
1. Value proposition – Give context about the technology (what it is, how it works)
1. User accomplishment – Briefly describe what users will learn or accomplish
1. Outcome reference – End with a reference to an example image showing the expected outcome

Example patterns:

```markdown
Welcome to the Explore logs using Logs Drilldown learning journey. [Feature name] is a feature designed to [core purpose] by providing [key benefit]. It lets you [specific capability] without the need for [what it replaces], making it accessible to [target audience].

[Product name] is a feature exclusive to [Grafana product]. [Product] is a [technical description] inspired by [comparison]. [Product] differs from [other technology] by [key differentiator].

The [description] image below shows [specific content example].
```

### 3. Product context section (when applicable)

- Give background information about the main technology or product.
- Explain how it fits into the broader Grafana ecosystem.
- Include links to relevant product documentation using inline links.
- Use technical details that help users understand the value proposition.

Pattern for product explanations:

```markdown
[Product name] is a [technical description]. [Product] is [architecture description] inspired by [reference technology]. [Product] differs from [comparison] by [key difference], and [additional differentiator].

[Additional technical context explaining how it works or fits in the ecosystem.]
```

Examples:

- "Loki is a horizontally scalable, highly available, multi-tenant log aggregation system inspired by Prometheus. Loki differs from Prometheus by focusing on logs instead of metrics, and collecting logs via push, instead of pull."
- "Prometheus is an open source systems monitoring and alerting toolkit that captures metrics as time series data and optional key-value pairs called labels."

### 4. Example image and caption

- Include a descriptive image showing the end result or main interface.
- Use detailed alt text that explains what the image shows.
- Place the image after the introduction and context paragraphs.
- Reference the image in the text with phrases like "The image below shows" or "The following image displays".

Image reference pattern:

```markdown
The [description] image below [shows/displays] [specific content].

![Descriptive alt text explaining what the image shows](/media/docs/learning-journey/[journey-folder]/[image-name].png)
```

Alt text patterns:

- `Example [feature] user interface`
- `Example [feature] visualization`
- `Image that shows [specific data] in [format] format`
- `Example [system type] node overview dashboard`

### 5. "Here's what to expect" section (required)

- Always include this section.
- Use the `{{< docs/box >}}` and `{{< docs/icon-heading >}}` shortcodes. - Always use the heading format: `{{< docs/icon-heading heading="## Here's what to expect" >}}`.
- Include 3-5 bullet points describing the main learning outcomes.
- Use action-oriented language starting with verbs.
- Order items logically from basic setup to advanced features.
- Focus on specific, measurable outcomes users will achieve.

Exact structure:

```markdown
{{< docs/box >}}

{{< docs/icon-heading heading="## Here's what to expect" >}}

When you complete this journey, you'll be able to:

- [Action verb] [specific outcome]
- [Action verb] [specific outcome]
- [Action verb] [specific outcome]
- [Action verb] [specific outcome]

{{< /docs/box >}}
```

Action verb patterns:

- Navigate to [feature/location]
- Search for and filter [specific items]
- Learn how to [specific skill]
- Understand the [concept/value]
- Install [product] and [configuration action]
- Configure [specific setting]
- Create [specific artifact]
- Use [tool] to [specific purpose]
- Verify that [outcome] appear in [location]
- Build a dashboard [with specific content]
- Interpret [interface] so that you can [goal]

Outcome examples:

- "Navigate to Logs Drilldown"
- "Search for and filter specific logs you want to investigate"
- "Learn how to drill down into a log and investigate related logs"
- "Add a log visualization to a dashboard"
- "Understand the value of observability and the advantages of Grafana [Product]"
- "Use pre-built dashboards and alerts to identify and troubleshoot problems in your environment"

### 6. Troubleshooting section (required)

- Always include this section.
- Always use the exact heading "## Troubleshooting".
- Include the standardized opening text: "If you get stuck, we've got your back! Where appropriate, troubleshooting information is just a click away."
- Don't include any specific issues here; mention them on the relevant step pages in the front matter.

Exact text:

```markdown
## Troubleshooting

If you get stuck, we've got your back! Where appropriate, troubleshooting information is just a click away.
```

### 7. More to explore section (required)

- Always include this section.
- Always use the exact heading "## More to explore".
- Include the standardized opening text: "We understand you might want to explore other capabilities not strictly on this path.
- We'll provide you opportunities where it makes sense."
- Link to related learning journeys using relative paths.
- Reference relevant product documentation with external links.
- Suggest logical next steps or advanced topics.
- Order items from most relevant to general resources.
- Use consistent link formatting for both internal and external links.

Exact text:

```markdown
## More to explore

We understand you might want to explore other capabilities not strictly on this path. We'll provide you opportunities where it makes sense.
```

### 8. Before you begin section (required)

- Always include this section.
- Always use the exact heading "## Before you begin".
- Always start with the phrase "Before you [action verb], ensure that you have:" or "Before you [action verb], you must:".
- List prerequisites in order of importance: account access, technical requirements, knowledge prerequisites.
- Use the exact text for Grafana Cloud account creation (often the first prerequisite): `- A Grafana Cloud account. To create an account, refer to [Grafana Cloud](https://grafana.com/signup/cloud/connect-account).`
- Include specific version requirements when relevant (for example, "Ubuntu 20.04 or later recommended").
- Include data availability requirements (metrics, logs, traces flowing into Grafana Cloud).
- Add familiarity requirements that help users succeed. Include shared snippets for common prerequisites using `{{< shared-snippet >}}` shortcode when available.

Opening patterns:

```markdown
Before you [action verb], ensure that you have:
Before you [action verb], you must:
```

Action verbs used:

- "use [Feature name]"
- "send [data type] to [destination]"
- "connect to [data source type]"
- "add and configure [plugin/component]"
- "visualize [data type] in a Grafana Cloud dashboard"
- "monitor [system type] using [tool]"

Standard prerequisites order:

1. **Grafana Cloud account (always first when applicable):**

```markdown
- A Grafana Cloud account. To create an account, refer to [Grafana Cloud](https://grafana.com/signup/cloud/connect-account).
```

1. **Technical requirements:**

- System access and permissions
- Software installation requirements
- Data availability requirements
- Network/connectivity requirements

1. **Knowledge prerequisites:**

- Familiarity with technologies
- Understanding of concepts
- Experience with tools

Common prerequisites patterns:

**Data availability:**

```markdown
- [Data type] up and running and connected as a data source, whether it's locally installed or using [Grafana Cloud Product], provided by Grafana Cloud.
- [Data type] ingested into [Product] using agents such as [Grafana Alloy](/docs/[product]/latest/send-data/alloy/), [OpenTelemetry Collector](/docs/[product]/latest/send-data/otel/), or other [third-party clients](/docs/[product]/latest/send-data/#third-party-clients).
```

**Experience requirements:**

```markdown
- Experience working with [Technology] and understand its basic setup and operation.
- Basic familiarity with [Product], [Query language], and basic Grafana navigation.
- A familiarity with the [data type] you plan to [action]. Understanding what these [items] measure or represent will help you [specific benefit].
```

**Access requirements:**

```markdown
- Access to the machine on which [Software] is installed.
- [Specific permission level] permissions.
- Access to your [data type] that does not contain private information.
```

**Related journey prerequisites:**

```markdown
- [Created a private connection to a data source](/docs/learning-journeys/private-data-source-connect) and the PDC agent is running.
  - While not required, setting up a private data source connection is highly recommended when you use Grafana Cloud to connect to a data source.
```
