---
title: Style conventions
description: Style conventions
aliases:
  - /docs/writers-toolkit/latest/style-guide/style-conventions/
weight: 200
keywords:
  - style conventions
---

# Style conventions

This section includes a non-exhaustive list of technical writing techniques and styles to consider when you write technical content. For questions you might have that are not addressed in these guidelines topic, refer to
[Google developer documentation style guide](https://developers.google.com/style).

## Focus on user goals

Before you begin writing, clearly identify the goal of the user, and write content that supports the user reaching that goal.

- Do not document implementation details, specifications, or back end system operations that have no clear consequence to the user.
  - Providing unnecessary information can lead to bloated content that forces the user to determine which content is relevant.
- Avoid marketing clichés and hyperbole.
  - Instead, use evidence-based or quantifiable language to focus and refine the information, and to offer value propositions.

## Address users clearly

To address users clearly and directly, write directives in the imperative second person where the unspoken 'you' is implied. Avoid making “the user” or their role (“system administrator”) the subject of a sentence, as in: “Users configure the Cloud using a secure shell.”.

Exception:

- You can use first person in UI elements that are specific to the user, such as “My profile" or “My account."

| Use                                                            | Don't use                                                          |
| -------------------------------------------------------------- | ------------------------------------------------------------------ |
| Click **Yes** to accept the license agreement.                 | The license agreement is accepted when you click **Yes**.    |
| To create a dashboard, add a panel and specify a data source. | To create a dashboard, you add a panel and specify a data source. |

### Write in active voice

When you write in _active voice_, you identify the subject of the sentence and the action that the subject performs. For example, "John drove the car" is active voice because it is clear that John (the subject) performed an action (drove). The passive voice variation is "The car was driven by John."

| Use                                               | Don't use                                                               |
| ------------------------------------------------- | ----------------------------------------------------------------------- |
| After you create a dashboard, add a panel.       | After the dashboard has been created, the panel can be added.          |
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

Make content relevant to the user's context. The more familiar you are with the user’s context, the better you can communicate without using a lot of words.

### Write in present tense

When you write in present tense, you avoid words such as have, has, had, been, should, would, and will.

| Use                                                                | Don't use                                                                    |
| ------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| The panel opens. | The panel will open. |
| The system **prompts** you to verify the deletion.                 | The system **will prompt** you to verify the deletion.                       |
| After you log in, your account **begins** the verification proces. | After you log in, your account **will then begin** the verification process. |

### Be positive

Write positive sentences instead of negative sentences. Positive sentences are easier for the user to understand and are usually shorter than negative sentences.

| Use                                                                    | Don't use                                                                                           |
| ---------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| The visualization updates with data after you click Apply. | The visualization won't update with data until you click Apply.                                   |
| Remember to involve your users in the dashboard creation process.     | Don't forget to involve your users in the dashboard creation process.                              |

## Write scannable content

Users often scan content rather than read. Long blocks of text minimize readability - they bury information and are particularly uninviting online.

Use the following techniques to enhance the scannability of content.

- Write important information first.
- Place actions before explanations.
- Use short, bulleted lists.
- Use headings to divide content.

## Unordered lists

Refer to the following guidelines when you write unordered lists.

For more guidance, refer to [Lists](https://developers.google.com/style/lists) in the [Google developer style guide](https://developers.google.com/style/).

- Begin list items with a capital letter unless there is a strong reason not to. For example, when you list case-sensitive parameters.
- If they are complete sentences, end list items with a period. If one item in a list ends with a period, then apply periods to all items in the list.

## Links and references

You should use "Refer to" instead of "See" or "Check out" when referencing another document.

Give the reader a sense of what to expect in the reference. Don't use blind references, such as "Refer to [this file]."

As much as possible, use the exact title of the page or section you are linking to as the link text.

**Example:** For more information about Grafana Labs products, refer to [Grafana documentation](https://grafana.com/docs/grafana/latest/).

## Admonitions

To focus a user's attention, Grafana Labs documentation includes notes, tips, cautions, and warnings.

### Notes

The most common admonition is a note. A note provides additional information that the user should be aware of. 

For example:

> **Note:** This page describes a feature for Grafana 9.0 beta.

### Tips

A tip describes a more efficient or alternate way of doing something. Tips are rarely used.

### Cautions

A caution warns the user to proceed with caution. A caution emphasizes a course of action's potential downsides.

> **Caution:** By disabling authentication requirements, anyone can access your Grafana instance. There is a considerable security risk associated with this.

### Warnings

A warning informs the user not to do something For example:

> **Warning:** You cannot back up your dashboards in Grafana. You might not be able to recover a dashboard if it is deleted.

## Command line

Use the following conventions when you include command line commands in technical content.

- Do not assume everyone is using Linux. Make sure instructions include enough information for Windows and Mac users to successfully complete procedures.

- Do not add `$` before commands. Make it easy for users to copy and paste commands.

  - **Right:** `sudo yum install grafana`
  - **Wrong:** `$ sudo yum install grafana`

- Include `sudo` before commands that require `sudo` to work.

For terminal examples and Grafana configuration, use a `bash` code block:

```bash
sudo yum install grafana
```

If your command-line instructions include a combination of input and output lines, use separate code blocks for input and output, and use a `console` code block for the output.

```bash
cat ~/.ssh/my-ssh-key.pub
```

The output is similar to the following:

```console
ssh-rsa KEY_VALUE USERNAME
```

For HTTP request/response, use an `http` code block:

```http
GET /api/dashboards/id/1/permissions HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
