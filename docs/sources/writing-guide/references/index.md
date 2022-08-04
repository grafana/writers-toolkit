---
title: Links and cross references
menuTitle: Links and cross references
description: Understand how Hugo determines references, the different types of references, and how to use them.
aliases:
  - /docs/writers-toolkit/latest/writing-guidelines/references/
weight: 600
keywords:
  - Hugo
  - references
  - relref
  - ref
---

# Links and cross references

References are hyperlinks between pages.
Hugo supports several types of references and the references themselves can have different forms.

You can split references into one of two categories:

- **Relative references**: often referred to as **relrefs**.
- **Absolute references**: typically just **references**.

The difference between the two categories of references is how Hugo resolves them to link between pages.
Hugo resolves relative references from the current page or file.
Hugo resolves absolute references from the root of the website.

> **Note:** For Hugo's purposes, other versions of the docs, such as a version-specific archived docs set (`https://grafana.com/docs/grafana/v8.5/`, etc.) or `/next/` docs for links in content residing in `/latest`, cannot be addressed using Hugo references.
> Use regular Markdown link syntax (`[link text](https://example.com/docs/etc)`) for such links.
> Unlike references, Hugo will _not_ confirm that these link destinations exist, so manually check the published links to confirm that they point correctly.

## Relative references

Relative references are the most common references in Grafana technical documentation.
This is the Hugo shortcode: `{{< relref "<RELATIVE FILE/URL PATH>" >}}`

There are two forms of relative references:

- **File path based**: resolved using file paths.
- **URL path based**: resolved using URL paths.

Typically, file path based relative references include a `.md` suffix and no trailing slash.
Conversely, URL path based relative references don't have a file suffix, but might have a trailing slash.

### Determine relrefs

