---
aliases:
  - /docs/writers-toolkit/writing-guide/documentation-structure/topic-types/tutorial
  - /docs/writers-toolkit/structure/topic-types/tutorial/
date: "2022-10-27T16:43:50-04:00"
description: Learn how to write a tutorial topic.
keywords:
  - topic types
  - template
  - tutorial
menuTitle: Tutorial
review_date: "2024-06-07"
title: Tutorial topic
---

# Tutorial topic

The purpose of a tutorial is to show the reader how to "learn by doing" in a safe environment.
A tutorial should build up quick successes.
The length of a tutorial can vary from a few steps to many subtasks.

If you have an idea for a tutorial you'd like to develop, contact the Grafana Labs documentation team.
Internal contributors can reach out on Slack and external contributors can send email to [`docs@grafana.com`](mailto:docs@grafana.com) or reach out on the #docs channel on the [Grafana Labs Community Slack](https://slack.grafana.com/).

## Tutorial structure

<!-- vale Grafana.GoogleWill = NO -->
<!-- "will" is commonly used for the future outcomes of tutorials -->

A tutorial topic includes the following elements:

- **Topic title:** Write a tutorial topic title that combines a verb and an object.

- **Overview:** Let the user know the goal they will achieve by completing the tutorial.
  Provide context and include a list of the tasks the user will complete.
  Suggested text: "In this tutorial, you will …".

  There can be conceptual material in this section of a tutorial topic.
  However, limit conceptual information to only what's relevant to the goal at hand.

  If you find yourself writing a long introduction, consider creating a concept topic, and then writing a shorter form of that concept in the tutorial introduction.
  You can link to the concept topic from the tutorial.

- **Before you begin (optional):** Describe or add links to tasks that a reader should complete before the tutorial.
  The links might sometimes be unrelated to the product, such as "Have this thing at hand."

  Additionally, this section can include decisions the user should make or permissions they need to confirm before starting the tutorial.
  Use a bulleted list if there is more than one prerequisite.

  If there are no prerequisites, don't include this section.

