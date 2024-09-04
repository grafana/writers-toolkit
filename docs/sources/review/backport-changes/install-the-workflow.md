---
title: Install the backport workflow
date: 2024-09-03
description: How to install the backport workflow in a new repository.
keywords:
  - backporting
  - backport
  - GitHub
  - workflow
menuTitle: Install the workflow
review_date: 2024-09-03
---

# Install the backport workflow

Before you can backport changes, the repository must have the workflow installed.

## Before you begin

- Clone the repository.

  For more information, refer to [Create a repository](/docs/writers-toolkit/write/tooling-and-workflows/#create-a-local-repository).

## Steps

To install the backport workflow:

1. Check out a new branch for your changes.

   To check out a branch, refer to [Create a branch from the default remote branch](/docs/writers-toolkit/write/tooling-and-workflows/#create-a-branch-from-the-default-remote-branch).

1. Create the `.github/workflows/backport.yml` file.

   Use the following YAML as the contents of the file.
   You must update _`REPOSITORY`_ to be the name of your repository.

   ```yaml
   name: Backport PR creator
   on:
     pull_request_target:
       types:
         - closed
         - labeled

   jobs:
     main:
       if: github.repository == 'grafana/<REPOSITORY>'
       runs-on: ubuntu-latest
       steps:
         - name: Checkout actions
           uses: actions/checkout@v4
           with:
             repository: grafana/grafana-github-actions
             path: ./actions
         - name: Install actions
           run: npm install --production --prefix ./actions
         - name: Run backport
           uses: ./actions/backport
           with:
             token: ${{ secrets.GITHUB_TOKEN }}
             labelsToAdd: backport
             title: "[{{base}}] {{originalTitle}}"
   ```

1. Commit the file.

   For more information, refer to [Commit changes to your branch](/docs/writers-toolkit/write/tooling-and-workflows/#commit-changes-to-your-branch).

1. Open a pull request for your changes.

1. After a maintainer merges your pull request, you can backport changes using labels by following the instructions in [Backport changes](/docs/writers-toolkit/review/backport-changes/).
