---
aliases:
  - /docs/writers-toolkit/structure/topic-types/learning-journey/
  - /docs/writers-toolkit/writing-guide/documentation-structure/topic-types/learning-journey/
date: "2024-12-19T00:00:00-00:00"
description: Learn how to write a learning journey topic.
keywords:
  - topic types
  - template
  - learning journey
  - milestone
  - pathfinder
menuTitle: Learning journey
review_date: "2024-12-19"
title: Learning journey topic
---

<!-- vale Grafana.Gerunds = NO -->
<!-- vale Grafana.GoogleWill = NO -->

# Learning journey topic

A learning journey is a structured and opinionated guide that takes users through a complete workflow or process to achieve a specific goal.
Learning journeys break down complex tasks into milestones.

Learning journeys are designed to:

- Guide users through a complete workflow or process
- Break complex tasks into sequential milestones
- Provide interactive, hands-on experiences where appropriate
- Build knowledge progressively from one milestone to the next
- Include optional side paths for users who want to explore related topics

## Learning journey structure

A learning journey topic includes the following elements:

- **Topic title:** Write a learning journey title that describes the goal or outcome for users.
  For example, "Monitor collector health using Grafana Fleet Management".

- **Introduction:** Provide an overview of what users will learn and accomplish.
  Include a video introduction if available, and explain the value of completing the journey.

- **Before you begin (optional):** Describe prerequisites, required permissions, or setup tasks users should complete before starting the journey.
  Use a bulleted list if there are multiple prerequisites.

