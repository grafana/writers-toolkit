---
title: Links and cross references
menuTitle: Links and cross references
description: Understand how Hugo determines references, the different types of references, and how to use them.
weight: 600
aliases:
  - /docs/writers-toolkit/writing-guide/references/
  - /docs/writers-toolkit/write/references/
keywords:
  - Hugo
  - references
  - relref
  - ref
---

# Links and cross references

Links are a mechanism for reusing content.
Instead of writing the same information twice, you can link to a definitive source of truth.

This page focuses only on HTTP-based URLs that have the scheme `http` or `https`.

## Choose the correct link

There are multiple ways to link to the same destination URL.
All of the following destinations link https://grafana.com/docs/grafana/latest/ to when followed from the page https://grafana.com/docs:

- https://grafana.com/docs/grafana/latest/: a fully qualified URL.
- /docs/grafana/latest: a partial URL with an absolute path.
- ./grafana/latest: a partial URL with a relative path.

**To choose the correct link:**

1. If the source is reused as described in [Reuse directories of content with Hugo mounts]({{< relref "../reuse-content/reuse-directories" >}}), use the `docs/reference` shortcode.

   For more information about the `docs/reference` shortcode, refer to [`docs/reference` shortcode]({{< relref "../shortcodes#docs-reference-shortcode" >}}).

1. If the destination is part of the current documentation set, consider using the `relref` shortcode.

   For example, `{{</* relref "./path/to/page" */>}}`.

   Hugo emits logs during the build for broken links defined with the `relref` shortcode.
   For more information about the `relref` shortcode, refer to [Build time link checking with Hugo](#build-time-link-checking-with-hugo).

1. Otherwise, use the fully qualified URL.

   For example, `[GitHub](https://github.com)`, or [Grafana](https://grafana.com/docs/grafana/latest/).

## Build time link checking with Hugo

Hugo has built-in shortcodes for creating links.
The `ref` and `relref` shortcodes display the absolute and relative permalinks to a page, respectively.
They both provide build time link checking to ensure that the destination file exists.

Relative references are the most common references in Grafana technical documentation.
This is the Hugo shortcode: `{{</* relref "<DESTINATION>" */>}}`.

{{% admonition type="note" %}}
Hugo link checking depends on having all the content available during the build.
In most projects, the only content available during local builds and CI is the current project documentation.
Therefore, the current advice is to only use the `relref` shortcode for links within the current project.
{{% /admonition %}}

### Determine `relref` shortcode destinations

The argument to the `relref` shortcode is the path to a source file in the Hugo content directory.
During local builds, the `docs/sources` directory is automatically mounted into the Hugo content directory.

Hugo has different kinds of source files for producing pages.
These include:

- page (`page.md`)
- leaf bundle (`page/index.md`)
- branch bundle (`page/_index.md`)

Each of those source files produce the same page.
To avoid a link breaking when the source file changes kind, you can ignore the file extension and index kind.
You can reference each of the preceding examples with the same argument -- `page`.

{{% admonition type="note" %}}
There is no trailing slash in the argument `page`.
Including a trailing slash prevents the argument working for some kinds of source files.
{{% /admonition %}}

{{% admonition type="note" %}}
If the destination file or its containing directory has a period (`.`) in the path, you must link to the source file directly.
{{% /admonition %}}

#### Example

In the Writers' Toolkit repository, with the following directory structure:

```
docs
└── sources
    ├── branch
    │   └── _index.md
    │   └── other.md
    └── leaf
        └── index.md
```

Hugo produces the following website pages:

```
/docs/writers-toolkit/branch/
/docs/writers-toolkit/branch/other/
/docs/writers-toolkit/leaf/
```

Refer to the following table for the correct `relref` shortcode to use to link between each of the example pages.

| Source page                           | Destination page                      | `relref` shortcode with relative path  | `relref` shortcode with absolute path                     |
| ------------------------------------- | ------------------------------------- | -------------------------------------- | --------------------------------------------------------- |
| `/docs/writers-toolkit/branch/`       | `/docs/writers-toolkit/branch/other/` | `{{</* relref "./other" */>}}`         | `{{</* relref "/docs/writers-toolkit/branch/other" */>}}` |
| `/docs/writers-toolkit/branch/`       | `/docs/writers-toolkit/leaf/`         | `{{</* relref "../leaf" */>}}`         | `{{</* relref "/docs/writers-toolkit/leaf" */>}}`         |
| `/docs/writers-toolkit/leaf/`         | `/docs/writers-toolkit/branch/`       | `{{</* relref "../branch" */>}}`       | `{{</* relref "/docs/writers-toolkit/branch" */>}}`       |
| `/docs/writers-toolkit/leaf/`         | `/docs/writers-toolkit/branch/other/` | `{{</* relref "../branch/other" */>}}` | `{{</* relref "/docs/writers-toolkit/branch/other" */>}}` |
| `/docs/writers-toolkit/branch/other/` | `/docs/writers-toolkit/branch/`       | `{{</* relref "." */>}}`               | `{{</* relref "/docs/writers-toolkit/branch" */>}}`       |
| `/docs/writers-toolkit/branch/other/` | `/docs/writers-toolkit/leaf/`         | `{{</* relref "../leaf" */>}}`         | `{{</* relref "/docs/writers-toolkit/leaf" */>}}`         |

## Anchors

In a reference, you can optionally include an anchor to a heading in the referenced page.
Specify and anchor at the end of the reference `#` followed by the normalized heading.

Hugo normalizes headings to make anchors.
To convert a heading to an anchor, Hugo makes the following changes:

1. Convert to lower case.
1. Remove any period characters (`.`).
1. Replace any character that's not a lower cased letter, a number, or an underscore (`_`) with dashes (`-`).
1. Trim any preceding or proceeding dashes (`-`).

If the anchor is in the current page, you don't need to use `relref` syntax.
The following Markdown links are equivalent:

- `[link text]({{</* relref "#anchor-in-current-page" */>}})`
- `[link text](#anchor-in-current-page)`

### Hugo error output

<!-- The output example is also used in review/run-a-local-webserver. -->

{{< docs/shared source="writers-toolkit" lookup="hugo-error-example-bad-link.md" version="" >}}

For additional information about Hugo error output, refer to [Test documentation changes]({{< relref "../../review/run-a-local-webserver" >}}).
