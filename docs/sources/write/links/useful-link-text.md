---
aliases:
  - /docs/writers-toolkit/write/style-guide/useful-link-text/
  - /docs/writers-toolkit/write/useful-link-text/
date: "2024-02-07T18:29:24+00:00"
description: Understand what makes useful link text.
review_date: "2024-09-03"
title: Useful link text
---

# Useful link text

Link text is important for search engines and users, and is critical for accessibility.
It's the visible, clickable text in an HTML hyperlink.

The primary purpose of link text is to give users a short summary of the link destination content.

## Summary

When writing link text, pay attention to the words used and the length of the link text.

### Words

- Make link text obviously reflect what the next page is about.
- Don't use words such as _Click here_, _More_, and _Read more_ as standalone link text, because they can be confusing when a screen reader reads them out of context.

- When linking to a GitHub repository, include `repository` in the sentence text or link text.

  For example:

  - To learn more, go to the `[faro-web-sdk]()` repository.
  - To learn more about Grafana Faro, go to its `[repository]()`.

- When linking from outside of a documentation page, include `documentation` in the text.

### Quantity

Aim for no more than two links per paragraph.

Some links are essential for additional context, and others are nice to have but not necessary.

Using too many links on a page makes it difficult for a user to read the content.

### Length

- Keep link text short and concise.
- Include all the words that help a user decide whether to follow the link.
- Never start with articles: _a_, _an_, or _the_.
- Don't link whole sentences; aim for a maximum of four to five words.

## Examples

| Use                                                                                    | Don't use                                      |
| -------------------------------------------------------------------------------------- | ---------------------------------------------- |
| `[Learn more about cluster navigation]()` or learn more about `[cluster navigation]()` | `[Learn more]()`                               |
| `[our Grafana products]()` or `[Grafana products]()`                                   | `[our products]()`                             |
| visit the `[Grafana website]()`                                                        | `[visit our website]()`                        |
| `[SNMP exporter]()`                                                                    | `[exporter]()`                                 |
| `[KubeCon + CloudNativeCon North America 2022]()`                                      | `[KubeCon event]()`                            |
| The basic `[definition of cardinality]()`                                              | The `[basic definition]()` of cardinality      |
| refer to `[Configure Kubernetes Monitoring]()`                                         | `[refer to Configure Kubernetes Monitoring]()` |
