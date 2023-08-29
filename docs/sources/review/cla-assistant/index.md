---
title: About CLA assistant
description: Understand the CLA assistant and its use on Github Pull Requests.
---

# About CLA assistant

[CLA assistant](https://cla-assistant.io/) is used by Grafana Labs to ensure that all contributors to a GitHub Pull Request (PR) have signed the Grafana Labs Contributor License Agreement (CLA).

Grafana Labs does not operate the CLA assistant infrastructure, instead it is provided by SAP.
The CLA assistant project is open source and the source code can be found in the [cla-assistant repository](https://github.com/cla-assistant/cla-assistant).

The workflow runs automatically on all PRs.
During periods of high load on the CLA assistant infrastructure, the workflow can fail to report a status.
If it does not report back and all other Continuous Integration (CI) checks have completed, refer to [Waiting for status to be reported](#waiting-for-status-to-be-reported) to rerun the workflow.

{{% admonition type="note" %}}
During prolonged periods of high load on the CLA assistant infrastructure, even rerunning the workflow doesn't result in a successful status.
In such a circumstance, wait a few hours and try again later.
{{% /admonition %}}

## Waiting for status to be reported

<!-- vale Grafana.WordList = NO -->
<!-- "check" here is used as a noun not a verb. TODO: update the WordList to only lint "check" as a verb. -->

Occasionally, the CLA assistant workflow is not reported on an open PR, preventing the PR from being merged.
You can request the CLA assistant to run again by browsing to the check URL for the PR.

<!-- vale Grafana.WordList = YES -->

The format of the URL is `https://cla-assistant.io/check/grafana/<REPOSITORY>?pullRequest=<PR>`.

- _`REPOSITORY`_ is the GitHub repository name.
- _`PR`_ is the Pull Request number in the repository.

To trigger the CLA assistant workflow for the Pull Request `https://github.com/grafana/grafana/pull/1`, the URL is `https://cla-assistant.io/check/grafana/agent?pullRequest=1`.

{{% admonition type="note" %}}
The URL redirects your browser back to https://github.com without any feedback suggesting the workflow has been triggered.
After a few minutes, check the PR to see if the status has been reported.
{{% /admonition %}}
