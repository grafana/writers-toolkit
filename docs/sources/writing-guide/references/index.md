---
title: Links and cross references
menuTitle: Links and cross references
description: Understand how Hugo determines references, the different types of references, and how to use them.
weight: 600
keywords:
  - Hugo
  - references
  - relref
  - ref
---

# Links and cross references

Links are a mechanism for reusing content.
Instead of writing the same information twice, you can link to a definitive source of truth.

When linking, keep in mind that the reader is directed away from the current content when following the link.
Just because a link can be made, it doesn't mean it should be made.

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
   For example, https://github.com.
1. If the source is reused as described in [Reuse directories of content with Hugo mounts]({{< relref "../reuse-directories" >}}):

   1. If the destination is also present in the destination mount, use a relative URL path.
      This keeps reader within the destination project.
      For example, `{{</* relref "./path/to/page" */>}}`.

      Use the Hugo `relref` shortcode for build time link checking.
      For more information about the `relref` shortcode, refer to [Build time link checking with Hugo](#build-time-link-checking-with-hugo).

   1. Otherwise use a partial URL with an absolute path.
      This always brings the reader back to the destination page in the source project.
      For example, `/docs/writers-toolkit/writing-guide/`.

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
If you do not use a Hugo relref for build time link checking, your links may be broken without you realizing it.
{{% /admonition %}}

Relative references are the most common references in Grafana technical documentation.
This is the Hugo shortcode: `{{</* relref "<DESTINATION>" */>}}`.

{{% admonition type="note" %}}
Hugo link checking depends on having all the content available during the build.
In most projects, the only content available during local builds and CI is the current project documentation.
Therefore, the current advice is that `relref`s should only be used for links within the current project.
{{% /admonition %}}

### Determine relref destinations

Hugo resolves URL path based relrefs from the page bundle, and not the page itself.

For example, with the following directory structure:

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

Refer to the following table for the correct relref to use to link between each of the example pages.

| Source page                           | Destination page                      | relref                                 |
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

### Anchors

In a reference, you can optionally include an anchor to a heading in the referenced page.
Specify and anchor at the end of the reference `#` followed by the normalized heading.

Hugo normalizes headings to make anchors.
To convert a heading to an anchor, Hugo makes the following changes:

1. Convert to lower case.
1. Remove any period characters (`.`).
1. Replace any character that's not a lower cased letter, a number, or an underscore (`_`) with dashes (`-`).
1. Trim any preceding or proceeding dashes (`-`).

### Hugo error output

Hugo emits `REF_NOT_FOUND` warnings indicating the filename and location of such references when building the docs, for example with `make docs` in `grafana/grafana` or `make server-quick` in `grafana/website`:

```
WARN 2022/08/04 21:35:37 [en] REF_NOT_FOUND: Ref "../../enterprise": "/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md:14:47": page not found
```

In this example,

- `Ref "../../enterprise"` is the destination of the reference that Hugo can't resolve
- `/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md` is the document containing the reference, where the path after `/next/` is relative to the documentation root of the component repository
- `:14` represents the line number containing the unresolved reference
- `:47` represents the character in that line where the unresolved reference begins

If the reference's destination appears to be invalid, for example due to a typo in the reference or relref directory traversal depth, then you should be able to resolve this by correcting the reference target.
