---
title: Links
description: Understand how to link between pages.
weight: 600
aliases:
  - /docs/writers-toolkit/write/links/
  - /docs/writers-toolkit/writing-guide/references/
  - /docs/writers-toolkit/write/references/
keywords:
  - Hugo
  - references
  - relref
  - ref
---

# Links

Choose your link type based on the applicable scenario:

- [Linking from source content that's used (or mounted) in multiple projects](#source-content-is-reused-in-multiple-projects)
- [Linking to grafana.com pages](#destination-page-is-on-grafanacom)
- [Linking to external pages](#destination-page-is-external)
- [Linking to page headings](#anchors)

Other link types exist in our documentation, but you shouldn't use them:

- [Hugo `relref` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#relref)

## Source content is reused in multiple projects

Use the `docs/reference` shortcode.

The source is reused as described in [Reuse directories of content with Hugo mounts](https://grafana.com/docs/writers-toolkit/write/reuse-content/reuse-directories/).
For more information and examples, refer to [`docs/reference` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#docsreference).

## Destination page is on grafana.com

Use a fully qualified URL with version substitution.
For example, `https://grafana.com/docs/grafana/<GRAFANA VERSION>/`

Version substitution is necessary for links to resolve to the correct version of documentation.
If you want to link to a specific version of documentation, you do not need to include the substitution syntax.

### About version substitution

Version substitution enables the use of absolute paths that resolve correctly, irrespective of version.
It uses special syntax using angle bracket delimiters like `<GRAFANA VERSION>`.

As a convention, use the name of the target project all upper-case.
For example, `grafana` becomes `GRAFANA`, `grafana-cloud` becomes `GRAFANA CLOUD`.

The special syntax `<SOMETHING VERSION>` is substituted by the version that is inferred from the page's URL.

You can override version inference by including additional metadata in the front matter of the file.
To override the value of `<GRAFANA VERSION>`, set the `GRAFANA VERSION` parameter in the page's front matter.
For example, to set the version to `next` irrespective of the source content version, add the following to the front matter: `GRAFANA VERSION: next`.

## Destination page is external

Use the fully qualified URL.
For example, `https://github.com`.

## Anchors

In a reference, you can optionally include an anchor to a heading in the referenced page.
Specify and anchor at the end of the reference `#` followed by the normalized heading.

Hugo normalizes headings to make anchors.
To convert a heading to an anchor, Hugo makes the following changes:

1. Convert to lower case.
1. Remove any period characters (`.`).
1. Replace any character that's not a lower cased letter, a number, or an underscore (`_`) with dashes (`-`).
1. Trim any preceding or proceeding dashes (`-`).
