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

- [Link from source content that's used in multiple projects](#link-from-source-content-thats-used-in-multiple-projects)
- [Link to grafana.com pages](#link-to-grafanacom-pages)
- [Link to external pages](#link-to-external-pages)
- [Link to page headings](#link-to-page-headings)

Although these other types of links still function, replace them with full URLs:

- [Hugo `relref` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#relref)

## Link from source content that's used in multiple projects

Use the `docs/reference` shortcode.

The source is reused as described in [Reuse directories of content with Hugo mounts](https://grafana.com/docs/writers-toolkit/write/reuse-content/reuse-directories/).
For more information and examples, refer to [`docs/reference` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#docsreference).

## Link to grafana.com pages

Use a full URL.

{{< admonition type="note" >}}
The `doc-validator` linter does not check links that use full URLs.
{{< /admonition >}}

If you are linking to versioned documentation, use a full URL with version substitution syntax.
Version substitution is necessary for full URLs to link to the correct version of documentation.
Usually, this is the current version of documentation.

In versioned documentation, set the correct version in the root `_index.md` file for your documentation.
The following YAML example merges with the existing front matter in the root `_index.md` file, and sets `GRAFANA_VERSION` to be `latest` for that page and all child pages.

```yaml
cascade:
  GRAFANA_VERSION: latest
```

### Examples

**Link to Grafana documentation**:

Start with `https://grafana.com/docs/grafana/<GRAFANA_VERSION>/`, and add the rest of the URL path. Include trailing slashes.

For example, to link to the [Developers](https://grafana.com/docs/grafana/latest/developers) page with version substitution, use:

```markdown
https://grafana.com/docs/grafana/<GRAFANA_VERSION>/developers/
```

- If you're linking from Grafana documentation, `<GRAFANA_VERSION>` is substituted with the version inferred from the page's URL.
- If you're linking from other documentation, `<GRAFANA_VERSION>` is substituted with the value of `GRAFANA_VERSION` from the source page's front matter.

**Link to Grafana Cloud documentation**:

Grafana Cloud documentation is not versioned so no version substitution syntax is needed.
Use the full URL.

For example, to link to the [Author and run tests](https://grafana.com/docs/grafana-cloud/k6/author-run/) page, use:

```markdown
https://grafana.com/docs/grafana-cloud/k6/author-run/
```

**Link to Mimir documentation**:

Start with `https://grafana.com/docs/grafana/<MIMIR_VERSION>/`, and add the rest of the URL path.

For example, to link to the [Release notes](https://grafana.com/docs/mimir/latest/release-notes/) page with version substitution, use:

```markdown
https://grafana.com/docs/mimir/<MIMIR_VERSION>/release-notes/
```

- If you're linking from Mimir documentation, `<MIMIR_VERSION>` is substituted with the version inferred from the page's URL.
- If you're linking from other documentation, `<MIMIR_VERSION>` is substituted with the value of `MIMIR_VERSION` from the source page's front matter.

## Link to external pages

Use the full URL. Copy the URL as it is from the address bar. If it includes a trailing slash, include it; if it doesn't, don't.

For example:

```markdown
https://github.com
```

## Link to page headings

Link to a heading on a page in one of two ways.

From within the same page:

```markdown
Read more in the [Configuration section](#configuration) of this page.
```

From a different page:

```markdown
Read more in the [Grafana Open Source section of the Introduction page](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/fundamentals/#grafana-open-source).
```

Include the trailing slash before the hash (#) that precedes the page heading.

To convert a heading to an anchor, make the following changes:

1. Convert to lower case.
1. Remove any period characters (`.`).
1. Replace any character that's not a lower cased letter, a number, or an underscore (`_`) with dashes (`-`).
1. Trim any preceding or proceeding dashes (`-`).
1. Prefix with a `#`.

The heading _Link to page headings_ becomes the anchor `#link-to-page-headings`.
