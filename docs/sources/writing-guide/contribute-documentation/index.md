---
title: Contribute to documentation
menuTitle: Contribute to documentation
description: This section describes the different ways of contributing to documentation.
aliases:
  - /docs/writers-toolkit/latest/writing-guide/contribute-documentation/
weight: 200
keywords:
  - contribute
  - request change
  - edit documentation
  - review documentation
---

# Contribute to documentation

How can I contribute to Grafana Labs documentation?

There are a number of different ways to contribute to documentation at Grafana Labs. Choose the one that applies best and learn how you can engage with the documentation team and help us by contributing your ideas.

## Request a change

Request a change when you want to make a suggestion or provide feedback about a topic, but don't want to provide an edit that generates a pull request.

Requesting a change gives you the freedom to express your ideas without committing language. Your suggestion can reflect a small change to wording or can reflect larger, more substantive changes.

GitHub captures your request as an Issue logged against the repository.

> **Note:** You can only request a change against the latest release of documentation.

### Before you begin

- Create a [GitHub](https://www.github.com) account.

**To request a change:**

1. From a topic on the documentation website, click **Request doc changes**.

   The **Issue** title auto-populates with the location of the file for which you are requesting a change.

1. Click **Submit new issue**.

The Doc Squad determines the priority and scope of the change, and schedules the change to be made.

## Edit a topic

If you want to recommend a small change, such as suggesting a correction to a topic, you can edit the topic directly in GitHub. You are not required to fork and clone the repo to use this approach.

Small changes might include:

- Adding steps to a task
- Adding clarifying language to a concept
- Providing an example

> **Note:** You can only edit a topic that is part of the latest release of documentation.

### Before you begin

- Create a [GitHub](https://www.github.com) account.

**To edit a topic:**

1. From a topic on the documentation website, click **Edit this page** (pencil icon).
1. Enter your changes.
1. Change the branch name, if required.
   The branch name is auto-populated.

1. Click **Propose changes**.

   A PR is created and then goes through the PR review and approval workflow.

## Develop a new topic

If you want to develop a new topic from scratch, you can create a documentation plan and collaborate with a member of the Docs team. According to the book [_Docs for Developers_](https://docsfordevelopers.com/), a documentation plan is a _flexible outline_ for anticipating where the writing process will lead you.

Your documentation plan helps you to:

- Identify existing information gaps and explain how you will fill them
- Get feedback from users and stakeholders before the writing process begins
- Consider different approaches you might take, and decide between them

> **Note:** You can only create new topics that are part of the latest release of documentation.

**To develop a new topic:**

1. Create a new document in Google Docs, Google Sheets, or a similar tool. For a small project, you may use an email message.
1. Copy and paste content from the [Document Planning Template]({{< relref "../../static/templates/document-planning-template.md) into your document. 
1. Fill in the requested information. 
1. Attach an outline of your proposed document. 
1. Schedule a review meeting with a member of the Docs team.

   If you don't know which writer will be reviewing your documentation plan, reach out to [docs@grafana.com](mailto:docs@grafana.com). 

Your writer on the Docs team will contact you with next steps. Our aim is to help you to be successful in improving the experience of our users and developers.

## Testing

It is a best practice to have someone else test any task you have written. If another user can successfully complete the task using _only_ the steps you have written, not guessing or using their inherent knowledge, then your task has passed the test. However, it is very common to find you have skipped steps because _you_ are very familiar with the topic you are explaining.

New users or members of other teams are very helpful for these tests.

## Review your changes

Prior to pushing your changes to Github, you can view a local build of the documentation so that you can review your work. For more information on using Git, refer to [Using Git]({{< relref "../tooling-and-workflows/#using-git" >}}).

**To view a local build:**

1. Install either Podman or Docker to manage containers on your system.

    <!-- vale Grafana.Colons = NO -->

   > **Note:** Podman has the upside that containers can either be run as root or in rootless mode.

    <!-- vale Grafana.Colons = YES -->

   - To install Podman, refer to [Podman Installation Instructions](https://podman.io/getting-started/installation).
   - To install Docker, refer to [Docker Engine installation overview](https://docs.docker.com/engine/install/).

1. If you are using Docker, start the Docker daemon if it is not already running.
1. Navigate to the project "docs" directory, this is typically `docs`.
1. Run `make docs`.
1. Browse to `localhost:3002/docs/` to review your changes.

## Push changes and create a PR

When you are ready for other people to review your work, perform the following tasks:

1. Add your changes, which prepares your content for the next commit.
1. Commit your changes.
1. Push your changes to Github.
1. Create a PR in Github.
1. Add the `type/docs` label.

   The docs build system automatically conducts a series of tests to ensure that the content doesn't conflict with other content in the docs repository.

## PR review and approval workflow

When you add a PR to the repo and assign the `type/docs` label, it will be reviewed by a member of the Docs Squad.

Depending on the size of the PR and the priority of other work, the PR will either be immediately reviewed and merged (minor fixes typically follow this pattern) or the PR will be triaged and placed in the backlog of work or moved into further development.

## Contributing across versions

When you edit the `main` branch of a project, it affects the content in the `next` directory of the website.
To edit a previous version, or `latest` (the most recent release), you must backport the changes into the long-lived version branches in the project repository.

To backport a change, use the `backport <vMAJOR.MINOR.x>` labels on the GitHub pull request.

`grafanabot` automatically creates a backport pull request after the original pull request is merged, if the merge commit can be cherry-picked without a conflict.
If this process fails due to a merge conflict, `grafanabot` posts a comment explaining how to manually backport the change.