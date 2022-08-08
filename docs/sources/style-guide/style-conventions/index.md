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

Effective technical communication is simple and direct.

- Don't use a long word where a short word will do, for example, “use,” not “utilize.”
  - When possible, "use" and its variants (utilize, make use of) as an imperative should be replaced with a more descriptive action.
- Write short sentences that are fewer than 25 words.
  - If it is possible to cut a word without losing meaning, do so (typical culprits: there is; there are; in order to; it is important to; keep in mind).
  - If you find yourself writing long sentences, consider writing smaller complete phrases or change the format, such as using a bulleted list.
- Use simple verbs and tenses.
- Use the simplest term possible, considering the characteristics of your audience.
- Avoid buzzwords and jargon.
- Write paragraphs that are three sentences or fewer.
  - Make the text more concise, use more headings, or both.

Focus on the user’s context and make content relevant. The more familiar you are with the user's context, the better you can communicate without using a lot of words.

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

When referencing another document, use "Refer to" rather than alternatives such as "See" or "Check out."

Always give the reader some idea of what to expect in the reference. Avoid blind references, such as, "Refer to [this file]."

When possible, use the exact title of the page or section you are linking to as the link text.

**Example**
Refer to the [Grafana documentation](https://grafana.com/docs/grafana/latest/) for information about how to use Grafana Labs products.

## Notes, tips, cautions, and warnings

Grafana documentation uses notes, tips, cautions, and warnings to focus the user's attention. Notes are the most commonly used admonition.

### Notes

Notes provide additional information that the user should be extra aware of. For example:

> **Note:** This page describes a feature for Grafana 8.0 beta.

### Tips

Tips describe alternate or more efficient ways of doing things. Rarely used.

### Cautions

Cautions warn the user that they should proceed with caution. Use cautions to emphasize the potential downside of a course of action.

> **Caution:** If you turn off authentication requirements, then anyone can access your Grafana instance. This poses a considerable security risk.

### Warnings

Warnings tell the user not to do something. For example:

> **Warning:** Grafana does not back up your dashboards. If you delete a dashboard, then you might not be able to recover it.

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
