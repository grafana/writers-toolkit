---
aliases:
  - /docs/writers-toolkit/writing-guide/documentation-structure/topic-types/landing-page
  - /docs/writers-toolkit/structure/topic-types/landing-page/
date: "2022-10-27T16:43:50-04:00"
description: Learn how to create a landing page.
keywords:
  - topic types
  - template
  - landing page
menuTitle: Landing page
review_date: "2024-06-07"
title: Landing page
---

# Landing page

A landing page is a page, which introduces a series of topics related to a particular product, product area, or set of features.

## Create landing pages

Create a landing page as a starting point for customers to get access to the documentation they are looking for.

To create a landing page, complete the following steps.

1. Go to the folder and \_index.md file for your landing page topic.
1. Double-check existing front matter. You can leave existing front matter as is.

   If there is useful text in the existing landing page - add it/create an Introduction topic nested underneath the landing page.

1. Add the `hero` shortcode to the front matter.

   This shortcode becomes the banner at the top of the landing page.

   Example:

   ```markdown
   hero:
   title: Grafana Alerting
   level: 1
   image: /media/docs/grafana-cloud/alerting-and-irm/grafana-icon-alerting.svg
    width: 100
    height: 100
    description: Grafana Alerting allows you to learn about problems in your systems moments after they occur.
   ```


   {{< admonition type="note" >}}
   You may have to adjust the width or height to 100/110 depending on the spacing.
   {{< /admonition >}}

1. Add the `card-grid` shortcode  to the front matter. This shortcode is for the tiles that appear below  ## Explore.

	Start each description with a verb.

   {{< admonition type="note" >}}
   If left empty, the description for the tile is automatically inherited from the short description front matter in the linked page. The descriptions in the front matter, however, are often short and it might be a good idea to add more context, making sure they all start with a verb for consistency.
   {{< /admonition >}}

   Example:

   ```markdown
   cards:
   title_class: pt-0 lh-1
     items:
     title: Grafana Alerting
    href: /docs/grafana-cloud/alerting-and-irm/alerting/
    description: Allows you to learn about problems in your systems moments after they occur. Monitor your incoming metrics data or log entries and set up your Alerting system to watch for specific events or circumstances and then send notifications when those things are found.
    logo: /media/docs/grafana-cloud/alerting-and-irm/grafana-icon-alerting.svg
    height: 24
   ```

   {{< admonition type="note" >}}

   If you are creating a landing page that appears in both Cloud and OSS, use a relative path (eg. ./set-up).

   Icons are only required for products. If you don’t have an icon, delete `logo` from the front matter.
   {{< /admonition >}}

1. To display the banner at the top of the page, add {{</* docs/hero-simple key="hero" */>}} after the front matter.
1. Add the ## Overview header and your content.
1. Add the ## Explore header and this syntax {{</* card-grid key="cards" type="simple" */>}} to display the tiles below.
1. Save your topic and build your documentation to review your changes.
