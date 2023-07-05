---
title: Introduction to documentation
menuTitle: Introduction to documentation
description: Learn about Grafana's documentation
weight: 100
aliases:
- /docs/writers-toolkit/writing-guide/about-grafana-docs
keywords:
  - Grafana
  - documentation
---

# Introduction to documentation

All Grafana Enterprise and OSS documentation is located in the Grafana open source project GitHub repository: https://github.com/grafana/grafana/tree/main/docs/sources.

These are the other Grafana OSS project repositories:

- Grafana Agent: https://github.com/grafana/agent
- Loki: https://github.com/grafana/loki
- Tempo: https://github.com/grafana/tempo
- Mimir: https://github.com/grafana/mimir

Every repository contains a `docs/sources` directory, which is where we store our documentation.

## Topic-based authoring

The technical writing team at Grafana Labs uses topic-based authoring. Topic-based authoring is a modular approach to content creation where content is structured around topics that can be mixed and reused in different contexts. The topic types we use at Grafana are concept, reference, and task.

For more information on topic types, refer to [Topic types]({{< relref "../structure/topic-types" >}}).

Why is topic-based authoring important?

- **Writing that isn’t topic-based is difficult to reuse.** If everyone frequently copies information to multiple locations and makes small modifications, the result is rework and increases in errors, which multiplies the costs associated with maintenance and translation.

- **Writing that isn’t topic-based is difficult for readers to understand.** Content, structure, terminology, and writing style differs and can confuse and frustrate readers.

- **Writing that is topic-based is more consistent and user-friendly.**
  Users can find the information they are looking for quickly and easily. It has a more consistent format and voice, clearly defining the Grafana brand.

## Markdown

We write our technical documentation using Markdown. For more information, refer to the [Markdown style guide]({{< relref "../write/markdown-guide" >}}).

## Ways to contribute

We're thrilled that you are considering contributing to the documentation. You can contribute content in the following ways:

- [Request a change]({{< relref "../contribute-documentation#request-a-change" >}})
- [Edit a topic]({{< relref "../contribute-documentation#edit-a-topic" >}})
- [Create a topic]({{< relref "../structure/topic-types" >}})

## Join our community

For general discussions about documentation, you’re welcome to join the [#docs](https://raintank-corp.slack.com/archives/C5PG2JK8W) channel on our public Grafana Slack team.
