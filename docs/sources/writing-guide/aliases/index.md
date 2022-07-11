---
title: "About Hugo aliases"
menuTitle: "About Hugo aliases"
description: "When and how to use aliases in Hugo in Grafana Labs technical documentation."
aliases: []
weight: 1
keywords:
  - Hugo
  - aliases
---

# About Hugo aliases

Technical writers use [Hugo aliases](https://gohugo.io/content-management/urls/#aliases) to create redirects to the current page from other URLs.

If you specify `aliases` in the frontmatter, Hugo creates a directory that matches the alias entry that contains a single `.html` file.

## Example

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
    const destination = 'https://grafana.com/intended-url/';
    console.log(window.location.search)
    document.head.innerHTML = `<meta http-equiv="refresh" content="0; url=${destination}${window.location.search}"/>`;
    </script>
    <title>https://grafana.com/intended-url/</title>
    <link rel="canonical" href="https://grafana.com/intended-url/"/>
    <meta name="robots" content="noindex">
    <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
    <noscript>
      <meta http-equiv="refresh" content="0; url={{ safeURL .Permalink }}"/>
    </noscript>
  </head>
</html>
```

The `http-equiv="refresh"` `meta` tag attribute, injected by JavaScript, performs an HTML redirect.
For more detail about HTML redirects, refer to [HTML redirections](https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections#html_redirections).

> **Note:** The redirect relies on first party JavaScript support which is common but not necessarily universal.

## Guideline

To allow content to be easily moved, include an `aliases` entry that refers to the initial published website directory.
Hugo doesn't create a redirect `.html` file when the directory is already populated with content.

> **Note:** The published directory is dependent on which `content` subdirectory documentation is synced to in the website repository.
> For example, documentation synced to a the `content/docs` directory requires the `/docs` prefix.
