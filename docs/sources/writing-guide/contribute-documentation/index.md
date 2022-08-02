---
title: Contribute to documentation
menuTitle: Contribute to documentation
description: This section describes the different ways of contributing to documentation.
aliases:
  - /docs/writers-toolkit/latest/writing-guidelines/contribute-documentation/
weight: 200
keywords:
  - contribute
  - request change
  - edit documentation
  - review documentation
---

# Contribute to documentation

How can I contribute to Grafana's documentation?

There are a number of different ways to contribute to documentation at Grafana Labs. Choose the one that applies best and learn how you can engage with the documentation team and help us by contributing your ideas.

## Request a change

Request a change when you want to make a suggestion or provide feedback about a topic, but don't want to provide an edit that generates a pull request.

Requesting a change gives you the freedom to express your ideas without committing language. Your suggestion can reflect a small change to wording or can reflect larger, more substantive changes.

GitHub captures your request as an Issue logged against the repository.

> **Note:** You can only request a change against the lastest release of documenation.

**Before you begin:**

- Create a [GitHub](https://www.github.com) account.

**To request a change:**

1. From a topic on the documentation website, click **Request doc changes**.

    The **Issue** title auto-populates with the location of the file for which you are requesting a change.

2. Click **Submit new issue**.

The Doc Squad determines the priority and scope of the change, and schedules the change to be made.

## Edit a topic

If you want to recommend a small change, such as suggesting a correction to a topic, you can edit the topic directly in GitHub. You are not required to fork and clone the repo to use this approach.

Small changes might include:

- Adding steps to a task
- Adding clarifying language to a concept
- Providing an example

> **Note:** You can only edit a topic that is part of the lastest release of documenation.

**Before you begin:**

- Create a [GitHub](https://www.github.com) account.

**To edit a topic:**

1. From a topic on the documentation website, click **Edit this page** (pencil icon).
2. Enter your changes.
3. Change the branch name, if required.
   The branch name is auto-populated.

4. Click **Propose changes**.

   A PR is created and then goes through the PR review and approval workflow.

## Testing

It is a best practice to have someone else test any task you have written. If another user can successfully complete the task using _only_ the steps you have written, not guessing or using their inherent knowledge, then your task has passed the test. However, it is very common to find you have skipped steps because _you_ are very familiar with the topic you are explaining.

New users or members of other teams are very helpful for these tests.

## Review your changes

Prior to pushing your changes to Github, you can view a local build of the documentation so that you can review your work. For more information on using Git, refer to [Using Git]({{< relref "../tooling-and-workflows/#using-git" >}}).

**To view a local build:**

1. Install Docker.
1. Run Docker.
1. Navigate to the `docs` root directory.
1. Run `make docs`.
1. Open `localhost:3002` to review your changes.

## Push changes and create a PR

When you are ready for other people to review your work, perform the following tasks:

1. Add your changes, which prepares your content for the next commit.
1. Commit your changes.
1. Push your changes to Github.
1. Create a PR in Github.
1. Add the `type/docs` label.

   For more information about the `type/docs` workflow, refer to [type/docs label workflow]({{< relref "../tooling-and-workflows/type-docs-label-workflow/" >}}).

   The docs build system automatically conducts a series of tests to ensure that the content doesn't conflict with other content in the docs repository.

## PR review and approval workflow

When you add a PR to the repo and assign the `type/docs` label, it will be reviewed by a member of the Docs Squad.

Depending on the size of the PR and the priority of other work, the PR will either be immediately reviewed and merged (minor fixes typically follow this pattern) or the PR will be triaged and placed in the backlog of work or moved into further development.
