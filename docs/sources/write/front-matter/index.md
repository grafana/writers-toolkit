---
aliases:
  - /docs/writers-toolkit/writing-guide/front-matter/
  - /docs/writers-toolkit/write/front-matter/
review_date: 2024-04-09
description: Learn about how Grafana builds front matter to properly enable the publication and search of our technical documentation.
keywords:
  - front matter
  - alias
title: Front matter
weight: 400
---

# Front matter

The source files of Grafana documentation use front matter to organize the content, order the project table of contents, and help users identify useful pages when searching or viewing the content in search engines or in social media, such as Twitter.

Use YAML for all front matter.
Unless a front matter field is documented as supporting Markdown, _do not_ include any special Markdown formatting, like italics, in that field.

Here's a correctly built example:

```
---
description: Learn more about Grafana Mimir's microservices-based architecture.
labels:
  products:
    - oss
keywords:
  - Mimir
  - microservices
  - architecture
menuTitle: Architecture
title: About Grafana Mimir architecture
weight: 100
---
```

## Reference

The following headings describe what each front matter field does and provides guidelines for using it.

### Aliases

Use `aliases` to create redirects from the previous URL to the new URL when a page changes or moves.

When you rename or move files, you must create an alias with a reference to the previous URL path to create a redirect from the old URL to the new URL.
In some cases, for example when you have deleted content or split a file into multiple topics, it may not be possible to create an alias for the moved content.

### Example

The following Markdown front matter snippet defines an alias `/original-url/`.
You must incorporate it with any existing front matter.

```markdown
---
aliases:
  - /original-url/
---
```

### Guidelines

The correct way to use aliases depends on whether the project is versioned or not.

#### Versioned projects

Aliases must be relative to avoid redirecting latest content to old versions.

If there is a page in the old documentation that has an alias that includes the version "latest", and the page referred to by that alias doesn't exist in the actual latest documentation, then Hugo creates a redirect at the page referred to by the alias.

That redirect redirects the user from latest documentation into the old documentation.

Aliases should include a YAML comment explaining the absolute URL path that the relative path redirects.
This helps a reviewer confirm that your alias works correctly.

A YAML comment starts with a hash (`#`).

For example, the following Markdown front matter snippet, in the file `new-url.md`, defines an alias to redirect `/docs/grafana/<GRAFANA_VERSION/original-url/`.
The comment `# /docs/grafana/<GRAFANA_VERSION>/original-url/` indicates the absolute URL path.

```
---
aliases:
  - ./original-url/ # /docs/grafana/<GRAFANA_VERSION>/original-url/
---
```

##### Determine the relative alias

To determine the relative alias, you must first understand the meaning of the current directory (`.`) and parent directory (`..`) path elements when they're used at the start of an alias.

For an alias in the page `/docs/grafana/latest/alerting/manage-notifications/`:

- The current directory element (`.`) refers to the directory containing the current page, _not_ the directory of the current page.

  In this example, this is the page `/docs/grafana/latest/alerting/`.

- The parent directory element (`..`) refers to the parent directory of the current directory element.

  In this example, this is the page `/docs/grafana/latest/`.

In the following table:

- **From page**: is the page that requires a redirect, for example because the page has been moved or no longer exists.
- **To page**: is the page where readers are redirected to, for example the renamed page or where the content has been moved.
- **Relative alias**: is the alias that must be added to the front matter of the file for **To page** to create the proper redirect.

