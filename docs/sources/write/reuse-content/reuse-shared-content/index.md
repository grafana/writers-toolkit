---
aliases:
  - /docs/writers-toolkit/writing-guide/reuse-shared-content/
  - /docs/writers-toolkit/write/reuse-content/reuse-shared-content/
date: "2023-02-22T16:13:40+00:00"
description: Learn to reuse chunks of content between pages.
keywords:
  - content reuse
  - shared content
review_date: "2024-06-27"
title: Reuse shared content
---

# Reuse shared content

Shared content is a sentence, paragraph, or page that makes sense in multiple contexts.
Reusing shared content propagates changes from a single source file to one or more destination pages.

This topic describes how to extract and share a chunk of content to multiple pages.

## Before you begin

- Identify a chunk of content that you want to reuse in multiple pages.

## Steps

To reuse shared content, follow these steps:

### Create a shared directory

1. Identify the sharing and consuming projects.

   - When sharing content within a single project, that project is both the sharing and consuming project.

   - However, when sharing content from one project to another then you must choose which is the sharing project and which is the consuming project.
     Because Grafana Labs values external contributions, prefer to share from an open source projects.
     For example, when sharing content between Tempo and Grafana Enterprise Traces, prefer Tempo to be the sharing project and Grafana Enterprise Traces to be the consuming project.

1. In the sharing project, create the `docs/sources/shared/` directory if it doesn't exist.

1. In the sharing project, create the file `docs/sources/shared/index.md` if it doesn't exist, with the following contents:

   ```markdown
   ---
   headless: true
   ---
   ```

   The `index.md` file tells Hugo that the `docs/sources/shared` directory is a leaf bundle which is necessary for the `docs/shared` shortcode to access files stored within.
   For more information about leaf bundles, refer to [Pages and page bundles](https://grafana.com/docs/writers-toolkit/structure/#pages-and-page-bundles).

   After performing the preceding instructions for the first time in a repository, the directory structure looks similar to the following:

   ```console
   docs/sources/shared
   └── index.md
   ```

### Create a shared file

1. In the sharing project, create a file for the shared content in the `docs/sources/shared/` directory.

   Name the file to reflect its contents.

   The file can be in a subdirectory, but that subdirectory must be in the `docs/sources/shared/` directory.
   The `docs/shared` shortcode can't look up files outside of this directory.

   The file's contents should be the chunk of writing you want to reuse and a front matter section with a title.
   You can include the labels in the front matter if you think it would be helpful information, but the labels won't be visible in the consuming project:

   ```markdown
   ---
   labels:
     products:
       - cloud
       - oss
   title: A shared file
   ---
   ```

1. In the consuming project, use the `docs/shared` shortcode to include the shared content.

   To consume the file `docs/sources/shared/common-introduction.md` from the latest version of the `tempo` project documentation, the shortcode would be the following:

   ```markdown
   {{</* docs/shared source="tempo" lookup="common-introduction.md" version="latest" */>}}
   ```

   For more information about the `docs/shared` shortcode parameters, refer to the [`docs/shared` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes#docsshared) reference.

1. Verify the include.

   To review the changes to the documentation, refer to the steps in [Review your changes](https://grafana.com/docs/writers-toolkit/contribute/#review-your-changes)
