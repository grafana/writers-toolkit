---
title: About our documentation design
menuTitle: Documentation design
description: Learn about Grafana's documentation docs pages
aliases:
  - /docs/writers-toolkit/latest/writing-guide/about-documentation-design/
weight: 100
keywords:
  - Grafana
  - documentation
  - page design
---

# About our documentation design

In late 2022, the documentation website was significantly modified to a more modern design. The goal of the redesign was to make our technical documentation more accessible, modern, and scalable as we grow.  

Our docs pages take advantage of our static site generator, Hugo. As a result, several elements of the page are automatically managed during the publication of the page using Hugo's taxonomy. Thus, the source markdown files **do not need to hand management** of these elements and **do not require** contributors to curate them.

We also include:

- **Improved navigation.** The left-hand sidebar broadly outlines key topics, with nested related topics underneath. This design supports the philosophy that "every page is page one" and creates an system of documentation around a topic that is easier to reference and navigate. 
- **More accessible table of contents.** The table of contents now floats on the page as you scroll to the content that's hidden beneath the fold. You can also view the upcoming topics, to enable a better user experience that helps you navigate to subtopics lower on the page.
- **Auto-generated _Related documentation_.** Using Hugo's taxonomy, our documentation now automatically finds other documentation that's pertinent to the page you're viewing.
- **Auto-generated _Related resources from Grafana Labs_.** Hugo's taxonomy again is used to automatically generate this content.
- **Feedback.** We added more prominent options for feedback from our community.

You can read more about the redesign in our [blog] [https://grafana.com/blog/2023/02/03/grafana-documentation-a-look-at-the-new-and-improved-design/].