<table>
  <thead>
    <tr>
      <th scope="col">FROM PAGE</th>
      <th scope="col">TO PAGE</th>
      <th scope="col">RELATIVE ALIAS</th>
    </tr>
  </thead>
  <tbody>
    {{< docs/alias from="/docs/grafana/latest/alerting.md" to="/docs/grafana/latest/alerting/manage-notifications/_index.md" output="row" >}}
    {{< docs/alias from="/docs/grafana/latest/alerting/silences/_index.md" to="/docs/grafana/latest/alerting/manage-notifications.md" output="row" >}}
    {{< docs/alias from="/docs/grafana/latest/alerting/manage-notifications/create-silence/index.md" to="/docs/grafana/latest/alerting/manage-notifications/_index.md" output="row" >}}
    {{< docs/alias from="/docs/grafana/latest/" to="/docs/grafana/latest/alerting/manage-notifications/" output="row" >}}
    {{< docs/alias from="/docs/grafana/latest/old-alerting/" to="/docs/grafana/latest/alerting/manage-notifications/" output="row" >}}
  </tbody>
</table>

###### Use the `docs/alias` shortcode

You can use the `docs/alias` shortcode to determine the relative alias but you can't use the shortcode in the front matter.

You must first use the shortcode in your source file and then copy the result from the page in the local web server into the front matter of the source file.

```markdown
{{</* docs/alias from="/docs/grafana/latest/old-alerting/" to="/docs/grafana/latest/alerting/manage-notifications/" */>}}
```

Produces:

{{< docs/alias from="/docs/grafana/latest/old-alerting/" to="/docs/grafana/latest/alerting/manage-notifications/" >}}

#### Other projects

- Include an `aliases` entry for the current URL path.
- Add an `aliases` entry to make it safer to move content around, as the redirect from old to new page location is already in place.
  Hugo doesn't create a redirect `.html` file when the directory is already populated with content.
- When a page is moved, update the `aliases` with the new URL path.

### Test an alias

To test an alias results in the correct redirect, use your browser or a command-line tool for making HTTP requests.

#### Use the browser

1. Start the documentation webserver with `make docs`.
1. Browse to the URL of the page that should be redirected.
1. Confirm that you are redirected to the desired page.

   For example, if you want the page `https://grafana.com/docs/grafana/latest/panels/working-with-panels/` to redirect to `https://grafana.com/docs/grafana/latest/panels-visualizations/panel-editor-overview/`, browse to the following URL in the browser to confirm the redirect is working: http://localhost:3002/docs/grafana/latest/panels/working-with-panels/.

#### Use `cURL`

1. Start the documentation webserver with `make docs`.
1. In a separate terminal, make an HTTP GET request to the URL of the page that should be redirected.
   For example, to request the page `localhost:3002/docs/grafana/latest/panels/working-with-panels/`

   ```bash
   curl localhost:3002/docs/grafana/latest/panels/working-with-panels/
   ```

   The output is similar to the following:

   ```console
   <!doctype html><html><head><script>const destination="http://localhost:3002/docs/grafana/latest/panels-visualizations/panel-editor-overview/";console.log(window.location.search),document.head.innerHTML=`<meta http-equiv="refresh" content="0; url=${destination}${window.location.search}"/>`</script><title>http://localhost:3002/docs/grafana/latest/panels-visualizations/panel-editor-overview/</title><link rel=canonical href=http://localhost:3002/docs/grafana/latest/panels-visualizations/panel-editor-overview/><meta name=robots content="noindex"><meta charset=utf-8><noscript><meta http-equiv=refresh content="0; url=http://localhost:3002/docs/grafana/latest/panels-visualizations/panel-editor-overview/"></noscript></head></html>
   ```

1. Confirm that the value of the `destination` `const` in the `<script>` tag is the pretty URL for the page with the alias.

### Canonical

The `canonical` front matter sets the preferred URL for duplicate or very similar pages.
Search engines use this information and only index the canonical URL.

The value of the `canonical` URL should be the full URL of the canonical page.
For example, `https://grafana.com/docs/writers-toolkit/`.

For content reused in Grafana Cloud, prefer the open source documentation as the canonical page.

### Cascade

`cascade` is a map of front matter fields.
The fields are passed down from the parent to the page descendants.

You can use `cascade` to define variables. For example:

```yaml
cascade:
  PRODUCT_VERSION: 10.1
  PRODUCT_NAME: Grafana
```

