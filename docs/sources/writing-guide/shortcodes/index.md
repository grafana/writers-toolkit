---
title: Shortcodes
menuTitle: Shortcodes
description: Understand what shortcodes are and how to use them in your markdown.
aliases:
  - /docs/writers-toolkit/latest/writing-guidelines/shortcodes/
weight: 800
keywords:
  - Hugo
  - shortcodes
---

# Shortcodes

Shortcodes are predefined templates used for rendering snippets in Hugo. 

## Why use shortcodes?

Markdown is limited in its ability to render complex elements. You might be tempted to insert HTML directly into content to make up for its limitations, but instead use shortcodes to ensure consistency across the Grafana repos.

The following sections describe shortcodes you can use in your markdown files. To learn about other shortcodes, refer to the Hugo [shortcode documentation]({{< relref "https://gohugo.io/content-management/shortcodes/" >}}).

## docs/shared shortcode

The `docs/shared` shortcode lets you reuse content across a site. To do so, you create a markdown file for sharing and store it in a shared folder. Then you insert the content into other markdown files using the `docs/shared` shortcode. 



Reuse content create a docs/shared folder: 

> **Note:** Hugo doesn't rebuild the destination file when a source file changes on disk.
> To trigger a rebuild after changes to a source file, perform a trivial change to the destination file and save that too.
`docs/shared` includes content from shared pages in source content repositories.

The source content repository must have explicitly shared the page by placing it into its shared directory.
Pages in the shared directory should set `headless: true` in the front matter to prevent the website publishing the page.

`docs/shared` has multiple named parameters:

- **lookup**: Path to the included content relative to the root of the shared directory.
- **source**: Name of the source content as shown on the website.
- **version**: Version of the source content to include.
  If not provided, the version is implicitly set to match the version of the destination content.
  If the including destination is at version `1.0.0`, then the version of included content is `1.0.0` also.
- **leveloffset**: Manipulates source content headings up to a maximum level of `h6`.
  Only positive offsets are currently supported.
  `leveloffset="+5"` ensures an `h1` in the source content is an `h6` in the destination content.

For example, to include the latest version of a page "shared-page.md", shared Grafana Enterprise Metrics, offsetting the headings by one level:

```markdown
{{</* docs/shared lookup="shared-page.md" source="enterprise-metrics" version="latest" leveloffset="+1" */>}}
```

## figure shortcode

`figure` renders an image with a caption using an HTML [`<figure>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure#usage_notes) element.

- **animated-gif**: If set, the HTML contains a div with an image link instead of a `<figure>` element.
  It's typically used for animated screenshots.
  Shortcode parameters other than the `caption` and `maxWidth` parameters are ignored.
- **caption**: Describes the figure using a [`<figcaption>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption) element.
- **class**: Can be optionally used to override the HTML class for the `<figure>` element.
- **lazy**: If set to `"false"`, the `lazyload` class is **not** also applied to the image.
  The `lazyload` class lets a browser render a page before the figure image loads.
  Once the image loads, the placeholder box transitions to the loaded image.
  Defaults to `"true"`.
- **lightbox**: If set to `"true"` the `figure-wrapper__lightbox` class is also applied to the `<figure>`.
- **link**: If set the value overrides the `src` shortcode parameter as the value to the `href` in the `<a>` element in the `<figure>`.
- **maxWidth**: If set, `maxWidth` controls the maximum width of the `<figure>` using the [`max-width`](https://developer.mozilla.org/en-US/docs/Web/CSS/max-width) CSS property. When specifying a length or percentage, value should include unit of measurement (e.g. '75px' or '25%').
- **showCaption**: If set to `"true"`, the rendered `<figure>` includes a `<figcaption>` element with the caption set in `caption`.
  Defaults to `"true"`.
- **src**: Sets the source of the image.

## section shortcode

`section` renders an unordered list of links to a page's child pages.

- **menuTitle**: If set to `"true"`, the menuTitle parameter modifies the template to use the `menuTitle` parameter of a child page's front matter instead of the page title as the text in the link.
  If the child page doesn't have a `menuTitle` parameter, the title is used instead.