---
title: Deprecate or remove content
menuTitle: Deprecate or remove
description: Learn about deprecating or removing content in your documentation.
weight: 800
aliases:
  - /docs/writers-toolkit/write/deprecate-remove
Keywords:
  - deprecate
  - remove
---

# Deprecate or remove content

**Deprecation** occurs when a feature or product is planned for removal in a future release. In the period between deprecation and removal, it will be in maintenance.

**Removal** is when the product or feature is removed and no longer supported.

The process for handling these scenarios varies according to how the product or feature is being phased out or maintained.

We should notify our users two minor releases in advance of any planned removal of a feature. For example, if we plan to remove a feature in Grafana v9.5, then we should begin communicating that information to users in Grafana v9.3.

To ensure that our customers are fully aware of the stages of deprecation for their products or features, we should:

- Inform customers of deprecation using notes in our What’s New and documentation and where necessary, in the UI.
- Inform the field by emailing [gtm-se-field-engineering@grafana.com](mailto:gtm-se-field-engineering@grafana.com) or by reaching out on Slack to # field-engineering-with-nomads
- Update these notes once removal has taken place.
- Announce the deprecation in our [community slack](https://grafana.slack.com/archives/C05675Y4F) and [forum](https://community.grafana.com/).
- Inform CX and Solution Engineering.

## Deprecation

Scenario: “We're planning on removing the feature or product in a future release.”

Example text:

```
{{% admonition type="caution" %}}
Starting with `<release>`, `<product or feature name>` is deprecated. It will be removed in a future release. *if the release is known, enter the release instead of “future”.
{{% /admonition %}}
```

### Removal

Scenario: “We've removed this feature.”

Example text:

```
{{% admonition type="warning" %}}
 `<Product or feature name>` is removed. It is no longer deployed, enhanced, or supported.
{{% /admonition %}}
```

## Docs deprecation

Scenario: There are situations in which we deprecate docs along with and/or independent of feature deprecation.

Example text:

```
{{% admonition type="caution" %}}
 Starting with release `<version #>`, the `<topic title>` documentation will no longer be published because `<provide rationale>`.  Link to docs, if they exist.
{{% /admonition %}}
```

Example: “As of Grafana 9.2, we are no longer publishing release notes, because they duplicate the content available in the What’s New document and the CHANGELOG."

## How to address deprecation

The following process describes how to address deprecation. This process applies when a feature is deprecated and associated docs are also deprecated and when docs are deprecated independent of feature deprecation.

You may need to add the above notes to both the documentation AND the UI.

1. Writers, Engineers, or Product Managers determine docs impacted by deprecation.

   If you have a writer assigned to your team, work with your writer. If you do not have an assigned writer, get in touch with [Fiona Peers Artiaga](mailto:fiona.artiaga@grafana.com).

1. If you do not have a writer assigned to your team, Engineers or Product Managers communicate with the Docs Squad and provide links to the impacted docs, the rationale for the deprecation, and timing. Post your message in the [#docs](https://raintank-corp.slack.com/archives/C5PG2JK8W) channel.

   The Docs Team triage process ensures that a writer is made aware of any pending docs deprecation.

1. In the PR that deprecates the docs, writers, engineers or Product Managers provide the rationale for deprecating the docs in the comments (using the docs deprecation language identified in the table above), and if there are alternative docs that a user can refer to, provide a link.

   Github automatically assigns the writer associated with those docs as a Reviewer of the PR.

1. The writer either drafts the PR or reviews and approves the PR and adds a notice of the deprecation (and links, if available) in the draft What’s New document.
