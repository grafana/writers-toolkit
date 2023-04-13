---
title: Write useful link text
menuTitle: Write useful link text
description: Learn how to write useful link text well.
weight: 100
keywords:
---

# Write useful link text

Link text is important for search engines and users, and is critical for accessibility. Wikipedia defines link text as, “The anchor text, link label or link text is the visible, clickable text in an HTML hyperlink. The term "anchor" was used in older versions of the HTML specification…”

The primary purpose of link text is to give users a short summary of the link destination content.

## Summary

When writing link text, pay attention to words, quantity of links, and the length of link text.

### Words

- Make link text abundantly obvious as to what the next page is about.
- Use or do not use words such as _Click here_, _More_, and _Read more_:

  - Use them if they are on a website card.

  - Do not use them as standalone link text, because they can be confusing when a screen reader reads them out of context.

- When linking to GitHub, include `GitHub repository` in the sentence text or link text.

  For example:

  - “To learn more, go to the `[faro-web-sdk]()` repository.”
  - ”To learn more about Grafana Faro, go to its `[repository]()`.”

- When linking from outside of a documentation page, include `documentation` in the text.

### Quantity

- Aim for no more than two links per paragraph.

  Some links are essential for additional context, and others are nice to have but not necessary.

  Using too many links on a page makes it difficult for a user to read the content.

### Length

- Keep link text short and concise.
- Never start with articles: _a_, _an_, or _the_.
- Do not link whole sentences; aim for a maximum of four to five words.

## Examples

| Use                                                                                    | Do not use                                          |
| -------------------------------------------------------------------------------------- | --------------------------------------------------- |
| `[Learn more about cluster navigation]()` or learn more about `[cluster navigation]()` | `[Learn more]()`                                    |
| To learn more read our `[Grafana Phlare blog post]()`                                  | To learn more about Phlare read our `[blog post]()` |
| read more about specific `[metrics and alerting rules]()`                              | `[read more]()`                                     |
| `[our <company name> products]()` or `[<company name> products]()`                     | `[our products]()`                                  |
| visit the `[<company name> website]()` or visit` [<company name>.com`]()               | `[visit our website]()`                             |
| `[SNMP exporter]()`                                                                    | `[exporter]()`                                      |
| `[KubeCon + CloudNativeCon North America 2022]()`                                      | `[KubeCon event]()`                                 |
| The basic `[definition of cardinality]()`                                              | The `[basic definition]()` of cardinality           |
| see `[configuring Kubernetes monitoring]()`                                            | `[see configuring Kubernetes monitoring]()`         |
