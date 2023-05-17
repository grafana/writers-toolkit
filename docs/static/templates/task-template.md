---
title: Perform the task
menuTitle: Perform the task
description: Use this template when you write a task topic.
aliases:
  - /docs/writers-toolkit/latest/templates/task-template
weight: 100
keywords:
  - keyword 1
  - keyword 2
  - keyword 3
---

<!-- For more information about how to populate front matter, see [Topic front matter]({{< relref "../../front-matter/" >}}). -->

# Perform the task

<!-- The task title is required. The task title succinctly describes the goal to accomplish, as the result of following the instructions. The task title contains a verb and an object. For example: Create a dashboard -->

Add an introduction to the task.

<!-- The introduction is required. Add an introduction to describe what the task is and why it’s important to the user. What is the goal the user accomplishes with this task and what context would it be used?

This section of a task topic can include conceptual material. However, limit conceptual information to only the task at hand.

If you find yourself writing a long introduction, consider creating a concept topic, and then write a shorter form of that concept in the task introduction. Finally, link to the longer concept topic for more information.

Some procedures, like configuring a data source, may have more than one task to accomplish a goal. Use the multiple-tasks-template.md instead of this template.

For guidelines about writing a task topic, refer to the [Task topic](https://grafana.com/docs/writers-toolkit/writing-guide/topic-types/task/) documentation.
-->

## Before you begin

- _System requirement_
- Software version _n_ is installed
- Valid _software_ license

<!-- This section is optional. Use it to identify any prerequisite conditions (such as a specific version, license, or system requirement), permissions, any necessary decision, or tasks to complete before proceeding. Sometimes you might want to include a tip, such as **Tip:** Run the commands within a `screen` session.

Write each prerequisite as a full sentence or sentence fragment, using parallel structures.

If you do not need this section, delete it.
 -->

## Steps

To [task name]:

<!--
The stem sentence introduces the steps and provides a visual cue for users who scan content, and it lets them know that the steps are about to begin.
A stem sentence begins with the word 'To' and includes the name of the task.
If you want to provide additional information about a step, add it to a separate line and indent it.

For example: To build a dashboard: -->

1. Open your web browser and go to http://localhost:3000/.

   The default HTTP port that Grafana listens to is `3000` unless you have configured a different port.

1. On the sign-in page, enter `admin` for the username and password.
1. Click **Sign in**.

   If successful, you will see a prompt to change the password.

1. Check the current context for your Kubernetes cluster and make sure that it is correct:

   ```bash
   kubectl config current-context
   ```

1. ...
<!-- Numbered steps provide a directive to the user. Steps explicitly tell the user what to do and formatted using 1. in Markdown so they get numbered automatically.

Write steps so that they contain one action, or possibly two related actions, such as _Copy and paste a value._ or _Save and quit the program._

If a sentence does not tell the reader to do something, then it is not a step.

If a step is not required but provides additional features, you can mark that step as optional and describe when it should be completed.

Text and code blocks need to be properly indented underneath a step in the markdown file to align with the step's display block. If the indent is not correct, then the code block doesn't display underneath the associated step. Incorrect indents can also cause auto-numbering to restart at 1.  
-->
