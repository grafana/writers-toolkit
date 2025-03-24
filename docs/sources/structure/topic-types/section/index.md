---
date: 2024-08-22
description: Learn how to write a section page.
keywords:
  - section
  - topic types
  - template
menuTitle: Section
review_date: 2024-08-22
title: Section page
---

# Section page

Section pages, also called landing pages, direct users to the most valuable resources contained within the section.

## Section structure

A _section_ page includes the following elements:

- **Hero:** A banner across the top of the page that includes the section title and a brief description.
- **Overview:** A heading and one or two paragraphs that describe the resources contained within the section.
- **Explore:** A heading that contains cards.
  - **Cards:** Defined in the front matter, cards provide overviews of important pages and include links to those pages.

## Write a section page

To write a section, follow these steps:

1. Determine where you want to add the section, or find an existing section page to update.
1. If you're creating a new section, create a directory that follows this naming convention:

   - The directory name matches the menu title planned for the section.
   - Use lowercase letters.
   - Add a hyphen between words.

1. Create an `_index.md` file within the section directory.
1. Add the usual front matter to the index file.

   For more information about front matter, refer to [Front matter](https://grafana.com/docs/writers-toolkit/write/front-matter/).

1. Add the elements specific to a section page:

   - **Hero:** A [hero](https://grafana.com/docs/writers-toolkit/write/shortcodes/#hero-simple) requires front matter and use of a shortcode.
   - **Cards:** A [card grid](https://grafana.com/docs/writers-toolkit/write/shortcodes/#card-grid) requires front matter and use of a shortcode.

### Section page examples

Refer to the following pages for section page examples:

- [Grafana Explore](https://grafana.com/docs/grafana/latest/explore/): A section page without an image in the hero.
- [Alerts and IRM](https://grafana.com/docs/grafana-cloud/alerting-and-irm/): A section page with an image in the hero.

## Section template

When you're ready to write, make a copy of the [Section template](https://github.com/grafana/writers-toolkit/blob/main/docs/static/templates/section-template.md) and add your content.
