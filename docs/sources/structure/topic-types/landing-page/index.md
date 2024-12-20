---
date: 2024-12-18
description: Learn how to create a landing page.
keywords:
  - topic types
  - template
  - landing page
review_date: 2024-12-18
title: Create landing pages
---

# Create landing pages

Create a landing page as a starting point for customers to get access to the documentation they're looking for.

To create a landing page, complete the following steps.

1. Go to the folder and `_index.md` file for your landing page topic.
1. Double-check existing front matter. You can leave existing front matter as is.

   If there is useful text in the existing landing page - add it to the landing page or create an Introduction topic nested underneath the landing page.

1. Add the `hero` shortcode to the front matter.

   This shortcode becomes the banner at the top of the landing page.

   Example:

   ```yaml
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

1. Add the `card-grid` shortcode to the front matter. This shortcode is for the tiles that appear below the Explore heading.

   Example:

   ```yaml
   cards:
     title_class: pt-0 lh-1
     items:
       - title: Grafana Alerting
         href: /docs/grafana-cloud/alerting-and-irm/alerting/
         description: Allows you to learn about problems in your systems moments after they occur. Monitor your incoming metrics data or log entries and set up your Alerting system to watch for specific events or circumstances and then send notifications when those things are found.
         logo: /media/docs/grafana-cloud/alerting-and-irm/grafana-icon-alerting.svg
         height: 24
   ```

   {{< admonition type="note" >}}

   - Start each description with a verb.

     If left empty, the description for the tile is automatically inherited from the short description front matter in the linked page. The descriptions in the front matter, however, are often short and it might be a good idea to add more context, making sure they all start with a verb for consistency.

   - If you are creating a landing page that appears in both Cloud and OSS, use a relative path, for example, `./set-up/`.

   - Icons are only required for products. If you don't have an icon, delete `logo` from the front matter.

   {{< /admonition >}}

1. To display the banner at the top of the page, add `{{</* docs/hero-simple key="hero" */>}}` after the front matter. This needs to come before the first heading.
1. Add the `## Overview` heading and your content.
1. Add the `## Explore` heading and this syntax `{{</* card-grid key="cards" type="simple" */>}}` to display the tiles below.
1. Save your topic and build your documentation to review your changes.
