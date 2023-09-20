---
description: Learn to reuse directories of content with Hugo mounts.
keywords:
  - content reuse
  - shared content
  - Hugo
menuTitle: Reuse directories of content
title: Reuse directories of content with Hugo mounts
aliases:
  - /docs/writers-toolkit/writing-guide/reuse-directories/
  - /docs/writers-toolkit/write/reuse-content/reuse-directories/
---

# Reuse directories of content with Hugo mounts

> **Note:** The following task can only be performed by a Grafana Labs employee.

Before Hugo performs a build, _Hugo mounts_ mount a source directory at a destination directory.
Use Hugo mounts to reuse a whole directory of content in another part of the published technical documentation.

## Before you begin

- Identify a directory of content that you want to reuse.
- Identify any pages to exclude.
  Typically, excluded pages do not make sense in the context of the target directory.
  > **Warning:** Pages in the directory might have relative links to other pages that would be broken by their exclusion.
  > If the excluded page is replaced by another in the target directory, relative links will continue to work.
  > Otherwise, the broken relative links should be made absolute.
  > That way, they will always refer to the page in the source directory.

## Steps

To reuse a shared directory:

1. Determine the source and target directories.

   The source directory is the path to a directory in the `website` repository.
   It always has the path prefix `content/docs/`.
   The source directory for all Grafana Cloud content at the URL https://grafana.com/docs/grafana-cloud/ is `content/docs/grafana-cloud`.

   Similarly, the target directory has the path prefix `content/docs`.
   For the destination URL `https://grafana.com/docs/target-directory/`, the path is `content/docs/target-directory`.

1. Update the website repository Hugo configuration.

   The configuration is in the in the `website` repository in the `config/_default/config.yaml` file.

   Append the mount to the `docs.manual_mount` key in the YAML configuration.

   For example, mount the source `content/docs/source-directory` at `content/docs/target-directory` and exclude the root `_index.md` file:

   ```yaml
   - source: content/docs/source-directory
     target: content/docs/target-directory
     excludeFiles:
       - /_index.md
   ```
