---
title: Perform multiple tasks
menuTitle: Perform multiple tasks
description: Use this template when you document a procedure with multiple task topics.
aliases:
  - /docs/writers-toolkit/latest/templates/multiple-tasks-template
weight: 100
keywords:
  - keyword 1
  - keyword 2
  - keyword 3
---

<!-- For more information about how to populate front matter, see [Topic front matter]({{< relref "../../front-matter/" >}}). -->

# Perform multiple tasks

<!-- The task title is required. The task title succinctly describes the goal to accomplish, as the result of following the instructions. The task title contains a verb and an object. For example: Create a dashboard -->

Add an introduction to the task.

<!-- The introduction is required. Add an introduction to describe what the task is and why it’s important to the user. What is the goal the user accomplishes with this task and in what context would it be used?

Use this template for procedures that contain a series of related tasks. If you only have one task, start from the [task-template.md](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/task-template.md) file instead.

This section of a task topic can include conceptual material. However, limit conceptual information to only the tasks at hand.

If you write a long introduction, consider creating a concept topic. Next, write a shorter form of that concept in the task introduction, and link to the longer concept topic for more information.

Some procedures, such as configuring a data source, have more than one task to accomplish a goal. Summarize the overall procedures in the introduction using a numbered list. In the following example, each step matches a task section heading.
Replace any text in _italics_ with content for your procedure.
-->

_Goal of procedure_:

1. [Configure service](#configure-service)
1. [Set up the environment](#set-up-the-environment)
1. [Optional: Install libraries](#optional-install-libraries)
1. [Verify the installation](#verify-the-installation)

## Before you begin

- _System requirement_
- Software version _n_ is installed
- Valid _software_ license

<!-- This section is optional. Use it to identify any prerequisite conditions (such as a specific version, license, or system requirement), permissions, any necessary decision, or tasks to complete before proceeding. Sometimes you might want to include a tip, such as **Tip:** Run the commands within a screen session.

Replace any text in _italics_ with content for your procedure and remove or add lines as needed.

Write each prerequisite as a full sentence or sentence fragment, using parallel structures.

If you have more than one task, include all prerequisites in this section. For example, if you have a page that configures a widget and several tasks have prerequisites, list all prerequisites in the Before you begin section. This way users can have everything they need before they start performing the tasks.

If you do not need this section, delete it.
 -->

## Configure service

<!-- Optional: Add an introductory sentence to this task. For example, summarize the purpose of this task in relation to the overall procedure.    -->

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
<!-- Numbered steps provide a directive to the user; they tell the user explicitly what to do. Format steps using 1. in Markdown so they get numbered automatically.

Write steps so that they contain one action, or possibly two related actions, such as _Copy and paste a value._ or _Save and quit the program._

If a sentence does not tell the reader to do something, then it is not a step.

To add context that is directly related to a step, or to add a code block, indent it underneath the step. Doing so properly scopes the added information to the step.
-->

## Set up the environment

<!-- This section provides an example of nested steps with code blocks. -->

1.  Create a file named `sample.yaml` and copy the following YAML configuration into it:

    ```yaml
    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
    ```

1.  Run the following command to apply the `sample.yaml` file:

        ```bash
        kubectl apply --namespace sample -f sample.yaml
        ```

1.  To check that MinIO is correctly configured, sign in and verify that a bucket has been created.

    1. Port-forward the service to port 9001:
       ```bash
        kubectl port-forward --namespace tempo service/sample 9001:9001
       ```
    1. Navigate to the admin bash using your browser: `https://localhost:9001`. The sign-in credentials are username `user` and password `user123`.
    1. Verify that the Buckets page lists `sample-data`.

    Without these buckets, no data will be stored.

## Optional: Install libraries

<!-- Optional: If a task is not required but provides additional features, you can mark that section as optional and describe when it should be completed. If this section is not needed, delete it.
-->

## Verify the installation

 <!-- 
 Optional: Use this section if you have more than one task that needs to be performed. Otherwise, delete this section. 
 
 Consider adding a verification section so users can validate the outcome of the tasks. 
 -->
