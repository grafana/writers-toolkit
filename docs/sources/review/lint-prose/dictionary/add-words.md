---
date: "2024-06-11T07:36:23+01:00"
description: How to add a word to the Grafana Labs dictionary template.
menuTitle: Add words
review_date: "2024-06-04"
title: Add words to the Grafana Labs dictionary
---

# Add words to the Grafana Labs dictionary

The Grafana Labs documentation team maintains a dictionary used for spell checking.
The same dictionary is used to generate some Vale rules from the metadata in the word definition.

The template file uses the [Jsonnet](https://jsonnet.org) programming language but you don't need to know Jsonnet to add a new word.
Unlike YAML, Jsonnet isn't sensitive to whitespace.

This topic explains how to perform the following tasks:

- [Add a general word (noun, verb, or adjective)](#add-a-general-word-noun-verb-or-adjective)
- [Add a product name](#add-a-product-name)
- [Add an abbreviation](#add-an-abbreviation)

For more complicated words, if you're comfortable with writing Jsonnet, refer to [word metadata reference](https://grafana.com/docs/writers-toolkit/review/lint-prose/dictionary/#word-metadata).
If you're not comfortable writing Jsonnet, [create an issue](https://github.com/grafana/writers-toolkit/issues/new?title=Grafana.Spelling%20%3A%20%3CWORD%3E), and a maintainer can add it for you.

## Before you begin

- Clone the [Writers' Toolkit repository](https://github.com/grafana/writers-toolkit).

  For more information, refer to [Create a local repository](https://grafana.com/docs/writers-toolkit/write/tooling-and-workflows/#create-a-local-repository).

- Create a branch for your change.

  For more information, refer to [Create a branch from the default remote branch](https://grafana.com/docs/writers-toolkit/write/tooling-and-workflows/#create-a-branch-from-the-default-remote-branch).

## Add a general word (noun, verb, or adjective)

A noun is a word that represents a concrete or abstract thing.
A verb generally describes an action.
Adjectives describe nouns.

### Steps

To add a general word:

1. Open the `vale/dictionary/<LETTER>.jsonnet` template file in your editor where _`LETTER`_ is the first letter of the word you want to add.
1. Add a line for your word definition.

   Your line goes in the array, between the other entries.
   Entries look like the following:

   ```jsonnet
   word.new(<STEM>, <AFFIXES>, <PART OF SPEECH>),
   ```

   The entries are ordered alphabetically.

1. Fill out the required fields _`<STEM>`_, _`<AFFIXES>`_, and _`<PART OF SPEECH>`_.

   1. Replace _`<STEM>`_ with the word stem.

      This is the word without any prefixes or suffixes.
      For the verb _downsampling_, the stem is _downsample_.

      You must put the word stem between single quotes (`'`).

      Your line should look similar to the following:

      ```jsonnet
      word.new('downsample', <AFFIXES>, <PART OF SPEECH>),
      ```

   1. Replace _`<AFFIXES>`_ with the concatenation of the Hunspell affixes.

      To learn which affixes you can add, refer to the [Hunspell affixes table](https://grafana.com/docs/writers-toolkit/review/lint-prose/dictionary/#hunspell-affixes).

      You must put the affixes between single quotes (`'`).

      To add affixes for the past tense and gerund forms, your line should look similar to the following:

      ```jsonnet
      word.new('downsample', 'DG', <PART OF SPEECH>),
      ```

   1. Replace _`<PART OF SPEECH>`_ with the part of speech.

      For verbs, this is `'verb'`.
      For nouns, this is `'noun'`.
      For adjectives, this is `'adjective'`.

      Your completed line should look similar to the following:

      ```jsonnet
      word.new('downsample', 'DG', 'verb'),
      ```

1. If the word isn't well known, extend the definition to include a description: `description: <DESCRIPTION>`.

   The description should define the word.

   Your line should look similar to the following:

   ```jsonnet
   word.new('downsample', 'DG', 'verb') { description: 'To reduce the sampling rate of a signal.' },
   ```

## Add a product name

A product can be a Grafana Labs' product, another company's product, or the name of a project.

### Steps

To add a product:

1. Open the `vale/dictionary/<LETTER>.jsonnet` template file in your editor where _`LETTER`_ is the first letter of the word you want to add.
1. Add a line for your word definition.

   Your line goes in the array, between the other entries.
   Entries look like the following:

   ```jsonnet
   word.new(<STEM>, <AFFIXES>, <PART OF SPEECH>),
   ```

   The entries are ordered alphabetically.

1. Fill out the required fields _`<STEM>`_, _`<AFFIXES>`_, and _`<PART OF SPEECH>`_.

   1. Replace _`<STEM>`_ with the word stem.

      For products, this is the product name.

      ```jsonnet
      word.new('CloudWatch', <AFFIXES>, <PART OF SPEECH>),
      ```

   1. Replace _`<AFFIXES>`_ with the concatenation of the Hunspell affixes.

      Products generally have no affixes.

      ```jsonnet
      word.new('CloudWatch', '', <PART OF SPEECH>),
      ```

   1. Replace _`<PART OF SPEECH>`_ with the part of speech.

      For products, this is `'noun'`.

      Your line should look similar to the following:

      ```jsonnet
      word.new('CloudWatch', '', 'noun'),
      ```

1. Extend the definition to indicate it's a product.

   Add the object `{ product: true }` between the right bracket (`)`), and the end of line comma (`,`).

   Your line should look similar to the following:

   ```jsonnet
   word.new('CloudWatch', '', 'noun') { product: true },
   ```

1. If the product is an Amazon product, extend the definition to include this.

   Update the object to have an additional field, `Amazon: true`.

   Your line should look similar to the following:

   ```jsonnet
   word.new('CloudWatch', '', 'noun') { Amazon: true, product: true },
   ```

1. Extend the definition to include a description: `description: <DESCRIPTION>`.

   The description should at least include a link to the product's primary documentation.

   Your line should look similar to the following:

   ```jsonnet
   word.new('CloudWatch', '', 'noun') { Amazon: true, description: 'https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/WhatIsCloudWatch.html', product: true },
   ```

## Add an abbreviation

An abbreviation is a shortened form of a phrase.
If the abbreviation is commonly known, like HTTP, you don't need to explain it in your writing.
You can say that the abbreviation is commonly known in your definition.

### Steps

To add an abbreviation:

1. Open the `vale/dictionary/<LETTER>.jsonnet` template file in your editor where _`LETTER`_ is the first letter of the word you want to add.
1. Add a line for your word definition.

   Your line goes in the array, between the other entries.
   Entries look like the following:

   ```jsonnet
   word.new(<STEM>, <AFFIXES>, <PART OF SPEECH>),
   ```

   The entries are ordered alphabetically.

1. Fill out the required fields _`<STEM>`_, _`<AFFIXES>`_, and _`<PART OF SPEECH>`_.

   1. Replace _`<STEM>`_ with the word stem.

      For abbreviation, this is the abbreviation letters.

      ```jsonnet
      word.new('SUT', <AFFIXES>, <PART OF SPEECH>),
      ```

   1. Replace _`<AFFIXES>`_ with the concatenation of the Hunspell affixes.

      To learn which affixes you can add, refer to the [Hunspell affixes table](https://grafana.com/docs/writers-toolkit/review/lint-prose/dictionary/#hunspell-affixes).

      You must put the affixes between single quotes (`'`).

      Some abbreviations have a plural suffix.
      To add the plural suffix, include `s`.
      All affixes are case sensitive.

      ```jsonnet
      word.new('SUT', 's', <PART OF SPEECH>),
      ```

   1. Replace _`<PART OF SPEECH>`_ with the part of speech.

      For abbreviations, this is `'noun'`.

      Your line should look similar to the following:

      ```jsonnet
      word.new('SUT', 's', 'noun'),
      ```

1. Extend the definition to indicate it's an abbreviation.

   Add the object `{ abbreviation: true }` between the right bracket (`)`), and the end of line comma (`,`).

   Your line should look similar to the following:

   ```jsonnet
   word.new('SUT', 's', 'noun') { abbreviation: true },
   ```

1. Extend the definition to include a description: `description: <DESCRIPTION>`.

   The description should at least include the expanded abbreviation.

   Your line should look similar to the following:

   ```jsonnet
   word.new('SUT', 's', 'noun') { abbreviation: true, description: 'System Under Test' },
   ```

1. If you don't need to expand the abbreviation for the general reader, extend the definition to include this.

   Update the object to have an additional field, `established_abbreviation: true`.

   For the well-known abbreviation HTTP, your line should look similar to the following:

   ```jsonnet
   word.new('HTTP', '', 'noun') { abbreviation: true, description: 'Hypertext Transfer Protocol', established_abbreviation: true },
   ```