How Hugo resolves relative references can cause confusion because of how they interact with [Page Bundles](https://gohugo.io/content-management/page-bundles/).
Hugo resolves URL path based relrefs from the page bundle, and not the page itself.

For example, the following directory structure:

```
docs
└── sources
    ├── branch
    │   └── _index.md
    │   └── other.md
    └── leaf
        └── index.md
```

It produces the following website pages:

```
/docs/technical-documentation/branch/
/docs/technical-documentation/branch/other/
/docs/technical-documentation/leaf/
```

All of the URL path based relrefs between these pages are are follows:

| Source page                                   | Destination page                              | relref                             |
| --------------------------------------------- | --------------------------------------------- | ---------------------------------- |
| `/docs/technical-documentation/branch/`       | `/docs/technical-documentation/branch/other/` | `{{< relref "other" >}}`           |
| `/docs/technical-documentation/branch/`       | `/docs/technical-documentation/leaf/`         | `{{< relref "../leaf" >}}`         |
| `/docs/technical-documentation/leaf/`         | `/docs/technical-documentation/branch/`       | `{{< relref "../branch" >}}`       |
| `/docs/technical-documentation/leaf/`         | `/docs/technical-documentation/branch/other/` | `{{< relref "../branch/other" >}}` |
| `/docs/technical-documentation/branch/other/` | `/docs/technical-documentation/branch/`       | `{{< relref "./" >}}`              |
| `/docs/technical-documentation/branch/other/` | `/docs/technical-documentation/leaf/`         | `{{< relref "../leaf" >}}`         |

You can refer to the table below for all file path based relrefs between these files.

| Source file                      | Destination file                 | relref                                 |
| -------------------------------- | -------------------------------- | -------------------------------------- |
| `/docs/sources/branch/_index.md` | `/docs/sources/branch/other.md`  | `{{< relref "other.md" >}}`            |
| `/docs/sources/branch/_index.md` | `/docs/sources/leaf/index.md`    | `{{< relref "../leaf/index.md" >}}`    |
| `/docs/sources/leaf/index.md`    | `/docs/sources/branch/_index.md` | `{{< relref "../branch/_index.md" >}}` |
| `/docs/sources/leaf/index.md`    | `/docs/sources/branch/other.md`  | `{{< relref "../branch/other.md" >}}`  |
| `/docs/sources/branch/other.md`  | `/docs/sources/branch/_index.md` | `{{< relref "_index.md" >}}`           |
| `/docs/sources/branch/other.md`  | `/docs/sources/leaf/index.md`    | `{{< relref "../leaf/index.md" >}}`    |

## Absolute references

Absolute references are less common in Grafana technical documentation.
This is the Hugo shortcode: `{{< ref "<ABSOLUTE FILE/URL PATH>" >}}`

There are two forms of absolute references:

- **File path based**: resolved using file paths.
- **URL path based**: resolved using URL paths.

Typically, file path based relative references include a `.md` suffix and no trailing slash.
Conversely, URL path based relative references don't have a `.md` file suffix.

> **Note:** Unlike relrefs, refs with a trailing slash aren't resolved by Hugo.

For example, the following directory structure:

```
docs
└── sources
    ├── branch
    │   └── _index.md
    │   └── other.md
    └── leaf
        └── index.md
```

It produces the following website pages:

```
/docs/technical-documentation/branch/
/docs/technical-documentation/branch/other/
/docs/technical-documentation/leaf/
```

All of the URL path based relrefs between these pages are are follows:

| File                             | Page                                          | File path ref                                                  | URL path ref                                               |
| -------------------------------- | --------------------------------------------- | -------------------------------------------------------------- | ---------------------------------------------------------- |
| `/docs/sources/branch/_index.md` | `/docs/technical-documentation/branch/`       | `{{< ref "/docs/technical-documentation/branch/_index.md" >}}` | `{{< ref "/docs/technical-documentation/branch" >}}`       |
| `/docs/sources/branch/other.md`  | `/docs/technical-documentation/branch/other/` | `{{< ref "/docs/technical-documentation/branch/other.md" >}}`  | `{{< ref "/docs/technical-documentation/branch/other" >}}` |
| `/docs/sources/leaf/index.md`    | `/docs/technical-documentation/leaf/`         | `{{< ref "/docs/technical-documentation/leaf/index.md" >}}`    | `{{< ref "/docs/technical-documentation/leaf" >}}`         |

## Anchors

In a reference, you can optionally include an anchor to a heading in the referenced page.
Specify and anchor at the end of the reference `#` followed by the normalized heading.

Hugo normalizes headings to make anchors.
To convert a heading to an anchor, Hugo makes the following changes:

1. Convert to lower case.
1. Remove any period characters (`.`).
1. Replace any character that's not a lower cased letter, a number, or an underscore (`_`) with dashes (`-`).
1. Trim any preceding or proceeding dashes (`-`).

## References across docs sets built by `grafana/website`

Product docs are collated as part of the `grafana/website` repository's publishing process, which generates the live website at grafana.com.
It might be necessary to create a reference from one of these component docs sets, such as docs in `grafana/grafana`, to another component docs set, such as docs in `grafana/cloud-docs` or tutorial content.
Regardless of how you reference such documents, Hugo generates `REF_NOT_FOUND` warnings about those references when building docs from only that component repository.

To link to such docs, use refs to reference those docs by their unique identifier as they exist in the other docs set. Avoid relrefs if possible. If you cannot uniquely address such docs, you might need to determine how to address those documents by troubleshooting the `grafana/website` build.

## Troubleshooting

### Links generated from references point to their own page/self-reference

Hugo generates HTML link tags for properly formatted but incorrectly addressed references, such as those targeting a document Hugo cannot resolve.
This does not break the docs build, neither locally nor in the publishing process's continuous integration (CI) pipeline.
In the generated source, Hugo leaves the hypertext reference (`href`) attribute unexpectedly empty:

```html
<a href="">Link text</a>
```

When clicking the resulting link in a browser, the browser loads the page that contains the link.
In other words, the user is taken nowhere.
Such links do not appear on 404 reports because the resulting links do not point to a technically invalid destination.

Hugo emits `REF_NOT_FOUND` warnings indicating the filename and location of such references when building the docs, for example with `make docs` in `grafana/grafana` or `make server-quick` in `grafana/website`:

```
WARN 2022/08/04 21:35:37 [en] REF_NOT_FOUND: Ref "../../enterprise/": "/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md:14:47": page not found
```

In this example,

- `Ref "../../enterprise/"` is the destination of the reference that Hugo cannot resolve
- `/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md` is the document containing the reference, where the path after `/next/` is relative to the documentation root of the component repository
- `:14` represents the line number containing the unresolved reference
- `:47` represents the character in that line where the unresolved reference begins

If the reference's destination appears to be invalid, for example due to a typo in the reference or relref directory traversal depth, then you should be able to resolve this by correcting the reference target.

However, if the reference's destination appears to be valid, it might not be referencing a unique document, or a sufficiently specific or correct path.
You might need to use a different or more specific destination, or use a ref to reference the document's unique identifier if it has one.

A document's filename or its aliases can serve as unique identifiers, but they must be unique across all documents Hugo is processing.
For the live grafana.com website, this means the document or an alias must be unique across all _component_ docs sets&mdash;for example, across the combination of `grafana/grafana` docs, and `grafana/mimir` docs, and `grafana/cloud-docs`, etc.
