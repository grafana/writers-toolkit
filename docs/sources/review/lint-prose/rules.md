---
date: "2024-06-25"
description: A description of every Grafana Labs prose linting rule.
menuTitle: Rules
review_date: "2025-11-13"
title: Vale rules
---

# Vale rules

<!-- These are our style rules. -->
<!-- vale Grafana.We = NO -->

Vale codifies our style guide into a series of rules that can be checked against your prose.
The following is a list of all the rules that we've defined.

<!-- This page breaks a number of rules in demonstrating them. -->
<!-- vale Grafana.Acronyms = NO -->
<!-- vale Grafana.Admin = NO -->
<!-- vale Grafana.Admonitions = NO -->
<!-- vale Grafana.Agentless = NO -->
<!-- vale Grafana.AllowsTo = NO -->
<!-- vale Grafana.AltText = NO -->
<!-- vale Grafana.AmazonCloudWatch = NO -->
<!-- vale Grafana.AmazonProductNames = NO -->
<!-- vale Grafana.AndOr = NO -->
<!-- vale Grafana.ApacheProjectNames = NO -->
<!-- vale Grafana.Archives = NO -->
<!-- vale Grafana.CHANGELOG = NO -->
<!-- vale Grafana.CommandLinePrompts = NO -->
<!-- vale Grafana.DatadogProxy = NO -->
<!-- vale Grafana.DialogBox = NO -->
<!-- vale Grafana.DocumentationTeam = NO -->
<!-- vale Grafana.DropDown = NO -->
<!-- vale Grafana.EndToEnd = NO -->
<!-- vale Grafana.Exclamation = NO -->
<!-- vale Grafana.Gerunds = NO -->
<!-- vale Grafana.GoogleAMPM = NO -->
<!-- vale Grafana.GoogleContractions = NO -->
<!-- vale Grafana.GoogleDateFormat = NO -->
<!-- vale Grafana.GoogleEllipses = NO -->
<!-- vale Grafana.GoogleEmDash = NO -->
<!-- vale Grafana.GoogleEnDash = NO -->
<!-- vale Grafana.GoogleFirstPerson = NO -->
<!-- vale Grafana.GoogleGender = NO -->
<!-- vale Grafana.GoogleGenderBias = NO -->
<!-- vale Grafana.GoogleHeadingPunctuation = NO -->
<!-- vale Grafana.GoogleLyHyphens = NO -->
<!-- vale Grafana.GoogleOptionalPlurals = NO -->
<!-- vale Grafana.GoogleOxfordComma = NO -->
<!-- vale Grafana.GooglePassive = NO -->
<!-- vale Grafana.GooglePeriods = NO -->
<!-- vale Grafana.GoogleProductNames = NO -->
<!-- vale Grafana.GoogleRanges = NO -->
<!-- vale Grafana.GoogleSemicolons = NO -->
<!-- vale Grafana.GoogleSlang = NO -->
<!-- vale Grafana.GoogleSpacing = NO -->
<!-- vale Grafana.GoogleSpelling = NO -->
<!-- vale Grafana.GoogleWill = NO -->
<!-- vale Grafana.GrafanaCom = NO -->
<!-- vale Grafana.Headings = NO -->
<!-- vale Grafana.Kubernetes = NO -->
<!-- vale Grafana.Latin = NO -->
<!-- vale Grafana.MetaMonitoring = NO -->
<!-- vale Grafana.OAuth = NO -->
<!-- vale Grafana.OK = NO -->
<!-- vale Grafana.Ordinal = NO -->
<!-- vale Grafana.Paragraphs = NO -->
<!-- vale Grafana.Parentheses = NO -->
<!-- vale Grafana.Please = NO -->
<!-- vale Grafana.ProductPossessives = NO -->
<!-- vale Grafana.PrometheusExporters = NO -->
<!-- vale Grafana.Quickstart = NO -->
<!-- vale Grafana.README = NO -->
<!-- vale Grafana.React = NO -->
<!-- vale Grafana.ReadabilityAutomatedReadability = NO -->
<!-- vale Grafana.ReadabilityColemanLiau = NO -->
<!-- vale Grafana.ReadabilityFleschKincaid = NO -->
<!-- vale Grafana.ReadabilityFleschReadingEase = NO -->
<!-- vale Grafana.ReadabilityGunningFog = NO -->
<!-- vale Grafana.ReadabilityLIX = NO -->
<!-- vale Grafana.ReadabilitySMOG = NO -->
<!-- vale Grafana.ReferTo = NO -->
<!-- vale Grafana.Relref = NO -->
<!-- vale Grafana.RepeatedWords = NO -->
<!-- vale Grafana.SQL = NO -->
<!-- vale Grafana.SelfManaged = NO -->
<!-- vale Grafana.Shortcodes = NO -->
<!-- vale Grafana.Simple = NO -->
<!-- vale Grafana.SmartQuotes = NO -->
<!-- vale Grafana.Spelling = NO -->
<!-- vale Grafana.Timeless = NO -->
<!-- vale Grafana.We = NO -->
<!-- vale Grafana.Wish = NO -->
<!-- vale Grafana.WordList = NO -->

## Errors

The following rules are considered errors and must be fixed.

### Grafana.GoogleAMPM

Extends: existence

