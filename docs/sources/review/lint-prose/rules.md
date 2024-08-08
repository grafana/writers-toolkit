---
date: "2024-06-25"
description: A description of every Grafana Labs prose linting rule.
menuTitle: Rules
review_date: "2024-08-08"
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
<!-- vale Grafana.Gerunds = NO -->
<!-- vale Grafana.GoogleAMPM = NO -->
<!-- vale Grafana.GoogleContractions = NO -->
<!-- vale Grafana.GoogleDateFormat = NO -->
<!-- vale Grafana.GoogleEllipses = NO -->
<!-- vale Grafana.GoogleEmDash = NO -->
<!-- vale Grafana.GoogleEnDash = NO -->
<!-- vale Grafana.GoogleExclamation = NO -->
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
<!-- vale Grafana.HTTP = NO -->
<!-- vale Grafana.Headings = NO -->
<!-- vale Grafana.Kubernetes = NO -->
<!-- vale Grafana.Latin = NO -->
<!-- vale Grafana.OAuth = NO -->
<!-- vale Grafana.OK = NO -->
<!-- vale Grafana.Ordinal = NO -->
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
<!-- vale Grafana.RepeatedWords = NO -->
<!-- vale Grafana.SQL = NO -->
<!-- vale Grafana.Shortcodes = NO -->
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

