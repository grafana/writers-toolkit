---
title: Deprecate or remove content
menuTitle: Deprecate or remove
description: Learn about deprecating or removing content in your documentation.
weight: 900
aliases:
  - /docs/writers-toolkit/write/deprecate-remove
Keywords:
  - deprecate
  - remove
---

# Deprecate or remove content

<!-- vale Google.Passive = NO -->

**Deprecation** occurs when a feature or product is planned for removal in a future release.
In the period between deprecation and removal, it's in maintenance.

**Removal** is when the product or feature is removed and no longer supported.

The process for handling these scenarios varies according to how the product or feature is being phased out or maintained.

<!-- vale Google.Passive = YES -->

You should notify users two minor releases in advance of any planned removal of a feature.
For example, if you plan to remove a feature in Grafana v9.5, then you should begin communicating that information to users in Grafana v9.3.

To ensure that users are fully aware of the stages of deprecation for their products or features, you should:

- Inform customers of deprecation using notes in the product's What’s New, documentation, and where necessary, in the UI.
- Inform the field by emailing [`gtm-se-field-engineering@grafana.com`](mailto:gtm-se-field-engineering@grafana.com) or by reaching out on Slack in the `#field-engineering-with-nomads` channel
- Update these notes once removal has taken place.
- Announce the deprecation in the Grafana [community slack](https://grafana.slack.com/archives/C05675Y4F) and [forum](https://community.grafana.com/).
- Inform CX and Solution Engineering.

## Deprecation

Scenario: "You're planning on removing the feature or product in a future release."

Example text:

```markdown
{{</* admonition type="caution" */>}}
Starting with <RELEASE>, <PRODUCT OR FEATURE NAME> is deprecated.
It will be removed in a future release.
{{</* /admonition */>}}
```

If you know the release, enter the release instead of "future".

### Removal

Scenario: "You've removed a feature."

Example text:

```markdown
{{</* admonition type="warning" */>}}
<PRODUCT OR FEATURE NAME> is removed.
It is no longer deployed, enhanced, or supported.
{{</* /admonition */>}}
```

## Documentation deprecation

Scenario: There are situations in which you may deprecate documentation along with and/or independent of feature deprecation.

Example text:

<!-- vale Google.Will = NO -->
<!-- vale Google.Passive = NO -->

```markdown
{{</* admonition type="caution" */>}}
Starting with release <VERSION>`, the <TOPIC TITLE> documentation will no longer be published because <RATIONALE>.
Link to documentation, if it exists.
{{</* /admonition */>}}
```

Concrete example:

{{< admonition type="caution" >}}
As of Grafana 9.2, release notes will no longer be published, because they duplicate the content available in the What’s New document and the CHANGELOG."
{{< /admonition >}}

<!-- vale Google.Passive = YES -->
<!-- vale Google.Will = YES -->

## How to address deprecation

The following process describes how to address deprecation.
This process applies when you deprecate a feature and associated documentation is also deprecated and also when you deprecate documentation independently of feature deprecation.

You may need to add the preceding notes to both the documentation _and_ the UI.

1. Writers, Engineers, or Product Managers determine documentation impacted by deprecation.
   <!-- vale Grafana.Spelling = NO -->

   If you have a writer assigned to your team, work with your writer.
   If you don't have an assigned writer, contact [Fiona Peers Artiaga](mailto:fiona.artiaga@grafana.com).

   <!-- vale Grafana.Spelling = YES -->

1. If you don't have a writer assigned to your team, Engineers or Product Managers communicate with the Docs Squad and provide links to the impacted documentation, the rationale for the deprecation, and timing.
   Post your message in the [#docs](https://raintank-corp.slack.com/archives/C5PG2JK8W) channel.

   The Docs Team triage process makes a writer aware of any pending documentation deprecation.

1. In the PR that deprecates the documentation, writers, engineers, or Product Managers provide the rationale for deprecating the documentation in the comments using the documentation deprecation language identified in the preceding table, and if there is alternative documentation that a user can refer to, provide a link.

   GitHub automatically assigns the writer associated with those documentation as a reviewer of the PR.

<!-- vale Google.Parens = NO -->

1. The writer either drafts the PR or reviews and approves the PR and adds a notice of the deprecation (including links, if available) in the draft What’s New document.

<!-- vale Google.Parens = YES -->
