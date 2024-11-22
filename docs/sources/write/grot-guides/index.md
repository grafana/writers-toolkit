---
title: Grot guides
review_date: "2024-08-29"
description: Understand what Grot guides are and how to use them in your Markdown.
keywords:
  - Hugo
  - shortcodes
weight: 501
---

# Grot guides

Grot guides are interactive guides embedded in a documentation page that people can follow to get help and guidance on topics and direction to documentation resources.

The `guide` shortcode embeds a Grot guide into a page.

```markdown
{{</* guide name="<NAME>" title="<TITLE>" text="<TEXT>" button="<BUTTON>" */>}}
```

The website looks up guide Markdown files from the website repository in the `content/guides` folder. The shortcode accepts the following parameters:

| Parameter | Description                                                                 | Required |
| --------- | --------------------------------------------------------------------------- | -------- |
| `name`    | The name of the guide Markdown file in the website `content/guides` folder. | yes      |
| `title`   | A custom guide title to override the default guide title.                   | no       |
| `text`    | Custom body text to override the default guide text.                        | no       |
| `button`  | Custom button copy to override the default guide button copy.               | no       |

If you don't set custom title, text, and button values, the guide introduction screen uses the following default text:

![Grot guide introduction screen with default copy](https://grafana.com/media/docs/writers-toolkit/grot-guide.png)

## Structure of a Grot guide

Guide content consists of various YAML structures in Markdown front matter for the guide meta, header, welcome, and screens.

### Meta

The meta sections covers general information about the guide. Customize the `name`, `api`, and `title` fields, and leave the rest as is:

```yaml
name: <GUIDE NAME>
api: <GUIDE NAME>
title: <GUIDE TITLE>
type: guides
layout: single
_build:
  render: false
  list: true
```

## Header

The header structure defines the guide's image and image spacing. The following example sets a `src`, `alt`, and image dimensions:

```yaml
header:
  image:
    src: /media/guides/grafana-guides-whichgrafana-header.svg
    alt: Grot metrics wizard
    width: 221
    height: 131
```

## Welcome

A guide needs a single welcome object, defined with a `welcome` field and the following attributes:

```yaml
welcome:
  type: welcome
  title: <WELCOME TITLE>
  body: <WELCOME BODY>
  ctas:
    - text: <CTA TEXT>
      screen_id: <SCREEN ID TO LINK TO>
```

## Screens

Define subsequent screens as a list of objects under a `screens` field:

```yaml
screens:
  - type: question
    id: <UNIQUE SCREEN ID>

  - type: result
    id: <UNIQUE SCREEN ID>
```

### Question

A question screen is a branch node in a decision tree and presents one or many options to further screens. question screen has the following structure:

```yaml
screens:
  - type: question
    id: <UNIQUE SCREEN ID>
    title: <SCREEN TITLE>
    options:
      - text: <OPTION TEXT>
        screen_id: <SCREEN ID TO LINK TO>
```

### Result

A result screen is a leaf node in a decision tree and terminates with one or many links. A question screen has the following structure, with examples for docs and play links:

```yaml
screens:
  - type: result
    id: <UNIQUE SCREEN ID>
    title: <SCREEN TITLE>
    body: <SCREEN BODY>
    links:
      - type: docs
        title: <LINK TITLE>
        link_text: <LINK TEXT>
        href: <ABSOLUTE DOCS LINK>
      - type: play
        title: <LINK TITLE>
        link_text: <LINK TEXT>
        href: <FULL GRAFANA PLAY LINK>
```
