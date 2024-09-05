---
aliases:
  - /docs/writers-toolkit/review/backport-changes/
  - /docs/writers-toolkit/review/backporting/
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/backporting/
date: "2024-02-16T14:04:14+00:00"
description: Understand how and when to backport changes to Grafana repositories.
keywords:
  - backporting
  - backport
menuTitle: Backport changes
review_date: "2024-05-22"
title: Backport changes
weight: 500
---

# Backport changes

Projects with versioned documentation typically maintain each version with the code in multiple long-lived branches.
The `main` branch has the most recent code and documentation.

Project releases typically use long-lived branches that include the major and minor versions of the release.
For example, in the `grafana/grafana` repository, the `v9.0.x` long-lived release branch contains code and documentation for all patched versions of the major version `9` and minor version `0` of Grafana.

Backporting takes a change from the `main` branch and ports it back to another long-lived release branch.

Every documentation pull request to a project with versioned documentation requires a decision about whether or not to backport it.

You should backport to all supported versions of the project affected by the pull request changes.

If you are unsure, ask for advice from a maintainer on the pull request.

## Before you begin

- To backport changes, the repository must have the backport workflow installed.
  To install the backport workflow, refer to [Install the backport workflow](/docs/writers-toolkit/review/backport-changes/install-the-workflow/).

## Backport a change

To backport a change, add the appropriate `backport <BRANCH>` label.
You can backport to more than one branch by using multiple labels.

For merged pull requests, Grot, the Grafana bot, creates a follow-up pull request for each of the `backport <BRANCH>` labels.
You can add the `backport <BRANCH>` labels either before or after you merge the pull request.
If Grot can't automatically backport the changes, it comments on the original pull request with instructions about how to backport the change manually.

In repositories such as `grafana/grafana`, engineers sometimes create a branch for a release well before the release has shipped.
If you intend to publish content against an imminent release, check for a backport label for the upcoming version before merging the pull request.
Apply the label if it exists to ensure the content is automatically backported to the upcoming version's documentation.

If you decide to _not_ backport a change, you don't need to add any label.

## When to backport

To decide whether to backport a pull request, use the following decision tree:

<!-- vale Grafana.Timeless = NO -->

<script type="module">
  import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
  mermaid.initialize({ startOnLoad: true });
</script>
<div class="mermaid">
  flowchart TD
  D1{Is the change documenting a new feature?}
  D2{Is the feature going to be released in a future version of the project that doesn't yet have a release branch?}
  D3{Is the change a fix for a typo?}
  D4{Is the change a documentation refactoring?}
  T1[Do nothing.]
  T2[Add a backport label for each affected release version.]
  T3[Update this flowchart as the decision making process is incomplete.]
  T4[Backport the change to release branch for the next version of the project.]
  D1 -- Yes --> D2
  D1 -- No -->  D3
  D2 -- Yes --> T1
  D2 -- No -->  T4
  D3 -- Yes --> T2
  D3 -- No -->  D4
  D4 -- Yes --> T4
  D4 -- No -->  T3
</div>

<!-- vale Grafana.Timeless = YES -->

## Backport tutorials

Depending on the location from which the website mounts tutorial content, you might need to backport changes to them.

For `grafana/grafana`, the website mounts tutorial content from the `next` version of the documentation, so you don't need to backport changes.

For other repositories, backport to the branch that corresponds with the version of the documentation on the website.
