---
aliases:
  - /docs/writers-toolkit/writing-guide/markdown-guide/
  - /docs/writers-toolkit/write/markdown-guide/
date: "2022-06-27T13:38:07-05:00"
description: Guidelines for writing technical documentation in Markdown.
keywords:
  - Markdown
  - headings
  - bold
  - tables
  - lists
review_date: "2024-06-26"
title: Markdown guide
weight: 300
---

# Markdown guide

The Grafana website uses the static site generator [Hugo](https://gohugo.io/) to generate pages for the documentation.

Hugo uses a Markdown parser named Goldmark, which supports the CommonMark flavor of Markdown including some extended features.
For more information, refer to the [CommonMark specification](https://spec.commonmark.org/), and a [quick reference guide](https://commonmark.org/help/) for CommonMark.

Write in sentence case throughout all technical documentation, be it long-form text or microcopy within a UI:

- This is sentence case
- This is Headline Case

## Headings

Similar to HTML headings (`<h1>`, `<h2>`, and `<h3>`), in Markdown, `#` symbols (or _hash tags_) create different heading levels:

### Example

- \# is a parent heading.
- \#\# is a child heading.
- \#\#\# is a child's child heading.

For the title of the page, use one `#`. For each child heading, use two `##` symbols.

<!-- vale Grafana.Gerunds = NO -->
<!-- A false positive because the noun heading looks like a gerund. -->

### Identifiers

Each heading has an auto-generated identifier that you can use to link to the heading within the page.
To derive the generated identifier from a heading, refer to [Link to page headings](https://grafana.com/docs/writers-toolkit/write/links/#link-to-page-headings).

You can also explicitly set the heading identifier.
The following Markdown example sets the heading identifier to be `alternative-heading-id`:

```markdown
# Heading {#alternative-heading-id}
```

### Heading don'ts

<!-- vale Grafana.Gerunds = YES -->

- Avoid stacked headings.
  Don't follow a heading with another without any content between the two.
- Avoid skipping heading levels.
  For example, after a single `#`, use `##`, rather than `###`.
- Avoid using hyphens in headings.
- With the exception of _(Optional)_, don't include parenthesized words such as _(Important)_.
- Avoid duplicate headings.
  If you need to reuse the same heading, try to keep the meaning consistent.
  Do what's best to avoid confusing the user.

## Bold and emphasis

- To make text **bold**, surround the text with `**two asterisks**`.
- To _emphasize_ text, surround the text with `_single underscores_`̣.

Don't use single asterisks (`*`), because they can be easily confused with the two asterisks used for bold.

To understand how to use bold and emphasis in documentation, refer to [Text formatting](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#text-formatting).

## Links

For information about creating links between topics inside and outside of a Grafana Labs repository, refer to [Links](https://grafana.com/docs/writers-toolkit/write/links/).

There are two forms of links in Markdown: inline and reference-style.

When you create an inline link, you define the link text and destination in the same location in the document.
The following snippet demonstrates an inline link with the text "Link text to display" and the destination https://example.com.

```markdown
[Link text to display](https://example.com)
```

When you create a reference-style link, you define your link text and then use a label to reference the link destination that's defined somewhere else in the document, usually at the end of the file.
Reference-style links let you define a link destination once, and then reuse the label multiple times in the document.

The following snippet demonstrates a reference-style link with the text "Link text to display", the destination https://example.com, and uses the label "label".

```markdown
[Link text to display][label]

[label]: https://example.com
```

You can also define reference-style links without an explicit label.
In such a case, the label is the link text.

The following snippet demonstrates the two different ways of writing reference-style links with implicit labels using an unordered list.

```markdown
- [Link text to display]
- [Link text to display][]

[Link text to display]: https://example.com
```

## Block quotes

Include block quotes within text by using a right-angle bracket (`>`).
For note, tip, warning, and caution admonitions, prefer the [admonition shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#admonition).

### Example

```markdown
> This text is in block quotes.
```

Produces:

> This text is in block quotes.

## Code blocks

Use three backticks to create a fenced code block.

The info string, after the first three backticks, describes the language contained within.
The website uses this information to apply syntax highlighting to code examples.
For more information, refer to [Highlight syntax](https://grafana.com/docs/writers-toolkit/write/style-guide/write-for-developers/#highlight-syntax).

## Tables

Construct a table by separating the table headings with pipe (`|`) characters.
Separate the table heading row from the data with a line of columns where each cell is a series of dashes (`-`) separated by another pipe characters.
Finally construct a series of table data rows by separating the cell data with pipe characters.

<!-- vale Grafana.Paragraphs = NO -->
<!-- The following explains what shouldn't be done (using <br> elements). -->

Don't use `<br>` elements to create paragraphs or lists.
Instead, use `<p>` elements for paragraphs, or the `<ol>` or `<ul>` elements for lists.

<!-- vale Grafana.Paragraphs = YES -->

For style guidance regarding tables, refer to the [Google Developer documentation style guide](https://developers.google.com/style/tables).

### Example

```markdown
| Heading one   | Heading two   |
| ------------- | ------------- |
| Cell one data | Cell two data |
```

Produces:

| Heading one   | Heading two   |
| ------------- | ------------- |
| Cell one data | Cell two data |

## Numbered lists

Use repetitive list numbering, where you prefix every list entry with `1.` instead of the actual number, to avoid inconsistent list numbering.
The Markdown renderer automatically increments the list.

For sub-steps, use repetitive numbering as well.

When writing paragraphs as list entries, you must use proper indentation:

- Each line in the entry must match the indentation of the preceding list item.
- Each paragraph must have an empty line before it.

For an numbered list in isolation, the indentation for the second sentence of a list entry is three spaces.

### Examples

```markdown
1. First
1. Second
1. Third
```

Produces:

1. First
1. Second
1. Third

```markdown
1. First
   1. Write a sub-step.
   1. Write another sub-step.
   1. Write yet another sub-step.
1. Second
1. Third
```

Produces:

1. First
   1. Write a sub-step.
   1. Write another sub-step.
   1. Write yet another sub-step.
1. Second
1. Third

```markdown
1. First paragraph in first entry.
   Second sentence in first paragraph.

   Second paragraph in first entry.

1. First paragraph in second entry.
```

Produces:

1. First paragraph in first entry.
   Second sentence in first paragraph.

   Second paragraph in first entry.

1. First paragraph in second entry.

## Unordered lists

Build a list of unordered items by using a hyphen (`-`).
Use unordered lists whenever the items have no particular sequence.

When writing paragraphs as list entries, you must use proper indentation:

- Each line in the entry must match the indentation of the preceding list item.
- Each paragraph must have an empty line before it.

For an unordered list in isolation, the indentation for the second sentence of a list entry is two spaces.

### Example

```markdown
- One item
- Another item
- And another list item
```

Produces:

- One item
- Another item
- And another list item

## Images

Include images in a document using the following syntax `![<ALT TEXT>](<URL> "<TITLE>")`.

{{< admonition type="note" >}}
Alt text doesn't appear when the user hovers their cursor over the image.
The title text does.
{{< /admonition >}}

### Examples

- `![The Grafana logo](/link/to/grafanalogo/logo.png "Grafana logo")`
- `![An alert test rule](/static/img/docs/folder_name/alert_test_rule.png "Example title")`

If you need more image options, such as adding captions or controlling the image size, you can use the [`figure` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#figure).
Within Markdown, HTML is valid, but you should avoid it.

If you are unable to achieve the desired styling with the `figure` shortcode, raise an [issue](https://github.com/grafana/writers-toolkit/issues/new).

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

You can include comments that don't display in published output:

`[comment]: <> (Comment text to display)`

## Shortcodes

Shortcodes are snippets you use in source files to calling built-in or custom templates.
Shortcodes templates avoid the need for HTML in Markdown and ensure consistency across the documentation set.
To learn how to use shortcodes, refer to [Shortcodes](https://grafana.com/docs/writers-toolkit/write/shortcodes/).
