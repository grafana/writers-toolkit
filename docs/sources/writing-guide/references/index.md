---
title: "Links and cross references"
menuTitle: "Links and cross references"
description: "Understand how Hugo determines references, the different types of references, and how to use them."
aliases: []
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
