---
title: Write useful link text
menuTitle: Write useful link text
description: Learn how to write useful link text well.
aliases:
  - /docs/writers-toolkit/latest/writing-guide/useful-link-text/
weight: 100
keywords:
---

# Write useful link text

Link text is important for search engines and users, and is critical for accessibility. Wikipedia defines link text as, “The anchor text, link label or link text is the visible, clickable text in an HTML hyperlink. The term "anchor" was used in older versions of the HTML specification…”

The primary purpose of link text is to give search engines and users a short summary of the link destination content. Search engines consider link text important for crosslinking relevancy.

> **Note:** If you refer to an OSS project in a blog post, first link to the project in GitHub, and then link to the https://grafana.com product page.

## Summary

When writing link text, pay attention to words, quantity of links, the length of link text, and localization concerns.

### Words
 
* Make link text abundantly obvious as to what the next page is about.
* Use the words that are in the destination URL, if possible.
* Use or do not use words such as *Click here*, *More*, and *Read more*:
  
  * Use them if they are on a website card.

  * Do not use them as standalone link text, because they can be confusing when a screen reader reads them out of context.
  
* When linking to GitHub, include `GitHub repository` in the sentence text or link text.

  For example: 

  * “To learn more, go to the `[Grafana faro-web-sdk GitHub repository]()`.”
  * ”To learn more about Grafana Faro, go to its `[GitHub repository]()`.”
* Do not use more than one *Learn more →* link.

  If you use more than one, it is not clear what the subsequent ones refers to because it’s preceded by a link. Will the user learn more about Grafana Cloud in general, Cloud sign-up options, or something else?

  Users can set screen readers to read the links on the page first to give the listener an idea of what’s to come and to assess the number of link topics on a page. Without context the listener will hear “learn more, learn more, try on grafana cloud, learn more”.

  An alternative and preferred approach is to add context to the link text or use, when possible, web available Design features, such as what appears on the <a href="https://grafana.com/about/press/">Grafana Labs Newsroom</a> and other pages. The entire box, called a *card*, is clickable.

* When linking from outside of a documentation page, include `documentation` in the text.

### Quantity

* Aim for no more than two links per paragraph.

  Some links are essential for additional context, and others are nice to have but not necessary.

  Using too many links on a page is considered SPAM, and doing so makes it difficult for a user to read the page.

* Limit linking to the same page as much as possible.
  
  Ideally, only once. Google does not track multiple occurrences of the same link on a single page.

### Length

* Keep link text short and as to the point as possible.
* Never start with articles: *a*, *an*, or *the*.
* Do not link whole sentences; aim for a maximum of four to five words.

### Localization

* On localized pages, indicate the language of the content on the destination page.

## Examples

| Use | Do not use |
|-|-|
| `[Learn more about cluster navigation]()` or learn more about `[cluster navigation]()` | `[Learn more]()` |
| To learn more read our `[Grafana Phlare blog post]()` | To learn more about Phlare read our `[blog post]()` |
| read more about specific `[metrics and alerting rules]()` | `[read more]()` |
| `[our <company name> products]()` or `[<company name> products]()` | `[our products]()` |
| visit the `[<company name> website]()` or visit` [<company name>.com`]() | `[visit our website]()` |
| `[SNMP exporter]()` | `[exporter]()` |
| `[KubeCon + CloudNativeCon North America 2022]()` | `[KubeCon event]()` |
| The basic `[definition of cardinality]()` | The `[basic definition]()` of cardinality |
| see `[configuring Kubernetes monitoring]()` | `[see configuring Kubernetes monitoring]()` |
