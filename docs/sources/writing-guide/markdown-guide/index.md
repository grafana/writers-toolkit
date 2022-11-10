---
title: Markdown guide
menuTitle: Markdown guide
description: Guidelines for writing technical documentation in Markdown.
aliases:
  - /docs/writers-toolkit/latest/writing-guide/markdown-guide/
weight: 500
keywords:
  - Markdown
  - headings
  - bold
  - tables
  - lists
---

# Markdown guide

This Markdown guide helps keep contributions consistent across all Grafana Labs documentation. Refer to the guide and update it as needed when a subject matter expert (SME) answers a question about Markdown syntax, or a decision is made about how to apply Markdown.

**Write in sentence case** throughout all technical documentation, be it long-form text or microcopy within a UI:

- This is sentence case
- This is Headline Case

## Headings

Similar to HTML headings (`<h1>`, `<h2>`, and `<h3>`), in Markdown, `#` symbols (or *hash tags*) create different heading levels:

**Example**

- \# is a parent heading.
- \#\# is a child heading.
- \#\#\# is a child’s child heading.

For the title of the page, use one `#`. For each child heading, use two `##` symbols.

### Heading don'ts

- Avoid stacked headings; do not follow a heading with another without any content between the two.
- Avoid skipping heading levels. For example, after a single `#`, use `##`, rather than `###`.
- Avoid having just one child-level heading:
  - Valid: `#`, `##`, `##`, `###`, `###`, `##`, `##`
  - Invalid: `#`, `##`, `###`, `##`, `###`, `##`
- Avoid using hyphens in headings.
- With the exception of `(Optional) `, do not include parenthesized words such as (Important).

## Bold and emphasis

- Make text **bold** using two asterisks. For example:

  > **Note:** It is important to use GitHub-flavored Markdown emojis consistently.

- To emphasize text, use single ` _underscores_`. Do not use single asterisks (`*`), because they can be easily confused with two (for bold).

  **For example:** The distributor only passes _valid_ data to the ingesters.

## Links and references

For information about creating links between topics inside and outside of a Grafana Labs repository, refer to [Links and cross references]({{< relref "../references/" >}}).

If you want to add a link to an external website, wrap the display text in square brackets, and the web URL in curly brackets.

\[Link text to display](https://website.com)

**Example:** For more information about including emojis in GitHub-flavored markdown, refer to the WebFX [Emoji Cheat Sheet](https://www.webfx.com/tools/emoji-cheat-sheet/). 

## Block quotes

Include block quotes within text by using a right-angle bracket:

**Example**:

> Any important information about emojis
> can be separated into a blockquote.

## Code blocks

Code blocks within Markdown can highlight syntax that is specific to a language. Use three back tics to create a code block. For example, ` ``` ` immediately followed by `javascript` produces the following highlights:

```javascript
function testNum(a) {
  if (a > 0) {
    return 'positive';
  } else {
    return 'NOT positive';
  }
}
```

## Tables

Construct a table by separating the table headings by a `|` (pipe) character. Then, add a second line of dashes (`-`) separated by another `|` character. When constructing the table cells, separate each cell’s data with a `|`.

**Example**:

Heading one | Heading two

\------------|------------

Cell one data| Cell two data

Displays as follows:

| Heading one   | Heading two   |
| ------------- | ------------- |
| Cell one data | Cell two data |

## Numbered lists

Use repetitive list numbering, to avoid inconsistent list numbering:

1. First
1. Second
1. Third

The preceding list displays as:

1. First
2. Second
3. Third

## Unordered lists

Build a list of unordered points by using a hyphen (`-`):

- First item
- Another item
- The final list item

> **Note:** If might be tempting to number a list because the verbiage that precedes it includes a number.
> Remember that if list items do not need to be performed in a particular order, use an unordered list.

**Example:**

There are three ways to ingest data:
- First way
- Second way
- Third way

## Images

_Do not_ use image shortcodes at this time. Instead, include images in a document using the following syntax:

```
![Alt text](link to image, starting with /static/img/docs/ if it is to an internal image "Title of image in sentence case")
```

> **Note:** Alt(ernative) text does not appear when the user hovers their cursor over the image. The title text does.

**Examples:**

- `![Grafana logo](/link/to/grafanalogo/logo.png "Grafana logo")`
- `![Example](/static/img/docs/folder_name/alert_test_rule.png "Example title")`

This follows the format `![alt text](URL)`.

Within Markdown, HTML is valid and to be used sparingly:

```
<img src="example.png"
     alt="Example image"
     style="float: left; margin-right: 5px;" />
```

In most cases, use Markdown syntax rather than HTML syntax. Only use the HTML if you need to change the image in ways that Markdown does not supported.

## Comments

You can include comments that do not display in published output:

`[comment]: <> (Comment text to display)`

## Shortcodes

Shortcodes are predefined templates that allow you to reuse snippets across the Grafana website. To learn how to use shortcodes, refer to [Shortcodes]({{< relref "../shortcodes/" >}}).