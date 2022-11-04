---
title: Tutorial topic
menuTitle: Tutorial
description: Learn how to write a tutorial topic.
aliases:
  - /docs/writers-toolkit/latest/writing-guidelines/tutorial/
weight: 400
keywords:
  - topic types
  - template
  - tutorial
---

# Tutorial topic

The purpose of a tutorial is to show the reader how to "learn by doing" in a safe environment. A tutorial should build up easy successes that inspire a user to say, “I just did that! Wow!” The length of a tutorial can vary from a few steps to many subtasks or even to many modules.

## Tutorial structure

A tutorial topic includes the following elements:

**Topic title:** Write a tutorial topic title that combines a verb and an object.

**Overview:** Let the user know the goal they will achieve by completing the tutorial. Suggested text: "In this tutorial, you will …" You can instead use: "By the end of this tutorial, you'll be able to" and include a list of learning objectives. 

- There can be conceptual material in this section of a tutorial topic. Limit conceptual information to only what is relevant to the goal at hand.
- If you find yourself writing a long introduction, consider creating a concept topic, and then writing a shorter form of that concept in the tutorial introduction. The longer concept topic can be accessed for more information by linking to it.

**Background (optional):** Provide context for the tutorial. Limit context information to only what is relevant to the goal at hand. Create a concept topic for more extensive content. 

**Before you begin (optional):** Describe or add links to tasks that need to be completed before the tutorial. The links might sometimes be unrelated to the product, such as “Have this thing at hand."
- Additionally, this section can include decisions the user should make or permissions they need to confirm before starting the tutorial.
- Use a bulleted list if there is more than one prerequisite.
- If there are no prerequisites, do not include this section.

**Task section (or sections)**: Create a section for each task needed to complete the tutorial. Follow the [task guidelines]({{< relref "../task/" >}}) to write the tasks. 

- Do not include written step numbers in the header, for example, "Step 1: Pick apples." Instead, include just the verb and object, for example "Pick apples."
- Include only the tasks required for a straight path to the tutorial's goal, not optional or alternative paths. 
- Minimize the explanation within task steps. Instead, link to supporting explanations in related concept, task, reference, and how-to topics. 

**Summary (optional):** Describe what the tutorial user has accomplished. 

**Next steps (optional):** Provide logical next steps, if they exist.  

![Tutorial structure](tutorial.png)

## Write a tutorial topic

To write a tutorial, complete these steps:

1. Determine where you want to add a tutorial to the Grafana Labs product documentation.
1. Create a child directory within the parent directory that follows this naming convention:
   
   - The directory name should include a verb and an object.
   - Use lowercase letters.
   - Add a hyphen between words.
  <br>
  <br>
   For example:
     - manage-dashboard-permissions
     - manage-organization-users
<br>
<br>

1. Create an `index.md` file within the task directory.
1. Add front matter to the `index` file.

   For more information about front matter, refer to [Front matter]({{< relref "../../front-matter" >}}).

1. Add the content to a copy of the [Tutorial template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/tutorial-template.md).

## Tutorial template

When you are ready to write, make a copy of the [Tutorial template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/tutorial-template.md) and add your content.

## Difference between tutorials and task topics

 The difference between a tutorial and a task topic is that a tutorial is for learning, and a task is for actual operational work. Another important distinction is that a tutorial typically provides a "sandbox" environment&mdash;a source of data that users can safely experiment with. Depending on the application, the  data might be:

  - In a sandboxes
  - On test servers
  - In demo repos that the user clones locally
 
 As an example, the [Play with Grafana Mimir](https://grafana.com/tutorials/play-with-grafana-mimir/) tutorial provides a repo that users can clone in order to complete the tutorial. As a comparison, the Mimir [Storing exemplars in Grafana Mimir](https://grafana.com/docs/mimir/latest/operators-guide/use-exemplars/storing-exemplars/) topic is a pure task that a user would follow to complete their work.   