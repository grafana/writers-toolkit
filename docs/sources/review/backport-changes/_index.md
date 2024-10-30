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

This guidance is general. Refer to [When and what to backport in `grafana/grafana`](#when-and-what-to-backport-in-grafanagrafana) for specific guidance about `grafana/grafana`.

<!-- prettier-ignore-start -->

| Change type                                       | Versions           |
|---------------------------------------------------|--------------------|
| Typo                                              | "Latest" version   |
| Copy edits                                        | "Latest" version   |
| Architecture change                               | "Latest" version   |
| Incorrect information                             | Supported versions |
| New content for version in upcoming release       | None               |
| New content for version after upcoming release    | None               |

<!-- prettier-ignore-end -->

## When and what to backport in `grafana/grafana`

This section outlines backport guidance specifically for the `grafana/grafana` repository.

The [later table](#guidance) outlines:

- Whether or not you must backport a change
- To which versions

One thing that can change that determination, however, is whether or not a version branch has been cut.

### Before the new version branch is cut

This is the time period between the date the "latest version was released and the upcoming version branch is cut. For example:

- Grafana v11.2 released August 27, 2024
- Grafana v11.3 version branch cut October 8, 2024

August 24 - October 8 is the period before the new version branch is cut.

This is the most common scenario.

### After the new version branch is cut but _before_ it's released

This is the time period covering approximately the last two weeks before a release. During this time, content that's intended for the upcoming release version needs to be backported. For example:

- Grafana v11.3 version branch cut October 8, 2024
- Grafana v11.3 released October 22, 2024

From October 8 - 22, the new version branch is cut but hasn't yet been released.

### Guidance

<!-- prettier-ignore-start -->

| Change type                                       | Before version branch is cut       | After version branch is cut                           |
|---------------------------------------------------|------------------------------------|-------------------------------------------------------|
| Typo                                              | "Latest" version                   | "Latest" version + upcoming version                   |
| Copy edits                                        | "Latest" version                   | "Latest" version + upcoming version                   |
| Architecture change                               | "Latest" version                   | "Latest" version + upcoming version                   |
| Incorrect information                             | [Supported versions](https://grafana.com/docs/grafana/latest/upgrade-guide/when-to-upgrade/#what-to-know-about-version-support) | [Supported versions](https://grafana.com/docs/grafana/latest/upgrade-guide/when-to-upgrade/#what-to-know-about-version-support) + upcoming version |
| New content for version in upcoming release       | None                               | Upcoming version                                      |
| New content for version after upcoming release    | None                               | None                                                  |

<!-- prettier-ignore-end -->

If you want to backport more than the guidance, that's at your discretion. This table outlines minimum standards.

## Backport tutorials

Depending on the location from which the website mounts tutorial content, you might need to backport changes to them.

For `grafana/grafana`, the website mounts tutorial content from the `next` version of the documentation, so you don't need to backport changes.

For other repositories, backport to the branch that corresponds with the version of the documentation on the website.
