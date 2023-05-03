---
title: Front matter
menuTitle: Front matter
description: Learn about how we build front matter to properly enable the publication and search of our technical documentation.
weight: 700
keywords:
  - front matter
  - alias
---

# Front matter

Grafana technical documentation includes front matter to help organize the content, develop the TOC (as published in the lefthand sidebar of the website), and help users identify useful pages when searching or viewing the content in search engines or in social media, such as Twitter.

Use YAML for all front matter.
In certain presentations, all front matter characters might render literally.
For this reason, _do not_ include any special Markdown formatting, like italics or monospace, in front matter.

Here’s a correctly built example:

```
---
aliases:
  - /docs/mimir/latest/old-architecture/
description: Learn more about Grafana Mimir’s microservices-based architecture.
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

# About Grafana Mimir architecture
```

## Reference

The following headings describe what each element does and provides guidelines for its content.

### aliases

Provides an HTML redirect from the pages in the list to the current page.
For more information, refer to [Hugo aliases](#hugo-aliases).

### date

Describes the initial publish date of the page.
Hugo produces XML page outputs for use by RSS feeds where users can be notified of updates.
Customers use RSS feeds of our release notes pages to be notified of new releases.
Therefore, the `date` front matter is recommended for release note pages.

The value to the `date` field should be a full ISO 8601 timestamp.
For example, `date: "2023-04-24T00:00:00Z"` is 12:00 AM, Apr 24 Coordinated Universal Time (UTC).

### description

**Required.**

On social media, such as Twitter, displays as a clue to users about what the page includes.

The number of characters vary by media, so make the description concise.
Provide enough information to guide users to the content by describing what content the link leads to.
Often, this doesn’t need to be original prose—you can often scan the first few paragraphs to pluck the appropriate terms or phrases into the description.
If it's too long, it is harmlessly truncated on social media.
Use double quotes (`"`) to surround the title. Do not use smart quotes.

### draft

When set to `true`, this option prevents Hugo from rendering the content.
Use the command line flag `--buildDrafts` to generate content marked as `draft: true`.

### keywords

The website uses keywords to link to related pages in the _Related content_ sections.
They do not appear in the resulting HTML source for the page and do not affect SEO.

Ideally, use single terms as opposed to phrases.

### labels

Use the `labels` key to add one or more values that you want to appear before the topic title on the published page.
Only certain labels are supported.

For `labels.products`, the supported values and the resulting published labels are as follows:

- `cloud`: "Grafana Cloud"
- `enterprise`: "Enterprise"
- `oss`: "Open source"

Labels are inherited through cascading front matter.
All labels, including project defaults, are set in the project's root index file.
This is almost always the `docs/sources/_index.md` file in the project repository.

If you add a new page, consider whether the default labels are appropriate.
If the default labels are incorrect for a page or directory of pages, update the cascading front matter in `docs/sources/_index.md`.

#### Set cascading front matter

Set cascading front matter by adding or updating the `cascade` mapping in the `docs/sources/_index.md` front matter:

```yaml
cascade: <TARGETS...>
```

Where `<TARGETS...>` is a sequence of mappings that looks similar to the following:

```yaml
- _target: { path: <PATH> }
  labels:
    products: <PRODUCTS...>
```

- `<PATH>` is a glob pattern that matches the affected source pages.
  For more information about the the `<PATH>` value, refer to [Specify a path](#specify-a-path).

- `<PRODUCTS...>` is a sequence of valid product labels.

The following example target mapping sets the product label to "Grafana Cloud" for all pages under `/docs/grafana/latest/troubleshooting/`:

```yaml
- _target: { path: /docs/grafana/*/troubleshooting/** }
  labels:
    products:
      - cloud
```

##### Specify a path

The value of the `_target.path` mapping is a string matching paths from the Hugo content root.
Glob patterns specified with `*` and `**` match multiple paths.

The `*` pattern matches any file or directory.
It is useful for matching any version directory for versioned documentation.

The `**` pattern matches all directories and files recursively.

### menuTitle

Becomes the link text for the page in website menus.

The `menuTitle` doesn't need to match the `title`.
It typically omits the product name and any other words that can be derived from the page or menu context.

Examples of omission:

- In Grafana Mimir documentation, it is not necessary to have "Grafana Mimir" in the `menuTitle` because the reader knows this from the page context.
- In a menu with a parent section of "Reference", it is not necessary to have "Reference" in the `menuTitle` because the reader knows this from the menu context.

### title

**Required.**

Becomes the page's HTML title element.
Often browsers display this in the tab for the page.

Optimize the title for search engines by including the product name.

If the `doc-validator` linter has been implemented on your repository, your topic heading must match the title in the metadata.

### weight

Determines the placement of the topic within the left-hand sidebar on https://grafana.com. Smaller numbers place the topic higher in the guide. Pages with the same weight have lexicographic ordering.

Use increments of `100` for all other content files. Doing so makes it easier for you to re-order existing topics when you add new topics. Weights are per directory.

## Example with different page and menu titles

```
---
title: About Grafana Mimir architecture
menuTitle: Architecture
---
```

## Description example

On Twitter:

![Twitter description](twitter.png)

For example:

- Add a panel using these steps.
- Understand the configuration options provided by…
- Learn more about hash rings and their usage

## Hugo aliases

Technical writers use [Hugo aliases](https://gohugo.io/content-management/urls/#aliases) to create redirects to the current page from other URLs.

If you specify `aliases` in the frontmatter, Hugo creates a directory that matches the alias entry that contains a single `.html` file.

### Example

The following example file `intended-url.md` contains the alias `/original-url` within its YAML frontmatter:

```markdown
---
aliases:
  - /original-url/
---
```

Assuming a `baseURL` of `grafana.com`, the auto-generated alias `.html` file found at `https://grafana.com/original-url/` contains something like the following:

```html
<!DOCTYPE html>
<html>
  <head>
    <script>
      const destination = "https://grafana.com/intended-url/";
      console.log(window.location.search);
      document.head.innerHTML = `<meta http-equiv="refresh" content="0; url=${destination}${window.location.search}"/>`;
    </script>
    <title>https://grafana.com/intended-url/</title>
    <link rel="canonical" href="https://grafana.com/intended-url/" />
    <meta name="robots" content="noindex" />
    <meta http-equiv="content-type" content="text/html; charset=utf-8" />
    <noscript>
      <meta http-equiv="refresh" content="0; url={{ safeURL .Permalink }}" />
    </noscript>
  </head>
</html>
```

The `http-equiv="refresh"` `meta` tag attribute, injected by JavaScript, performs an HTML redirect.
For more detail about HTML redirects, refer to [HTML redirections](https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections#html_redirections).

> **Note:** The redirect relies on first party JavaScript support which is common but not necessarily universal.

### Guidelines

The correct way to use aliases depends on whether the project is versioned or not.

#### Unversioned projects

Include an `aliases` entry that refers to the initial published website directory.
Adding an `aliases` entry makes it safer to move content around as the redirect from old to new page location is already in place.
Hugo doesn't create a redirect `.html` file when the directory is already populated with content.
Aliases can be relative or absolute paths.

> **Note:** The published directory is dependent on which `content` subdirectory documentation is synced to in the website repository.
> For example, documentation synced to a the `content/docs` directory requires the `/docs` prefix.

#### Versioned projects

Do not include an `aliases` entry that refers to the initial published website directory.
The version in the URL path can cause undesirable redirects, such as a redirect from latest content to an old version.
Aliases should be relative and not absolute paths so that old versions do not steal redirects from "latest" content when it is moved around.

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
