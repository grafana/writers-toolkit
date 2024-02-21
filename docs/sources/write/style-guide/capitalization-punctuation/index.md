---
title: Capitalization and punctuation
description: Guidelines for use of capitalization and punctuation.
weight: 400
aliases:
  - /docs/writers-toolkit/style-guide/capitalization-punctuation
  - /docs/writers-toolkit/write/style-guide/capitalization-punctuation
keywords:
  - capitalization
  - punctuation
---

# Capitalization and punctuation

This section includes capitalization and punctuation guidelines.

## Capitalization

Consult the following capitalization guidelines when you write.

### Sentence-case capitalization

Use sentence-case capitalization to capitalize the first word of a sentence, the first word of a heading, and all proper nouns.

This style is primarily lowercase and is considered the quickest form to read.

### All caps capitalization

Use all-caps capitalization exclusively for abbreviations, such as `API`, `HTTP`, `ID`, `JSON`, `SQL`, or `URL`.

### Capitalizing proper nouns

The names of people, places, and products take initial capitals because they are proper nouns.

If you're not sure of whether something is a product or not, consult the product manager, or if there isn't one, the responsible squad.

### Grafana-specific capitalization guidelines

- Menu and submenu titles always use sentence case: capitalize the first word, and lowercase the rest.
  - _Dashboards_ when referring to the submenu title.
  - _Keyboard shortcuts_ when referring to the submenu topic.
- Generic and plural versions are always lowercase.
  - Lowercase _dashboard_ when referring to a dashboard generally.
  - Lowercase _dashboards_ when referring to multiple dashboards.
- **Exceptions:** If a term is lowercase in the Grafana UI, then match the UI.

### Kubernetes objects

When referring to Kubernetes objects, such as Jobs, Pods, and StatefulSets, follow the guidance in the [Kubernetes documentation style guide](https://kubernetes.io/docs/contribute/style/style-guide/#use-upper-camel-case-for-api-objects).
Capitalizing objects makes it clear that you aren't talking about generic jobs or deployments.

In the first use, introduce the object as _Kubernetes XX_, then use it alone in subsequent uses.

**Examples:**

- Create the Kubernetes Job and check the logs to retrieve the generated token:
- The Job requires the token be submitted as …

## Punctuation guidelines

Refer to the following punctuation guidelines when you write technical content.

- After a period, add one space, not two.
- When listing a series of items, insert a comma before _and_ or _or_. This is known as using serial commas or the Oxford comma.
  - Example: "During lunch, we enjoyed quiche, quinoa, _and_ kale salad.”
- Do not abbreviate _and_ with an ampersand (_&_).
  - Exception: If the UI uses an ampersand, match the UI.
- At the end of a paragraph, remove extra space characters.
- Use bold to indicate paths within a web application, and greater-than symbols (>) to indicate path separators.
  - Example: “To add an administrator to the list of local users, navigate to **Appliance** > **Configuration** > **Access**.”