[More information ->](https://developers.google.com/style/word-list)

### Grafana.GoogleDateFormat

Extends: existence

Use 'July 31, 2016' format, not '%s'.

[More information ->](https://developers.google.com/style/dates-times)

### Grafana.GoogleEmDash

Extends: existence

Don't put a space before or after a dash.

[More information ->](https://developers.google.com/style/dashes)

### Grafana.GoogleEnDash

Extends: existence

Use an em dash ('—') instead of '–'.

[More information ->](https://developers.google.com/style/dashes)

### Grafana.GoogleExclamation

Extends: existence

Don't use exclamation points in text.

[More information ->](https://developers.google.com/style/exclamation-points)

### Grafana.GoogleGender

Extends: existence

Don't use '%s' as a gender-neutral pronoun.

[More information ->](https://developers.google.com/style/pronouns#gender-neutral-pronouns)

### Grafana.GoogleGenderBias

Extends: substitution

Consider using '%s' instead of '%s'.

[More information ->](https://developers.google.com/style/inclusive-documentation)

### Grafana.GoogleLyHyphens

Extends: existence

'%s' doesn't need a hyphen.

[More information ->](https://developers.google.com/style/hyphens)

### Grafana.GoogleOptionalPlurals

Extends: existence

Don't use plurals in parentheses such as in '%s'.

[More information ->](https://developers.google.com/style/plurals-parentheses)

### Grafana.GooglePeriods

Extends: existence

Don't use periods with acronyms or initialisms such as '%s'.

[More information ->](https://developers.google.com/style/abbreviations)

### Grafana.GoogleSlang

Extends: existence

Don't use internet slang abbreviations such as '%s'.

[More information ->](https://developers.google.com/style/abbreviations)

### Grafana.GoogleSpacing

Extends: existence

'%s' should have one space.

[More information ->](https://developers.google.com/style/sentence-spacing)

### Grafana.HTTP

Extends: substitution

Use '%s' instead of '%s'.

The HTTP scheme is insecure and all grafana.com links must use HTTPS.

### Grafana.Latin

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://developers.google.com/style/abbreviations#dont-use)

### Grafana.Ordinal

Extends: existence

For ordinals, write out first through ninth. For 10th on, use numerals.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#numbers)

### Grafana.Please

Extends: existence

It's great to be polite, but using 'please' in a set of instructions is overdoing the politeness.

[More information ->](https://developers.google.com/style/tone#politeness)

### Grafana.ReferTo

Extends: substitution

When linking in Markdown, use '%s' instead of '%s'.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/#links-and-references)

### Grafana.RepeatedWords

Extends: repetition

'%s' is repeated

### Grafana.Spelling

Extends: spelling

Did you really mean '%s'?

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

Use administrator instead of admin unless it's the name of the UI label like in the Grafana 'Admin' role.

[More information ->](https://developers.google.com/style/word-list#admin)

### Grafana.Admonitions

Extends: script

Prefer the `admonition` shortcode over blockquotes.

The admonition shortcode renders its content in a blockquote with consistent styling across the website.

[More information ->](https://grafana.com/docs/writers-toolkit/write/shortcodes/#admonition)

### Grafana.Agentless

Extends: substitution

Grafana Agent has been replaced by Grafana Alloy, so you shouldn't use agent-based terminology.

If you're talking about why and how to send signals directly from an application to Grafana Cloud, prefer no-collector to agentless.

This is consistent with [OTel documentation](https://opentelemetry.io/docs/collector/deployment/no-collector/).

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#no-collector)

### Grafana.AllowsTo

Extends: substitution

Did you mean '%s' instead of '%s'?

Allows to is a common wording error.

### Grafana.AltText

Extends: script

All images must have alt text.

[More information ->](https://grafana.com/docs/writers-toolkit/write/image-guidelines/#alt-text)

### Grafana.AmazonCloudWatch

Extends: substitution

Use '%s' instead of '%s'.

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

[More information ->](https://developers.google.com/style/slashes#and-or)

### Grafana.ApacheProjectNames

Extends: conditional

Use the full Apache project name in the first instance.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#apache-projects)

### Grafana.CHANGELOG

Extends: substitution

Use '%s' instead of '%s' unless you're referring to a specific file which has that spelling.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#changelog)

### Grafana.CommandLinePrompts

Extends: script

Don't add `$` or `#` as prompts before commands.
Make it easy for users to copy and paste commands.

Also, don't use `#` to include comments in commands.
That explanation should be outside of the code block.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/write-for-developers/#command-lines)

### Grafana.DatadogProxy

Extends: substitution

Use '%s' instead of '%s'.

### Grafana.DialogBox

Extends: substitution

Use '%s' rather than '%s'.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#dialog-box)

### Grafana.DocumentationTeam

Extends: substitution

Use '%s' rather than '%s'.

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

[More information ->](https://developers.google.com/style/ellipses)

### Grafana.GoogleFirstPerson

Extends: existence

Avoid first-person pronouns such as '%s'.

[More information ->](https://developers.google.com/style/pronouns#personal-pronouns)

### Grafana.GoogleHeadingPunctuation

Extends: existence

Don't put a period at the end of a heading.

[More information ->](https://developers.google.com/style/capitalization#capitalization-in-titles-and-headings)

### Grafana.GoogleOxfordComma

Extends: existence

Use the Oxford comma in '%s'.

[More information ->](https://developers.google.com/style/commas)

### Grafana.GoogleProductNames

Extends: conditional

Use the full Google product name in the first instance.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#google-products)

### Grafana.GoogleRanges

Extends: existence

Don't add words such as 'from' or 'between' to describe a range of numbers.

[More information ->](https://developers.google.com/style/hyphens)

### Grafana.GoogleSpelling

Extends: existence

In general, use American spelling instead of '%s'.

[More information ->](https://developers.google.com/style/spelling)

### Grafana.GoogleWill

Extends: existence

Avoid using '%s'.

Use present tense for statements that describe general behavior that's not associated with a particular time.

[More information ->](https://developers.google.com/style/tense)

### Grafana.Headings

Extends: capitalization

Use sentence-style capitalization for '%s'.

If your heading contains capitalized words that represent product names, you need to add those words as exceptions in https://github.com/grafana/writers-toolkit/blob/main/vale/Grafana/Headings.yml for them to be considered correctly cased.

Vale considers multi-word exceptions such as _Grafana Enterprise Metrics_ as a single correctly cased word.

[More information ->](https://developers.google.com/style/capitalization#capitalization-in-titles-and-headings)

### Grafana.Kubernetes

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#kubernetes-objects)

### Grafana.OK

Extends: existence

Don't use any variation of okay in prose.
The exceptions are when you’re referencing or quoting:

- A user interface
- HTTP status codes or other code

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#ok-okay)

### Grafana.ProductPossessives

Extends: existence

Don't form a possessive from a feature name, product name, or trademark, regardless of who owns it.
Instead, use the name as a modifier or rewrite to use a word like of to indicate the relationship.

[More information ->](https://developers.google.com/style/possessives#product,-feature,-and-company-names)

### Grafana.PrometheusExporters

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#node-exporter)

### Grafana.Quickstart

Extends: substitution

Use the compound adjective without a hyphen whether the noun is implied or explicit. For example, you can use _quickstart guide_ or just _quickstart_.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#quickstart)

### Grafana.README

Extends: substitution

Use '%s' instead of '%s' unless you're referring to a specific file which has that spelling.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#readme)

### Grafana.React

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#react)

### Grafana.SQL

Extends: substitution

Use '%s' instead of '%s'.

The article—a or an—that you use before the acronym SQL depends on how the word is pronounced.

When referring to the product Microsoft SQL Server, SQL should be pronounced "sequel".
In this case, use the article 'a', as in "a SQL Server analysis".

When referring to the term in any other context, such as SQL databases, errors, or servers, SQL should be pronounced "ess-cue-el".
In this case, use the article 'an', as in "an SQL error".

[More information ->](https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/#sql-structured-query-language)

### Grafana.Shortcodes

Extends: script

Prefer `{{</* admonition type="<TYPE>" */>}}`.

It has the most consistent semantics.

The percent syntax is used for special behavior that isn't required with this shortcode.

[More information ->](https://grafana.com/docs/writers-toolkit/write/shortcodes/#admonition)

### Grafana.SmartQuotes

Extends: script

Avoid smart quotes in the source file, especially in code blocks.

Replace all smart double quotes like `“` or `”` with `"`.
Replace all smart single quotes like `‘`, `’`, or `ʼ` with `'`.

To disable smart quotes on Mac, refer to [Use smart quotes and dashes](https://support.apple.com/en-gb/guide/mac-help/mh35735/mac#:~:text=the%20command%20again.-,Use%20smart%20quotes%20and%20dashes,-Automatically%20convert%20straight).

In some contexts, Unicode characters aren't supported and break configurations.

The website renders paired quotes using smart quotes in paragraphs.

### Grafana.We

Extends: existence

Use first person plural pronouns like '%s' carefully.

Don't use 'we' when you're talking about the reader, instead use 'you'.

It's OK to use 'we' when you're talking about Grafana Labs.

[More information ->](https://developers.google.com/style/person#use-first-person-plural-pronouns-carefully)

### Grafana.Wish

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://developers.google.com/style/word-list#wish)

### Grafana.WordList

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://developers.google.com/style/word-list)

## Suggestions

The following rules are suggestions to consider a certain point of style.

### Grafana.Acronyms

Extends: conditional

Spell out '%s', if it's unfamiliar to the audience.

[More information ->](https://developers.google.com/style/abbreviations)

### Grafana.Archives

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://developers.google.com/style/word-list#extract)

### Grafana.GoogleContractions

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://developers.google.com/style/contractions)

### Grafana.GooglePassive

Extends: existence

In general, use active voice instead of passive voice ('%s').

[More information ->](https://developers.google.com/style/voice)

### Grafana.GoogleSemicolons

Extends: existence

Use semicolons judiciously.

[More information ->](https://developers.google.com/style/semicolons)

### Grafana.OAuth

Extends: substitution

Use '%s' instead of '%s'.

[More information ->](https://developers.google.com/style/word-list#oauth-20)

### Grafana.Parentheses

Extends: existence

Use parentheses judiciously.

[More information ->](https://developers.google.com/style/parentheses)

### Grafana.ReadabilityAutomatedReadability

Extends: metric

%s aim for below 8.

[More information ->](https://en.wikipedia.org/wiki/Automated_readability_index)

### Grafana.ReadabilityColemanLiau

Extends: metric

%s aim for below 9.

[More information ->](https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index)

### Grafana.ReadabilityFleschKincaid

Extends: metric

%s aim for below 8.

[More information ->](https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests)

### Grafana.ReadabilityFleschReadingEase

Extends: metric

%s aim for above 70.

[More information ->](https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests)

### Grafana.ReadabilityGunningFog

Extends: metric

%s aim for below 10.

[More information ->](https://en.wikipedia.org/wiki/Gunning_fog_index)

### Grafana.ReadabilityLIX

Extends: metric

%s aim for below 35.

[More information ->](<https://en.wikipedia.org/wiki/Lix_(readability_test)>)

### Grafana.ReadabilitySMOG

Extends: metric

%s aim for below 10.

[More information ->](https://en.wikipedia.org/wiki/SMOG)

### Grafana.Timeless

Extends: existence

Avoid using '%s' to keep the documentation timeless.

In general, document the current version of a product or feature.

It reduces the maintenance required to keep documentation up to date.
It avoids assuming the reader is familiar with earlier versions of the product.

If you're writing procedural or time-stamped content such as press releases, blog posts, or release notes, such time-based words and phrases are OK.

[More information ->](https://developers.google.com/style/timeless-documentation)
