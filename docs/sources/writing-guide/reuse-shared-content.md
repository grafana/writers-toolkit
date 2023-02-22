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

1. Identify the sharing project.

   If the shared content is to be reused in a single project, then the sharing project is that project.
   However, if shared content is to be reused across multiple projects, for example Tempo and Grafana Enterprise Traces, then you must choose which project shares the content.

   Prefer sharing from open source projects.

1. In the sharing project, create the `docs/sources/shared/` directory if it does not exist.

1. In the sharing project, create the file `docs/sources/shared/index.md` with the following contents:

   ```markdown
   ---
   headless: true
   ---
   ```

   The `index.md` file tells Hugo that the `docs/sources/shared` directory is a leaf bundle which is necessary for the `docs/shared` shortcode to access files stored within.

   For more information about leaf bundles, refer to [Pages and page bundles]({{< relref "./documentation-structure/#pages-and-page-bundles" >}}).

1. In the sharing project, create a file for the shared content in the `docs/sources/shared/` directory.

   The file should be named to reflect its contents.

   The file can be in a subdirectory, but that subdirectory must be in the `docs/sources/shared/` directory.
   The `docs/shared` shortcode cannot lookup files outside of this directory.

   The contents of the file should be the chunk of writing you wish to reuse.

1. In the consuming project, use the `docs/shared` shortcode to include the shared content.

   To consume the file `docs/sources/shared/common-introduction.md` from the latest version of the `tempo` project documentation, the shortcode would be the following:

   ```markdown
   {{</* docs/shared source="tempo" lookup="common-introduction.md" version="latest" */>}}
   ```

   For more information about the `docs/shared` shortcode parameters, refer to [docs/shared shortcode]({{< relref "./shortcodes/#docsshared-shortcode" >}}).

1. Verify the include.

   To review the changes to the documentation, refer to the steps in [Review your changes]({{< relref "./contribute-documentation/#Review-your-changes" >}})
