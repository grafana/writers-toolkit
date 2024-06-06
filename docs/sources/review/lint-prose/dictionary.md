---
description: Learn about the Grafana Labs dictionary template format.
menuTitle: Dictionary
review_date: 2024-06-04
title: Grafana Labs' dictionary
---

# Grafana Labs' dictionary

The Grafana Labs documentation team maintains a dictionary used for spell checking.
The same dictionary is used to generate some Vale rules from the metadata in the word definition.

The dictionary uses the [Hunspell format](https://github.com/hunspell/hunspell) generated from a Jsonnet template.

If a word doesn't yet exist in the dictionary, you can add one by modifying the [dictionary template file](https://github.com/grafana/writers-toolkit/blob/main/vale/dictionary.jsonnet).
The template file uses the Jsonnet programming language but you don't need to know Jsonnet to add a new word.

## Word metadata

Each word in the dictionary has metadata that describes it.
The structure of that metadata is a Jsonnet object with the following fields:

| Key                   | Value type | Description                                                                                                                            |
| --------------------- | ---------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `word`                | `string`   | The spelling of the word.                                                                                                              |
| `affixes`             | `string`   | A concatenation of the letters corresponding to Hunspell affixes. For more information, refer to [Hunspell affixes](#hunspell-affixes) |
| `po`                  | `string`   | The part of speech descriptor for the word. Known acceptable values include `'noun'`, `'adjective'`, and `'verb'`.                     |
| `description`         | `string`   | A description of the word.                                                                                                             |
| `acronym`             | `boolean`  | If `true`, the word is an acronym. Defaults to `false`.                                                                                |
| `established_acronym` | `boolean`  | If `true`, the acronym doesn't generally need explaining, like the acronym HTTP. Defaults to `false`                                   |
| `product`             | `boolean`  | If `true`, the word is the name of any product, like Mimir or Facebook. Defaults to `false`.                                           |
| `Amazon`              | `boolean`  | If `true`, the word is the name of an Amazon product, like CloudWatch. Defaults to `false`.                                            |

You don't create the object directly in the dictionary file.
Instead, you use the `newWord` function which prescribes the required fields.

To create new word, add the following line to the dictionary.
The line goes in the `words` array, after an existing entry.
The entries are ordered alphabetically.

```jsonnet
newWord(<WORD>, <AFFIXES>, <PO>)
```

For example, to add the verb downsample, with affixes for past tense and the gerund, add the following line:

```jsonnet
newWord('downsample', 'DG', 'verb')
```

To understand the Hunspell affixes, refer to [Hunspell affixes](#hunspell-affixes).

## Hunspell affixes

Hunspell affixes teach the dictionary multiple words by applying affix rules to the stem word.
The affix rules are defined in [`en_US-grafana.aff`](https://github.com/grafana/writers-toolkit/blob/master/vale/dictionaries/en_US-grafana.aff).

| Letter | Kind   | Description                                                                   |
| ------ | ------ | ----------------------------------------------------------------------------- |
| `d`    | Prefix | Adds the `de` prefix to the stem, making deprovision from provision.          |
| `u`    | Prefix | Adds the `un` prefix to the stem, making unregister from register.            |
| `D`    | Suffix | Adds the past tense suffix `ed` to the stem, making ingested from ingest.     |
| `G`    | Suffix | Adds the gerund suffix `ing` to the stem, making singing from sing.           |
| `M`    | Suffix | Adds the possessive suffix `'s` to the stem, making ingester's from ingester. |
| `R`    | Suffix | Adds the `r` suffix to the stem, making profiler from profile.                |
| `S`    | Suffix | Adds the plural suffix `s` to the stem, making namespaces from namespace.     |
