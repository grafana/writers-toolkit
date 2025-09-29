---
aliases:
  - /docs/writers-toolkit/style-guide/inclusive-writing
  - /docs/writers-toolkit/write/style-guide/inclusive-writing
date: "2022-06-27T11:51:13-05:00"
description:
  Understand the importance of inclusive writing at Grafana Labs and learn
  how to write inclusively.
keywords:
  - inclusive writing
review_date: 2024-06-27
title: Inclusive writing
weight: 400
---

# Inclusive writing

<!-- vale Grafana.We = NO -->
<!-- This page talks about the voice and tone of our organization. -->

When writing for Grafana Labs, always use inclusive terminology and phrasing.
Avoid all statements that perpetuate gender, racial, or cultural stereotypes.
Don't use gendered words or demographically oriented terms that are irrelevant to the content.

## Write about people

When writing about people, be compassionate, inclusive, and respectful.

No matter the context, we always make sure to view people as people.
With that in mind, there are important distinctions between the ways that different people interact with our product and, thus, how we refer to those people.

- Users are every person who uses our products, whether or not they use a commercial version of one of our products.
  Customers are people and companies who pay to use our products.
- Visitors are people who visit our website for any reason, whether it's as a potential user, customer, or employee.
  We appreciate all interest in our hard work.
- Our customers (and users) often have their own customers (and users).
  Generally speaking, we refer to these customers of our customers as users.
- Employees are people who work at Grafana Labs.

### Best practices

Consider the following best practices to ensure inclusivity in your writing.

- Group or audience: Don't refer to a group or audience as _it_.
  Instead, use _they_.
- Age: Don't refer to someone's age unless it's relevant.
  Don't use age-related words to describe people.
  Avoid: _young_, _old_, or _elderly_.
- Hearing: Do use the adjectives _deaf_, _partially deaf_, or _hard of hearing_ when describing a person with hearing loss.
- Vision: Do use the adjective _blind_ or _low vision_ when describing a person with limited vision.

## Write for an international audience

Grafana Labs uses US English as our standard for spelling, grammar, punctuation, and more.
We write for an international audience with the following guidelines in mind:

### Avoid acronyms

If you use an acronym, first write out each word in the acronym on first use, and then continue using just the acronym.

- **Example**: Grafana Cloud is software as a service (SaaS).
  Using SaaS typically requires setting up an account.

### Avoid idioms

Understanding idioms typically requires cultural context.
Users of Grafana Labs products and projects are all around the world, and many are non-native English speakers.
Therefore, avoid idioms and idiomatic puns.

For example, some common American English idioms:

- Costs an arm and a leg
- Hit the nail on the head
- Bite off more than you can chew

Here are some more obscure British idioms:

- A few sandwiches short of a picnic
- Chuffed to bits
- Over-egg the pudding

### Avoid cultural references

Cultural references are references that relate to the culture of a community, country, continent, and so on.

**Examples:**

- Her ambition is a clear reflection of the American dream.
- You came to help like a good Samaritan.
- They grew up in a typical nuclear family.

## Avoid charged language

<!-- vale Grafana.WordList = NO -->
<!-- This is demonstrating improper usage. -->

Avoid using charged language such as _blacklist_, _master_, and _slave_.

<!-- vale Grafana.WordList = YES -->

### Allow, block

<!-- vale Grafana.WordList = NO -->
<!-- This is demonstrating improper usage. -->

Avoid _whitelist_ or _blacklist_.

<!-- vale Grafana.WordList = YES -->

When referring to _allowing_ or _blocking_ content or traffic, use a form of _allow_ or _block_:

- For the noun form, use _allowlist_ or _blocklist_
- For the verb form, use _allow_ or _block_

Example: _To **allow** outgoing traffic, add the IP to the **allowlist**._

### Primary and secondary

Avoid _master_ or _slave_.

Use the following approach to describe relationships between nodes or processes:

- Use _primary_, _main_, or _parent_, instead of _master_.
- Use _secondary_, _replica_, or _child_, instead of _slave_.

### Right and left

Use the terms _right_ and _left_ with a qualifier like _top_ or _lower_ [to help individuals with cognitive disabilities or those using screen-reading software](https://learn.microsoft.com/en-us/style-guide/a-z-word-list-term-collections/l/left-leftmost-left-hand).

Hyphenate when you're modifying a noun. For example:

- in the top-right corner
- at the bottom left of the dialog box
- the upper-left section
- scroll toward the lower right

Don't use the word _hand_ as a qualifier like in _right-hand corner_.

### Avoid directional language

Language that relies only on spatial direction doesn't help users with visual disabilities or those who use screen readers. Remove directional language from your documentation when possible, and add language that doesn't depend on spatial cues to help users navigate.

Use:

- In the previous section
- In the following table
- The next step
- Earlier in this topic

Don't use:

- In the section above
- The table below
- To the right
- See the sidebar on the left

