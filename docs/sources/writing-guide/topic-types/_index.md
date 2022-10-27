---
title: Topic types
menuTitle: Topic types
description: Topic types that we use at Grafana Labs.
aliases:
  - /docs/writers-toolkit/latest/writing-guidelines/topic-types/
weight: 400
keywords:
  - topic types
  - template
  - concept
  - task
  - reference
---

# Topic types

The Grafana Labs documentation is divided into different topic types. The main types are concepts, tasks, and references. We encouage you to use these basic topic types when you write content. Avoid mixing topic types within a single topic file. 

The Grafana Labs documentation includes other types of topics depending on the needs of particular product areas. Select a topic type from the following table to learn about each.

Topic type | Description
---|---
[Concept]({{< relref “concept/index” >}}) | Provides an overview and background information. Answers the question "What is it?".
[Task]({{< relref “task/index” >}}) | Provides numbered steps that describe how to achieve an outcome. Answers the question "How do I?".
[Reference]({{< relref “reference/index” >}}) | Provides users with the information they might need to refer to during a task. Answers the question "What details do I need to accomplish this task?".
[Tutorial]({{< relref “tutorial/index” >}}) |  

For your convenience, we have created topic templates. Please refer to [Templates](https://github.com/grafana/writers-toolkit/tree/main/docs/static/templates) for more information.



## Concepts

A concept provides an overview and background information to help end users understand a product, interface, or task. Concepts answer the question “what is it?”. Readers learn about features through concepts.

The following types of content can be included in concepts:

- Detailed overviews of Grafana's features with benefits and clearly defined terms
- Diagrams that help users understand the components of a system
- Process flow diagrams
- Best practice guidelines
- An example of how a feature is used. Examples might include screenshots or other supporting visuals

A concept topic does not include:

- Step-by-step instructions
- Reference information, such as lookup tables or lists of values

### Concept topic structure

A _concept_ topic includes the following elements:

- **Topic title:** Topic titles should be nouns, for example, Grafana panels. By using this naming convention, readers are able to distinguish between conceptual topics and tasks that begin with verbs.
- **Introduction:** Include an introduction that explains what this topic is about.
- **Body:** Provide as much content as needed to explain the concept thoroughly. There can be sections, visuals, and text in the body of a concept.

![Concept structure](concept.png)

### Write a concept topic

To write a concept topic, follow these steps.

1. Decide which top-level entity you want to add documentation to by reviewing Grafana Labs' product documentation.
1. Within the top-level entity, create a parent directory with the following naming convention:

   - Use a noun
   - Use lowercase letters
   - Add a hyphen between words
  <br>
  <br>
   For example:
     - organization-management
     - alert-groups
     - installation
     - service-accounts
<br>
<br>

1. Within the parent directory, create an `_index.md` file.
1. Add front matter to the `_index` file.

   For more information about front matter, refer to [Front matter]({{< relref "../front-matter" >}}).

1. Add the content to a copy of the [Concept template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/concept-template.md).

   For more information about the kinds of content you can add to a concept topic, refer to [Concepts](#concepts).

### Concept topic examples

Refer to the following topics for concept topic examples:

- [Roles and permissions](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/)
- [Deployment modes](https://grafana.com/docs/loki/next/fundamentals/architecture/deployment-modes/)
- [Best practices for managing dashboards](https://grafana.com/docs/grafana/latest/best-practices/best-practices-for-managing-dashboards/)

### Concept template

When you are ready to write, make a copy of the [Concept template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/concept-template.md) and add your content.

## Tasks

Task topics include numbered steps that describe how to achieve an outcome. 

### Task structure

A _task_ topic includes the following elements:

**Topic title:** Write a task topic title that combines a verb and an object.

**Introduction:** Provide an introduction that explains why the end user should care about the task.

- There might be conceptual material in this section of a task topic. Limit conceptual information to only what is relevant to the task at hand.
- If you find yourself writing a long introduction, consider creating a concept topic, and then writing a shorter form of that concept in the task introduction. The longer concept topic can be accessed for more information by linking to it.

**Before you begin: (optional)** Add links to tasks that need to be completed before the current one. The links might sometimes be unrelated to the product, such as “Have this thing at hand”.
- Additionally, this section can include decisions the user should make or permissions they need to confirm before starting the task.
- If there are no prerequisites, do not include this section.

**Stem sentence: (optional)** The stem sentence introduces the steps and signals to users who scan content that the steps are about to begin. You should include a stem sentence only when you include a before you begin section.

**Steps:** Users are provided with a directive through numbered steps.
- Write steps so that they contain one action, or possibly two related actions, such as _Copy and paste a value_ or _Save and quit the program._
- Unless a sentence instructs the reader to act, it isn't a step.

![Task structure](task.png)

### Write a task topic

To write a task, complete these steps:

1. Determine where you want to add task documentation to the Grafana Labs product documentation.
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

   For more information about front matter, refer to [Front matter]({{< relref "../front-matter" >}}).

1. Add the content to a copy of the [Task template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/task-template.md).

### Task topic examples

Refer to the following topics for task topic examples:

- [Create a service account](https://grafana.com/docs/grafana/latest/administration/service-accounts/#create-a-service-account-in-grafana)
- [Create an organization](https://grafana.com/docs/grafana/latest/administration/organization-management/#create-an-organization)
- [Create a dashboard and add a panel](https://grafana.com/docs/grafana/latest/dashboards/add-organize-panels/#create-a-dashboard-and-add-a-panel)

### Task template

When you are ready to write, make a copy of the [Task template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/task-template.md) and add your content.

### When to combine tasks into a single topic

In some cases, task topics are standalone and do not contain any other content. Other times, multiple task topics can be combined into a single Markdown file. By combining tasks into a single topic, the table of contents entities can be reduced in number, which reduces scrolling and clicking for users.

> **Note:** It is not a good idea to combine content in the same Markdown file at random. If you combine content incorrectly, you may inadvertently hide information from the user.

When combining multiple topics into one, follow these guidelines:

- When you document more than one approach to accomplishing the same user goal.

  In the [Assign RBAC roles](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/assign-rbac-roles/) topic, a user can use the user interface or provisioning to assign roles. There is no need to create two task topic files in this case.

- When tasks are likely to be completed around the same time.

  Users might find it useful to have all task documentation available on the same page if they are likely to complete a number of tasks simultaneously.

  In the [Data source management](https://grafana.com/docs/grafana/latest/administration/data-source-management/) topic, it is likely that an Admin user will enable permissions immediately after adding a data source.

- When you document CRUD operations.

  Create, read, update, and delete tasks can be combined into one topic. The [Manage organizations](https://grafana.com/docs/grafana/latest/administration/organization-management/) topic includes tasks such as viewing, creating, editing, and deleting organizations all under the umbrella topic title of **Manage**.

- When you document a user workflow.

  Combine tasks when the user should start at the beginning, complete the first task, and then complete the remaining tasks in sequence.

  In [Activate a Grafana Enterprise license from AWS Marketplace on EKS](https://grafana.com/docs/grafana/latest/administration/enterprise-licensing/activate-aws-marketplace-license/activate-license-on-eks/), the user is guided through all the tasks necessary to activate their license.

## References

A reference topic provides users with the information they might need to refer to during a task. An effective reference should provide a comprehensive listing of data, such as functions and parameters, error messages, and return codes. A reference is usually presented as a table, a bulleted list, or a sample script.

API information is also included in reference topics.

Because reference topics contain information the user needs to accomplish a task, they are often linked to task topics.

> **Note:** Do not include steps or conceptual information in reference topics.

### Reference structure

- **Topic title:** Reference topic titles should contain a qualifier and noun, for example, *Grafana CLI*. This helps readers distinguish between reference topics and tasks.
- **Introduction:** Provide an introduction that explains what to expect from this topic.
- **Body:** Tables or lists are often used to provide information in reference topics.

![Reference structure](reference.png)

### Write a reference topic

To write a reference, complete these steps:

1. Determine where you want to add reference documentation to the Grafana Labs product documentation.
1. Create a child directory within the parent directory that follows this naming convention:
   
   - Begin the directory name with a qualifier followed by an noun.
   - Use lowercase letters.
   - Add a hyphen between words.
  <br>
  <br>
   For example:
     - calculation-types
     - standard-field-definitions
<br>
<br>

1. Create an `index.md` file within the reference directory.
1. Add front matter to the `index` file.

   For more information about front matter, refer to [Front matter]({{< relref "../front-matter" >}}).

1. Add the content to a copy of the [Reference template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/reference-template.md).

### Reference topic examples

Refer to the following topics for a reference topic examples:

- [Calculation types](https://grafana.com/docs/grafana/latest/panels/calculation-types/)
- [Standard field definitions](https://grafana.com/docs/grafana/latest/panels/standard-field-definitions/)
- [Grafana CLI](https://grafana.com/docs/grafana/latest/administration/cli/)

### Reference template

When you are ready to write, make a copy of the [Reference template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/task-template.md) and add your content.
