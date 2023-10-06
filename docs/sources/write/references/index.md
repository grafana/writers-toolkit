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

Choose your link type based on the applicable scenario:

- [Linking from source content that's used (or mounted) in multiple projects](#source-content-is-reused-in-multiple-projects)
- [Linking only within the same project](#linking-only-to-content-within-the-same-project)
- [Linking that isn't the first two types](#any-other-kind-of-link)
- [Linking within a page](#anchors)

## Source content is reused in multiple projects

Use:

- the `docs/reference` shortcode
- partial URL with a relative path

The source is reused as described in [Reuse directories of content with Hugo mounts]({{< relref "../reuse-content/reuse-directories" >}}). For more information, refer to the [`docs/reference` shortcode]({{< relref "../shortcodes#docsreference" >}}).

For example:
add examples here


## Linking only to content within the same project

Use:

- the `relref` shortcode
- partial URL with relative path

For example: `{{</* relref "./path/to/page" */>}}`.

### Buildtime link checking and Hugo error output

The `relref` shortcode provides build-time link checking to ensure that the destination file exists.

Hugo link checking only applies to the content available during the build. 
In most projects, the only content available during local builds and CI is the current project documentation,
so you should be aware that just because a link uses a `relref`, it doesn't automatically follow that the link can be checked.

Emitted Hugo errors look like this:

<!-- The output example is also used in review/run-a-local-webserver. -->

{{< docs/shared source="writers-toolkit" lookup="hugo-error-example-bad-link.md" version="" >}}

For additional information about Hugo error output, refer to [Test documentation changes]({{< relref "../../review/run-a-local-webserver" >}}).

For information about determining the relative path included in a `relref` shortcode, refer to ### Determine relref

<!-- why don't we document relrefs on the shortcodes page? -->

### Determine `relref` shortcode destinations

To determine the path between the source and destination content, do the following:

Find the folder where the destination content lives. 
Find the folder that the source and destination folders have in common. 
Note the pathway from the destination folder to the common folder.
Count the number of folders from the source to the common folder and that number equals the number of aliases (../) you need to add to your relative path.

For example, with the following folder structure:
Vehicles
|_Trucks
  |_F150
    |_1999 F150
Vans

In this case, the source content is in the 1999 F150 folder and the destination content is in the Vans folder.
The common folder for the two pieces of content is the Vehicles folder.
The number of folders between the source folder, 1999 F150 and the common folder Vehicles is 3 folders, so this requires three aliases: ../../../
The pathway from the common Vehicles to destination folder Vans /vans
The relative path is ../../../vans

If the source folder was Vans and the desitination was 1999 F150, the pathway would be ../trucks/F150/1999-F150

## Any other kind of link

Use a fully qualified URL. For example: https://grafana.com/docs/grafana/latest/

<!-- is the version inferred or do you need to use a version inference thing? -->

This includes links to:

- Other projects where content isn't shared
- Other parts of grafana.com
- All sites external to grafana.com


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



