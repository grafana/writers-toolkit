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

{{% admonition type="note" %}}
When linking to specific versions or across repositories, use standard markdown links. Read the [Versions and cross-repository linking]({{< relref "#versions-and-cross-repository-linking" >}}) section for details.
{{% /admonition %}}

## Understanding hyperlinks

Links can be written in many forms that are enumerated in [<a>: The Anchor element - HTML: HyperText Markup Language | MDN](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a#href).

This document focuses only on HTTP-based URLs that have the scheme `http` or `https`.

Link destinations can be specified in different ways.
All of the following destinations link https://grafana.com/docs/grafana/latest to when followed from the page https://grafana.com/docs:

- https://grafana.com/docs/grafana/latest: fully specified HTTPS URL.
- /docs/grafana/latest: partial URL with absolute path.
- ./grafana/latest: relative path.

To choose the correct link destination type, follow these steps:

1. If the destination is external to the https://grafana.com website, use the fully specified HTTPS URL.
   For example, `[GitHub](https://github.com)`.
1. If the source is reused as described in [Reuse directories of content with Hugo mounts]({{< relref "../reuse-content/reuse-directories" >}}):

   1. If the destination is also present in the destination mount, use a relative URL path.
      This keeps reader within the destination project.
      For example, `{{</* relref "./path/to/page" */>}}`.

      Use the Hugo `relref` shortcode for build time link checking.
      For more information about the `relref` shortcode, refer to [Build time link checking with Hugo](#build-time-link-checking-with-hugo).

   1. Otherwise use a partial URL with an absolute path.
      This always brings the reader back to the destination page in the source project.
      For example, `/docs/writers-toolkit/write/`.

1. If the destination is internal to the https://grafana.com website, but external to the project documentation,
   use a partial URL with an absolute path.
   For example, `/blog/`.
1. Otherwise, use a relative path.
   For example, `{{</* relref "./path/to/page" */>}}`.

   Use the Hugo `relref` shortcode for build time link checking.
   For more information about the `relref` shortcode, refer to [Build time link checking with Hugo](#build-time-link-checking-with-hugo).

## Build time link checking with Hugo

Hugo has built-in shortcodes for creating links to documents.
The `ref` and `relref` shortcodes display the absolute and relative permalinks to a document, respectively.
They both provide build time link checking to ensure that the destination exists.

{{% admonition type="warning" %}}
If you do not use a Hugo `relref` shortcode for build time link checking, your links may be broken without you realizing it.
{{% /admonition %}}

Relative references are the most common references in Grafana technical documentation.
This is the Hugo shortcode: `{{</* relref "<DESTINATION>" */>}}`.

{{% admonition type="note" %}}
Hugo link checking depends on having all the content available during the build.
In most projects, the only content available during local builds and CI is the current project documentation.
Therefore, the current advice is that `relref`s should only be used for links within the current project.
{{% /admonition %}}

### Determine `relref` shortcode destinations

With the following directory structure:

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

| Source page                           | Destination page                      | `relref` shortcode                     |
| ------------------------------------- | ------------------------------------- | -------------------------------------- |
| `/docs/writers-toolkit/branch/`       | `/docs/writers-toolkit/branch/other/` | `{{</* relref "./other" */>}}`         |
| `/docs/writers-toolkit/branch/`       | `/docs/writers-toolkit/leaf/`         | `{{</* relref "../leaf" */>}}`         |
| `/docs/writers-toolkit/leaf/`         | `/docs/writers-toolkit/branch/`       | `{{</* relref "../branch" */>}}`       |
| `/docs/writers-toolkit/leaf/`         | `/docs/writers-toolkit/branch/other/` | `{{</* relref "../branch/other" */>}}` |
| `/docs/writers-toolkit/branch/other/` | `/docs/writers-toolkit/branch/`       | `{{</* relref "." */>}}`               |
| `/docs/writers-toolkit/branch/other/` | `/docs/writers-toolkit/leaf/`         | `{{</* relref "../leaf" */>}}`         |

{{% admonition type="warning" %}}
If the destination file or its containing directory has a period (`.`) in the path, you must link to the source file directly.
{{% /admonition %}}

## Versions and cross-repository linking

For Grafana's webserver environments, you can't address other versions of the docs, such as a version-specific archived docs set (`https://grafana.com/docs/grafana/v8.5/` and so forth) or `/next/` docs for links in content residing in `/latest/`, using Hugo references.

Hugo references addressed across different products' docs, such as from `/docs/grafana/` to `/docs/loki/` and vice-versa, as well as references from docs addressed to other Hugo-published content on grafana.com, can also be predictably addressed.

To avoid broken links in these situations on grafana.com, use regular Markdown link syntax (`[link text](/docs/repo/version/folder/file/)`) instead of Hugo references (`relref`). To ensure the links work in local builds, staging environments, and the live website, you **shouldn't** use a fully qualified URL with `https://grafana.com` for links to other content on grafana.com.

For cross-repository links, use a standard markdown link, with the link structured like this: `/docs/repo/page`.

For example:

```markdown
This is an [example cross-repository link](/docs/grafana/whatsnew) to the Grafana repository.
```

Using a Hugo `relref` in a cross-repository link or a link to a specific version can result in a page not found error message when running `make docs` if the linked content isn't mounted when using the script.

Unlike references, Hugo does _not_ confirm that these link destinations exist during its build, so manually confirm that the published links in a local build and on the published website point correctly.
With partial URIs, you also cannot check these links without the content mounted. For example, `/docs/grafana/latest/` from `/docs/tempo/latest` won't resolve unless you have both projects mounted in the webserver.

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

{{< docs/shared source="writers-toolkit" lookup="hugo-error-example-bad-link.md" version="latest" >}}

For additional information about Hugo error output, refer to [Test documentation changes]({{< relref "../../review/run-a-local-webserver" >}}).
