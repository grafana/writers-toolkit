---
aliases:
  - /docs/writers-toolkit/structure/topic-types/
  - /docs/writers-toolkit/writing-guide/documentation-structure/topic-types/
date: "2022-10-27T16:43:50-04:00"
description: Learn to write different types of topics.
keywords:
  - topic types
  - template
  - concept
  - task
  - reference
review_date: "2024-05-30"
title: Topic types
weight: 400
---

# Topic types

Grafana Labs documentation uses different topic types: _concept_, _task_, _reference_, _tutorial_, and _scenario_.
When you write content, you should use one of these topic types.

Depending on the needs of a particular product area, select a topic type from the following table to learn about each.

<!-- vale Grafana.GoogleFirstPerson = NO -->

**[Concept](https://grafana.com/docs/writers-toolkit/structure/topic-types/concept/)**
: Provides an overview and background information. Answers the question "What is it?".

**[Task](https://grafana.com/docs/writers-toolkit/structure/topic-types/task/)**
: Provides numbered steps that describe how to achieve an outcome. Answers the question "How do I?".

**[Reference](https://grafana.com/docs/writers-toolkit/structure/topic-types/reference/)**
: Provides users with the information they might need to refer to during a task. Answers the question "What details do I need to accomplish this task?".

**[Tutorial](https://grafana.com/docs/writers-toolkit/structure/topic-types/tutorial/)**
: Provides procedures that users can safely reproduce and learn from. Answers the question: "Can you teach me to …?"

**[Scenario](https://grafana.com/docs/writers-toolkit/structure/topic-types/scenario/)**
: Provides guidance to apply knowledge to solve a real problem in context. Answers the question "How should I approach this situation?".

**[Section](https://grafana.com/docs/writers-toolkit/structure/topic-types/section/)**
: Provides a landing page for users to find content.

<!-- vale Grafana.GoogleFirstPerson = YES -->

For your convenience, there are topic [templates](https://github.com/grafana/writers-toolkit/tree/main/docs/static/templates).

{{< admonition type="tip" >}}
You can create a topic from its template with the `topic/<TYPE>` GNU Make target.

For example, to create a task topic at the path `sources/task.md`:

```console
make topic/task TOPIC_PATH=sources/task.md
```

Types include:

- **Concept**: `make topic/concept`
- **Reference**: `make topic/reference`
- **Task**: `make topic/task`
- **Multiple tasks**: `make topic/multiple-tasks`
- **Section**: `make topic/section`
- **Visualization**: `make topic/visualization`

{{< /admonition >}}

## Templates for standardized topics

In addition to the primary topic types, there are also templates for specific topic types to ensure that pages documenting the same subject matter have a standard format.

These [templates](https://github.com/grafana/writers-toolkit/tree/main/docs/static/templates) are in the same directory as the topic type templates.

| Type                                                                                           | Description                                                                          |
| ---------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| [Visualization](https://grafana.com/docs/writers-toolkit/structure/topic-types/visualization/) | Describes a visualization type. May include conceptual, task, and reference content. |
