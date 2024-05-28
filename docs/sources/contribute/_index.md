---
aliases:
  - /docs/writers-toolkit/contribute/
  - /docs/writers-toolkit/contribute-documentation/
  - /docs/writers-toolkit/writing-guide/contribute-documentation/
review_date: 2024-05-16
description: Learn how you can contribute to Grafana Labs documentation.
keywords:
  - contribute
  - request change
  - edit documentation
  - review documentation
menuTitle: Contribute
title: Contribute to documentation
weight: 200
---

# Contribute to documentation

How can you contribute to Grafana Labs documentation?

There are a number of different ways to contribute to documentation at Grafana Labs.
Choose the one that applies best and learn how you can engage with the documentation team and help by contributing your ideas.

## Report a problem

Report a problem when you want to make a suggestion or provide feedback about a topic, but don't want to provide an edit that generates a pull request.

Reporting a problem gives you the freedom to express your ideas without committing language.
Your suggestion can reflect a small change to wording or can reflect larger, more substantive changes.

You make your request with an email to [`docs@grafana.com`](mailto:docs@grafana.com).

To report a problem:

1. From a topic on the documentation website, click **Report a problem**.
   Your mail client opens.
1. Provide your feedback in the email body.
1. Include a link to your current page in the email.
1. Send your email.

The Grafana Labs documentation team checks the email inbox regularly and responds to emails in a timely fashion.

## Suggest an edit

If you want to recommend a small change, such as suggesting a correction to a topic, you can edit the topic directly in GitHub.

Small changes might include:

- Adding steps to a task
- Adding clarifying language to a concept
- Providing an example

### Before you begin

- Create a [GitHub](https://www.github.com) account.
- Find the source repository.
  To find the source repository, refer to [Find the source repository](#find-the-source-repository).

### Find the source repository

Technical documentation published from public projects have a **Suggest an edit** link with a pencil icon.
Click this link to directly edit the page in GitHub.

<!-- vale Grafana.Timeless = NO -->

{{< admonition type="warning" >}}
Because development happens in the `main` branch on GitHub which generally corresponds to the next version of documentation.

The latest version of documentation is typically published from a different _version_ branch, and the **Suggest an edit** link can result in a 404 error from GitHub.

In that case, you can use the GitHub code navigation to try and find the new location or reach out to the Grafana Labs documentation team for support.
{{< /admonition >}}

<!-- vale Grafana.Timeless = YES -->

If pages don't have a **Suggest an edit** link, the documentation isn't open source.
Only Grafana Labs employees can update closed source documentation.

For example, [Grafana Cloud](https://grafana.com/docs/grafana-cloud/) is in the [website repository](https://github.com/grafana/website).

{{< admonition type="note" >}}
The website repository is private and only accessible to Grafana Labs employees.
{{< /admonition >}}

Some Grafana Cloud content comes from other projects, listed below:

- `/docs/grafana-cloud/alerting`: https://github.com/grafana/grafana/tree/main/docs/sources/alerting
- `/docs/grafana-cloud/api-reference/http-api`: https://github.com/grafana/grafana/tree/main/docs/sources/developers/http_api
- `/docs/grafana-cloud/incident`: https://github.com/grafana/incident/tree/main/docs/sources
- `/docs/grafana-cloud/oncall`: https://github.com/grafana/oncall/tree/main/docs/sources
- `/docs/grafana-cloud/monitor-infrastructure/integrations/integration-reference/`:
  Automatically generated and published from https://github.com/grafana/cloud-onboarding (Grafana Labs only).

  Integration documentation is generated from the metadata in https://github.com/grafana/cloud-onboarding/tree/main/integrations and the template file https://github.com/grafana/cloud-onboarding/blob/main/pkg/integrations-api/integrations/docs/templates/layout.tmpl.

### To edit a topic

1. From a topic on the documentation website, click **Suggest an edit** (pencil icon).
1. Enter your changes.
1. Change the branch name, if required.
   The branch name is auto-populated.

1. Click **Propose changes**.

   GitHub creates a pull request which then goes through the review and approval workflow.

<!-- vale Grafana.Timeless = NO -->
<!-- vale Grafana.GoogleWill = NO -->

## Develop a new topic

If you want to develop a new topic from scratch, you can create a documentation plan and collaborate with a member of the technical writing team.
According to the book [_Docs for Developers_](https://docsfordevelopers.com/), a documentation plan is a _flexible outline_ for anticipating where the writing process will lead you.

<!-- vale Grafana.Timeless = YES -->
<!-- vale Grafana.GoogleWill = YES -->

Your documentation plan helps you to:

- Identify information gaps, and explain how to fill them.
- Get feedback from users and stakeholders before the writing process begins.
- Consider different approaches you might take, and decide on one of them.

## Testing

It's a best practice to have someone else test any task you have written.
If another user can successfully complete the task using _only_ the steps you have written, not guessing or using their inherent knowledge, then your task has passed the test.
However, it's very common to find you have skipped steps because _you_ are very familiar with the topic you are explaining.

New users or members of other teams are very helpful for these tests.

## Review your changes

Prior to pushing your changes to GitHub, you can view a local build of the documentation so that you can review your work.
For more information on using Git, refer to [Use Git](https://grafana.com/docs/writers-toolkit/write/tooling-and-workflows/#use-git).

To view a local build, refer to [Test documentation changes](https://grafana.com/docs/writers-toolkit/review/test-documentation-changes/)

## Push changes and create a pull request

When you are ready for other people to review your work, perform the following tasks:

1. Add your changes, which prepares your content for the next commit.
1. Commit your changes.
1. Push your changes to GitHub.
1. Create a pull request in GitHub.
1. When writing the description for your pull request, use [GitHub keywords](https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/using-keywords-in-issues-and-pull-requests#linking-a-pull-request-to-an-issue), for example "Fixes #1234", to link your pull request to the issue and take advantage of GitHub automation for status updates and closing resolved issues.
1. Add the `type/docs` label, so the Grafana Labs documentation team can track the issue.

   The docs build system automatically conducts a series of tests to ensure that the content doesn't conflict with other content in the docs repository.

## Pull request review and approval workflow

When you add a pull request to a repository and assign the `type/docs` label, it's added to a queue that's regularly reviewed by a member of the Grafana Labs documentation team.

The Grafana Labs documentation team aims to review all PRs in a timely fashion.

## Contribute across versions

When you edit the `main` branch of a project, it affects the content in the `next` directory of the website.
To edit a previous version, or `latest` (the most recent release), you must backport the changes into the long-lived version branches in the project repository.

To backport a change, use the `backport <BRANCH>` labels on the GitHub pull request.
For more information, refer to [Backport changes](https://grafana.com/docs/writers-toolkit/review/backport-changes/).

Grot, the Grafana bot, automatically creates a backport pull request if the merge commit can be cherry-picked without a conflict.
If this process fails due to a merge conflict, Grot posts a comment explaining how to manually backport the change.
