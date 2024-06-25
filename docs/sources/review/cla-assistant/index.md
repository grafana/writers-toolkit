---
date: "2023-08-30T11:24:32+01:00"
description:
  Understand the CLA assistant and its use on GitHub pull requests in Grafana
  repositories.
menuTitle: CLA assistant
review_date: "2024-05-23"
title: Contributor License Agreement (CLA) assistant
---

# Contributor License Agreement (CLA) assistant

Grafana Labs uses [CLA assistant](https://cla-assistant.io/) to ensure that all contributors to a GitHub pull request have signed the [Grafana Labs Contributor License Agreement (CLA)](https://grafana.com/docs/grafana/latest/developers/cla/).

Grafana Labs doesn't operate the CLA assistant infrastructure.
The CLA assistant project is open source, and you can find the source code in the [`cla-assistant` repository](https://github.com/cla-assistant/cla-assistant).

The workflow runs automatically on all PRs.
During periods of high load on the CLA assistant infrastructure, the workflow can fail to report a status.
If it doesn't report back and all other continuous integration (CI) checks have completed, follow the directions in [Waiting for status to be reported](#waiting-for-status-to-be-reported) to rerun the workflow.

{{< admonition type="note" >}}
During prolonged periods of high load on the CLA assistant infrastructure, even rerunning the workflow might not result in a successful status.

In this case, wait a few hours and try again later.
{{< /admonition >}}

<!-- vale Grafana.Gerunds = NO -->
<!-- vale Grafana.GooglePassive = NO -->
<!-- This heading matches the text that is displayed when the check hasn't run -->

## Waiting for status to be reported

<!-- vale Grafana.Gerunds = YES -->
<!-- vale Grafana.GooglePassive = YES -->

Occasionally, the CLA assistant workflow isn't reported on an open pull request, preventing you from merging the pull request.
You can request that the CLA assistant run again by browsing to the check URL for the pull request.

The format of the URL is `https://cla-assistant.io/check/grafana/<REPOSITORY>?pullRequest=<PULL REQUEST>`.

- _`<REPOSITORY>`_ is the GitHub repository name.
- _`<PULL REQUEST>`_ is the pull request number in the repository.

For example, to trigger the CLA assistant workflow for the pull request `https://github.com/grafana/grafana/pull/1`, the URL is `https://cla-assistant.io/check/grafana/grafana?pullRequest=1`.

{{< admonition type="note" >}}
The URL redirects your browser back to https://github.com without any feedback suggesting the workflow has run.
{{< /admonition >}}
