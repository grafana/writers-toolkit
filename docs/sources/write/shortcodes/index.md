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
VARIABLE: test
---

# Shortcodes

Shortcodes are snippets you use in source files to calling built-in or custom templates.
Shortcodes templates avoid the need for HTML in Markdown and ensure consistency across the documentation set.

The following sections describe shortcodes available for use in Grafana Markdown files.
To learn about other shortcodes, refer to the Hugo [shortcode documentation](https://gohugo.io/content-management/shortcodes/).

{{< admonition type="note" >}}
The website team maintains shortcode templates in the `layouts/shortcodes` folder of the website repository which is only accessible to Grafana Labs employees.
To request custom shortcodes, [create an issue](https://github.com/grafana/writers-toolkit/issues).
{{< /admonition >}}

## Admonition

The `admonition` shortcode renders its content in a blockquote or stylized banner.
The style depends on the admonition type as defined in Writers' Toolkit [Style conventions]({{< relref "../style-guide/style-conventions" >}}).

The content of the admonition must be within opening and closing tags,
and the type of admonition must be within quotes.

| Parameter | Description                                                            | Required |
| --------- | ---------------------------------------------------------------------- | -------- |
| `type`    | The type of admonition. One of `caution`, `note`, `tip`, or `warning`. | yes      |

Use a tip when you want to show the reader _how_ to do something that isn’t necessarily obvious.
Tips should be helpful, additional information.
You can think of some tips as tricks.
Your reader can feel free to skip them if they want because they don't contribute to core understanding.

{{< admonition type="warning" >}}
Reference style links such as `[link text][label]` or `[link text][]` don't work in the inner text of shortcodes if you use reference links defined at the topic level.
To use reference links within an admonition, you must use standard inline Markdown links or define the reference links within the admonition shortcode.

For more information, refer to [Markdown Reference Links in Shortcodes](https://discourse.gohugo.io/t/markdown-reference-links-in-shortcodes/5770/3).
{{< /admonition >}}

### Examples

The following example renders an admonition of type `note` with the message `Kingston is the capital of Jamaica.`:

```markdown
{{</* admonition type="note" */>}}
Kingston is the capital of Jamaica.
{{</* /admonition */>}}
```

Produces:

{{< admonition type="note" >}}
Kingston is the capital of Jamaica.
{{< /admonition >}}

The following example renders an admonition of type `tip` with the message `This also applies to headings that contain a forward slash or parentheses or square brackets.`:

```markdown
{{</* admonition type="tip" */>}}
This also applies to headings that contain a forward slash or parentheses or square brackets.
{{</* /admonition */>}}
```

Produces:

{{< admonition type="tip" >}}
This also applies to headings that contain a forward slash or parentheses or square brackets.{{< /admonition >}}

## Card grid

The `card-grid` shortcode renders a responsive grid of card elements that fits the width of its container.

### Grid parameters

| Parameter       | Description                                                                                                                                                                                                                                                   | Required |
| :-------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :------- |
| `key`           | Front matter parameter name of hero fields. Default: `hero`.                                                                                                                                                                                                  | Yes      |
| `items`         | Front matter array of card parameters                                                                                                                                                                                                                         | Yes      |
| `type`          | The type of card to use. Only current option is: `simple`. Default: `simple`.                                                                                                                                                                                 | No       |
| `min`           | Sets the minimum card width. This affects the number of cards in each row, as well as the breakpoints at which the cards wrap. Options: `xs`, `sm`, `md`, `lg`. These correspond to minimum card widths of `100px`, `250px`, `350px`, `500px`. Default: `sm`. | No       |
| `wrapper_class` | Optional CSS class for the wrapper element.                                                                                                                                                                                                                   | No       |
| `grid_class`    | Optional CSS class for the grid element.                                                                                                                                                                                                                      | No       |
| `card_class`    | Optional CSS class for the cards.                                                                                                                                                                                                                             | No       |

### Card parameters (type="simple")

| Parameter     | Description                                                                                                                                                                                                                                    | Required |
| :------------ | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------- |
| `title`       | Card title text.                                                                                                                                                                                                                               | No       |
| `href`        | URL of card target. Use relative path for links within "grafana.com" domain (e.g. `/docs/grafana/latest/`)                                                                                                                                     | Yes      |
| `description` | Description text. Accepts Markdown.                                                                                                                                                                                                            | No       |
| `logo`        | Logo image URL.                                                                                                                                                                                                                                | No       |
| `width`       | For raster images (`png`, `jpg`, `webp`), this is the image's natural width. For vector images (`svg`), this is the desired display width. Accepts a number (pixels) or a percentage. Pixel values must _not_ include `px`. Default: `auto`.   | No       |
| `height`      | For raster images (`png`, `jpg`, `webp`), this is the image's natural height. For vector images (`svg`), this is the desired display height. Accepts a number (pixels) or a percentage. Pixel values must _not_ include `px`. Default: `auto`. | No       |

### Examples

Render a card grid with a minimum card width of `sm` and a `simple` card type:

```markdown
---
my_card_grid:
  type: simple
  min: sm
  items:
    - title: Grafana Alerting
      href: /docs/grafana-cloud/alerting-and-irm/alerting/
      description: >-
        Allows you to learn about problems in your systems moments after they occur. Monitor your incoming metrics data or log entries and set up your Alerting system to watch for specific events or circumstances and then send notifications when those things are found.
      logo: /media/docs/grafana-cloud/alerting-and-irm/grafana-icon-alerting.svg
      height: 24
    - title: Grafana SLO
      href: /docs/grafana-cloud/alerting-and-irm/slo/
      description: >-
        Provides a framework for measuring the quality of service you provide to users. Use SLOs to collect data on the reliability of your systems over time and as a result, help engineering teams reduce alert fatigue, focus on reliability, and provide better service to your customers.
      logo: /media/docs/grafana-cloud/alerting-and-irm/grafana-icon-slo.svg
      height: 24
---

{{</* card-grid key="my_card_grid" */>}}
```

## Code

The `code` shortcode provides the ability to show multiple snippets of code in different languages.
When a user selects a language, the website sets other code blocks on the page to that language.
The website saves the selected language to browser storage and persists it across navigation.

### Example

{{< admonition type="note" >}}
If your repository uses `prettier` to format the files, use the HTML comments `<!-- prettier-ignore-start -->` and `<!-- prettier-ignore-end -->` around the shortcode tags to ensure correct rendering.
{{< /admonition >}}

<!-- prettier-ignore-start -->

````markdown
{{</* code */>}}

```bash
curl "https://your-stack.grafana.net/api/plugins/grafana-incident-app/resources/api/v1/ActivityService.AddActivity"
```

```go
package main

import (
  "context"
)

func main() {

  ...
```

```javascript
import { GrafanaIncidentClient, ActivityService } from '@grafana/incident-node';

// https://grafana.com/docs/grafana-cloud/incident/api/auth/#get-a-service-account-token
const serviceAccountToken = process.env.GRAFANA_CLOUD_SERVICE_ACCOUNT_TOKEN;
const client = new GrafanaIncidentClient(
  "https://your-stack.grafana.net",
  serviceAccountToken
);

...
```

```json
POST https://your-stack.grafana.net/api/plugins/grafana-incident-app/resources/api/v1/ActivityService.AddActivity
Content-Type="application/json; charset=utf-8"
Authorization: Bearer glsa_HOruNAb7SOiCdshU9algkrq7F...
```

{{</* /code */>}}
````
<!-- prettier-ignore-end -->

Produces:

<!-- prettier-ignore-start -->

{{< code >}}

```bash
curl "https://your-stack.grafana.net/api/plugins/grafana-incident-app/resources/api/v1/ActivityService.AddActivity"
```

```go
package main

import (
  "context"
)

func main() {

  ...
```

```javascript
import { GrafanaIncidentClient, ActivityService } from '@grafana/incident-node';

// https://grafana.com/docs/grafana-cloud/incident/api/auth/#get-a-service-account-token
const serviceAccountToken = process.env.GRAFANA_CLOUD_SERVICE_ACCOUNT_TOKEN;
const client = new GrafanaIncidentClient(
  "https://your-stack.grafana.net",
  serviceAccountToken
);

...
```

```json
POST https://your-stack.grafana.net/api/plugins/grafana-incident-app/resources/api/v1/ActivityService.AddActivity
Content-Type="application/json; charset=utf-8"
Authorization: Bearer glsa_HOruNAb7SOiCdshU9algkrq7F...
```

{{< /code >}}

## Collapse

The `collapse` shortcode toggles visibility of sections of content, often helpful when hiding and showing large amounts of content.

| Parameter | Description                         | Required |
| --------- | ----------------------------------- | -------- |
| `title`   | Text explaining the hidden content. | yes      |

### Example

```markdown
{{</* collapse title="Title of hidden content" */>}}
Kingston is the capital of Jamaica.
{{</* /collapse */>}}
```

Produces:

{{< collapse title="Title of hidden content" >}}
Kingston is the capital of Jamaica.
{{< /collapse >}}

Use this shortcode for:

<!-- vale Grafana.GooglePassive = NO -->

- [Deprecated content](#deprecation-example): Features that have been deprecated, but still need to be documented for some time.
- [Configuration options](#configuration-options-example): Features that have several ways they can be configured.

<!-- vale Grafana.GooglePassive = YES -->

Add a lead-in sentence and a title that, taken together, are descriptive enough for the reader to guess what's included in the collapsed content.
Don't duplicate headings in the title parameter.

You can't do the following with this shortcode:

- Use them as page headings
- Control the size of the title text
- Add images or videos between the shortcode tags

### Deprecation example

```markdown
#### BoltDB (deprecated)

The following example is for a deprecated store and you shouldn't use it for new Loki deployments:

{{</* collapse title="boltdb-shipper" */>}}
Also known as _boltdb-shipper_ during development.
The single store configurations for Loki utilize the chunk store for both chunks and the index, requiring just one store to run Loki.
{{</* /collapse */>}}
```

Produces:

#### BoltDB (deprecated)

The following example is for a deprecated store and you shouldn't use it for new Loki deployments:

{{< collapse title="boltdb-shipper" >}}
Also known as _boltdb-shipper_ during development.
The single store configurations for Loki to utilize the chunk store for both chunks and the index, which requires just one store to run Loki.
{{< /collapse >}}

### Configuration options example

````markdown
#### `6-Compactor-Snippet.yaml`

The following partial configuration sets the compactor to use S3 and run the compaction every five minutes.

{{</* collapse title="Example" */>}}

```yaml
compactor:
  working_directory: /tmp/loki/compactor
  shared_store: s3
  compaction_interval: 5m
```

{{</* /collapse */>}}
````

Produces:

#### `6-Compactor-Snippet.yaml`

The following partial configuration sets the compactor to use Amazon S3 and runs the compaction every 5 minutes.

{{< collapse title="Example" >}}

```yaml
compactor:
  working_directory: /tmp/loki/compactor
  shared_store: s3
  compaction_interval: 5m
```

{{< /collapse >}}

## Docs/experimental-deployment

The `docs/experimental-deployment` shortcode produces a note admonition with the preferred copy for explaining that the described deployment is experimental.

It takes no parameters.

### Example

```markdown
{{</* docs/experimental-deployment */>}}
```

Produces:

{{< docs/experimental-deployment >}}

## Docs/experimental

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

## Docs/private-preview

The `docs/private-preview` shortcode produces a note admonition with the preferred copy for explaining that the described product or feature is in private preview.

| Parameter | Description                         | Required |
| --------- | ----------------------------------- | -------- |
| `product` | The name of the product or feature. | yes      |

```markdown
{{</* docs/private-preview product="private-preview-feature" */>}}
```

Produces:

{{< docs/private-preview product="private-preview-feature" >}}

## Docs/public-preview

The `docs/public-preview` shortcode produces a note admonition with the preferred copy for explaining that the described product or feature is in public preview.

| Parameter     | Description                                                              | Required |
| ------------- | ------------------------------------------------------------------------ | -------- |
| `product`     | The name of the product or feature.                                      | yes      |
| `featureFlag` | The name of the feature flag users use to enable the product or feature. | no       |

```markdown
{{</* docs/public-preview product="public-preview-feature" */>}}
```

Produces:

{{< docs/public-preview product="public-preview-feature" >}}

```markdown
{{</* docs/public-preview product="public-preview-feature" featureFlag="its-feature-flag" */>}}
```

Produces:

{{< docs/public-preview product="public-preview-feature" featureFlag="its-feature-flag" >}}

## Docs/reference

The `docs/reference` shortcode lets you specify different destinations for the same link that depend on where you publish the source file.
Use this shortcode for links when you reuse content from one repository to more than one documentation set.

You set all possible destinations in one part of the file, usually at end of the file, like a footer.

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

`[LABEL]: "PROJECT PATH PREFIX -> REFERENCE"`

- _`LABEL`_ - The label you'll use in the reference-style links in the file.
  In the preceding example, the label is `dashboards`.
  The label can be multiple words like `[dashboard docs]` and can include spaces.

- _`PROJECT PATH PREFIX`_ - Designates the target project.
  In the preceding example, the path prefixes are `/docs/grafana/` for Grafana and `/docs/grafana-cloud/` for Cloud.

- _`REFERENCE`_ - The path to the destination file.
  This shortcode supports version substitution using values like `<GRAFANA VERSION>`.
  To learn about version substitution, refer to [About version substitution](#about-version-substitution).
  Don't include trailing slashes in the path.

Then add the link in the body of the file in the following format:

```markdown
For more information about Grafana dashboards, refer to [Dashboards][dashboards].
```

- If the page you're on is `/docs/grafana-cloud/alerting/`, the returned link is `/docs/grafana-cloud/dashboards/`.
- If the page you're on is `/docs/grafana/latest/alerting/`, the inferred version is `latest`, and the returned link is `/docs/grafana/latest/dashboards/`.
- If the page you're on is `/docs/grafana/next/alerting/`, the inferred version is `next`, and the returned link is `/docs/grafana/next/dashboards/`.

### Other use cases

Markdown reference-style links are also useful when you want to link to the same destination multiple times in one file.
It allows you to specify the link destination once while you use the label multiple times.
For example:

**Reference:**

```markdown
[Grafana website]: www.grafana.com
```

**Body text:**

```markdown
Find more information on [Grafana][Grafana website].
```

## Docs/shared

The `docs/shared` shortcode lets you reuse content across the Grafana website by including shared pages from source content repositories.
The source content repository must explicitly share the page by placing it into its `shared` directory.

To share content, follow these steps:

1. Create a Markdown file containing the shared source and include `headless: true` in the front matter to prevent the website from publishing the page.
1. Store the file in a shared folder.
1. To include the shared content in a Markdown file, insert the `docs/shared` shortcode with the following named parameters:

For more detailed instructions, refer to [Reuse shared content](https://grafana.com/docs/writers-toolkit/write/reuse-content/reuse-shared-content).

| Parameter     | Description                                                                                                                                                                                                                                                                                                                    | Required |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| `lookup`      | Path to the included content relative to the root of the shared directory.                                                                                                                                                                                                                                                     | yes      |
| `source`      | Name of the source content as shown on the website. For example, for https://grafana.com/docs/enterprise-metrics/ content, the _source_ is `enterprise-metrics`.                                                                                                                                                               | yes      |
| `version`     | Version of the source content to include. For source content that doesn't have a version, use the empty string `""` as the value. This shortcode supports version substitution using values like `<GRAFANA VERSION>`. To learn about version substitution, refer to [About version substitution](#about-version-substitution). | yes      |
| `leveloffset` | Manipulates source content headings up to a maximum level of `h6`. Only positive offsets are currently supported. `leveloffset="+5"` ensures an `h1` in the source content is an `h6` in the destination content.                                                                                                              | no       |

{{< admonition type="note" >}}
Hugo doesn't rebuild the destination file when a source file changes on disk.
To trigger a rebuild after changes to a source file, perform a trivial change to the destination file and save that, too.
{{< /admonition >}}

### Examples

The following shortcode inserts the content from the `oauth2-block.md` file.
The _`lookup`_ path is relative to the `shared` folder in the `agent` source repository.

```markdown
{{</* docs/shared lookup="flow/reference/components/oauth2-block.md" source="agent" version="<AGENT VERSION>" */>}}
```

The following shortcode inserts the latest version of `shared-page.md` from the `shared` folder in the `enterprise-metrics` project.
Headings are offset by one level, so if the source content contains an `h1`, the resulting heading is an `h2`.

```markdown
{{</* docs/shared lookup="shared-page.md" source="enterprise-metrics" version="<GEM VERSION>" leveloffset="+1" */>}}
```

## Figure

The `figure` shortcode renders an image with a caption using an HTML [`<figure>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure#usage_notes) element.
This shortcode allows you more control over how to render an image, but if you don't need these options, you can use [basic Markdown to add images]({{< relref "../markdown-guide#images" >}}).

To add a figure, insert the `figure` shortcode with the following named parameters:

| Parameter       | Description                                                                                                                                                                                                                                                                                     | Required |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| `alt`           | If set, `alt` specifies the alt text for the image.                                                                                                                                                                                                                                             | no       |
| `animated-gif`  | If set, the HTML contains a div with an image link instead of a `<figure>` element. It's typically used for animated screenshots. When you set this parameter, the shortcode ignores other parameters except _`caption`_ and _`maxWidth`_.                                                      | no       |
| `caption`       | Describes the figure using a [`<figcaption>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption) element. If you don't set `alt`, the shortcode uses `caption` as the alt text.                                                                                              | no       |
| `caption-align` | Change the alignment of the `caption` property. Accepted values are `left`, `center`, and `right`.                                                                                                                                                                                              | no       |
| `class`         | Override the HTML class for the `<figure>` element.                                                                                                                                                                                                                                             | no       |
| `link-class`    | Override the HTML class for the `<a>` element.                                                                                                                                                                                                                                                  | no       |
| `lazy`          | If set to `"false"`, an additional `lazyload` class is **not** applied to the image. The `lazyload` class lets a browser render a page before the figure image loads. After the image loads, the placeholder box transitions to the loaded image. Defaults to `"true"`.                         | no       |
| `lightbox`      | If set to `"true"`, the shortcode applies an additional `figure-wrapper__lightbox` class to the `<figure>`.                                                                                                                                                                                     | no       |
| `link`          | If set the value overrides the `src` shortcode parameter as the value to the `href` in the `<a>` element in the `<figure>`.                                                                                                                                                                     | no       |
| `height`        | If set, `_height_` controls the height of the `<img>` element using the [`height`](https://developer.mozilla.org/en-US/docs/Web/CSS/height) CSS property. When specifying a value, it should be an integer representing pixels without a `"px"` string at the end, for example `"500"`.         | no       |
| `width`         | If set, `_width_` controls the width of the `<img>` element using the [`width`](https://developer.mozilla.org/en-US/docs/Web/CSS/width) CSS property. When specifying a value, it should be an integer representing pixels without a `"px"` string at the end, for example `"500"`.             | no       |
| `max-width`     | If set, `_max-width_` controls the maximum width of the `<figure>` element using the [`max-width`](https://developer.mozilla.org/en-US/docs/Web/CSS/max-width) CSS property. When specifying a length or percentage, value should include unit of measurement, for example `"75px"` or `"25%"`. | no       |
| `show-caption`  | If set to `"true"`, the rendered `<figure>` includes a `<figcaption>` element with the caption set in _caption_. Defaults to `"true"`.                                                                                                                                                          | no       |
| `src`           | Sets the source of the image.                                                                                                                                                                                                                                                                   | yes      |

{{< admonition type="note" >}}
Including the original image dimensions as the 'width' and 'height' properties is highly recommended, as it improves page performance and SEO.

These values are _only_ used for determining the image aspect ratio and don't equate to the final displayed size.
{{< /admonition >}}

### Center an image example

You can center an image using the `figure` shortcode by adding the following properties:

- Add the `class="w-100p"` property.
- Add the `link-class="w-fit mx-auto d-flex flex-direction-column"` property.
- Add the `max-width="WIDTHpx"` property, replacing `WIDTH` with the value you want for your image. The value should be lower than the original image's width.
- Optionally, you can add the `width` and `height` properties with the original image values, without pixels or percentages. For example, for a 800x600 image, you can add `width="800" height="600"`.

This is an example for centering an image that's 1275x738 pixels:

```markdown
{{</* figure src="/static/img/docs/grafana-cloud/k8sPods.png" width="1275" height="738" max-width="500px" class="w-100p" link-class="w-fit mx-auto d-flex flex-direction-column" caption="Pod view in Grafana Kubernetes Monitoring" caption-align="center" */>}}
```

For more details about what each parameter does, refer to the section below.

{{< collapse title="Center image properties" >}}

When centering an image on a page, each property you have to add to the `<figure>` shortcode has a different purpose:

- The `class="w-100p"` property specifies that the `<figure>` element should take up 100% of the width of its container.
- The `link-class` attribute values are:
  - `w-fit` specifies that the width of the <a> element that wraps around the image should fit its contents. For example, if the image's width is 500px, the width of its parent <a> element will be set to match.
  - `mx-auto` sets the x-axis margin on either side of the link to automatically take up whatever space remains in the container. For example, if the figure's container is 750px wide and the image is 500px wide, the margin on either side of the image will be 125px ((750 - 500) / 2) . This is what actually does the centering.
  - `d-flex` sets the figure element to be a flex container.
  - `flex-direction-column` specifies the direction of the flex container's layout. In this case, the child elements are stacked vertically, as opposed to flex-direction-row, which displays the child elements horizontally.
- The `width`, or `max-width` property, is necessary to calculate the difference between the image's width and the container size. While only the `width` or `max-width` value is necessary for the purposes of centering, it's best practice to set both `width` and `height`.

Setting the `width` and `height` properties is recommended because it allows the browser to optimize delivery of the image and account for the user's device and display dimensions, and service up a smaller image if the full-sized image isn't needed. Providing both values also allows the browser to correctly calculate a placeholder space when images are lazy loaded. This helps with [preventing reflow](https://css-tricks.com/preventing-content-reflow-from-lazy-loaded-images/), which is the cause of most instances where links in an article's Table of Contents don't land on the correct section.

For raster images, such as PNG or JPG files, width and height should usually be set to the original dimensions of the image. That provides the greatest flexibility for image optimization. However, for SVG images, width should typically be set to 100% and height left blank. That will expand the image to fill its container while maintaining the image's aspect ratio.

In cases where the desired display width of an image is smaller than the max width of the container, it's important to calculate the correct height so as to maintain the image's aspect ratio without warping the image. To do this, you can divide the desired image width by the image's full width and multiply by the original height: (new-width / full-width) * full-height = new-height . For example, if the original image dimensions are 1000px x 750px and you want the image displayed at 500px wide, you'd calculate (500 / 1000) * 750 = 325, or width="500" height="325".

{{< /collapse >}}

### Example

In this example, the image has a CSS class that makes the image display floated to the right.

```markdown
{{</* figure class="float-right"  src="/static/img/docs/grafana-cloud/k8sPods.png" caption="Pod view in Grafana Kubernetes Monitoring" */>}}
```

This example sets the image's display size to have a maximum width of 50%.
The `max-width` value must have a unit of measurement, such as pixels or percentages.

```markdown
{{</* figure max-width="50%" src="/static/img/docs/grafana-cloud/k8sPods.png" caption="Pod view in Grafana Kubernetes Monitoring" */>}}
```

## Hero (simple)

A hero section is a large section that contains a title, description, and image, usually placed at the top of a page.

The `hero-simple` shortcode renders a hero section with an optional title, description, and image.
To add a simple hero, insert the `hero-simple` shortcode using the following named parameters:

| Parameter           | Description                                                                                                                                                                                                                                    | Required |
| :------------------ | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------- |
| `key`               | Front matter parameter name of hero fields. Default: `hero`.                                                                                                                                                                                   | No       |
| `title`             | Heading title text.                                                                                                                                                                                                                            | No       |
| `level`             | Heading level. Default: `3`.                                                                                                                                                                                                                   | No       |
| `image`             | Image URL.                                                                                                                                                                                                                                     | No       |
| `width`             | For raster images (`png`, `jpg`, `webp`), this is the image's natural width. For vector images (`svg`), this is the desired display width. Accepts a number (pixels) or a percentage. Pixel values must _not_ include `px`. Default: `auto`.   | No       |
| `height`            | For raster images (`png`, `jpg`, `webp`), this is the image's natural height. For vector images (`svg`), this is the desired display height. Accepts a number (pixels) or a percentage. Pixel values must _not_ include `px`. Default: `auto`. | No       |
| `description`       | Description text. Accepts Markdown.                                                                                                                                                                                                            | No       |
| `wrapper_class`     | Optional CSS class for the wrapper element.                                                                                                                                                                                                    | No       |
| `hero_class`        | Optional CSS class for the hero element.                                                                                                                                                                                                       | No       |
| `img_class`         | Optional CSS class for the image container element.                                                                                                                                                                                            | No       |
| `title_class`       | Optional CSS class for the heading element.                                                                                                                                                                                                    | No       |
| `description_class` | Optional CSS class for the description element.                                                                                                                                                                                                | No       |
<!-- prettier-ignore-end -->

You can provide shortcode arguments by:

- Adding front matter and referencing it front matter with the `key` argument.
- Adding them directly to the shortcode itself.

Shortcode arguments override those in the front matter.
If you don't provide any arguments, the shortcode uses default values.

### Examples

Insert a simple hero using front matter:

```markdown
---
my_hero:
  title: Alerts and IRM
  level: 1
  image: /media/docs/grafana-cloud/alerting-and-irm/grafana-cloud-docs-hero-alerts-irm.svg
  width: 110
  height: 110
  description: >-
    Alerts & IRM is Grafana Cloud’s Incident Response Management (IRM) solution, which enables you to detect, respond, and learn from incidents in one centralized platform.
---

{{</* hero-simple key="my_hero" */>}}
```

Insert a simple hero using shortcode arguments:

```markdown
{{</* hero-simple title="Alerts and IRM" level="1" image="/media/docs/grafana-cloud/alerting-and-irm/grafana-cloud-docs-hero-alerts-irm.svg" width="110" height="110" description="Alerts & IRM is Grafana Cloud’s Incident Response Management (IRM) solution, which enables you to detect, respond, and learn from incidents in one centralized platform." */>}}
```

<!-- vale Grafana.Spelling = NO -->

## Param

<!-- vale Grafana.Spelling = YES -->

The `param` shortcode provides build-time variable substitution.

To add a new variable definition:

1. Define a [`cascade` variable](https://grafana.com/docs/writers-toolkit/write/front-matter/#cascade) in the parent topic.
1. Insert the `param` variable where it's required in the parent and child topics.

{{% admonition type="note" %}}
If you use the `param` shortcode in headings, you must use `%` in place of `<` and `>`.

For example: `{﻿{% param VARIABLE %}}`.
{{% /admonition %}}

### Example

The front matter definition for a product version:

```yaml
cascade:
  PRODUCT_VERSION: 10.2
```

The `param` shortcode in the topic body text:

```markdown
Welcome to Grafana {{</* param PRODUCT_VERSION */>}}. Read on to learn about changes to dashboards and visualizations, data sources, security and authentication, and more.
```

The `param` shortcode in a topic heading:

```markdown
## What's new in Grafana {{%/* param PRODUCT_VERSION */%}}.
```

<!-- vale Grafana.Spelling = NO -->

## Relref

<!-- vale Grafana.Spelling = YES -->

{{< admonition type="warning" >}}
This shortcode is present in the documentation, but you should prefer full URLs.
Don't use it when creating new or updating existing documentation.

For more information, refer to [Links](https://grafana.com/docs/writers-toolkit/write/links).
{{< /admonition >}}

The `relref` shortcode provides build-time link checking to ensure that the destination file exists.

For example: `{{</* relref "./path/to/file" */>}}`.

| Parameter  | Description      | Required |
| ---------- | ---------------- | -------- |
| position 0 | The file lookup. | yes      |

The argument is an absolute or relative file lookup.
An absolute lookup begins with a slash (`/`) and starts from the site root.
For example, to link to the Grafana Cloud index page, the `relref` shortcode argument is `{{</* relref "/docs/grafana-cloud/_index.md" */>}}`.

Using the full path to file, including the extension, can break the lookup if the destination file changes from a page to a leaf or branch bundle.
To avoid this, omit the file extension and any `/_index` or `/index` part of the lookup.
Using the previous example of the Grafana Cloud index page, the preferred `relref` shortcode argument is `{{</* relref "/docs/grafana-cloud" */>}}`.

Hugo link checking only applies to the content available during the build.
In most projects, the only content available during local builds and CI is the current project documentation.

{{< docs/shared source="writers-toolkit" lookup="hugo-error-example-bad-link.md" version="" >}}

For additional information about Hugo error output, refer to [Test documentation changes](https://grafana.com/docs/writers-toolkit/review/test-documentation-changes/).

### Determine `relref` shortcode arguments

To determine the path between the source and destination content, do the following:

Find the directory where the destination content lives.
Find the directory that the source and destination directories have in common.
Note the relative path from the common directory to the destination directory.
Count the number of folders from the source to the common directory and that number equals the number of parent directory path elements (`..`) you need to add to your relative path.
Join all the path elements together with forward slashes (`/`).

For example, with the following folder structure:

```
Vehicles
├── Trucks
│   ├── F150
│   └── 1999 F150
└── Vans
```

In this case, the source content is in the `1999 F150` directory and the destination content is in the `Vans` directory.
The common folder for the two pieces of content is the `Vehicles` directory.

The parent directory of `1999 F150` is `Trucks`, requiring one `..` path element.
To parent directory of `Trucks` is `Vehicles`, requiring another `..` path element.
Therefore, the relative path from the source directory, `1999 F150`, and the common directory, `Vehicles`, is `../..`

The pathway from the common directory `Vehicles` to destination directory `Vans` is `vans`
The relative path is `../../vans`

If the source directory was `Vans` and the destination was `1999 F150`, the relative path would be `../trucks/F150/1999-F150`.

## Section

The `section` shortcode renders an unordered list of links to a page's child pages. To add a section, insert the `section` shortcode with the following optional parameters:

| Parameter          | Description                                                                                                                                                                                                                                                             | Required |
| ------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| `menuTitle`        | If set to `"true"`, the _menuTitle_ parameter modifies the template to use the child page's `menuTitle` front matter instead of the page title as the text in the link. If the child page doesn't have a `menuTitle` parameter, the shortcode uses the `title` instead. | no       |
| `ordered`          | If set to `"true"`, the _ordered_ parameter modifies the template to use an ordered list instead of an unordered list, displaying each item with a number marker                                                                                                        | no       |
| `withDescriptions` | If set to `"true"`, the _withDescriptions_ parameter modifies the template to include the front matter descriptions for child pages that have them.                                                                                                                     | no       |
| `depth`            | Controls how many levels of child pages to include in the output. The default depth is `"1"` which only shows the immediate children.                                                                                                                                   | no       |

### Examples

The following shortcode inserts a list of links to child pages.
The link text uses the value of `menuTitle` from the front matter of the child page.

```markdown
{{</* section menuTitle="true"*/>}}
```

The following shortcode inserts a lists of links to child pages and includes the `description` content from the front matter of each child page.

```markdown
{{</* section withDescriptions="true"*/>}}
```

## Table of contents

The `table-of-contents` shortcode renders the page's table of contents in the page body.
You should generally avoid using this shortcode because every documentation page already has a table of contents rendered with the page body.

```markdown
{{</* table-of-contents */>}}
```

## Term

The `term` shortcode enables a tooltip when a user hovers over text surrounded by the shortcode.

| Parameter  | Description         | Required |
| ---------- | ------------------- | -------- |
| position 0 | Glossary lookup key | yes      |

### Examples

```markdown
{{</* term "data source" */>}}data sources{{</* /term */>}}
```

```markdown
Grafana comes with built-in support for many {{</* term "data source" */>}}data sources{{</* /term */>}}.
```

Produces:

Grafana comes with built-in support for many {{< term "data source" >}}data sources{{< /term >}}.

The Grafana Labs technical documentation team maintains the associated definitions internally.
If you are a Grafana Labs employee and want to make changes, edit [`glossary.yaml`](https://github.com/grafana/website/blob/master/data/glossary.yaml).
If you aren't a Grafana Labs employee, request changes by [creating an issue](https://github.com/grafana/writers-toolkit/issues/new).

### Guidance

For terms with multiple definitions, follow the common dictionary practice of numbering each alternative.
For an example, refer to the definition of [graph](https://www.dictionary.com/browse/graph).

## Vimeo

The `vimeo` shortcode embeds videos hosted on vimeo.com.

| Parameter  | Description | Required |
| ---------- | ----------- | -------- |
| position 0 | Video ID    | yes      |

The shortcode requires a single argument which is the video ID.
For example, `{{</* vimeo 1111111*/>}}`.

You can find the video ID at the end of the URL.
In this example, the video is a Preview of Tempo 2.0 and TraceQL: `https://vimeo.com/773194063`.

```
{{</* vimeo 773194063 */>}}
```

## YouTube

The `youtube` shortcode embed videos hosted on YouTube.

| Parameter  | Description                                     | Required |
| ---------- | ----------------------------------------------- | -------- |
| `id`       | Video ID                                        | yes      |
| `start`    | Start time in seconds                           | no       |
| `end`      | End time in seconds                             | no       |
| `autoplay` | Set to `"true"` to automatically play the video | no       |

The `id` is the `v` URL parameter in the YouTube URL.
For example, for the YouTube URL `https://www.youtube.com/watch?v=g97CjKOZqT4`, the shortcode is the following:

```markdown
{{</* youtube id="g97CjKOZqT4" */>}}
```

You can configure `start` and `end`, in seconds, with:

```markdown
{{</* youtube id="g97CjKOZqT4" start="100" end="200" */>}}
```

You can configure automatic playback with:

```markdown
{{</* youtube id="g97CjKOZqT4" autoplay="true" */>}}
```

## Escaping Hugo shortcodes

If you need to display the syntax for a shortcode, you can escape it using this syntax:

![Escaped shortcode](./writers-toolkit-escaped-shortcode.png)

Produces:

```markdown
{{</* myshortcode */>}}
```

## About version substitution

Version substitution enables the use of absolute paths that resolve correctly, irrespective of version.
It uses special syntax using angle bracket delimiters like `<GRAFANA_VERSION>`.

As a convention, use the name of the target project all upper-case.
For example, `grafana` becomes `GRAFANA`, `grafana-cloud` becomes `GRAFANA CLOUD`.

The shortcode substitutes the special syntax `<SOMETHING_VERSION>` with the version inferred from the page's URL.
If the page's URL has the prefix `/docs/grafana/latest/`, the shortcode replaces the syntax `<SOMETHING_VERSION>` with `latest` in the final URL.

You can override version inference by including additional metadata in the front matter of the file.
To override the value of `<GRAFANA_VERSION>`, set the `GRAFANA_VERSION` parameter in the page's front matter.
For example, to set the version to `next` irrespective of the source content version, add the following to the front matter: `GRAFANA_VERSION: next`.
