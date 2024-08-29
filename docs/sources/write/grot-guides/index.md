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

The [`guide` shortcode](/docs/writers-toolkit/write-shortcodes/#guide) embeds a Grot guide into a page.

```markdown
{{</* guide name="instrument" title="title" text="text" button="button" */>}}
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

Guide content consists of various YAML structures in Markdown front matter for the guide meta, header, welcome, and screens.

## Meta

Customize the `name`, `api`, and `title` fields, and leave the rest as is:

```yaml
# name that matches the filename and name used to embed
name: instrument
# match the name field
api: instrument
title: Applicating Instrumentation
type: guides
layout: single
# Hugo rendering options
_build:
  render: false
  list: true
```

## Header

The header structure defines the guide's image and image spacing:

```yaml
header:
  image:
    src: /media/guides/grafana-guides-whichgrafana-header.svg
    alt: Grot metrics wizard
    width: 221
    height: 131
```

## Welcome

A guide needs a single welcome object, defined with a `welcome` field and the following attributes. Define a list of `ctas` with at least one item which points to the next `screen_id`.

```yaml
welcome:
  type: welcome
  title: Let's get started with application instrumentation
  body: Answer a few questions and Grot can help you find the documentation you need to instrument your application.
  ctas:
    - text: Let's go!
      screen_id: instrument
```

## Screens

Define subsequent screens as a list of objects under a `screens` field:

```yaml
screens:
  - type: question
    id: ...

  - type: result
    id: ...
```

### Question

A question screen is a branch node in a decision tree and presents one or many options to further screens. A question screen has the following structure:

```yaml
- type: question
    id: instrument
    title: What programming language do you want to instrument?
    options:
      - text: JVM (Java, Scala, Kotlin)
        screen_id: jvm
      - text: .Net
        screen_id: dotnet
      - text: Node.js
        screen_id: nodejs
```

### Result

A result screen is a leaf node in a decision tree and terminates with one or many links. A question screen has the following structure:

```yaml
  - type: result
    id: beyla
    title: Grafana Beyla
    body: |
      Grafana Beyla is a Linux eBPF kernel module application to auto-instrument applications without modifying them by monitoring an executable or port.
    links:
      - type: docs
        title: Grafana Beyla
        link_text: Visit docs page
        href: /docs/beyla/latest/
```
