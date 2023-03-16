---
title: Front matter
menuTitle: Front matter
description: Learn about how we build front matter to properly enable the publication and search of our technical documentation.
aliases:
  - /docs/writers-toolkit/latest/writing-guide/front-matter/
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

Labels can be inherited through cascading front matter.
Each project has a set of default labels that are defined in the root `_index.md` file of the project.

For versioned projects, the `_index.md` file resides in the `website` repository.
For unversioned projects, the `_index.md` file resides in the project’s repository.

If the default labels are incorrect for a page or directory of pages, update the labels.
For each page, include a label in the `labels.products` sequence for every product that the page relates to.

For a single page, if the page describes a feature available in Grafana Cloud and Grafana Enterprise, the source file front matter should include the following:

```yaml
labels:
  products:
    - cloud
    - enterprise
```

For a directory of pages, if all pages describe a feature only available in Grafana Cloud, the branch bundle `_index.md` file front matter should include the following:

```yaml
cascade:
  labels:
    products:
      - cloud
```

### title

**Required.**

Becomes the document title element. Often browsers display this in the tab for the page.

It doesn't need to precisely match the `menuTitle`.
Optimize the title for search engines. Use double quotes (`"`) to surround the title. Do not use smart quotes.

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

**Unversioned projects**
: Include an `aliases` entry that refers to the initial published website directory.
Adding an `aliases` entry makes it safer to move content around as the redirect from old to new page location is already in place.
Hugo doesn't create a redirect `.html` file when the directory is already populated with content.
: > **Note:** The published directory is dependent on which `content` subdirectory documentation is synced to in the website repository.
: > For example, documentation synced to a the `content/docs` directory requires the `/docs` prefix.

**Versioned projects**
: Do not include an `aliases` entry that refers to the initial published website directory. The version in the URL path can cause undesirable redirects, such as a redirect from latest content to an old version.
