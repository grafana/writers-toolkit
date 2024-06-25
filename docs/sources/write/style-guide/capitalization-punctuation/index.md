---
aliases:
  - /docs/writers-toolkit/style-guide/capitalization-punctuation
  - /docs/writers-toolkit/write/style-guide/capitalization-punctuation
date: "2022-06-27T11:51:13-05:00"
description: Guidelines for use of capitalization and punctuation.
keywords:
  - capitalization
  - punctuation
review_date: "2024-04-15"
title: Capitalization and punctuation
weight: 400
---

# Capitalization and punctuation

This section includes capitalization and punctuation guidelines.

## Capitalization

Write headings in sentence case.

Consult the following capitalization guidelines when you write.

### Sentence case

Sentence case capitalizes the first word of a heading and all proper nouns.

This style is primarily lowercase and is considered the quickest form to read.

### Capitalize abbreviations

Use all-caps capitalization for abbreviations, such as `API`, `HTTP`, `ID`, `JSON`, `SQL`, or `URL`.

### Capitalize proper nouns

The names of people, places, and products take initial capitals because they're proper nouns.

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
- The Job requires the token

### Amazon products

When referring to Amazon products such as [Amazon CloudWatch](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/WhatIsCloudWatch.html), include the "Amazon" name prefix.
After the first use, it's OK to use just the product name.

In headings, always use the full name including "Amazon".

## Punctuation guidelines

Refer to the following punctuation guidelines when you write technical content.

- After a period, add one space, not two.
- When listing a series of items, insert a comma before _and_ or _or_.
  This is known as using serial commas or the Oxford comma.
  - Example: "During lunch, they enjoyed quiche, quinoa, _and_ kale salad.”
- Don't abbreviate _and_ with an ampersand (_&_).
  - Exception: If the UI uses an ampersand, match the UI.
