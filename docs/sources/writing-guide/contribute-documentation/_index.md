---
title: "Contribute to documentation"
menuTitle: "Contribute to documenation"
description: "This section describes the different ways of contributing to documentation."
aliases: ["/docs/writers-toolkit/latest/writing-guidelines/contribute-documentation/"]
weight: 100
Keywords:
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

### Before you begin:

Create a GitHub account.

To request a change, complete the following steps:

1. From a topic on the documentation website, click **Request doc changes**.

    The Issue title auto-populates with the location of the file for which you are requesting a change.
2. Click **Submit new issue**.

## Edit a topic

If you want to recommend a small change, such as suggesting a correction to a topic, you can edit the topic directly in GitHub. You are not required to fork and clone the repo to use this approach.
Small changes might include:

* Adding steps to a task
* Adding clarifying language to a concept
* Providing an example

### Before you begin:
Create a GitHub account.

To edit a topic, complete the following steps:

1. From a topic on the documentation website, click **Edit this page** (pencil icon).
1. Enter your changes.
1. Change the branch name, if required.
   The branch name is auto-populated.

2. Click **Propose changes**.

    A PR is created and then goes through the PR review and approval workflow.

## Create a topic

PLACEHOLDER
decide whether your topic is a concept, task, or reference
add frontmatter
write a clear title
use the concept, task, and reference templates to help structure your content

## Review your changes

Prior to pushing your changes, you can view a local build of the documentation so that you can review your work.

To view a local build:

1. Install Docker.
1. Run Docker.
1. Navigate to the docs root directory.
1. Run make docs.
1. Open localhost:3002 to review your changes.

## Push changes and create a PR

When you are ready for other people to review your work, perform the following tasks:

1. Add your changes, which prepares your content for the next commit.
1. Commit your changes.
1. Push your changes to Github.
1. Create a PR in Github.

    The docs build system automatically conducts a series of tests to ensure that the content doesn't conflict with other content in the docs repository.

## PR review and approval workflow

When a PR is added to the repo with a label associated with docs, it will be reviewed by a member of the technical writing team.

Depending on the size of the PR and the priority of other work, the PR will either be immediately reviewed and merged (minor fixes typically follow this pattern) or the PR will be triaged and placed in the backlog of work or moved into further development.
