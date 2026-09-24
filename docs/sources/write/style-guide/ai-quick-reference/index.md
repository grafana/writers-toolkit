---
title: AI quick reference
menuTitle: AI quick reference
description: Concise style rules for AI agents and LLM-powered documentation skills.
weight: 700
---

# AI quick reference

This page is a concise reference for AI agents writing or reviewing Grafana documentation.
It summarizes the [style guide](https://grafana.com/docs/writers-toolkit/write/style-guide/) and the [Google developer documentation style guide](https://developers.google.com/style).
If a rule isn't here or a case is ambiguous, refer to the full page before you guess.

## Ground every claim

Back every claim about product behavior, defaults, UI, or permissions with a source file or page.
If you can't verify a claim, mark it `[UNVERIFIED]` rather than guessing.

## Audience

- Write for Grafana Labs users, not staff.
- Don't document how to develop on the project.
- Don't document deployment for Grafana Cloud products.

## Product names

- Use long product names with "Grafana" in overviews; use short names without "Grafana" in the body.
- Always write "Grafana Cloud," never just "Cloud."
- List signal types in this order: metrics, logs, traces, profiles.

## Voice and tone

- Write in second person, active voice, and present tense.
- Use contractions: "isn't", "don't", "you're."
- Be confident, not boastful. Avoid "easy", "simple", "just", and marketing clichés.
- Prefer positive framing over negative.
- Follow "Every Page is Page One": each page stands on its own.
- Define or plainly describe a term before you use it; don't assume the reader knows Grafana jargon.

## Sentences and paragraphs

- Prefer short sentences and paragraphs.
- Cut filler: "there is", "there are", "in order to", "it is important to."
- Don't use lists as a substitute for paragraphs.

## Capitalization and punctuation

- Use sentence case for titles, headings, and UI element references.
- Use the serial (Oxford) comma: "metrics, logs, and traces."
- Don't abbreviate "and" with "&" unless matching the UI.

## Headings

- Start task headings with a verb. Don't use "Step X:" in headings.
- Don't start headings with a gerund.
- Include a short introduction after each heading.
- Structure most content under h2 headings; use h3 for related subsections.

## UI elements

- Bold the label text, not the element type: "Click **Apply**", not "click the **Apply** button."
- Match UI casing exactly.
- Write navigation paths with `>`: "Go to **Alerting** > **Notification policies**."
- Capitalize roles as adjectives; don't bold them.
- Use code formatting for file names, config options, CLI commands, and status codes.

## Links

- Use the exact title of the linked page as link text — not "click here" or "this file."
- Use "refer to," not "see," "check out," or "consult."
- Use relative links for internal pages. End links in `/`, not `.md`.

## Word list and conventions

- Follow Grafana's preferred terms — for example, "data source" (not "datasource"), "self-managed" (not "self-hosted"), and "allowlist"/"blocklist" (not "whitelist"/"blacklist"). For the full list, refer to [Word list](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/).
- Use `<VARIABLE_NAME>` in code blocks and _VARIABLE_NAME_ in prose.
- Write for an international audience: avoid idioms and directional language such as "on the left."

## Related resources

- [Style guide](https://grafana.com/docs/writers-toolkit/write/style-guide/)
- [Style conventions](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/)
- [Voice and tone guidelines](https://grafana.com/docs/writers-toolkit/write/style-guide/voice-tone-guidelines/)
- [Google developer documentation style guide](https://developers.google.com/style)
