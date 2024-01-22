---
title: CLA assistant
description: Understand the CLA assistant and its use on GitHub pull requests.
---

# CLA assistant

[CLA assistant](https://cla-assistant.io/) is used by Grafana Labs to ensure that all contributors to a GitHub pull request (PR) have signed the Grafana Labs Contributor License Agreement (CLA).

Grafana Labs doesn't operate the CLA assistant infrastructure, instead it's provided by SAP.
The CLA assistant project is open source and you can find the source code in the [cla-assistant repository](https://github.com/cla-assistant/cla-assistant).

The workflow runs automatically on all PRs.
During periods of high load on the CLA assistant infrastructure, the workflow can fail to report a status.
If it doesn't report back and all other continuous integration (CI) checks have completed, follow the directions in [Waiting for status to be reported](#waiting-for-status-to-be-reported) to rerun the workflow.

{{% admonition type="note" %}}
During prolonged periods of high load on the CLA assistant infrastructure, even rerunning the workflow might not result in a successful status.

In this case, wait a few hours and try again later.
{{% /admonition %}}

## Waiting for status to be reported

<!-- vale Grafana.WordList = NO -->
<!-- "check" here is used as a noun not a verb. TODO: update the WordList to only lint "check" as a verb. -->

Occasionally, the CLA assistant workflow isn't reported on an open PR, preventing the PR from being merged.
You can request that the CLA assistant run again by browsing to the check URL for the PR.

<!-- vale Grafana.WordList = YES -->

The format of the URL is `https://cla-assistant.io/check/grafana/<REPOSITORY>?pullRequest=<PR>`.

- _`REPOSITORY`_ is the GitHub repository name.
- _`PR`_ is the pull request number in the repository.

For example, to trigger the CLA assistant workflow for the pull request `https://github.com/grafana/grafana/pull/1`, the URL is `https://cla-assistant.io/check/grafana/grafana?pullRequest=1`.

{{% admonition type="note" %}}
The URL redirects your browser back to https://github.com without any feedback suggesting the workflow has run.
{{% /admonition %}}
