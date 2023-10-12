---
title: Links
description: Understand how to link between pages.
weight: 600
aliases:
  - /docs/writers-toolkit/write/links/
  - /docs/writers-toolkit/write/references/
  - /docs/writers-toolkit/writing-guide/references/
keywords:
  - Hugo
  - link
  - linking
  - links
  - ref
  - references
  - relref
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

Use a fully qualified URL with version substitution syntax (if needed).
Version substitution is necessary for fully qualified URLs to link to the correct version of documentation.
Usually, this is the current version of documentation.

### Examples

**Link to Grafana documentation**:

Start with `https://grafana.com/docs/grafana/<GRAFANA VERSION>/`, and add the rest of the URL path.
For example, to link to the [Developers](https://grafana.com/docs/grafana/latest/developers) page with version substitution,
use `https://grafana.com/docs/grafana/<GRAFANA VERSION>/developers/`.

- If you are from other documentation, `<GRAFANA VERSION>` is substituted with the value of `GRAFANA VERSION` set in the page's front matter.

  Ensure that you set the appropriate version in the root `_index.md` file for your documentation.
  The following YAML, merged with the existing front matter in the root `_index.md` file sets `GRAFANA VERSION` to be `latest` for that page and all child pages.

  ```yaml
  cascade:
    GRAFANA VERSION: latest
  ```

**Link to Grafana Cloud documentation**:

Grafana Cloud documentation is not versioned so no version substitution syntax is needed.
Use the fully qualified URL.
For example, to link to the [Author and run tests](https://grafana.com/docs/grafana-cloud/k6/author-run/) page, use `https://grafana.com/docs/grafana-cloud/k6/author-run/`.

**Link to Mimir documentation**:

Start with `https://grafana.com/docs/grafana/<MIMIR VERSION>/`, and add the rest of the URL path.
For example, to link to the [Release notes](https://grafana.com/docs/mimir/latest/release-notes/) page with version substitution,
use `https://grafana.com/docs/mimir/<MIMIR VERSION>/release-notes/`.

- If you are linking from Mimir documentation, `<MIMIR VERSION>` is substituted with the version inferred from the page's URL.

- If you are from other documentation, `<MIMIR VERSION>` is substituted with the value of `MIMIR VERSION` set in the page's front matter.

  Ensure that you set the appropriate version in the root `_index.md` file for your documentation.
  The following YAML, merged with the existing front matter in the root `_index.md` file sets `MIMIR VERSION` to be `latest` for that page and all child pages.

  ```yaml
  cascade:
    MIMIR VERSION: latest
  ```

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
