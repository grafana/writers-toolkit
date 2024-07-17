---
aliases:
  - /docs/writers-toolkit/write/links/
  - /docs/writers-toolkit/write/references/
  - /docs/writers-toolkit/writing-guide/references/
date: "2023-11-02T14:11:04+00:00"
description: Understand how to link between pages.
keywords:
  - Hugo
  - link
  - linking
  - links
  - ref
  - references
  - relref
review_date: "2024-06-24"
title: Links
weight: 600
---

# Links

Choose your link type based on your goal:

- [Link from source content that's reused as multiple pages](#link-from-source-content-thats-reused-as-multiple-pages)
- [Link to `grafana.com` pages](#link-to-grafanacom-pages)
- [Link to external pages](#link-to-external-pages)
- [Link to page headings](#link-to-page-headings)
  - Use this with one of the preceding options.

Although these other types of links still function, replace them with one of the preceding options:

- [Hugo `relref` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#relref)
- [`docs/reference` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#docsreference)

## Link from source content that's reused as multiple pages

Use `ref` URIs.
`Ref` URIs have two components:

- [Link](#link)
- [Front matter](#front-matter)

{{< admonition type="note" >}}
If you're using `doc-validator` in CI, you must upgrade to v5.0.0 to use `ref` URIs.

`doc-validator` no longer checks link destinations.
You must manually check link destinations in the local preview or fix broken links identified by the weekly website crawl.
{{< /admonition >}}

`ref` URIs look up destinations based upon the page's URL path and the definitions in the page's front matter.

### Link

A link with a `ref` URI looks like:

```markdown
[LINK TEXT](ref:<KEY>)
```

For the _`<KEY>`_ , enter an alphabetic term.
It can include hyphens (`-`).

Hugo looks up _`<KEY>`_ in the value for the `refs` field in the page's front matter.
If there is no _`<KEY>`_ in the `refs` field, or there is no `refs` field in the front matter, Hugo logs a build error.

### Front matter

{{< docs/shared source="writers-toolkit" lookup="refs-example.md" leveloffset="+2" >}}

## Link to `grafana.com` pages

Use a full URL.

{{< admonition type="note" >}}
The `doc-validator` linter doesn't check links that use full URLs.
{{< /admonition >}}

If you are linking to versioned documentation, use a full URL with version substitution syntax instead of the version path element.
For example, in Grafana, use `<GRAFANA_VERSION>` instead of `latest` in the URL `https://grafana.com/docs/grafana/latest/`.

When Hugo renders links with version substitution, it replaces the `<SOMETHING_VERSION>` syntax with the version inferred from the current page.

To understand the behavior in more detail, refer to [About version substitution](https://grafana.com/docs/writers-toolkit/write/shortcodes/#about-version-substitution)
You use full URLs with the substitution variable so that links resolve to the correct version of documentation without requiring the writer to update the version for each release.
For examples of behavior, refer to [Examples](#examples).

To override the version inferred by version substitution, set the preferred version in the root `_index.md` file for your documentation.
The following YAML snippet sets `GRAFANA_VERSION` to be `latest` for that page and all child pages.
You must merge the following YAML example with the front matter in the root `_index.md` file.

```yaml
cascade:
  GRAFANA_VERSION: latest
```

### Examples

**Link to Grafana documentation**:

Start with `https://grafana.com/docs/grafana/<GRAFANA_VERSION>/`, and add the rest of the URL.
Include trailing slashes.

For example, to link to the [Developers](https://grafana.com/docs/grafana/latest/developers/) page with version substitution, use:

```markdown
https://grafana.com/docs/grafana/<GRAFANA_VERSION>/developers/
```

- If you're linking from Grafana documentation, `<GRAFANA_VERSION>` is substituted with the version inferred from the page's URL.
- If you're linking from other documentation, `<GRAFANA_VERSION>` is substituted with the value of `GRAFANA_VERSION` from the source page's front matter.

**Link to Grafana Cloud documentation**:

Grafana Cloud documentation isn't versioned and doesn't require version substitution syntax.
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

Use the full URL.
Copy the URL exactly from the address bar.
If it includes a trailing slash, include it.
If it doesn't include a trailing slash, don't add one.

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
