---
aliases:
  - /docs/writers-toolkit/writing-guide/reuse-directories/
  - /docs/writers-toolkit/write/reuse-content/reuse-directories/
date: "2023-03-29T09:30:24+01:00"
description: Learn to reuse directories of content with Hugo mounts.
keywords:
  - content reuse
  - shared content
  - Hugo
menuTitle: Reuse directories of content
review_date: "2024-06-26"
title: Reuse directories of content with Hugo mounts
---

# Reuse directories of content with Hugo mounts

{{< admonition type="note" >}}
Only a Grafana Labs employee can perform the following task.
{{< /admonition >}}

Before Hugo performs a build, it reads the [_Hugo module mounts_](https://gohugo.io/hugo-modules/configuration/#module-configuration-mounts) configuration to construct a virtual filesystem.
Each mount mounts a source directory at a destination directory within that filesystem.
You can use Hugo mounts to reuse a whole directory of content in another part of the published technical documentation.

## Before you begin

- Identify a directory of content that you want to reuse.
- Identify any pages to exclude.
  Typically, excluded pages are ones that don't make sense in the context of the target directory.

  {{< admonition type="warning" >}}
  Pages in the directory might have relative links to other pages that would be broken by their exclusion.

  If the excluded page is replaced by another in the target directory, relative links continue to work.

  Otherwise, you should make the links absolute so that they always refer to the page in the source directory.
  {{< /admonition >}}

## Steps

To reuse a shared directory:

1. Determine the source and target directories.

   The source directory is the path to a directory in the `website` repository.
   It always has the path prefix `content/docs/`.
   The source directory for all Grafana Cloud content at the URL https://grafana.com/docs/grafana-cloud/ is `content/docs/grafana-cloud`.

   Similarly, the target directory has the path prefix `content/docs`.
   For the destination URL `https://grafana.com/docs/target-directory/`, the path is `content/docs/target-directory`.

1. For every page in the source directory, set the canonical URL using the `canonical` front matter, to the published open source page URL.

   The `canonical` front matter indicates the preferred URL for duplicate or very similar pages.
   For more information, refer to [Canonical](https://grafana.com/docs/writers-toolkit/write/front-matter/#canonical).

1. Update the website repository Hugo configuration.

   The configuration is in the `website` repository in the `config/_default/params.yaml` file.

   Append the mount to the `manual_mounts` key in the YAML configuration.

   For example, mount the source `content/docs/source-directory` at `content/docs/target-directory` and exclude the `content/docs/source-directory/_index.md` file:

   ```yaml
   - source: content/docs/source-directory
     target: content/docs/target-directory
     excludeFiles:
       - /_index.md
   ```
