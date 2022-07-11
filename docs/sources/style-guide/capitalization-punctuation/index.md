---
title: "Capitalization and punctuation"
description: "Capitalization and punctuaton"
aliases: []
weight: 50
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
  - "Dashboards" when referring to the submenu title.
  - "Keyboard shortcuts" when referring to the submenu topic.
- Generic and plural versions are always lowercase.
  - Lowercase "dashboard" when referring to a dashboard generally.
  - Lowercase "dashboards" when referring to multiple dashboards.
- **Exceptions:** If a term is lowercased in the Grafana UI, then match the UI.

### Kubernetes objects

Capitalize Kubernetes objects such as Job, Pod, and StatefulSet when it is clear you are specifically talking about them and not generic jobs and pods.

Introduce the object as "Kubernetes XX" on the first usage, then just the object in subsequent uses.

**Examples:**

- Create the Kubernetes Job and check the logs to retrieve the generated token:
- The Job requires the token be submitted as …

## Punctuation guidelines

Refer to the following punctuation guidelines when you write technical content.

- Add one space after a period, not two.
- Use serial commas, which are commas before "and" or "or" in a series.
  - Example:  “Through lunch, we suffered quiche, quinoa, and kale salad.”
- Do not use an ampersand (&) as an abbreviation for _and_.
  - **Exception:** If an ampersand is used in the Grafana UI, then match the UI.
- Remove extra space characters at the end of a paragraph.
- Navigation: When referring to a location within the web application, use italics to indicate a path, and greater-than symbols (>) to indicate path separators.
  - Example: “To add an administrator to the list of local users, navigate to *Appliance > Configuration > Access*.”
