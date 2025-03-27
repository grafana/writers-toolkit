---
date: "2024-04-04T12:05:52+01:00"
description: An example of the `refs` front matter.
review_date: "2024-06-24"
title: An example of the `refs` front matter
---

A partial front matter containing the `refs` field looks like:

```yaml
refs:
  <KEY>:
    - pattern: <URL PATH PREFIX>
      destination: <FULL URL>
    - pattern: <URL PATH PREFIX>
      destination: <FULL URL>
```

- _`<URL PATH PREFIX>`_ - Enter the part of the page URL that represents the project in which the documentation is published.
- `destination` - Enter the full URL of the destination page for that project including trailing slashes.

If the page's URL matches the pattern _`<URL PATH PREFIX>`_, then Hugo uses `destination` as the link destination.
If no pattern matches the current page, Hugo logs a build error.

The _`<FULL URL>`_ destination has the same behavior as described in [Link to `grafana.com` pages](https://grafana.com/docs/writers-toolkit/write/links/#link-to-grafanacom-pages)

# Example

The following Markdown snippet demonstrates the `refs` front matter and link that uses a `ref` URI.

```markdown
---
refs:
  find-plugins:
    - pattern: /docs/grafana/
      destination: /docs/grafana/<GRAFANA_VERSION>/administration/plugin-management/#browse-plugins
    - pattern: /docs/grafana-cloud/
      destination: /docs/grafana-cloud/introduction/find-and-use-plugins/
---

# Manage plugins

## Before you begin

- Find the plugin you want to install. To find a plugin, refer to [Find and use plugins](ref:find-plugins).
```

In the latest version of a Grafana documentation page, the link destination is <https://grafana.com/docs/grafana/latest/administration/plugin-management/#browse-plugins>.
Hugo replaces the [version substitution syntax](https://grafana.com/docs/writers-toolkit/write/shortcodes/#about-version-substitution) `<GRAFANA_VERSION>` with the version inferred from the current page.

In a Grafana Cloud page, the link destination is <https://grafana.com/docs/grafana-cloud/introduction/find-and-use-plugins/>.
