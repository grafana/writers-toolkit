---
date: "2025-09-10T00:00:00+01:00"
description: Learn where data source documentation files live, how they publish, and how to write front matter and links.
keywords:
  - data source
  - plugin
  - front matter
  - links
menuTitle: File structure and front matter
review_date: "2027-09-04"
title: File structure and front matter
weight: 100
---

# File structure and front matter

Document every data source as if it lives in a standalone plugin repository.
This page describes the file structure, front matter, and linking rules that follow from that model.

## File structure

Use a flat file structure under `docs/sources/`: one Markdown file per guide, with no subfolders.
Each guide is a top-level `.md` file, not an `index.md` inside a folder.

```
docs/
└── sources/
    ├── _index.md              # Overview and getting started
    ├── install.md             # Installation and upgrade (Enterprise plugins)
    ├── configure.md           # Configuration and authentication
    ├── query-editor.md        # Query editor usage and examples
    ├── template-variables.md  # Template variable support
    ├── annotations.md         # Annotation support
    ├── alerting.md            # Grafana Alerting support
    └── troubleshooting.md     # Common issues and solutions
```

Create only the guides a data source needs.
For a small data source, a single `_index.md` is enough.
Add separate guides, such as `troubleshooting.md`, as the content grows.

## Published URLs

Each flat `<page>.md` file publishes to `/docs/plugins/<plugin-id>/latest/<page>/`.
The `_index.md` file publishes to the bare `/docs/plugins/<plugin-id>/latest/` URL, with no `_index` segment.

## Front matter

Every data source documentation file requires YAML front matter.
Use the following template:

```yaml
---
aliases:
  - ../data-sources/<data-source-name>/
description: Guide for using <Data Source Name> in Grafana
keywords:
  - grafana
  - <data-source-name>
  - <additional-keywords>
labels:
  products:
    - cloud
    - enterprise
    - oss
menuTitle: <Short Name>
title: <Data Source Name> data source
weight: <number>
review_date: <YYYY-MM-DD>
---
```

The following table describes the required fields:

| Field             | Purpose                                                                      |
| ----------------- | ---------------------------------------------------------------------------- |
| `description`     | Appears in search results. Keep it under 160 characters.                     |
| `keywords`        | Helps users find the page. Include common terms users search for.            |
| `labels.products` | Controls which product docs include this page: `cloud`, `enterprise`, `oss`. |
| `menuTitle`       | Short navigation title.                                                      |
| `title`           | The H1 heading for the page. It must match the `# Title` in the content.     |
| `weight`          | Controls sort order in navigation. Lower numbers sort first.                 |
| `review_date`     | Date of the last content review, in `YYYY-MM-DD` format.                     |

## Links

Because you document every data source as a standalone plugin, always use fully qualified links.

- Link to sibling pages in the same plugin docs: `https://grafana.com/docs/plugins/<plugin-id>/latest/<page>/`.
- Link to Grafana product docs: `https://grafana.com/docs/grafana/<GRAFANA_VERSION>/<path>/`.
- Always include trailing slashes. Append `#anchor` after the trailing slash to link to a heading.

Don't use the `refs:` front matter section or `ref:` links.
The ref system only works within the main `grafana` repository, so it breaks when a plugin moves to its own repository.

Don't use relative links or `.md`-extension links.
The Hugo build for plugin docs published on the Grafana website doesn't rewrite them, so they return a 404 or render as raw Markdown text.

| Don't                              | Do                                                                                 |
| ---------------------------------- | ---------------------------------------------------------------------------------- |
| `[Configure](configure.md)`        | `[Configure](https://grafana.com/docs/plugins/<plugin-id>/latest/configure/)`      |
| `[GCE](configure.md#gce)`          | `[GCE](https://grafana.com/docs/plugins/<plugin-id>/latest/configure/#gce)`        |
| `[Quota](_index.md#quota)`         | `[Quota](https://grafana.com/docs/plugins/<plugin-id>/latest/#quota)`              |
| `Refer to [Explore](ref:explore).` | `Refer to [Explore](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/explore/).` |

## Aliases for redirects

Always add aliases when you change the documentation structure.
When you rename, move, or restructure a page, add aliases to the new page that point to the old URLs so existing links and bookmarks keep working.

Add aliases in these common scenarios:

- Renaming a page, for example `setup/` to `configure/`.
- Moving a page to a different location.
- Consolidating multiple pages into one.
- Restructuring the documentation hierarchy.

When you rename a page, add aliases for both the `latest/<old-page>/` and the bare `<old-page>/` form so that bookmarks to both the versioned and the bare URLs keep working.
For example, when you rename `troubleshoot` to `troubleshooting`, add the following to the new `troubleshooting` page:

```yaml
aliases:
  - /docs/plugins/<plugin-id>/latest/troubleshoot/
  - /docs/plugins/<plugin-id>/troubleshoot/
```
