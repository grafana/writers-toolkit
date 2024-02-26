---
aliases:
  - /docs/writers-toolkit/writing-guide/about-documentation-design/
  - /docs/writers-toolkit/structure/about-documentation-design/
date: 2024-02-26
description: Learn about the design of Grafana's documentation pages
keywords:
  - Grafana
  - documentation
  - page design
menuTitle: Documentation design
title: About documentation design
weight: 100
---

# About documentation design

The documentation website uses a modern design approach to make technical documentation accessible and scalable.

Documentation pages take advantage of the static site generator Hugo.
As a result, several elements of the page are automatically managed during the publication of the page using Hugo's taxonomy.
Thus, the source Markdown files _don't need to hand management_ of these elements and _don't require_ contributors to curate them.

Pages also include:

- **Navigation to preview primary topics.**
  The left-hand sidebar broadly outlines key topics, with nested related topics underneath.
  This design supports the philosophy that "every page is page one" and creates an system of documentation around a topic that's easier to reference and navigate.
- **Floating table of contents.**
  The table of contents floats on the page as you scroll to the content that's hidden beneath the fold.
  You can also view the upcoming topics, to enable a better user experience that helps you navigate to subtopics lower on the page.
- **Auto-generated _Related documentation_.**
  Using Hugo's taxonomy, documentation automatically finds other documentation that's related to the page you're viewing.
- **Auto-generated _Related resources from Grafana Labs_.**
  Hugo's taxonomy again automatically generates this content.
- **Feedback.**
  Thumbs up and thumbs down feedback.

You can read about the redesign of the documentation pages in [Grafana documentation: A look at the new and improved design](https://grafana.com/blog/2023/02/03/grafana-documentation-a-look-at-the-new-and-improved-design/).
