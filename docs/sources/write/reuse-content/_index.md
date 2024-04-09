---
aliases:
  - /docs/writers-toolkit/write/reuse-content
date: 2022-02-22
description: Learn about reusing content in your documentation.
keywords:
  - sharing
  - reuse content
menuTitle: Reuse
title: Reuse content
weight: 800
---

# Reuse content

Grafana's publication process can handle reusing content from snippets to whole directories.

There are some limited cases where it may be useful to reuse content in more than one location.
For example, a procedure for adding a panel to a dashboard applies to an open source and enterprise product.

## Before you begin

Before you begin reusing content, it's important to understand the tradeoffs.
Reusing content introduces maintenance overheads that can be greater than the overhead of maintaining separate content.

You can reuse content in the following ways:

1. Link to definitive content
1. Share chunks of content
1. Mount directories of content

The following sections discuss each option in detail.

### Link to definitive content

Instead of directly including content in your documentation, you can link to the definitive source.

#### Pros

- You can reuse content immediately.

#### Cons

- You have no control over the content.

  If the definitive content changes, your link may need to be updated.

- The linked content might lead to a more complex page that could potentially confuse readers.

- Users might lose context by following the link to another page or section.

  This concern hasn't been validated with UX research.

### Share chunks of content

You can reuse chunks of content using the `docs/shared` shortcode.
To do this, refer to [Reuse shared content](https://grafana.com/docs/writers-toolkit/write/reuse-content/reuse-shared-content/).

#### Pros

- You can reuse small amounts of content.
- You can maintain the content in one place.
- You can compose the shared content with other content.

#### Cons

- You need to perform the initial setup to create the shared content.

### Mount directories of content

You can reuse whole directories of content using Hugo mounts.
To do this, refer to [Reuse directories of content with Hugo mounts](https://grafana.com/docs/writers-toolkit/write/reuse-content/reuse-directories/).

With this option, you share content and structure.
Each page has exactly the same heading structure.
If you want flexibility in your page structure, you need to use another option instead.

#### Pros

- You can reuse large amounts of content with minimal effort.
- You can maintain the content in one place.

#### Cons

- There is no conditional text within pages.

  Conditional text complicates the source file and makes it harder for external contributors to understand, so it isn't implemented.

- You have to maintain two destinations for each link.

  Notably, you want to stay within each documentation set for content reused between OSS and Grafana Cloud documentation rather than have the reused content link back to the OSS documentation.

  To do so, use [`ref` URIs](https://grafana.com/docs/writers-toolkit/write/links/#link-from-source-content-thats-used-in-multiple-projects).

  If both pages link to the same place, then that link probably shouldn't exist in the reused documentation, and this indicates that the documentation isn't appropriate for reuse.

- You may have to exclude certain files that aren't appropriate in the mount destination.
