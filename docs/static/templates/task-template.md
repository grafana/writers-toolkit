---
title: Task title
menuTitle: Task
description: Use this template when you write a task topic.
aliases:
  - /docs/writers-toolkit/latest/templates/task-template
weight: 100
keywords:
  - keyword
  - key
  - word
---
<!-- For more information about how to populate front matter, see [Topic front matter]({{< relref "../../front-matter/" >}}). -->

# Task title
<!-- vale Grafana.Quotes = NO -->
<!-- The task title is required. The task title succinctly describes the goal to accomplish, as the result of following the instructions. The task title contains a verb and an object. For example: "Create a dashboard". -->
<!-- vale Grafana.Quotes = YES -->

Add an introduction to the task.

<!-- The introduction is required. Add an introduction to describe what the task is and why it’s important to the user.

This section of a task topic can include conceptual material. However, limit conceptual information to only the task at hand.

If you find yourself writing a long introduction, consider creating a concept topic, and then write a shorter form of that concept in the task introduction. Finally, link to the longer concept topic for more information.
-->

**Before you begin:**

- Item
- Another item
- Yet another item

<!-- The _Before you begin_ section is optional. Use it to identify any prerequisite conditions or tasks to complete before proceeding with the task at hand. The links might sometimes be unrelated to the product, such as “Have this thing on hand.”.

This area can also include decisions that the user should make, or permissions that they might need to verify that they have, before they begin. -->

**To [task name]:**
<!-- Add a stem sentence only when you have _Before you begin_ items.
The stem sentence introduces the steps and provides a visual cue for users who scan content, and it lets them know that the steps are about to begin.
A stem sentence begins with the word 'To' and includes the name of the task.
If you want to provide additional information about a step, add it to a separate line and indent it.

For example: To build a dashboard: -->

1. Open your web browser and go to http://localhost:3000/.

   The default HTTP port that Grafana listens to is `3000` unless you have configured a different port.
1. On the sign-in page, enter `admin` for the username and password.
1. Click **Sign in**.

   If successful, you will see a prompt to change the password.

1. ...
<!-- Numbered steps provide a directive to the user. Steps explicitly tell the user what to do and formatted using 1. in Markdown so they get numbered automatically.

Write steps so that they contain one action, or possibly two related actions, such as _Copy and paste a value._ or _Save and quit the program._

If a sentence does not tell the reader to do something, then it is not a step.
-->
