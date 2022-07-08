---
title: "About the type/docs label workflow"
menuTitle: "About the `type/docs` label workflow"
description: "How `type/docs` label workflow automation works."
aliases: ["/about-the-type-docs-label-workflow/"]
keywords:
---

# About the type/docs label workflow

There are many Grafana code repos, and only one docs team.
When a docs-related issue or pull request comes up in one of the repos, there needs to be an easy, automated way to notify the docs team.

This workflow makes all issues with the `type/docs` label visible in the [Technical documentation project](https://github.com/orgs/grafana/projects/69) which is an organization wide GitHub Project.
Whenever someone makes an issue or PR with that label, the resource ends up in the **Needs triage** column of that project.

For each repository that the Docs Squad wants to monitor, this workflow needs to be in a file in the repository's `.github/workflows` directory:

```yaml
name: Run commands when issues or PRs are labeled
on:
  issues:
    types: [labeled]
  pull_request:
    types: [labeled]
jobs:
  main:
    if: ${{ github.event.pull_request == false || github.event.pull_request.head.repo.full_name == github.repository }}
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Actions
        uses: actions/checkout@v2
        with:
          repository: "grafana/grafana-github-actions"
          path: ./actions
          ref: main
      - name: Install Actions
        run: npm install --production --prefix ./actions
      - name: Run Commands
        uses: ./actions/commands
        with:
          token: ${{secrets.GH_BOT_ACCESS_TOKEN}}
          configPath: issue_and_pr_commands
```

> **Warning:** The workflow uses the `GH_BOT_ACCESS_TOKEN` secret that must not be exposed to pull requests from forked repositories.
> The `if: ${{ github.event.pull_request == false || github.event.pull_request.head.repo.full_name == github.repository }}` protects against this.

The commands run by the workflow are defined in the [`grafana/grafana-github-actions`](https://github.com/grafana/grafana-github-actions/) repository.
