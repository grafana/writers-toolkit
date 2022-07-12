---
title: "Content structure and naming conventions"
description: "Content structure and naming conventions"
aliases: []
weight: 400
---

# Content structure and naming conventions

Due to our robust documentation contribution model, it is important that contributors have the same understanding when it comes to building out the documentation in any of Grafana Lab’s many repositories.

This topic addresses how to structure and name document folders and files, title topics, and includes information about specific topic types.

## Directory structure

Use the following guidelines when you create the directory.

**(A)**: Nest an `_index.md` file below the main topic directory:

- This file serves as the topic area home page.
- This topic includes an overview of what’s covered in the topic area.
- Use the `{{< section >}}` short code to display child topic links on a page; remove any introductory phrase or sentence, such as “This section includes the following topics.” This will eliminate future bugs and manual bugs.

**(B)**: Create a directory for each topic.

**(C)**: Add an `_index.md` file:

- Functionally, the `_index.md` file creates pretty URLs. For example, `eat-pie.html` ends up as `eat-pie/`, and you do not have to remember the file extension (`.htm` or `.html` or something else).
- Content-wise, use the `_index.md` file to provide an overview of what’s covered in that section and links to child pages.

**(D)**: Add one concept topic and all related task topics to a directory or subdirectory.

**(E)**: Nest reference files directly in the main topic directory; do not add them to subdirectories.

Directory structure image here

## Topic types

Technical content comprises three types of topics:

- Concept
- Task
- Reference

Use the following guidelines when you write concept, task, and reference topics.