- **Milestones:** Each milestone represents a step in the journey.
  Milestones are organized sequentially and build upon each other.
  Each milestone has its own directory with an `index.md` file.

  Milestones can be:

  - **Standard milestones:** Regular Markdown content.
  - **Interactive milestones:** Use interactive JSON stored in `pathfinder.json` and the `pathfinder/json` shortcode to provide guided, hands-on experiences.

    For more information about interactive milestones, refer to [Interactive milestones](#interactive-milestones).

- **End journey milestone:** A final milestone that congratulates users on completing the journey and summarizes what they've accomplished.
  Optionally, provide links to related journeys or next steps.

- **Side journeys (optional):** Provide opportunities for users to explore related topics at specific points in the journey.
  These appear in milestone front matter and allow users to take detours without losing their place.

## Learning journey front matter

The main learning journey `_index.md` file requires specific front matter fields:

```yaml
---
menuTitle: <SHORT TITLE>
title: <FULL TITLE>
description: <DESCRIPTION>
weight: <NUMBER>
journey:
  group: <GROUP NAME>
  skill: <Beginner|Intermediate|Advanced>
  source: <SOURCE TYPE>
  logo:
    src: <LOGO PATH>
    background: <COLOR>
    width: <NUMBER>
    height: <NUMBER>
step: 1
layout: single-journey
cascade:
  layout: single-journey
cta:
  type: start
  title: <CTA TITLE>
  cta_text: <CTA TEXT>
related_journeys:
  title: <TITLE>
  heading: <HEADING>
  items:
    - title: <JOURNEY TITLE>
      link: <JOURNEY LINK>
---
```

### Front matter fields

- **`menuTitle`:** Short title for navigation menus.
- **`title`:** Full title of the learning journey.
- **`description`:** Brief description of the journey.
- **`weight`:** Numeric value for ordering journeys.
- **`journey.group`:** Category or group for the journey.
- **`journey.skill`:** Skill level: "Beginner", "Intermediate", or "Advanced".
- **`journey.source`:** Source type
- **`journey.logo`:** Logo configuration with source path, background color, width, and height.
- **`step`:** Always `1` for the main journey page.
- **`layout`:** Always `single-journey` for learning journey pages.
- **`cascade.layout`:** Ensures all child pages use the `single-journey` layout.
- **`cta`:** Call-to-action configuration for the start button.
- **`related_journeys`:** Optional list of related journeys to suggest to users.

## Milestone structure

Each milestone is a separate directory with an `index.md` file.
Milestone directories should use descriptive, lowercase names with hyphens like `create-token` or `check-health-status`.

### Milestone front matter

Each milestone requires the following front matter:

```yaml
---
menuTitle: <MILESTONE_TITLE>
title: <FULL_MILESTONE_TITLE>
description: <DESCRIPTION>
weight: <NUMBER>
step: <STEP_NUMBER>
layout: single-journey
cta:
  type: continue
side_journeys:
  title: More to explore (optional)
  heading: <HEADING>
  items:
    - title: <SIDE_JOURNEY_TITLE>
      link: <SIDE_JOURNEY_LINK>
---
```

### Milestone front matter fields

- **`menuTitle`:** Short title for the milestone.
- **`title`:** Full title of the milestone.
- **`description`:** Brief description of what users will accomplish in this milestone.
- **`step`:** Sequential step number in the journey (starts at 2 for the first milestone after the introduction).
- **`layout`:** Always `single-journey`.
- **`cta.type`:**
  - `continue` for intermediate milestones
  - `conclusion` for the final milestone
- **`side_journeys`:** Optional list of related topics users can explore at this point.

### End journey milestone

The final milestone uses `cta.type: conclusion` and typically includes:

- A congratulatory message
- A summary of what users accomplished
- Links to related journeys or next steps

Example:

```yaml
---
cta:
  type: conclusion
  image:
    src: /media/docs/learning-journey/journey-conclusion-header-1.svg
    width: 735
    height: 175
---
```

## Interactive milestones

Interactive milestones use JSON blocks to define the interactive behavior.

### Creating an interactive milestone

To create an interactive milestone:

1. Create a milestone directory with an `index.md` file as described in [Milestone structure](#milestone-structure).

1. Add the `pathfinder/json` shortcode to the milestone's `index.md` file content.
   Content is everything after the front matter:

   ```markdown
   {{< pathfinder/json >}}
   ```

1. Create a `pathfinder.json` file in the same directory as the milestone's `index.md` file.

1. Define the interactive content in `pathfinder.json` using the Pathfinder JSON format.

   The JSON structure includes:

   - **`id`:** Unique identifier for the guide section
   - **`title`:** Title of the interactive guide
   - **`blocks`:** Array of content blocks that can include:
     - **`type: "video"`:** Embedded video content
     - **`type: "markdown"`:** Markdown content
     - **`type: "section"`:** Grouped blocks
     - **`type: "interactive"`:** Interactive UI guidance with actions like `highlight`, `hover`, `formfill`
     - **`type: "guided"`:** Multi-step guided interactions

### Pathfinder JSON example

```json
{
  "id": "guide-section-36937",
  "title": "Create an access token",
  "blocks": [
    {
      "type": "video",
      "provider": "youtube",
      "title": "Create access token",
      "src": "https://www.youtube.com/embed/PrjTxzc27Ns",
      "start": 166,
      "end": 216
    },
    {
      "type": "markdown",
      "content": "If your Grafana Alloy collector isn't shown on the Fleet Management page, you can manually register it.\n\nIn this milestone, you'll generate an access token that registers your collector with Fleet Management."
    },
    {
      "type": "section",
      "blocks": [
        {
          "type": "interactive",
          "action": "highlight",
          "reftarget": "a[data-testid='data-testid Nav menu item'][href='/a/grafana-collector-app/alloy']",
          "requirements": ["navmenu-open", "exists-reftarget"],
          "content": "In your Grafana Cloud instance, navigate to **Connections** > **Collector** > **Configure**."
        }
      ]
    }
  ]
}
```

For detailed information about the interactive JSON format, refer to the [JSON guide format reference](https://github.com/grafana/grafana-pathfinder-app/blob/main/docs/developer/interactive-examples/json-guide-format.md).

## Write a learning journey topic

To write a learning journey, complete these steps:

1. Determine the goal and scope of your learning journey.
   Work with a Subject Matter Expert (SME) to identify:

   - The end-to-end workflow users need to complete
   - The sequential milestones required to achieve the goal
   - Which milestones should be interactive vs. standard
   - Prerequisites users need before starting

1. Create a directory structure in your project repository.

   You contribute learning journeys to the website repository at `content/docs/learning-journey/<JOURNEY>`.

1. Create the main journey `_index.md` file with the required front matter.

1. Create milestone directories, each with an `index.md` file.

   - Use descriptive, lowercase directory names with hyphens
   - Number milestones sequentially (weight increments of 100, step increments of 1)
   - Start the first milestone at `step: 2` (the introduction is `step: 1`)

1. For interactive milestones, create `pathfinder.json` files and add the `{{< pathfinder/json >}}` shortcode to the milestone's `index.md`.

1. Add content to each milestone following the structure described in [Milestone structure](#milestone-structure).

1. Create an end journey milestone that congratulates users and summarizes accomplishments.

1. Add front matter to all files.
