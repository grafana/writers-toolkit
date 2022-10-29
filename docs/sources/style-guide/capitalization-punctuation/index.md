---
title: Capitalization and punctuation
description: Capitalization and punctuaton
aliases:
  - /docs/writers-toolkit/latest/style-guide/capitalization-punctuation/
weight: 400
keywords:
  - capitalization
  - punctuation
---

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

### Grafana-specific capitalization guidelines
- Menu and submenu titles always use sentence case: capitalize the first word, and lowercase the rest.
  - _Dashboards_ when referring to the submenu title.
  - _Keyboard shortcuts_ when referring to the submenu topic.
- Generic and plural versions are always lowercase.
  - Lowercase _dashboard_ when referring to a dashboard generally.
  - Lowercase _dashboards_ when referring to multiple dashboards.
- **Exceptions:** If a term is lowercased in the Grafana UI, then match the UI.

### Kubernetes objects

When referring to Kubernetes objects, such as Jobs, Pods, and StatefulSets, capitalizing them makes it clear that you are not talking about generic jobs and pods.

In the first use, introduce the object as _Kubernetes XX_, then use it alone in subsequent uses.

**Examples:**

- Create the Kubernetes Job and check the logs to retrieve the generated token:
- The Job requires the token be submitted as …

## Punctuation guidelines

Refer to the following punctuation guidelines when you write technical content.

- After a period, add one space, not two.
- In a series, use serial commas before _and_ or _or_.
  - Example: "During lunch, we enjoyed quiche, quinoa, _and_ kale salad.”
- Do not abbreviate _and_ with an ampersand (_&_).
  - Exception: If the UI uses an ampersand, match the UI.
- At the end of a paragraph, remove extra space characters.
- Use italics to indicate paths within a web application, and greater-than symbols (>) to indicate path separators.
  - Example: “To add an administrator to the list of local users, navigate to Appliance > Configuration > Access.”
