---
title: Shortcodes
menuTitle: Shortcodes
description: Understand what shortcodes are and how to use them in your Markdown.
weight: 500
aliases:
  - /docs/writers-toolkit/writing-guide/shortcodes/
  - /docs/writers-toolkit/write/shortcodes/
keywords:
  - Hugo
  - shortcodes
---

# Shortcodes

Shortcodes are predefined templates used for rendering snippets in Hugo.

## Why use shortcodes?

Markdown is limited in its ability to render complex elements. Although you might be tempted to insert HTML directly into content to make up for its limitations, you can instead use shortcodes to ensure consistency across the Grafana website.

The following sections describe shortcodes available for use in Grafana Markdown files. To learn about other shortcodes, refer to the Hugo [shortcode documentation](https://gohugo.io/content-management/shortcodes/).

{{% admonition type="note" %}}
The Grafana shortcode templates are defined in the `layouts/shortcodes` folder of the website repository which is only accessible to Grafana Labs employees.
To request custom shortcodes, [create an issue](https://github.com/grafana/writers-toolkit/issues).
{{% /admonition %}}

## `admonition` shortcode

The `admonition` shortcode renders its content in a blockquote or stylized banner.
The style depends on the admonition type as defined in Writers' Toolkit [Style conventions]({{< relref "../style-guide/style-conventions" >}}).

The content of the admonition must be within opening and closing tags.

| Parameter | Description                                                           | Required |
| --------- | --------------------------------------------------------------------- | -------- |
| `type`    | The type of admonition. One of `"note"`, `"caution"`, or `"warning"`. | yes      |

### Example

The following snippet renders an admonition of _type_ `"note"` with the message `Kingston is the capital of Jamaica`.

```markdown
{{%/* admonition type="note" */%}}
Kingston is the capital of Jamaica.
{{%/* /admonition */%}}
```

## `docs/experimental-deployment` shortcode

The `docs/experimental-deployment` shortcode produces a note admonition with the preferred copy for explaining that the described deployment is experimental.

It takes no parameters.

### Example

```markdown
{{</* docs/experimental-deployment */>}}
```

Produces:

{{< docs/experimental-deployment >}}

## `docs/experimental` shortcode

The `docs/experimental` shortcode produces a note admonition with the preferred copy for explaining that the described product or feature is experimental.

| Parameter     | Description                                                              | Required |
| ------------- | ------------------------------------------------------------------------ | -------- |
| `product`     | The name of the product or feature.                                      | yes      |
| `featureFlag` | The name of the feature flag users use to enable the product or feature. | yes      |

### Example

```markdown
{{</* docs/experimental product="experimental-feature" featureFlag="its-feature-flag" */>}}
```

Produces:

{{< docs/experimental product="experimental-feature" featureFlag="its-feature-flag" >}}

## `docs/private-preview` shortcode

The `docs/private-preview` shortcode produces a note admonition with the preferred copy for explaining that the described product or feature is in private preview.

| Parameter | Description                         | Required |
| --------- | ----------------------------------- | -------- |
| `product` | The name of the product or feature. | yes      |

```markdown
{{</* docs/private-preview product="private-preview-feature" */>}}
```

Produces:

{{< docs/private-preview product="private-preview-feature" >}}

## `docs/public-preview` shortcode

The `docs/public-preview` shortcode produces a note admonition with the preferred copy for explaining that the described product or feature is in public preview.

| Parameter | Description                         | Required |
| --------- | ----------------------------------- | -------- |
| `product` | The name of the product or feature. | yes      |

```markdown
{{</* docs/public-preview product="public-preview-feature" */>}}
```

Produces:

{{< docs/public-preview product="public-preview-feature" >}}

## `docs/shared` shortcode

The `docs/shared` shortcode lets you reuse content across the Grafana website by including shared pages from source content repositories. The source content repository must explicitly share the page by placing it into its `shared` directory.

To share content, follow these steps:

1. Create a Markdown file containing the source to be shared and include `headless: true` in the front matter to prevent the website from publishing the page.
1. Store the file in a shared folder.
1. To include the shared content in a Markdown file, insert the `docs/shared` shortcode with the following named parameters:

| Parameter     | Description                                                                                                                                                                                                                                                                             | Required |
| ------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| `lookup`      | Path to the included content relative to the root of the shared directory.                                                                                                                                                                                                              | yes      |
| `source`      | Name of the source content as shown on the website. For example, for https://grafana.com/docs/enterprise-metrics/ content, the _source_ is `enterprise-metrics`.                                                                                                                        | yes      |
| `version`     | Version of the source content to include. For unversioned source content, use the empty string `""` as the value. Version inference is supported using values like `<GRAFANA VERSION>`. To learn about version inference, refer to [About version inference](#about-version-inference). | yes      |
| `leveloffset` | Manipulates source content headings up to a maximum level of `h6`. Only positive offsets are currently supported. `leveloffset="+5"` ensures an `h1` in the source content is an `h6` in the destination content.                                                                       | no       |

{{% admonition type="note" %}}
Hugo doesn't rebuild the destination file when a source file changes on disk.
To trigger a rebuild after changes to a source file, perform a trivial change to the destination file and save that, too.
{{% /admonition %}}

### Examples

The following shortcode inserts the content from the `oauth2-block.md` file. The _lookup_ path is relative to the `shared` folder in the `agent` source repository.

```markdown
{{</* docs/shared lookup="flow/reference/components/oauth2-block.md" source="agent" version="<AGENT VERSION>" */>}}
```

The following shortcode inserts the latest version of `shared-page.md` from the `shared` folder in the `enterprise-metrics` project.
Headings are offset by one level, so if the source content contains an `h1`, the resulting heading is an `h2`.

```markdown
{{</* docs/shared lookup="shared-page.md" source="enterprise-metrics" version="<GEM VERSION>" leveloffset="+1" */>}}
```

## `figure` shortcode

The `figure` shortcode renders an image with a caption using an HTML [`<figure>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure#usage_notes) element. This shortcode allows you more control over how an image is rendered, but if you don't need these options, you can use [basic Markdown to add images]({{< relref "../markdown-guide#images" >}}).

To add a figure, insert the `figure` shortcode with the following named parameters:

| Parameter      | Description                                                                                                                                                                                                                                                                             | Required |
| -------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| `alt`          | If set, `alt` specifies the alt text for the image.                                                                                                                                                                                                                                     | no       |
| `animated-gif` | If set, the HTML contains a div with an image link instead of a `<figure>` element. It's typically used for animated screenshots. Shortcode parameters other than the _caption_ and _maxWidth_ parameters are ignored.                                                                  | no       |
| `caption`      | Describes the figure using a [`<figcaption>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption) element.                                                                                                                                                            | no       |
| `class`        | Can be optionally used to override the HTML class for the `<figure>` element.                                                                                                                                                                                                           | no       |
| `lazy`         | If set to `"false"`, an additional `lazyload` class is **not** applied to the image. The `lazyload` class lets a browser render a page before the figure image loads. Once the image loads, the placeholder box transitions to the loaded image. Defaults to `"true"`.                  | no       |
| `lightbox`     | If set to `"true"`, an additional `figure-wrapper__lightbox` class is applied to the `<figure>`.                                                                                                                                                                                        | no       |
| `link`         | If set the value overrides the `src` shortcode parameter as the value to the `href` in the `<a>` element in the `<figure>`.                                                                                                                                                             | no       |
| `max-width`    | If set, `_max-width_` controls the maximum width of the `<figure>` using the [`max-width`](https://developer.mozilla.org/en-US/docs/Web/CSS/max-width) CSS property. When specifying a length or percentage, value should include unit of measurement, for example `"75px"` or `"25%"`. | no       |
| `show-caption` | If set to `"true"`, the rendered `<figure>` includes a `<figcaption>` element with the caption set in _caption_. Defaults to `"true"`.                                                                                                                                                  | no       |
| `src`          | Sets the source of the image.                                                                                                                                                                                                                                                           | yes      |

### Example

In this example, the image has a CSS class that makes the image display floated to the right.

```markdown
{{</* figure class="float-right"  src="/static/img/docs/grafana-cloud/k8sPods.png" caption="Pod view in Grafana Kubernetes Monitoring" */>}}
```

In this example, the image's display size is changed to have a maximum width of 50%. The `max-width` value must have a unit of measurement, such as pixels or percentages.

```markdown
{{</* figure max-width="50%" src="/static/img/docs/grafana-cloud/k8sPods.png" caption="Pod view in Grafana Kubernetes Monitoring" */>}}
```

## `responsive-table` shortcode

The `responsive-table` shortcode wraps the table within the shortcode tags with a class that makes the table responsive to the browser window.
This results in a table with horizontal scrolling that is fixed to the width of the containing element.

Without the `responsive-table` shortcode, a table can often overflow its containing element and text can be hidden by neighboring elements like the table of contents.

### Example

```markdown
{{%/* responsive-table */%}}
| Heading, column one | Heading, column two |
| ------------------- | ------------------- |
| Row one, column one | Row one, column two |
{{%/* /responsive-table */%}}
```

## `section` shortcode

The `section` shortcode renders an unordered list of links to a page's child pages. To add a section, insert the `section` shortcode with the following optional parameters:

| Parameter          | Description                                                                                                                                                                                                                                                | Required |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| `menuTitle`        | If set to `"true"`, the _menuTitle_ parameter modifies the template to use the child page's `menuTitle` front matter instead of the page title as the text in the link. If the child page doesn't have a `menuTitle` parameter, the title is used instead. | no       |
| `ordered`          | If set to `"true"`, the _ordered_ parameter modifies the template to use an ordered list instead of an unordered list, displaying each item with a number marker                                                                                           | no       |
| `withDescriptions` | If set to `"true"`, the _withDescriptions_ parameter modifies the template to include the front matter descriptions for child pages that have them.                                                                                                        | no       |

### Examples

The following shortcode inserts a list of links to child pages.
The links are named using the value of `menuTitle` from the front matter of the child page.

```markdown
{{</* section menuTitle="true"*/>}}
```

The following shortcode inserts a lists of links to child pages and includes the `description` content from the front matter of each child page.

```markdown
{{</* section withDescriptions="true"*/>}}
```

## `docs/reference` shortcode

The `docs/reference` shortcode offers more flexible linking than the Hugo built-in `relref` shortcode.

Use this shortcode when content from one repository is published to more than one documentation set, because it lets you specify appropriate links for each doc set in one part of the file (usually at end of the file, like a footer) while using the link label in the body text.

For example, a page in versioned Grafana documentation is also mounted in the Grafana Cloud documentation.
The page in Grafana should link to the Grafana dashboards page but the page in Grafana Cloud should link to the Grafana Cloud dashboards page.

Set the reference at the end of the page as follows:

```markdown
{{%/* docs/reference */%}}
[dashboards]: "/docs/grafana/ -> /docs/grafana/<GRAFANA VERSION>/dashboards"
[dashboards]: "/docs/grafana-cloud/ -> /docs/grafana-cloud/dashboards"
{{%/* /docs/reference */%}}
```

The content within the shortcode tags is as follows:

- `label` - The label you'll use in the reference-style links in the file. In the example above, the label is `dashboards`. The label can be multiple words (for example, [dashboard docs]) and can include spaces.
- `project path prefix` - Designates the target project. In the example above, the path prefixes are `/docs/grafana/` for Grafana and `/docs/grafana-cloud/` for Cloud.
- `reference` - The path to the destination file. Version inference is supported using values like `<GRAFANA VERSION>`. To learn about version inference, refer to [About version inference](#about-version-inference).

Then add the link in the body of the file in the following format:

```markdown
For more information about Grafana dashboards, refer to the [Dashboards documentation][dashboards].
```

- If the page you're on is `/docs/grafana/latest/alerting/`, the inferred version is `latest`, and the returned reference is `/docs/grafana/latest/dashboards`.
- If the page you're on is `/docs/grafana/next/alerting/`, the inferred version is `next`, and the returned reference is `/docs/grafana/next/dashboards`.

### Other use cases

The `docs/reference` shortcode is also useful when you want to link to the same destination multiple times in one file.
It allows you to specify the link destination once while you use the label multiple times. For example:

**Reference:**

```markdown
{{%/* docs/reference */%}}
[Grafana website]: "/ -> www.grafana.com"
{{%/* /docs/reference */%}}
```

**Body text:**

```markdown
Find more information on [Grafana][Grafana website].
```

## Escaping Hugo shortcodes

If you need to display the syntax for a shortcode, you can escape it using this syntax:

![Escaped shortcode](./writers-toolkit-escaped-shortcode.png)

This Markdown renders as:

```markdown
{{</* myshortcode */>}}
```

## About version inference

Version inference enables the use of absolute paths that resolve correctly, irrespective of version.
It uses special syntax using angle bracket delimiters like `<GRAFANA VERSION>`.

As a convention, use the name of the target project, with spaces but no hyphens or underscores, all upper-case.
For example, `grafana` becomes `GRAFANA`, `grafana-cloud` becomes GRAFANA CLOUD.

The special syntax `<SOMETHING VERSION>`, is inferred from the version of the page.

You can override version inference by including additional metadata in the front matter of the file.
To override the value of `<GRAFANA VERSION>`, set the `GRAFANA VERSION` parameter in the page's front matter.
For example, with the front matter `GRAFANA VERSION: next`, the shortcode replaces `<GRAFANA VERSION>` with `next` irrespective of the source content version.
