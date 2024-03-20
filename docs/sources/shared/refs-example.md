---
date: 2024-03-18
description: An example of the `refs` front matter.
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

Hugo checks each `pattern` and `destination` pair, in order, against the page's URL.
If the page's URL matches the pattern _`URL PATH PREFIX`_, then Hugo uses `destination` as the link destination.
If no pattern matches the current page, Hugo logs a build error.

The _`FULL URL`_ destination has the same behavior as described in [Link to `grafana.com` pages](https://grafana.com/docs/writers-toolkit/write/links/#link-to-grafanacom-pages)

# Example

The following Markdown snippet demonstrates the `refs` front matter and link that uses a `ref` URI.

In the [Grafana page](https://grafana.com/docs/grafana/latest/alerting/set-up/), the link destination is https://grafana.com/docs/grafana/latest/alerting/fundamentals/data-source-alerting/.
Hugo replaces the [version substitution syntax](https://grafana.com/docs/writers-toolkit/write/shortcodes/#about-version-substitution) `<GRAFANA_VERSION>` with the version inferred from the current page.

In the page [Grafana Cloud page](https://grafana.com/docs/grafana/latest/alerting/set-up/), the link destination is https://grafana.com/docs/grafana-cloud/alerting-and-irm/alerting/fundamentals/alert-rules/.

```markdown
---
refs:
  grafana-alerting:
    - pattern: /docs/grafana/
      destination: https://grafana.com/docs/grafana/<GRAFANA_VERSION>/alerting/fundamentals/data-source-alerting/
    - pattern: /docs/grafana-cloud/
      destination: https://grafana.com/docs/grafana-cloud/alerting-and-irm/alerting/fundamentals/data-source-alerting/
---

# Set up Alerting

## Before you begin

- Check which data sources are compatible with and supported by [Grafana Alerting][ref:grafana-alerting].
```
