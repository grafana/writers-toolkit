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

Other link types exist in our documentation but their usage is discouraged:

- [Hugo `relref` shortcode](#hugo-relref-shortcode)

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

## Hugo `relref` shortcode

The `relref` shortcode provides build-time link checking to ensure that the destination file exists.

For example: `{{</* relref "./path/to/page" */>}}`.

Hugo link checking only applies to the content available during the build.
In most projects, the only content available during local builds and CI is the current project documentation,
so you should be aware that just because a link uses a `relref`, it doesn't automatically follow that the link can be checked.

Emitted Hugo errors look like this:

<!-- The output example is also used in review/run-a-local-webserver. -->

{{< docs/shared source="writers-toolkit" lookup="hugo-error-example-bad-link.md" version="" >}}

For additional information about Hugo error output, refer to [Test documentation changes](https://grafana.com/docs/writers-toolkit/review/run-a-local-webserver/).

### Determine `relref` shortcode arguments

To determine the path between the source and destination content, do the following:

Find the directory where the destination content lives.
Find the directory that the source and destination directories have in common.
Note the relative path from the common directory to the destination directory.
Count the number of folders from the source to the common directory and that number equals the number of parent directory path elements (`..`) you need to add to your relative path.
Join all the path elements together with forward slashes (`/`).

For example, with the following folder structure:

```
Vehicles
├── Trucks
│   ├── F150
│   └── 1999 F150
└── Vans
```

In this case, the source content is in the `1999 F150` directory and the destination content is in the `Vans` directory.
The common folder for the two pieces of content is the `Vehicles` directory.

The parent directory of `1999 F150` is `Trucks`, requiring one `..` path element.
To parent directory of `Trucks` is `Vehicles`, requiring another `..` path element.
Therefore, the relative path from the source directory, `1999 F150`, and the common directory, `Vehicles`, is `../..`

The pathway from the common directory `Vehicles` to destination directory `Vans` is `vans`
The relative path is `../../vans`

If the source directory was `Vans` and the destination was `1999 F150`, the relative path would be `../trucks/F150/1999-F150`.
