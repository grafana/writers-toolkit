---
aliases:
  - /docs/writers-toolkit/structure/topic-types/scenario/
  - /docs/writers-toolkit/writing-guide/documentation-structure/topic-types/scenario/
date: "2022-10-27T16:43:50-04:00"
description: Learn how to write a scenario topic.
keywords:
  - topic types
  - template
  - scenario
menuTitle: Scenario
review_date: "2025-11-01"
title: Scenario topic
---

# Scenario topic

A scenario shows how to apply product knowledge to solve a real problem in a realistic context. Scenarios help you make decisions, solve problems, and combine concepts and tasks to reach an outcome. Unlike a tutorial, a scenario doesn't teach a single step-by-step procedure. It guides judgment and problem-solving in real situations.

## Scenario structure

Use this structure to keep scenarios consistent across products while allowing flexibility for scope and narrative.

- Topic title: Name the situation and goal. Use either a verb phrase or a noun phrase. For example, "Investigate a CPU spike during a deployment" or "CPU spike investigation during deployment." Choose the form that best communicates the real-world context.

- Overview: State the situation, the goal, and what the reader can do after completing the scenario. Highlight key decisions they'll make. Example: "By the end, you can identify the source of a CPU spike and decide whether to scale or optimize."

- Before you begin: List prerequisites such as assumptions, access, roles, data available, and any baseline knowledge expected.

- Approach (stages or decision points): Organize the body as a sequence of stages or micro-scenarios. For each stage:
  - Context: What you observe and why it matters.  
  - Actions: What to try first. Link to tasks for steps and to concepts for background. Keep inline steps minimal.  
  - Decision: What to conclude or choose based on evidence.  
  - Principle (optional): The underlying takeaway or method you can apply to other problems, such as "begin with the alert signal, check dashboards for trends, and drill into logs or traces to confirm the root cause".


- Verify outcomes: Describe what "good" looks like and how to tell you're done. For example, "p95 latency returns under target and the alert clears".

- Next steps: Link to related tutorials for learning and to tasks for implementation.

## Audience and scope

Scenarios are for users who have completed onboarding and want to apply what they've learned to real environments. Assume readers have a functioning setup with access, permissions, and live data. Unlike tutorials, which use canned datasets, scenarios use the reader's own data and context.

Scenario topics can vary in scope. You can write micro scenarios that focus on a single decision, such as "Alert fired: Is it noisy or real?", or end-to-end scenarios that cover broader workflows, such as "Respond to a latency regression after a deploy". Choose the smallest scope that produces a clear, transferable outcome.

## Narrative and style

Follow these guidelines for narrative and style:

- Use direct second person ("you") and keep the voice instructional, not fictional. A light narrative frame is fine when it clarifies context, but avoid over-storytelling.  
- Focus on decisions, signals, and trade-offs. State the principle behind each step to help readers apply it in new situations.  
- Keep step lists short and link to tasks for detail. Keep explanations concise and link to concepts for depth.

## Placement

Place scenarios under a **Scenarios** section within each product's documentation, or near **Learn** content depending on your information architecture (IA).  
Link scenarios from related concepts, tutorials, and tasks to create a clear learning path through the documentation.

## Write a scenario topic

Follow these steps to write a scenario topic:

1. Choose a real situation users face, such as on-call, rollout, migration, optimization, or investigation.  
1. Identify the few key principles and decisions that drive the path to resolution.  
1. Draft the "Overview" and "Before you begin" sections to set expectations and assumptions.  
1. Map 2-4 stages. In each, link to existing tasks and concepts instead of repeating steps.  
1. Add "Verify outcomes" and "Next steps" to reinforce learning and help readers apply it to their own environment.

## Difference between scenarios and other topic types

This table shows how scenarios differ from other content types:

| Type | Purpose | Data & path | Reader outcome |
|------|----------|--------------|----------------|
| Scenario | Apply knowledge to solve a realistic problem and practice judgment. | Often uses the reader's own environment, has multiple valid paths, and provides links to tasks and concepts. | Reader can solve a similar problem and adapt the approach to their context. |
| Tutorial | Teach a defined skill through reproducible steps. | Uses controlled inputs and expected results in a linear path. | Reader can replicate the taught workflow. |
| Task | Complete a specific action with numbered steps. | Minimal context, single path. | Reader can perform the action once they're in the right place. |
| Concept | Explain what something is and why it matters. | Explanatory only. No procedures. | Reader understands the idea and components. |
| Example / Use case | Show a common pattern, configuration, or system interaction, such as how alerting rules and dashboards work together. | Brief and abstract. May be non-procedural or focus on relationships rather than steps. | Reader recognizes when and where to apply a similar pattern in their own environment. |

{{< admonition type="tip" >}}
If you're writing many steps with canned data and guaranteed outputs, write a tutorial. If you're helping readers decide what to do with their own data to reach a goal, write a scenario.
{{< /admonition >}}

## Scenario template

Use this template to write a scenario:

````markdown
# <SCENARIO_TITLE>

## Overview
You're <CONTEXT>. Your goal is to <GOAL>. By the end, you can <CAPABILITY_OR_OUTCOME>.

## Before you begin
- Access/roles:
- Data available:
- Baseline knowledge:

## Stage 1: <DECISION_OR_INVESTIGATION_STEP>
Context:  
Actions: (link tasks)  
Decision:  
Principle: (optional)

## Stage 2: <NEXT_STEP>
...

## Verify outcomes

You know you're done when <SIGNALS_OR_THRESHOLDS>.

## Variations and what-ifs
- If <CONDITION>, see <LINK>.
- If <CONDITION>, try <LINK>.

## Next steps
- <TASK_LINK>
- <CONCEPT_LINK>
- <TUTORIAL_LINK>
````

## Use AI to write scenarios

You can use AI to speed up planning and drafting of scenarios. 

When you prompt an AI tool:
- Add context files such as your product docs, this scenario template, and the Grafana style guide.  
- Ask the AI to anonymize any transcripts and focus on the user's goal, decisions, and outcomes.  
- Review the draft for accuracy and adjust the tone to match Grafana documentation.

For more example prompts and workflows, refer to the [Docs AI Toolkit repository](https://github.com/grafana/docs-ai/).

## Examples

- [Drop low-value traces to reduce noise](https://grafana.com/docs/grafana-cloud/adaptive-telemetry/adaptive-traces/guides/drop-traces-scenario/)
- [Onboard a group of services](https://grafana.com/docs/grafana-cloud/adaptive-telemetry/adaptive-traces/guides/onboard-services-scenario/)




