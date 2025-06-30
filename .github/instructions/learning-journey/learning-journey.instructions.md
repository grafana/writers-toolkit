---
applyTo: "learning-journeys/**/*.md"
---

# Copilot instructions: Learning journeys

A learning journey is a multi-step tutorial that helps users achieve a specific product goal using Grafana Labs products.
Each learning journey is a folder in `learning-journeys/`, with an `_index.md` file for the journey overview and an `index.md` file in each subfolder for individual steps.

## Folder and file structure

- Each learning journey is a folder under `learning-journeys/`.
- The journey overview is in `_index.md` at the root of the journey folder.
- Each step is a subfolder with its own `index.md`.
- Use kebab-case for folder names (for example, `monitor-linux-servers`).
- Use descriptive step folder names (for example, `install-alloy`, `configure-grafana`).

Example:

```
learning-journeys/
  monitor-linux-servers/
    _index.md
    install-alloy/
      index.md
    configure-grafana/
      index.md
    create-dashboard/
      index.md
```

## Steps

### Naming patterns

Step folder names use descriptive kebab-case that indicates the action or topic:

**Business value and introduction steps:**

- `business-value` / `business-value-[product]` - Explains why use the product
- `advantages` / `advantages-[product]` - Lists product benefits
- `value-of-[feature]` - Explains feature value

**Installation and setup steps:**

- `install-[product]` - Install software or components
- `add-data-source` - Add data source configuration
- `create-token` – Generate authentication tokens
- `config-authentication` / `enter-authentication-credentials` – Set up authentication

**Configuration and data steps:**

- `select-[item]` / `choose-[item]` – Make configuration choices
- `search-[item]` / `search-and-filter` – Find and filter data
- `send-[data-type]` – Configure data sending
- `verify-[item]` – Confirm setup works

**Build and test steps:**

- `build-[component]` / `create-[artifact]` – Build applications or components
- `run-[application]` / `execute-[script]` – Run applications or execute scripts
- `test-[functionality]` / `validate-[setup]` – Test functionality or validate configuration
- `deploy-[service]` / `publish-[artifact]` – Deploy services or publish artifacts

**Analysis and exploration steps:**

- `analyze-data` / `investigate-data` – Examine collected data
- `view-[item]` – Display specific views or lists
- `open-[location]` / `open-in-explore` – Navigate to analysis tools

**Dashboard and visualization steps:**

- `add-visualization` / `create-dashboard` – Build visualizations
- `write-query` – Create queries for data
- `add-to-dashboard` / `add-[item]-dashboard` – Save to dashboards
- `time-range-refresh` – Configure time settings
- `save-dashboard` – Save dashboard configurations

**Completion steps:**

- `end-journey` / `destination-reached` – Journey conclusion

### Ordering patterns

Steps follow logical progression patterns:

**Infrastructure monitoring journeys:**

1. Business value/case for observability (weight: 100)
2. Product advantages (weight: 200)
3. Installation/setup (weight: 300-400)
4. Configuration (weight: 500-600)
5. Verification (weight: 700)
6. Usage/analysis (weight: 800)
7. Conclusion (weight: 900)

**Data exploration journeys:**

1. Business value (weight: 100)
2. Navigation/access (weight: 200)
3. Search and filter (weight: 300)
4. Analysis/investigation (weight: 400-500)
5. Open in Explore (weight: 600)
6. Add to dashboard (weight: 700)
7. Conclusion (weight: 800)

**Visualization journeys:**

1. Business value (weight: 100)
2. Create dashboard (weight: 300)
3. Write query (weight: 400)
4. Configure time/refresh (weight: 500)
5. Save dashboard (weight: 600)
6. Conclusion (weight: 700)

**Data source connection journeys:**

1. Business value (weight: 100)
2. Product advantages (weight: 200)
3. Install plugin (weight: 300)
4. Add data source (weight: 400)
5. Configure authentication (weight: 500)
6. Test connection (weight: 600-700)
7. Build dashboard (weight: 800)
8. Conclusion (weight: 900)

### Weight increments

- Use increments of 100 for weight values
- Start at weight 100 for the first step after the overview
- Allow gaps for future insertion of steps
- Conclusion steps typically have the highest weight

## Shared Content Integration

**Guidelines:**

- Use `{{< shared-snippet >}}` shortcodes for reusable content blocks
- Use a shared snippet shortcode to reuse and embed eixistin documentation snippets
- Include the code snippet, and if the snippet came from an attached context file,
  also include the shared snippet shortcode linking to the file
- Reference shared prerequisites from other documentation sections
- Include `{{% shared-snippet %}}` for content that needs to be processed as markdown

**Examples:**

```markdown
{{< shared-snippet path="/docs/grafana-cloud/monitor-infrastructure/integrations/get-started.md" id="grafana-cloud-setup" >}}

{{% shared-snippet path="/docs/fleet-management/next/_index.md" id="flt-mgt-lj-intro" %}}
```