- **Task section (or sections)**: Create a section for each task needed to complete the tutorial.
  Follow the [task guidelines](https://grafana.com/docs/writers-toolkit/structure/topic-types/task/) to write the tasks.

  To determine what tasks and steps you should include in your tutorial, perform a goal analysis and determine the valuable outcome the user wants to achieve.
  Limit the tutorial to the tasks needed to satisfy that goal.
  Work with a Subject Matter Expert (SME) to determine the goal and the minimal set of tasks.
  If possible, record the SME completing the tasks needed to accomplish the goal or ask the SME to record a demo of the tasks if that's preferable.

  Work with a Subject Matter Expert (SME) to:

  - Provide steps that explain how to access or set up the data needed to complete the task.
    For more information, refer to [Data for your tutorial](#data-for-your-tutorial).
  - Don't include written step numbers in the header, for example, "Step 1: Pick apples.".
    Instead, include just the verb and object, for example "Pick apples."
  - Include only the tasks required for a straight path to the tutorial's goal, not optional or alternative paths.
  - Minimize the explanation within task steps.
    Instead, link to supporting explanations in related concept, task, and reference topics.

- **Summary (optional):** Describe what the tutorial user has accomplished.

- **Next steps (optional):** Provide logical next steps, if they exist.

{{< figure src="/media/docs/writers-toolkit/tutorial.png" alt="Annotated example of a tutorial page's structure" >}}

## Write a tutorial topic

To write a tutorial, complete these steps:

1. Add a `docs/sources/tutorials` directory to your project repository if one doesn't yet exist.

   The tutorial is committed alongside the other documentation in your repository, and after it's published, it's displayed on the Grafana [Tutorials](https://grafana.com/tutorials/) page.
   For more information, refer to [Publish your tutorial](#publish-your-tutorial).

1. Create a child directory within the `tutorials` directory that follows this naming convention:

   - The directory name should include a verb and an object.
   - Use lowercase letters.
   - Add a hyphen between words.

1. Create an `index.md` file within the tutorial's directory.

1. Add the content to a copy of the [Tutorial template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/tutorial-template.md).

1. Add front matter to the `index.md` file.

   For more information about front matter, refer to [Front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/).

## Tutorial template

When you are ready to write, make a copy of the [Tutorial template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/tutorial-template.md) and add your content.

## Difference between tutorials and task topics

The difference between a tutorial and a task topic is that a tutorial is for learning, and a task is for actual operational work.
Another important distinction is that a tutorial typically provides a "sandbox" environment—a source of data that users can safely experiment with.

## Data for your tutorial

Depending on the application, your tutorial's data might be:

- In a sandbox
- On test servers
- In demo repositories that the user clones locally

For example, the [Play with Grafana Mimir](https://grafana.com/tutorials/play-with-grafana-mimir/) tutorial provides a repository that users can clone to complete the tutorial.
In contrast, the [Store exemplars in Grafana Mimir](https://grafana.com/docs/mimir/latest/manage/use-exemplars/store-exemplars/) topic is a pure task that a user would follow to complete their work.
For guidance on writing tasks, refer to [Tasks](https://grafana.com/docs/writers-toolkit/structure/topic-types/task/).

If getting access to the tutorial data is complex, include the instructions in the steps of the tutorial.
If getting access to the data is straightforward, include it in the "Before you begin" section.

## Publish your tutorial

You store your tutorial source in your project repository in a `docs/sources/tutorials` directory and mounted to the tutorials repository so that it's displayed on the [Tutorials](https://grafana.com/tutorials/) page.
You store the source in your project repository for team members to review and edit the content.

The following sections describe how to hide the tutorial from your project's table of contents and to display it on the Tutorials page.

### Hide your tutorial from your table of contents

Tutorials are for learning, so it's best to keep them together on the Tutorials page, accessible directly from the Grafana website's **Learn** menu.
As such, you need to hide the tutorial so that it doesn't appear in your project's table of contents.

#### Before you begin

Before completing these steps, you need to create a `tutorials` directory under `docs/sources` and add your tutorial into its own subdirectory as described in the [Write a tutorial topic](#write-a-tutorial-topic) section.

To hide your tutorial from your documentation's table of contents:

1. Open your project's `docs/sources/_index.md` file.

1. Add the following YAML to the page's front matter, replacing _`<PROJECT>`_ with your project URL path.

   If the `cascade` front matter already exists, you must merge this snippet with the existing front matter.

   {{< admonition type="note" >}}
   Hugo `cascade` front matter can have two forms, _array_ and _mapping_.

   The following snippet uses the array form of the `cascade` front matter.
   If your front matter isn't already in the array form, you will need to change it to that form.

   For more information, refer to the [`cascade` front matter documentation](/docs/writers-toolkit/write/front-matter/#cascade)
   {{< /admonition >}}

   ```yaml
   cascade:
     - _target:
         path: /docs/<PROJECT>/*/tutorials/**
       _build:
         list: false
         render: false
   ```

### Add your tutorial to the Tutorials page

{{< admonition type="note" >}}
This procedure is for writers who have permissions to update the Grafana website repository.
{{< /admonition >}}

To add your tutorial to the Tutorials page:

1. Add the following YAML to the `manual_mounts` field in the `config/_default/params.yaml` file in the website repository.

   Replace _`<PROJECT>`_ with your project name, _`<TUTORIAL>`_ with the directory of your tutorial, and _`<VERSION>`_ with the version of documentation you want to mount.
   Typically _`<VERSION>`_ is either "next" or "latest".

   ```yaml
   - source: content/docs/<PROJECT>/<VERSION>/tutorials/<TUTORIAL>
     target: content/tutorials/k8s-monitoring-app
   ```

1. Add the following YAML to the `list` field in the `data/tutorials.yaml` file in the website repository.

   Replace _`<TUTORIAL>`_ with the directory of your tutorial and _`<LEVEL>`_ with one of "beginner", "intermediate", or "advanced" depending on the difficulty of your tutorial.

   ```yaml
   - page: /tutorials/<TUTORIAL>
     level: <LEVEL>
     type: tutorial
   ```

## Tutorial topic examples

Refer to the following for tutorial examples:

- [Tutorials page](https://grafana.com/tutorials/)
- [Tutorial example (Monitor an app with Kubernetes Monitoring)](https://grafana.com/tutorials/k8s-monitoring-app/)
