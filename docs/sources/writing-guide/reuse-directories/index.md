---
description: Reuse directories of content with Hugo mounts
keywords:
  - content reuse
  - shared content
  - Hugo
menuTitle: Reuse directories of content
title: Reuse directories of content with Hugo mounts
---

# Reuse directories of content with Hugo mounts

> **Note:** The following task can only be performed by a Grafana Labs employee.

Hugo mounts mount a source directory at a destination directory before the Hugo build.
Use Hugo mounts to reuse a whole directory of content in another part of the website.

## Before you begin

- Identify a directory of content that should be reused.
- Identify any pages to exclude.
  Typically this is used to exclude pages that do not make sense in the context of the target directory.
  > **Note:** Excluding files breaks any relative links to that page unless there is an equivalent in the target directory.

## Steps

To reuse a shared directory:

1. Determine the source and target directories.

   The source directory is the path to a directory in the `website` repository.
   It always has the path prefix `content/docs/`.
   The source directory for all Grafana Cloud content at the URL https://grafana.com/docs/grafana-cloud/ is `content/docs/grafana-cloud`.

   Similarly, the target directory always has the path prefix `content/docs`.
   For the destination URL https://grafana.com/docs/target-directory/, the path is `content/docs/target-directory`.

1. Update the website repository Hugo configuration.

   The configuration is in the `config/_default/config.yaml` file in the `website` repository.

   Append the mount to the `docs.manual_mount` key in the YAML configuration.

   For example, to mount the source `content/docs/source-directory` at `content/docs/target-directory`, ignoring the root `_index.md` file:

   ```yaml
   - source: content/docs/source-directory
     target: content/docs/target-directory
     excludeFiles:
       - /_index.md
   ```