Use 'AM' or 'PM' (preceded by a space).

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\d{1,2}[AP]M\b`
- `\d{1,2} ?[ap]m\b`
- `\d{1,2} ?[aApP]\.[mM]\.`

[More information ->](https://developers.google.com/style/word-list)

### Grafana.GoogleDateFormat

Extends: existence

Use 'July 31, 2016' format, not _`<CURRENT TEXT>`_.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\d{1,2}(?:\.|/)\d{1,2}(?:\.|/)\d{4}`
- `\d{1,2} (?:Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)|May|Jun(?:e)|Jul(?:y)|Aug(?:ust)|Sep(?:tember)?|Oct(?:ober)|Nov(?:ember)?|Dec(?:ember)?) \d{4}`

[More information ->](https://developers.google.com/style/dates-times)

### Grafana.GoogleEmDash

Extends: existence

Don't put a space before or after a dash.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\s[—–]\s`

[More information ->](https://developers.google.com/style/dashes)

### Grafana.GoogleEnDash

Extends: existence

Use an em dash ('—') instead of '–'.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `–`

[More information ->](https://developers.google.com/style/dashes)

### Grafana.GoogleGender

Extends: existence

Don't use _`<CURRENT TEXT>`_ as a gender-neutral pronoun.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `he/she`
- `s/he`
- `\(s\)he`

[More information ->](https://developers.google.com/style/pronouns#gender-neutral-pronouns)

### Grafana.GoogleGenderBias

Extends: substitution

Consider using _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text                  | Replacement text              |
| ----------------------------- | ----------------------------- |
| `(?:alumnae\|alumni)`         | `graduates`                   |
| `(?:alumna\|alumnus)`         | `graduate`                    |
| `air(?:m[ae]n\|wom[ae]n)`     | `pilot(s)`                    |
| `anchor(?:m[ae]n\|wom[ae]n)`  | `anchor(s)`                   |
| `authoress`                   | `author`                      |
| `camera(?:m[ae]n\|wom[ae]n)`  | `camera operator(s)`          |
| `door(?:m[ae]\|wom[ae]n)`     | `concierge(s)`                |
| `draft(?:m[ae]n\|wom[ae]n)`   | `drafter(s)`                  |
| `fire(?:m[ae]n\|wom[ae]n)`    | `firefighter(s)`              |
| `fisher(?:m[ae]n\|wom[ae]n)`  | `fisher(s)`                   |
| `fresh(?:m[ae]n\|wom[ae]n)`   | `first-year student(s)`       |
| `garbage(?:m[ae]n\|wom[ae]n)` | `waste collector(s)`          |
| `lady lawyer`                 | `lawyer`                      |
| `ladylike`                    | `courteous`                   |
| `mail(?:m[ae]n\|wom[ae]n)`    | `mail carriers`               |
| `man and wife`                | `husband and wife`            |
| `man enough`                  | `strong enough`               |
| `mankind`                     | `human kind\|humanity`        |
| `manmade`                     | `manufactured`                |
| `manpower`                    | `personnel`                   |
| `middle(?:m[ae]n\|wom[ae]n)`  | `intermediary`                |
| `news(?:m[ae]n\|wom[ae]n)`    | `journalist(s)`               |
| `ombuds(?:man\|woman)`        | `ombuds`                      |
| `oneupmanship`                | `upstaging`                   |
| `poetess`                     | `poet`                        |
| `police(?:m[ae]n\|wom[ae]n)`  | `police officer(s)`           |
| `repair(?:m[ae]n\|wom[ae]n)`  | `technician(s)`               |
| `sales(?:m[ae]n\|wom[ae]n)`   | `salesperson or sales people` |
| `service(?:m[ae]n\|wom[ae]n)` | `soldier(s)`                  |
| `steward(?:ess)?`             | `flight attendant`            |
| `tribes(?:m[ae]n\|wom[ae]n)`  | `tribe member(s)`             |
| `waitress`                    | `waiter`                      |
| `woman doctor`                | `doctor`                      |
| `woman scientist[s]?`         | `scientist(s)`                |
| `work(?:m[ae]n\|wom[ae]n)`    | `worker(s)`                   |

[More information ->](https://developers.google.com/style/inclusive-documentation)

### Grafana.GoogleLyHyphens

Extends: existence

_`<CURRENT TEXT>`_ doesn't need a hyphen.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\b[^\s-]+ly-\w+\b`

[More information ->](https://developers.google.com/style/hyphens)

### Grafana.GoogleOptionalPlurals

Extends: existence

Don't use plurals in parentheses such as in _`<CURRENT TEXT>`_.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\b\w+\(s\)`

[More information ->](https://developers.google.com/style/plurals-parentheses)

### Grafana.GooglePeriods

Extends: existence

Don't use periods with acronyms or initialisms such as _`<CURRENT TEXT>`_.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\b(?:[A-Z]\.){3,}`

[More information ->](https://developers.google.com/style/abbreviations)

### Grafana.GoogleSlang

Extends: existence

Don't use internet slang abbreviations such as _`<CURRENT TEXT>`_.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `tl;dr`
- `ymmv`
- `rtfm`
- `imo`
- `fwiw`

[More information ->](https://developers.google.com/style/abbreviations)

### Grafana.GoogleSpacing

Extends: existence

_`<CURRENT TEXT>`_ should have one space.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `[a-z][.?!] {2,}[A-Z]`
- `[a-z][.?!][A-Z]`

[More information ->](https://developers.google.com/style/sentence-spacing)

### Grafana.Latin

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text | Replacement text |
| ------------ | ---------------- |
| `e\.?g[,.]?` | `for example`    |
| `i\.?e[,.]?` | `that is`        |

[More information ->](https://developers.google.com/style/abbreviations#dont-use)

### Grafana.Ordinal

Extends: existence

For ordinals, write out first through ninth. For 10th on, use numerals.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `1st`
- `2nd`
- `3rd`
- `4th`
- `5th`
- `6th`
- `7th`
- `8th`
- `9th`

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#numbers)

### Grafana.Paragraphs

Extends: script

br elements must be used only for line breaks that are actually part of the content, as in poems or addresses.

In tables, instead of br elements, use p for paragraphs, and ul or ol for list items.

[More information ->](https://html.spec.whatwg.org/multipage/text-level-semantics.html#the-br-element)

### Grafana.Please

Extends: existence

It's great to be polite, but using 'please' in a set of instructions is overdoing the politeness.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `please`

[More information ->](https://developers.google.com/style/tone#politeness)

### Grafana.ReferTo

Extends: substitution

When linking in Markdown, use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text   | Replacement text |
| -------------- | ---------------- |
| `Check out \[` | `Refer to [`     |
| `See \[`       | `Refer to [`     |
| `check out \[` | `refer to [`     |
| `see \[`       | `refer to [`     |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#links-and-references)

### Grafana.Relref

Extends: script

Don't use the relref shortcode for any links as they don't follow redirects and have confusing semantics.

Instead use one of the links described in [links](https://grafana.com/docs/writers-toolkit/write/links/).

[More information ->](https://grafana.com/docs/writers-toolkit/write/links/)

### Grafana.RepeatedWords

Extends: repetition

_`<CURRENT TEXT>`_ is repeated

### Grafana.Spelling

Extends: spelling

Did you really mean _`<CURRENT TEXT>`_?

The Grafana dictionary might not know of this word yet.

To add a new word, refer to [Add words to the Grafana Labs dictionary](https://grafana.com/docs/writers-toolkit/review/lint-prose/dictionary/add-words/).
Alternatively, raise an [issue](https://github.com/grafana/writers-toolkit/issues/new?title=Grafana.Spelling%%3A%%20%[1]s) and a maintainer will add it for you.

For UI elements, use [bold formatting](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#bold).
The spell checker doesn't check words with bold formatting.

For paths; configuration; user input; code; class, method, and variable names; statuscodes; and console output, use [code formatting](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#bold).
The spell checker doesn't check words with code formatting.

## Warnings

The following rules are warnings and may need to be fixed or otherwise require consideration.

### Grafana.Admin

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_ unless it's the name of the UI label like in the Grafana 'Admin' role.

| Current text | Replacement text |
| ------------ | ---------------- |
| `admin`      | `administrator`  |

[More information ->](https://developers.google.com/style/word-list#admin)

### Grafana.Admonitions

Extends: script

Prefer the `admonition` shortcode over blockquotes.

The admonition shortcode renders its content in a blockquote with consistent styling across the website.

[More information ->](https://grafana.com/docs/writers-toolkit/write/shortcodes/#admonition)

### Grafana.Agentless

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

Grafana Agent has been replaced by Grafana Alloy, so you shouldn't use agent-based terminology.

If you're talking about why and how to send signals directly from an application to Grafana Cloud, prefer no-collector to agentless.

This is consistent with [OTel documentation](https://opentelemetry.io/docs/collector/deployment/no-collector/).

| Current text | Replacement text |
| ------------ | ---------------- |
| `agentless`  | `no-collector`   |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#no-collector)

### Grafana.AllowsTo

Extends: substitution

Did you mean _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_?

Allows to is a common wording error.

| Current text | Replacement text                      |
| ------------ | ------------------------------------- |
| `allows to`  | `allows you to\|makes it possible to` |

### Grafana.AltText

Extends: script

All images must have alt text.

[More information ->](https://grafana.com/docs/writers-toolkit/write/image-guidelines/#alt-text)

### Grafana.AmazonCloudWatch

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text     | Replacement text    |
| ---------------- | ------------------- |
| `AWS CloudWatch` | `Amazon CloudWatch` |
| `Cloudwatch`     | `CloudWatch`        |
| `aws CloudWatch` | `Amazon CloudWatch` |
| `cloudWatch`     | `CloudWatch`        |
| `cloudwatch`     | `CloudWatch`        |

[More information ->](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/WhatIsCloudWatch.html)

### Grafana.AmazonProductNames

Extends: conditional

Use the full Amazon product name in the first instance.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#amazon-products)

### Grafana.AndOr

Extends: existence

Avoid writing and/or except when space is limited, such as in tables.

Often, 'and' implies 'or', so you don't need to write both words.

If you need to specify both in your content, write something like "You can export raw events, processed events, or both."

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `and/or`

[More information ->](https://developers.google.com/style/slashes#and-or)

### Grafana.ApacheProjectNames

Extends: conditional

Use the full Apache project name in the first instance.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#apache-projects)

### Grafana.CHANGELOG

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_ unless you're referring to a specific file which has that spelling.

| Current text       | Replacement text |
| ------------------ | ---------------- |
| `[Cc]hangelog`     | `CHANGELOG`      |
| `[Cc]hangelog\.md` | `CHANGELOG.md`   |
| `[Cc]hangelogs`    | `CHANGELOGs`     |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#changelog)

### Grafana.CommandLinePrompts

Extends: script

Don't add `$` or `#` as prompts before commands so users can copy and paste them.

Also, don't use `#` to include comments in commands.
That explanation should be outside of the code block.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/write-for-developers/#command-lines)

### Grafana.DatadogProxy

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text              | Replacement text |
| ------------------------- | ---------------- |
| `Datadog Proxy`           | `Datadog proxy`  |
| `[tT]he [Dd]atadog proxy` | `Datadog proxy`  |

### Grafana.DialogBox

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ rather than _`<CURRENT TEXT>`_.

| Current text         | Replacement text   |
| -------------------- | ------------------ |
| `dialog box appears` | `dialog box opens` |
| `dialog(?! box)`     | `dialog box`       |
| `modal`              | `dialog box`       |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#dialog-box)

### Grafana.DocumentationTeam

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ rather than _`<CURRENT TEXT>`_.

| Current text                          | Replacement text                      |
| ------------------------------------- | ------------------------------------- |
| `[Dd]ocs? (?:[Ss]quad\|[Tt]eam)`      | `the Grafana Labs documentation team` |
| `[Dd]ocumentation (?:[Ss]quad\|Team)` | `the Grafana Labs documentation team` |

### Grafana.EndToEnd

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text | Replacement text |
| ------------ | ---------------- |
| `[eE]2[eE]`  | `end-to-end`     |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#end-to-end)

### Grafana.Exclamation

Extends: existence

Avoid exclamation points in text, except in rare really exciting moments.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\w+!(?:\s|$)`

[More information ->](https://developers.google.com/style/tone#some-things-to-avoid-where-possible)

### Grafana.Gerunds

Extends: script

For a task-based heading, start with a [bare infinitive](https://en.wikipedia.org/wiki/Infinitive#English), also known as a plain form or [base form](https://en.wikipedia.org/wiki/English_verbs#Base_form) verb.
In English, the imperative mood also uses the base form verb, so it looks the same as the bare infinitive.

Task-based headings are frequently used in quickstarts, how-to documents, and tutorials.

For a conceptual or non-task-based heading, use a [noun phrase](https://en.wikipedia.org/wiki/Noun_phrase) that doesn't start with an -ing verb.

Noun-phrase headings are frequently used in concept documentation.

[More information ->](https://developers.google.com/style/headings#heading-and-title-text)

### Grafana.GoogleEllipses

Extends: existence

In general, don't use an ellipsis.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\.\.\.`

[More information ->](https://developers.google.com/style/ellipses)

### Grafana.GoogleFirstPerson

Extends: existence

Avoid first-person pronouns such as _`<CURRENT TEXT>`_.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `(?:^|\s)I\s`
- `(?:^|\s)I,\s`
- `\bI'm\b`
- `\bme\b`
- `\bmy\b`
- `\bmine\b`

[More information ->](https://developers.google.com/style/pronouns#personal-pronouns)

### Grafana.GoogleHeadingPunctuation

Extends: existence

Don't put a period at the end of a heading.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `[a-z0-9][.]\s*$`

[More information ->](https://developers.google.com/style/capitalization#capitalization-in-titles-and-headings)

### Grafana.GoogleProductNames

Extends: conditional

Use the full Google product name in the first instance.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#google-products)

### Grafana.GoogleRanges

Extends: existence

Don't add words such as 'from' or 'between' to describe a range of numbers.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `(?:from|between)\s\d+\s?-\s?\d+`

[More information ->](https://developers.google.com/style/hyphens)

### Grafana.GoogleSpelling

Extends: existence

In general, use American spelling instead of _`<CURRENT TEXT>`_.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `(?:\w+)nised?`
- `colour`
- `labour`
- `centre`

[More information ->](https://developers.google.com/style/spelling)

### Grafana.GoogleWill

Extends: existence

Avoid using _`<CURRENT TEXT>`_.

Use present tense for statements that describe general behavior that's not associated with a particular time.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `will`

[More information ->](https://developers.google.com/style/tense)

### Grafana.GrafanaCom

Extends: existence

Don't use `grafana.com`, instead use one of the following:

- If you're talking about Grafana Cloud, use `Grafana Cloud`.
- If you're talking about the company, use `Grafana Labs`.
- If you're linking to a page on the website, use the page title or the full URL including scheme. For example, `https://grafana.com/`.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `grafana\.com`

### Grafana.Headings

Extends: capitalization

Use sentence-style capitalization for _`<CURRENT TEXT>`_.

Vale considers multi-word exceptions such as _Grafana Enterprise Metrics_ as a single correctly cased word.

If your heading contains capitalized words that represent product names, you need to add those words to the Grafana dictionary or the list of static exceptions in https://github.com/grafana/writers-toolkit/blob/main/vale/Headings.jsonnet for them to be considered correctly cased.

[More information ->](https://developers.google.com/style/capitalization#capitalization-in-titles-and-headings)

### Grafana.Kubernetes

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text               | Replacement text        |
| -------------------------- | ----------------------- |
| `[Kk]ubectl`               | `kubectl`               |
| `[Kk]ubelet`               | `kubelet`               |
| `[Kk]ubernetes deployment` | `Kubernetes Deployment` |
| `cron job`                 | `CronJob`               |
| `d[ae][ae]mon[Ss]et`       | `DaemonSet`             |
| `pod`                      | `Pod`                   |
| `replica[Ss]et`            | `ReplicaSet`            |
| `stateful[Ss]et`           | `StatefulSet`           |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#kubernetes-objects)

### Grafana.MetaMonitoring

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text       | Replacement text  |
| ------------------ | ----------------- |
| `meat ?monitoring` | `meta-monitoring` |
| `meta ?monitoring` | `meta-monitoring` |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#meta-monitoring)

### Grafana.OK

Extends: existence

Don't use any variation of okay in prose.
The exceptions are when you’re referencing or quoting:

- A user interface
- HTTP status codes or other code

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `O.K.`
- `OK`
- `ok`
- `Ok`
- `Okay`
- `okay`
- `A-OK`
- `hokay`
- `k`
- `keh`
- `kk`
- `M'kay`
- `oka`
- `okeh`
- `Okie dokie`
- `Okily Dokily`

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#ok-okay)

### Grafana.ProductPossessives

Extends: existence

Don't form a possessive from a feature name, product name, or trademark, regardless of who owns it.
Instead, use the name as a modifier or rewrite to use a word like of to indicate the relationship.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `ADOT's`
- `AI Observability's`
- `Agent's`
- `Alloy's`
- `ARN's`
- `Asserts'`
- `AWS's`
- `AWS Distro for OpenTelemetry Collector's`
- `AWS X-Ray's`
- `Beyla's`
- `BoringCrypto's`
- `CentOS's`
- `CloudWatch's`
- `Codespaces'`
- `Data Firehose's`
- `Databricks'`
- `Datadog's`
- `Dynatrace's`
- `EKS's`
- `Elastic Kubernetes Service's`
- `Entra's`
- `Figma's`
- `Firehose's`
- `FreeBSD's`
- `GEM's`
- `GKE's`
- `Git's`
- `GitHub's`
- `GitLab's`
- `GNU's`
- `Grafana's`
- `Gravatar's`
- `Graylog's`
- `Gunicorn's`
- `hashmod's`
- `IBM's`
- `InfluxDB's`
- `Jaeger's`
- `Jira's`
- `JMESPath's`
- `journald's`
- `Jsonnet's`
- `Keycloak's`
- `Kibana's`
- `Killercoda's`
- `Kinesis'`
- `Kotlin's`
- `KQL's`
- `Kubernetes'`
- `Kubernetes Engine's`
- `Kusto's`
- `Kustomize's`
- `LangChain's`
- `launchd's`
- `Logs Drilldown's`
- `Loki's`
- `Lucene's`
- `Markdown's`
- `Memcached's`
- `Metrics Drilldown's`
- `Mesos'`
- `Mimir's`
- `Moodle's`
- `MySQL's`
- `Netlink's`
- `Okta's`
- `OnCall's`
- `OpenAI's`
- `OpenShift's`
- `OpenTelemetry's`
- `Opsgenie's`
- `OTel's`
- `PagerDuty's`
- `Parca's`
- `PDC's`
- `Phlare's`
- `Pinecone's`
- `Podman's`
- `Postgres'`
- `PostgreSQL's`
- `pprof's`
- `Profiles Drilldown's`
- `Prometheus'`
- `Promtail's`
- `Pyroscope's`
- `RCA workbench's`
- `RDS's`
- `Relational Database Service's`
- `React's`
- `Redis'`
- `RHEL's`
- `Rollup's`
- `RudderStack's`
- `Sensu's`
- `Sensu Go's`
- `Splunk's`
- `SSM's`
- `SUSE's`
- `Team Sync's`
- `Tempo's`
- `Thanos'`
- `Threema's`
- `Traces Drilldown's`
- `Velero's`
- `Vite's`
- `VMware's`
- `Webex's`
- `WildFly's`
- `windows_exporter's`
- `YugabyteDB's`
- `Zipkin's`

[More information ->](https://developers.google.com/style/possessives#product,-feature,-and-company-names)

### Grafana.PrometheusExporters

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text       | Replacement text |
| ------------------ | ---------------- |
| `[Nn]ode exporter` | `Node Exporter`  |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#node-exporter)

### Grafana.Quickstart

Extends: substitution

Use the compound adjective _`<REPLACEMENT TEXT>`_ without a hyphen instead of _`<CURRENT TEXT>`_ whether the noun is implied or explicit. For example, you can use _quickstart guide_ or just _quickstart_.

| Current text  | Replacement text |
| ------------- | ---------------- |
| `quick start` | `quickstart`     |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#quickstart)

### Grafana.README

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_ unless you're referring to a specific file which has that spelling.

| Current text    | Replacement text |
| --------------- | ---------------- |
| `[Rr]eadme`     | `README`         |
| `[Rr]eadme\.md` | `README.md`      |
| `[Rr]eadmes`    | `READMEs`        |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#readme)

### Grafana.React

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text            | Replacement text |
| ----------------------- | ---------------- |
| `[Rr]eact[. ]?[Jj][Ss]` | `React`          |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#react)

### Grafana.SQL

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

The article—a or an—that you use before the acronym SQL depends on how the word is pronounced.

When referring to the product Microsoft SQL Server, SQL should be pronounced "sequel".
In this case, use the article 'a', as in "a SQL Server analysis".

When referring to the term in any other context, such as SQL databases, errors, or servers, SQL should be pronounced "ess-cue-el".
In this case, use the article 'an', as in "an SQL error".

| Current text             | Replacement text              |
| ------------------------ | ----------------------------- |
| `[Aa] SQL server`        | `an SQL server\|a SQL Server` |
| `[Aa] SQL(?! [Ss]erver)` | `an SQL`                      |
| `[Aa]n SQL Server`       | `a SQL Server`                |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#sql-structured-query-language)

### Grafana.SelfManaged

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_ when talking about Grafana deployment methods.

| Current text      | Replacement text |
| ----------------- | ---------------- |
| `on-prem(?:ise)?` | `self-managed`   |
| `self-hosted`     | `self-managed`   |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#self-managed)

### Grafana.Shortcodes

Extends: script

Prefer `{{</*` and `*/>}}` instead of `{{%/*` and `*/%}}`

It has the most consistent semantics.

The percent syntax is used for special behavior that isn't required with this shortcode.

[More information ->](https://grafana.com/docs/writers-toolkit/write/shortcodes/#admonition)

### Grafana.Simple

Extends: existence

Avoid the word easy or simple -- what might be simple for you might not be simple for others.

Try eliminating this word from the sentence because usually you can convey the same meaning without it.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `easy(?![ -]to[ -]understand)`
- `easily`
- `simple`
- `simply`

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#simple)

### Grafana.SmartQuotes

Extends: script

Avoid smart quotes in the source file, especially in code blocks.

Replace all smart double quotes like `“` or `”` with `"`.
Replace all smart single quotes like `‘`, `’`, or `ʼ` with `'`.

In some contexts, Unicode characters aren't supported and break configurations.

The website renders paired quotes using smart quotes in paragraphs.

### Grafana.Wish

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text | Replacement text |
| ------------ | ---------------- |
| `wish`       | `need\|want`     |

[More information ->](https://developers.google.com/style/word-list#wish)

### Grafana.WordList

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text                                                      | Replacement text           |
| ----------------------------------------------------------------- | -------------------------- |
| `(?:(?<!Data )Firehose\|Kinesis Data Firehose\|Kinesis Firehose)` | `Data Firehose`            |
| `(?:SHA-1\|HAS-SHA1)`                                             | `SHA-1`                    |
| `(?:WiFi\|wifi)`                                                  | `Wi-Fi`                    |
| `(?:[Oo]penshift\|openShift)`                                     | `OpenShift`                |
| `(?:[eE]-mail)`                                                   | `email`                    |
| `(?:[jJ][mM][eE][sS]p\|jmesP)ath`                                 | `JMESPath`                 |
| `(?:[oO]pentelemetry\|openTelemetry)`                             | `OpenTelemetry`            |
| `(?:alert[Mm]anager\|[Aa]lert [Mm]anager\|AlertManager)`          | `Alertmanager`             |
| `(?:cell ?phone\|smart ?phone)`                                   | `phone\|mobile phone`      |
| `(?:content\|media)-?type`                                        | `media type`               |
| `(?:file ?path\|path ?name)`                                      | `path`                     |
| `(?:file ?path\|path ?name)s`                                     | `paths`                    |
| `(?:github\|gitHub\|Github)`                                      | `GitHub`                   |
| `(?:gitlab\|gitLab\|Gitlab)`                                      | `GitLab`                   |
| `(?:hamburger menu\|kebab menu)`                                  | `menu icon`                |
| `(?:java[Ss]cript\|Javascript)`                                   | `JavaScript`               |
| `(?:kill\|terminate\|abort)`                                      | `stop\|exit\|cancel\|end`  |
| `(?<!kube-)prometheus`                                            | `Prometheus`               |
| `(?<!lambda-)promtail`                                            | `Promtail`                 |
| `GME`                                                             | `GEM`                      |
| `Grafana AI observability`                                        | `Grafana AI Observability` |
| `HTTPs`                                                           | `HTTPS`                    |
| `Influx[Dd]b`                                                     | `InfluxDB`                 |
| `Influxd[Bb]`                                                     | `InfluxDB`                 |
| `Once`                                                            | `After`                    |
| `Pagerduty`                                                       | `PagerDuty`                |
| `RCA Workbench`                                                   | `RCA workbench`            |
| `Rudderstack`                                                     | `RudderStack`              |
| `VMWare`                                                          | `VMware`                   |
| `Vmware`                                                          | `VMware`                   |
| `[Ww]orld [Ww]ide [Ww]eb`                                         | `web`                      |
| `[cC]entos`                                                       | `CentOS`                   |
| `\b(?:[aA]daptive metrics\|adaptive Metrics)\b`                   | `Adaptive Metrics`         |
| `ad[- ]?hoc`                                                      | `free-form\|user-written`  |
| `back[ -]end`                                                     | `backend`                  |
| `blacklist`                                                       | `blocklist`                |
| `blacklisted`                                                     | `blocklisted`              |
| `blacklisting`                                                    | `blocklisting`             |
| `blacklists`                                                      | `blocklists`               |
| `cadvisor`                                                        | `cAdvisor`                 |
| `check[- ]box`                                                    | `checkbox`                 |
| `content type`                                                    | `media type`               |
| `data-?source`                                                    | `data source`              |
| `data-?sources`                                                   | `data sources`             |
| `data[- ]?set`                                                    | `dataset`                  |
| `data[- ]?sets`                                                   | `datasets`                 |
| `datacenter`                                                      | `data center`              |
| `datacenters`                                                     | `data centers`             |
| `de-duplicate`                                                    | `deduplicate`              |
| `de-duplicated`                                                   | `deduplicated`             |
| `de-duplicates`                                                   | `deduplicates`             |
| `de-duplication`                                                  | `deduplication`            |
| `fewer data`                                                      | `less data`                |
| `figma`                                                           | `Figma`                    |
| `file name`                                                       | `filename`                 |
| `file names`                                                      | `filenames`                |
| `firewalls`                                                       | `firewall rules`           |
| `front[ -]end`                                                    | `frontend`                 |
| `front[ -]ends`                                                   | `frontends`                |
| `git`                                                             | `Git`                      |
| `grafana`                                                         | `Grafana`                  |
| `grayed-out`                                                      | `unavailable`              |
| `gunicorn`                                                        | `Gunicorn`                 |
| `in order to`                                                     | `to`                       |
| `influx[Dd][Bb]`                                                  | `InfluxDB`                 |
| `jsonnet`                                                         | `Jsonnet`                  |
| `kotlin`                                                          | `Kotlin`                   |
| `langchain`                                                       | `LangChain`                |
| `left[- ]hand[- ]side`                                            | `left-side`                |
| `log(?:ql\|QL)`                                                   | `LogQL`                    |
| `loki`                                                            | `Loki`                     |
| `lucene`                                                          | `Lucene`                   |
| `markdown`                                                        | `Markdown`                 |
| `memcached`                                                       | `Memcached`                |
| `meta[- ]data`                                                    | `metadata`                 |
| `mix[- ]in`                                                       | `mixin`                    |
| `mysql`                                                           | `MySQL`                    |
| `network IP address`                                              | `internal IP address`      |
| `open-source`                                                     | `open source`              |
| `otel`                                                            | `OTel`                     |
| `otlp`                                                            | `OTLP`                     |
| `pager[dD]uty`                                                    | `PagerDuty`                |
| `phlare`                                                          | `Phlare`                   |
| `postgres`                                                        | `Postgres`                 |
| `postgresql`                                                      | `PostgreSQL`               |
| `prom(?:ql\|QL)`                                                  | `PromQL`                   |
| `redis`                                                           | `Redis`                    |
| `regex[ep]?s`                                                     | `regular expression`       |
| `regexp?`                                                         | `regular expression`       |
| `repo`                                                            | `repository`               |
| `repos`                                                           | `repositories`             |
| `right[- ]hand[- ]side`                                           | `right-side`               |
| `rudderstack`                                                     | `RudderStack`              |
| `sensu`                                                           | `Sensu`                    |
| `sign into`                                                       | `sign in to`               |
| `sqlite`                                                          | `SQLite`                   |
| `style sheet`                                                     | `stylesheet`               |
| `style sheets`                                                    | `stylesheet`               |
| `synch`                                                           | `sync`                     |
| `synched`                                                         | `synced`                   |
| `synching`                                                        | `syncing`                  |
| `tempo`                                                           | `Tempo`                    |
| `the Grafana Agent`                                               | `Grafana Agent`            |
| `the RCA [Ww]orkbench`                                            | `RCA workbench`            |
| `threema`                                                         | `Threema`                  |
| `timeseries`                                                      | `time series\|time-series` |
| `trace(?:ql\|QL)`                                                 | `TraceQL`                  |
| `un(?:check\|select)`                                             | `clear`                    |
| `url`                                                             | `URL`                      |
| `urls`                                                            | `URLs`                     |
| `vmware`                                                          | `VMware`                   |
| `vs\.`                                                            | `versus`                   |
| `webex`                                                           | `Webex`                    |
| `whitelist`                                                       | `allowlist`                |
| `whitelisted`                                                     | `allowlisted`              |
| `whitelisting`                                                    | `allowlisting`             |
| `whitelists`                                                      | `allowlists`               |
| `yugabyte`                                                        | `YugabyteDB`               |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/)

## Suggestions

The following rules are suggestions to consider a certain point of style.

### Grafana.Acronyms

Extends: conditional

Spell out _`<CURRENT TEXT>`_, if it's unfamiliar to the audience.

[More information ->](https://developers.google.com/style/abbreviations)

### Grafana.Archives

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text                           | Replacement text           |
| -------------------------------------- | -------------------------- |
| `[Uu]n(?:archive\|compress\|tar\|zip)` | `extract`                  |
| `[Zz][Ii][Pp](?: file)?`               | `archive\|compressed file` |
| `unzip`                                | `extract`                  |

[More information ->](https://developers.google.com/style/word-list#extract)

### Grafana.DropDown

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

Use drop-down as a modifier rather than as a standalone noun. For example: _drop-down menu_.

| Current text | Replacement text |
| ------------ | ---------------- |
| `drop ?down` | `drop-down`      |

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#drop-down)

### Grafana.GoogleContractions

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text | Replacement text |
| ------------ | ---------------- |
| `are not`    | `aren't`         |
| `cannot`     | `can't`          |
| `could not`  | `couldn't`       |
| `did not`    | `didn't`         |
| `do not`     | `don't`          |
| `does not`   | `doesn't`        |
| `has not`    | `hasn't`         |
| `have not`   | `haven't`        |
| `how is`     | `how's`          |
| `is not`     | `isn't`          |
| `it is`      | `it's`           |
| `should not` | `shouldn't`      |
| `that is`    | `that's`         |
| `they are`   | `they're`        |
| `was not`    | `wasn't`         |
| `we are`     | `we're`          |
| `we have`    | `we've`          |
| `were not`   | `weren't`        |
| `what is`    | `what's`         |
| `when is`    | `when's`         |
| `where is`   | `where's`        |
| `will not`   | `won't`          |

[More information ->](https://developers.google.com/style/contractions)

### Grafana.GoogleOxfordComma

Extends: existence

Use the Oxford comma in _`<CURRENT TEXT>`_.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `(?:[^,]+,){1,}\s\w+\s(?:and|or)`

[More information ->](https://developers.google.com/style/commas)

### Grafana.GooglePassive

Extends: existence

In general, use active voice instead of passive voice (_`<CURRENT TEXT>`_).

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `[\w]+ed`
- `awoken`
- `beat`
- `become`
- `been`
- `begun`
- `bent`
- `beset`
- `bet`
- `bid`
- `bidden`
- `bitten`
- `bled`
- `blown`
- `born`
- `bought`
- `bound`
- `bred`
- `broadcast`
- `broken`
- `brought`
- `built`
- `burnt`
- `burst`
- `cast`
- `caught`
- `chosen`
- `clung`
- `come`
- `cost`
- `crept`
- `cut`
- `dealt`
- `dived`
- `done`
- `drawn`
- `dreamt`
- `driven`
- `drunk`
- `dug`
- `eaten`
- `fallen`
- `fed`
- `felt`
- `fit`
- `fled`
- `flown`
- `flung`
- `forbidden`
- `foregone`
- `forgiven`
- `forgotten`
- `forsaken`
- `fought`
- `found`
- `frozen`
- `given`
- `gone`
- `gotten`
- `ground`
- `grown`
- `heard`
- `held`
- `hidden`
- `hit`
- `hung`
- `hurt`
- `kept`
- `knelt`
- `knit`
- `known`
- `laid`
- `lain`
- `leapt`
- `learnt`
- `led`
- `left`
- `lent`
- `let`
- `lighted`
- `lost`
- `made`
- `meant`
- `met`
- `misspelt`
- `mistaken`
- `mown`
- `overcome`
- `overdone`
- `overtaken`
- `overthrown`
- `paid`
- `pled`
- `proven`
- `put`
- `quit`
- `read`
- `rid`
- `ridden`
- `risen`
- `run`
- `rung`
- `said`
- `sat`
- `sawn`
- `seen`
- `sent`
- `set`
- `sewn`
- `shaken`
- `shaven`
- `shed`
- `shod`
- `shone`
- `shorn`
- `shot`
- `shown`
- `shrunk`
- `shut`
- `slain`
- `slept`
- `slid`
- `slit`
- `slung`
- `smitten`
- `sold`
- `sought`
- `sown`
- `sped`
- `spent`
- `spilt`
- `spit`
- `split`
- `spoken`
- `spread`
- `sprung`
- `spun`
- `stolen`
- `stood`
- `stridden`
- `striven`
- `struck`
- `strung`
- `stuck`
- `stung`
- `stunk`
- `sung`
- `sunk`
- `swept`
- `swollen`
- `sworn`
- `swum`
- `swung`
- `taken`
- `taught`
- `thought`
- `thrived`
- `thrown`
- `thrust`
- `told`
- `torn`
- `trodden`
- `understood`
- `upheld`
- `upset`
- `wed`
- `wept`
- `withheld`
- `withstood`
- `woken`
- `won`
- `worn`
- `wound`
- `woven`
- `written`
- `wrung`

[More information ->](https://developers.google.com/style/voice)

### Grafana.GoogleSemicolons

Extends: existence

Use semicolons judiciously.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `;`

[More information ->](https://developers.google.com/style/semicolons)

### Grafana.OAuth

Extends: substitution

Use _`<REPLACEMENT TEXT>`_ instead of _`<CURRENT TEXT>`_.

| Current text        | Replacement text |
| ------------------- | ---------------- |
| `O[Aa]uth 2(?!\.0)` | `OAuth 2.0`      |
| `O[Aa]uth(?! 2\.0)` | `OAuth 2.0`      |

[More information ->](https://developers.google.com/style/word-list#oauth-20)

### Grafana.Parentheses

Extends: existence

Use parentheses judiciously.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `\(.{4,}\)`

[More information ->](https://developers.google.com/style/parentheses)

### Grafana.ReadabilityAutomatedReadability

Extends: metric

_`<CURRENT TEXT>`_ aim for below 8.

[More information ->](https://en.wikipedia.org/wiki/Automated_readability_index)

### Grafana.ReadabilityColemanLiau

Extends: metric

_`<CURRENT TEXT>`_ aim for below 9.

[More information ->](https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index)

### Grafana.ReadabilityFleschKincaid

Extends: metric

_`<CURRENT TEXT>`_ aim for below 8.

[More information ->](https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests)

### Grafana.ReadabilityFleschReadingEase

Extends: metric

_`<CURRENT TEXT>`_ aim for above 70.

[More information ->](https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests)

### Grafana.ReadabilityGunningFog

Extends: metric

_`<CURRENT TEXT>`_ aim for below 10.

[More information ->](https://en.wikipedia.org/wiki/Gunning_fog_index)

### Grafana.ReadabilityLIX

Extends: metric

_`<CURRENT TEXT>`_ aim for below 35.

[More information ->](<https://en.wikipedia.org/wiki/Lix_(readability_test)>)

### Grafana.ReadabilitySMOG

Extends: metric

_`<CURRENT TEXT>`_ aim for below 10.

[More information ->](https://en.wikipedia.org/wiki/SMOG)

### Grafana.Timeless

Extends: existence

Avoid using _`<CURRENT TEXT>`_ to keep the documentation timeless.

In general, document the current version of a product or feature.

It reduces the maintenance required to keep documentation up to date.
It avoids assuming the reader is familiar with earlier versions of the product.

If you're writing procedural or time-stamped content such as press releases, blog posts, or release notes, such time-based words and phrases are OK.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `as of this writing`
- `currently`
- `does not yet`
- `eventually`
- `existing`
- `future`
- `in the future`
- `latest`
- `new`
- `newer`
- `now`
- `old`
- `older`
- `presently`
- `at present`
- `soon`

[More information ->](https://developers.google.com/style/timeless-documentation)

### Grafana.We

Extends: existence

Use first person plural pronouns like _`<CURRENT TEXT>`_ carefully.

Don't use 'we' when you're talking about the reader, instead use 'you'.

It's OK to use 'we' when you're talking about Grafana Labs.

_`<CURRENT TEXT>`_ was matched by one or more of the following regular expressions:

- `we`
- `we'(?:ve|re)`
- `ours?`
- `us`
- `let's`

[More information ->](https://developers.google.com/style/person#use-first-person-plural-pronouns-carefully)