Use the [`param`](https://grafana.com/docs/writers-toolkit/write/shortcodes#param) shortcode in the topic body text wherever you need to insert the variable.

### Date

`date` describes the initial publish date of the page.
Hugo produces XML page outputs for use by RSS feeds where users can be notified of updates.
Customers use RSS feeds of release notes pages to be notified of new releases.
Therefore, the `date` front matter is recommended for release note pages.

The value of the `date` field should be a full ISO 8601 timestamp.
For example, `date: "2023-04-24T00:00:00Z"` is 12:00 AM, Apr 24 Coordinated Universal Time (UTC).

The `date` front matter also impacts menu ordering.
Pages with more recent dates are lower in the menu.

### Description (required)

Use `description` to provide the short description of the topic to search engines, including the search engine used in the Grafana documentation site.
The description is also displayed on social media, such as Twitter, to provide a clue to users about the page contents.

Since the reader isn't on the Grafana website, your description should include contextual information, such as the product name.

The number of characters vary by media, so make the description concise.
Provide enough information to guide users to the content by describing what content the link leads to.
Often, this doesn't need to be original text.
You can scan the first few paragraphs to pluck the appropriate terms or phrases into the description.
If the description is too long, it's truncated on social media.

### Draft

When `draft` is set to `true`, this option prevents Hugo from rendering the content.
Use the command-line flag `--buildDrafts` to generate content marked as `draft: true`.

### Keywords

The website uses `keywords` to generate links to related pages in the _Related content_ sections.
They don't appear in the resulting HTML source for the page and don't affect search engine optimization (SEO).

Ideally, use single terms as opposed to phrases.

### Labels

Use the `labels` key to add one or more values that you want to appear before the topic title on the published page.
Only certain labels are supported.

For `labels.products`, the supported values and the resulting published labels are as follows:

- `cloud`: <span class="badge docs-labels__product-cloud docs-labels__item">Grafana Cloud</span>
- `enterprise`: <span class="badge docs-labels__product-enterprise docs-labels__item">Enterprise</span>
- `oss`: <span class="badge docs-labels__product-oss docs-labels__item">Open source</span>

For `labels.stages`, the supported values and the resulting published labels are as follows:

- `beta`: <span class="badge docs-labels__stage docs-labels__item">Beta</span>
- `alpha`: <span class="badge docs-labels__stage docs-labels__item">Alpha</span>
- `experimental`: <span class="badge docs-labels__stage docs-labels__item">Experimental</span>

Labels can be inherited through cascading front matter.

For versioned projects, the `_index.md` file resides in the `website` repository.
For other projects, the `_index.md` file resides in the project’s repository.

If the default labels are incorrect for a page or directory of pages, update the labels.
Also, if you are adding a new page, consider whether the default labels are appropriate.
For each page, include a label in the `labels.products` sequence for every product that the page relates to.

For example, if a _single page_ describes a feature available in Grafana Cloud and Grafana Enterprise, the source file front matter should include the following:

```yaml
labels:
  products:
    - cloud
    - enterprise
```

For a _directory of pages_ that describe a feature only available in Grafana Cloud, the branch bundle `_index.md` file front matter should include the following:

```yaml
cascade:
  labels:
    products:
      - cloud
```

<!-- vale Grafana.Headings = NO -->
<!-- False positive due to the PascalCase that results from capitalizing camelCased shortcodes -->

### MenuTitle

<!-- vale Grafana.Headings = YES -->

Use `menuTitle` to specify a different title in the sidebar navigation than in the `title` element.
For example, if you want to abbreviate the topic title in the table of contents.

Don't remove the verb from [task topic](https://grafana.com/docs/writers-toolkit/structure/topic-types/task/) titles.
The verb helps the reader know that they're navigating to a task topic before they follow the link.

It is OK to remove the verb from the `menuTitle` if it's implied by the containing section.
For example, the page 'Install Grafana Alloy in a Docker container' can be shortened to 'Docker' if it's in a section named 'Install'.

### Refs

Use the `refs` front matter with `ref` URIs to link to different pages in reused content.

{{< docs/shared source="writers-toolkit" lookup="refs-example.md" leveloffset="+3" >}}

For more information, refer to [Link from source content that’s used in multiple projects](https://grafana.com/docs/writers-toolkit/write/links/#link-from-source-content-thats-used-in-multiple-projects).

### Review date

Use the `review_date` front matter to set the date you last reviewed a page for correctness.

Set the date using the `YYYYY-MM-DD` format, and separate the elements by using hyphens.
For example, to set the last review date to June 6, 2024, use `2024-06-06`.

The website includes the review date at the foot of the page's content.
You can see how this renders on the [Writers' Toolkit home page](https://grafana.com/docs/writers-toolkit/#:~:text=Last%20reviewed%3A).

### Slug

The `slug` front matter overrides the last segment of the URL path.
It's ineffective on `_index.md` files which are also known as _section pages_ or _branch bundles_.
For more information, refer to [Slug](https://gohugo.io/content-management/urls/#slug).

You should prefer to update the filename instead of using the `slug` front matter because it makes it easier to find the correct source file for a URL.

### Title (required)

Hugo uses the `title` to generate the sidebar table of contents if there is no `menuTitle` specified in the front matter.
Your `title` should match your first heading for search engine optimization (SEO).
The `doc-validator` linter enforces this.

The `title` becomes the document title element in the HTML.
Often, browsers display this in the tab for the page.

Optimize the title for search engines.

### Weight

By default, topics are displayed in alphabetical order by `title`.

Use `weight` to specify a different topic order within the left-hand sidebar on https://grafana.com.
Smaller numbers place the topic earlier in the guide or section of the guide.
Pages with the same weight are displayed in alphabetical order.

Use increments of `100` for content files.
Doing so makes it easier for you to re-order existing topics when you add new topics.
Weights are per directory.

## Tutorials

There is additional front matter that you only need for tutorials.
Tutorials should also include all the regular front matter.

### Associated technologies

The `associated_technologies` front matter is a sequence of strings that refer to taxonomies in the website data directory.
If you are a Grafana Labs employee, you can view the associated technologies in the [website data directory](https://github.com/grafana/website/tree/master/data/taxonomies/associated_technologies).

The `associated_technologies` value is the filename without the file extension.
For example, to refer to author defined in the `mimir.yaml` file, use `mimir`.

If you don't set the `associated_technologies` front matter, `grafana` is the default.

The following YAML example demonstrates setting a single associated technology of `mimir`.
You must incorporate it with the rest of your front matter.

```yaml
associated_technologies:
  - mimir
```

### Authors

The `authors` front matter is a sequence of strings that refer to author files defined in the website data directory.
If you are a Grafana Labs employee, you can view and add authors to the [website data directory](https://github.com/grafana/website/tree/master/data/authors).

The `authors` value is the filename without the file extension.
For example, to refer to author defined in the `grafana_labs.yaml` file, use `grafana_labs`.

If no appropriate author file exists, `grafana_labs` is a good default.

The following YAML example demonstrates setting a single author of `grafana_labs`.
You must incorporate it with the rest of your front matter.

```yaml
authors:
  - grafana_labs
```

### Summary

The `summary` front matter defines a short summary used on the tutorial's card on https://grafana.com/tutorials/.

The following YAML example demonstrates summary front matter.
You must incorporate it with the rest of your front matter.

```yaml
summary: Use Telegraf to stream live metrics to Grafana.
```

### Tags

The `tags` front matter is a sequence of strings displayed as tags under the author section on the tutorials page.

Typically, at least one of the tags is an expertise level.
The expertise levels are:

- Beginner
- Intermediate
- Advanced

The following YAML example demonstrates setting a single tag of the expertise level `Beginner`.
You must incorporate it with the rest of your front matter.

```yaml
tags:
  - Beginner
```
