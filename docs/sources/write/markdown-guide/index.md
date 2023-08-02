---
title: Markdown guide
menuTitle: Markdown guide
description: Guidelines for writing technical documentation in Markdown.
weight: 300
aliases:
  - /docs/writers-toolkit/writing-guide/markdown-guide/
  - /docs/writers-toolkit/write/markdown-guide/
keywords:
  - Markdown
  - headings
  - bold
  - tables
  - lists
---

# Markdown guide

This Markdown guide helps keep contributions consistent across all Grafana Labs documentation. Refer to the guide and update it as needed when a subject matter expert (SME) answers a question about Markdown syntax, or a decision is made about how to apply Markdown.

We use the static site generator [Hugo](https://gohugo.io/) to generate the web site for the documentation.

Hugo uses a Markdown parser named Goldmark, which supports the CommonMark flavor of Markdown including some extended features. For more information, see the [CommonMark specification](https://spec.commonmark.org/), and a [quick reference guide](https://commonmark.org/help/) for CommonMark.

**Write in sentence case** throughout all technical documentation, be it long-form text or microcopy within a UI:

- This is sentence case
- This is Headline Case

## Headings

Similar to HTML headings (`<h1>`, `<h2>`, and `<h3>`), in Markdown, `#` symbols (or _hash tags_) create different heading levels:

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

- To make text **bold**, surround the text with `**two asterisks**`. For example:

  > **Note:** It is important to use GitHub-flavored Markdown emojis consistently.

- To emphasize text, surround the text with `_single underscores_`̣.
  Do not use single asterisks (`*`), because they can be easily confused with two (for bold).

```markdown
**Note:** The distributor only passes _valid_ data to the ingesters.
```

Displays as:

**Note:** The distributor only passes _valid_ data to the ingesters.

## Links and references

For information about creating links between topics inside and outside of a Grafana Labs repository, refer to [Links and cross references]({{< relref "../references" >}}).

There are two forms of links in Markdown: inline and reference-style.

When you create an inline link, you define the link text and destination in the same location in the document.
The following snippet demonstrates an inline link with the text "Link text to display" and the destination https://example.com.

```markdown
[Link text to display](https://example.com)
```

When you create a reference-style link, you define your link text and then use a label to reference the link destination that is defined somewhere else in the document, usually at the end of the file.
Reference-style links let you define a link destination once, and then reuse the label multiple times in the document.

The following snippet demonstrates a reference-style link with the text "Link text to display", the destination https://example.com, and uses the label "label".

The destination can be defined anywhere in the page but is typically put at the bottom of the page to imitate footnotes.

```markdown
[Link text to display][label]

[label]: https://example.com
```

## Block quotes

Include block quotes within text by using a right-angle bracket:

```markdown
> This text is in block quotes.
```

**Example:**

> Any important information about emojis
> can be separated into a blockquote.

## Code blocks

Code blocks within Markdown can highlight syntax that is specific to a language. Use three back tics to create a code block. For example, ` ``` ` immediately followed by `javascript` produces the following highlights:

```javascript
function testNum(a) {
  if (a > 0) {
    return "positive";
  } else {
    return "NOT positive";
  }
}
```

## Tables

Construct a table by separating the table headings by a `|` (pipe) character. Then, add a second line of dashes (`-`) separated by another `|` character. When constructing the table cells, separate each cell’s data with a `|`.

**Example**:

```markdown
| Heading one   | Heading two   |
| ------------- | ------------- |
| Cell one data | Cell two data |
```

When rendered, the preceding table displays as follows:

| Heading one   | Heading two   |
| ------------- | ------------- |
| Cell one data | Cell two data |

## Numbered lists

Use repetitive list numbering, to avoid inconsistent list numbering:

```markdown
1. First
1. Second
1. Third
```

The preceding list displays as:

1. First
2. Second
3. Third

## Unordered lists

Build a list of unordered items by using a hyphen (`-`):

```markdown
- One item
- Another item
- And another list item
```

> **Note:** Use unordered lists whenever the items have no particular sequence.

The preceding snippet displays as follows:

- One item
- Another item
- And another list item

## Images

Include images in a document using the following syntax:

```markdown
![Alt text](link to image, starting with /static/img/docs/ if it is to an internal image "Title of image in sentence case")
```

> **Note:** Alternative text (alt text) doesn't appear when the user hovers their cursor over the image. The title text does.

**Examples:**

- `![Grafana logo](/link/to/grafanalogo/logo.png "Grafana logo")`
- `![Example](/static/img/docs/folder_name/alert_test_rule.png "Example title")`

This follows the format `![alt text](URL)`.

Alternatively, you can use the [figure shortcode]({{< relref "../shortcodes#figure-shortcode" >}}) if you need more image options, such as adding captions or controlling the image size.

Within Markdown, HTML is valid, but should be used sparingly:

```html
<img
  src="example.png"
  alt="Example image"
  style="float: left; margin-right: 5px;"
/>
```

In most cases, use Markdown syntax rather than HTML syntax. Only use HTML if you need to change the image in ways that Markdown does not support.

## Description list

The Markdown parser that Hugo uses, Goldmark, has built-in support for description lists.
You can use description lists for terms and their definitions, or core concepts.
The syntax is as follows:

```markdown
term
: description_text
```

You can add more markup in a description list.
For example, you can format the definition terms as bold text.

```markdown
Reasons you might want to write programs in Go include the following:

**Fast compile times**
: The Go compiler is fast!

**Ecosystem**
: Go tooling is excellent.
```

The preceding description list displays as follows:

**Fast compile times**
: The Go compiler is fast.

**Ecosystem**
: Go tooling is excellent.

## Comments

You can include comments that do not display in published output:

`[comment]: <> (Comment text to display)`

## Shortcodes

Shortcodes are predefined templates that let you reuse snippets of technical documentation. To learn how to use shortcodes, refer to [Shortcodes]({{< relref "../shortcodes" >}}).
