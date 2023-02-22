---
title: Reuse shared content
description: How to reuse chunks of content between pages
keywords:
  - content reuse
  - shared content
---

# Reuse shared content

Shared content is a sentence, paragraph, or page that makes sense in multiple contexts.
Reusing shared content propagates changes from a single source file to one or more destination pages.

This topic describes how to extract and share a chunk of content to multiple pages.

## Before you begin

- Identify a chunk of content that should be reused in multiple pages.

## Steps

To reuse shared content:

1. Identify the sharing and consuming projects.

   When sharing content within a single project, that project is both the sharing and consuming project.   
   
   However, when sharing content from one project to another then you must choose which is the sharing project and which is the consuming project.
   Because we rely on external contributions, sharing from open-source projects is preferred.
   For example, when sharing content between Tempo and Grafana Enterprise Traces, prefer Tempo to be the sharing project and Grafana Enterprise Traces to be the consuming project.

1. In the sharing project, create the `docs/sources/shared/` directory if it does not exist.

1. In the sharing project, create the file `docs/sources/shared/index.md` if it does not exist, with the following contents:

   ```markdown
   ---
   headless: true
   ---
   ```

   The `index.md` file tells Hugo that the `docs/sources/shared` directory is a leaf bundle which is necessary for the `docs/shared` shortcode to access files stored within.
   For more information about leaf bundles, refer to [Pages and page bundles]({{< relref "./documentation-structure/#pages-and-page-bundles" >}}).

   After performing the preceding instructions for the first time in a new repository, the directory structure looks similar to the following:

   ```console
   $ tree docs/sources/shared
   docs/sources/shared
   └── index.md

   0 directories, 1 file
   ```

1. In the sharing project, create a file for the shared content in the `docs/sources/shared/` directory.

   The file should be named to reflect its contents.

   The file can be in a subdirectory, but that subdirectory must be in the `docs/sources/shared/` directory.
   The `docs/shared` shortcode cannot lookup files outside of this directory.

   The file's contents should be the chunk of writing you want to reuse.

1. In the consuming project, use the `docs/shared` shortcode to include the shared content.

   To consume the file `docs/sources/shared/common-introduction.md` from the latest version of the `tempo` project documentation, the shortcode would be the following:

   ```markdown
   {{</* docs/shared source="tempo" lookup="common-introduction.md" version="latest" */>}}
   ```

   For more information about the `docs/shared` shortcode parameters, refer to [docs/shared shortcode]({{< relref "./shortcodes/#docsshared-shortcode" >}}).

1. Verify the include.

   To review the changes to the documentation, refer to the steps in [Review your changes]({{< relref "./contribute-documentation/#Review-your-changes" >}})
