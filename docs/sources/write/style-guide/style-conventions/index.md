---
title: Style conventions
description: Style conventions
weight: 200
aliases:
  - /docs/writers-toolkit/style-guide/style-conventions
  - /docs/writers-toolkit/write/style-guide/style-conventions
keywords:
  - style conventions
---

# Style conventions

This section includes a non-exhaustive list of technical writing techniques and styles to consider when you write technical content. For questions you might have that are not addressed in these guidelines topic, refer to
[Google developer documentation style guide](https://developers.google.com/style).

## Focus on user goals

Before you begin writing, clearly identify the goal of the user, and write content that supports the user reaching that goal.

- Do not document implementation details, specifications, or backend system operations that have no clear consequence to the user.
  - Providing unnecessary information can lead to bloated content that forces the user to determine which content is relevant.
- Avoid marketing clichés and hyperbole.
  - Instead, use evidence-based or quantifiable language to focus and refine the information, and to offer value propositions.

## Address users clearly

To address users clearly and directly, write directives in the imperative second person where the unspoken 'you' is implied. Avoid making “the user” or their role (“system administrator”) the subject of a sentence, as in: “Users configure the Cloud using a secure shell.”.

Exception:

- You can use first person in UI elements that are specific to the user, such as “My profile" or “My account."

| Use                                                           | Don't use                                                         |
| ------------------------------------------------------------- | ----------------------------------------------------------------- |
| Click **Yes** to accept the license agreement.                | The license agreement is accepted when you click **Yes**.         |
| To create a dashboard, add a panel and specify a data source. | To create a dashboard, you add a panel and specify a data source. |

### Write in active voice

When you write in _active voice_, you identify the subject of the sentence and the action that the subject performs. For example, "John drove the car" is active voice because it is clear that John (the subject) performed an action (drove). The passive voice variation is "The car was driven by John."

| Use                                               | Don't use                                                               |
| ------------------------------------------------- | ----------------------------------------------------------------------- |
| After you create a dashboard, add a panel.        | After the dashboard has been created, the panel can be added.           |
| Click **OK** to save the dashboard configuration. | The dashboard configuration is saved when the **OK** button is clicked. |

### Write simple words, sentences, and paragraphs

Simple, direct communication is the key to effective technical communication.

- Use short words whenever possible, such as "use," not "utilize."
  - If possible, replace "use" and its variants (utilize, make use of) with a more descriptive verb.
- Make your sentences shorter than 25 words.
  - If you can remove a word without losing meaning, do so (typical culprits: there is; there are; in order to; it is important to; keep in mind).
  - Consider writing shorter sentences or using a bulleted list if you find yourself writing long sentences.
- Use simple verbs and tenses.
- Consider the characteristics of your audience when choosing a term.
- Don't use buzzwords or jargon.
- Keep paragraphs to three sentences or less.
  - Condense the text, add more headings, or do both.
- Use contractions in common and negative cases to express a conversational style: _you’re_, _that’s_, _isn’t_ or \_don’t, for example.

Make content relevant to the user's context. The more familiar you are with the user’s context, the better you can communicate without using a lot of words.

### Write in present tense

When you write in present tense, avoid words such as have, has, had, been, should, would, and will.

| Use                                                                 | Don't use                                                                    |
| ------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| The panel opens.                                                    | The panel will open.                                                         |
| The system **prompts** you to verify the deletion.                  | The system **will prompt** you to verify the deletion.                       |
| After you log in, your account **begins** the verification process. | After you log in, your account **will then begin** the verification process. |

### Be positive

Write positive sentences instead of negative sentences. Positive sentences are easier for the user to understand and are usually shorter than negative sentences.

| Use                                                               | Don't use                                                             |
| ----------------------------------------------------------------- | --------------------------------------------------------------------- |
| The visualization updates with data after you click Apply.        | The visualization won't update with data until you click Apply.       |
| Remember to involve your users in the dashboard creation process. | Don't forget to involve your users in the dashboard creation process. |

## Write scannable content

Users often scan content rather than read. Long blocks of text minimize readability - they bury information and are particularly uninviting online.

Use the following techniques to make content easier to scan.

- Write important information first.
- Place actions before explanations.
- Use short, bulleted lists.
- Use headings to divide content.

## Lists

For a discussion about lists and tables, read the [Lists](https://developers.google.com/style/lists) page in the Google developer documentation style guide.

### Ordered lists

Refer to the following guidelines when you write ordered lists.

In general, Grafana Labs guidance matches the [Google Markdown style guide for lists](https://google.github.io/styleguide/docguide/style.html#lists).

Because any list may change, we recommend using lazy numbering for ordered lists.
Markdown is smart enough to let the resulting HTML render your numbered lists correctly.

```markdown
1.  Foo.
1.  Bar.
    1.  Foofoo.
    1.  Barbar.
1.  Baz.
```

The resulting HTML "numbering" type for lazy numbering depends on the website style.
Presently, all levels of nested numbering results in numbers.
For more information about the available types of "numbering" types, refer to [The Ordered List element - HTML: HyperText Markup Language | MDN](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol#attributes).

### Unordered lists

Refer to the following guidelines when you write unordered lists.

- Begin list items with a capital letter unless there is a strong reason not to. For example, when you list case-sensitive parameters.
- If they're complete sentences, end list items with a period. If one item in a list ends with a period, then apply periods to all items in the list.

For more guidance, refer to [Lists](https://developers.google.com/style/lists) in the [Google developer style guide](https://developers.google.com/style/).

### Definition lists

Definition lists often used for providing a list of a term and indented definition. For example, you can use a definition list to document commands and their meanings:

Feature
: Creates a new feature

Features
: Creates more than one feature
: Sample additional features

To create a definition list, add the term on a new line then add a new line with a colon (`:`) along with the definition.

For additional information, refer to the [Definition Lists](https://www.markdownguide.org/extended-syntax/#definition-lists) from the Markdown Guide.

### Sort lists

Sort lists and table rows alphabetically unless the order is important to understanding the information in the list or table.

## Links and references

<!-- vale Grafana.ReferTo = NO -->
<!-- This usage is the advice that the rule is based on. -->

You should use _refer to_ instead of _see_ or _check out_ when referencing another document.

<!-- vale Grafana.ReferTo = YES -->

Give the reader a sense of what to expect in the reference.
Don't use generic references, such as "Refer to [this file]."

As much as possible, use the exact title of the page or section you are linking to as the link text.

For example:

```markdown
For more information about Grafana Labs products, refer to [Grafana documentation](/docs/grafana/latest/).
```

For more comprehensive guidance, refer to [Write useful link text](https://grafana.com/docs/writers-toolki/write/style-guide/useful-link-text/)

## Admonitions

To focus a user's attention, Grafana Labs documentation includes notes, cautions, and warnings.

To standardize styling, each admonition has a special shortcode declaration.
The following sections provide examples how to write each type.
For the complete syntax reference, refer to [Shortcodes]({{< relref "../../../write/shortcodes" >}}).

### Notes

A note provides additional information that the user should be aware of.
Notes are the most common admonition.

The syntax for a note admonition is as follows:

```markdown
{{%/* admonition type="note" */%}}
This page describes a feature for Grafana 9.0 beta.
{{%/* /admonition */%}}
```

On the published page, this note renders as follows:

{{% admonition type="note" %}}
This page describes a feature for Grafana 9.0 beta.
{{% /admonition %}}

### Cautions

A caution warns the user to proceed with caution.
A caution emphasizes a course of action's potential downsides.

The syntax for a caution admonition is as follows:

```markdown
{{%/* admonition type="caution" */%}}
By disabling authentication requirements, anyone can access your Grafana instance.
There is a considerable security risk associated with this.
{{%/* /admonition */%}}
```

On the published page, this caution renders as follows:

{{% admonition type="caution" %}}
By disabling authentication requirements, anyone can access your Grafana instance. There is a considerable security risk associated with this.
{{% /admonition %}}

### Warnings

A warning informs the user not to do something.
Warnings are reserved for actions that could cause harm to hardware, software, or data.

The syntax for a warning admonition is as follows:

```markdown
{{%/* admonition type="warning" */%}}
Don't back up your dashboards in Grafana.
You might not be able to recover a dashboard if it's deleted.
{{%/* /admonition */%}}
```

On the published page, this warning renders as follows:

{{% admonition type="warning" %}}
Don't back up your dashboards in Grafana.
You might not be able to recover a dashboard if it's deleted.
{{% /admonition %}}

## Semantic line breaks

The [Semantic Line Breaks organization](https://sembr.org/) suggests adding line breaks in your writing.
We find this can help with reviews.

Note that we don't require that this suggestion be followed.
We share the Semantic Line Breaks rules so that you can discuss with your review team and adapt to your circumstance.

With line breaks:

```markdown
When you write in active voice, you identify the subject of the sentence and the action that the subject performs.
For example, “John drove the car” is active voice because it is clear that John (the subject) performed an action (drove).
The passive voice variation is “The car was driven by John.”
```

Without line breaks:

```markdown
When you write in active voice, you identify the subject of the sentence and the action that the subject performs. For example, “John drove the car” is active voice because it is clear that John (the subject) performed an action (drove). The passive voice variation is “The car was driven by John.”
```

The output is the same in both cases.
However, the first is easier to review and edit and is less subject to the screen and text editor settings of each contributor.

## Text formatting

It's a good idea to take a consistent approach to bold, italic, and other text formats.

### Bold

Use bold formatting (\*\* in Markdown), for UI elements.
Don't use bold to draw attention to a word or phrase within a sentence, instead use italic emphasis.

It's OK to use bold for a the first sentence in unordered lists that are followed by more information.

For example,

- **Thing**: about the thing.

### Italic

Use italics formatting (\_ in Markdown), to emphasize a specific word or phrase. This is particularly useful when defining a term for the first time.

For example:

> The Prometheus data model is arranged into _metrics_ that consist of a _timestamp_ and a _sample_.

### Code

Use code formatting (\` in Markdown) to refer to:

- File names
- Configuration options
- User input
- Code in text/inline text
- Class and method names, status codes, and console output
