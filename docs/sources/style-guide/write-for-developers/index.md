---
title: Write for developers
description: Learn to write documentation for software developers and engineers.
aliases:
  - /docs/writers-toolkit/latest/style-guide/write-for-developers/
weight: 500
keywords:
  - Grafana
  - Docs for developers
---

# Write for developers

The guidelines that follow provide suggestions for writing documentation for software developers and engineers. 
Follow these tips to write useful API documentation, code examples, and other technical material.
To learn how to communicate effectively with the developers who enhance and work with the code of Grafana Labs projects or products, read these guidelines in the context of the [Style guide]({{< relref "../../style-guide/" >}}).

## Developer documentation basics

Writing technical content for software developers is similar to writing content for users or administrators of software products, so the same general guidelines apply.
However, developer-facing documentation tends to be more technical, and relies on certain important conventions.
Because this type of documentation contains details about code, it's important to know how to structure, format, and identify common problems that might arise.

When your readers are developers, you can assume that they are familiar with general programming concepts.
There is no need to explain elementary ideas. Instead, introduce those concepts and features that are specific to Grafana Labs products.
For example, instead of covering the fundamentals of UI design  _in general_, explain how Grafana Labs software or APIs _interpret_ those principles.

## Code comments

The foundation of strong documentation is well-written comments in code that are concise, relevant, and current. 

For do's and don'ts of writing comments, refer to the [Guidelines for code comments in grafana-* packages](https://github.com/grafana/grafana/blob/main/contribute/style-guides/code-comments.md).

For more general advice, consult one of the reputable [Google Style Guides](https://google.github.io/styleguide/) for your favorite programming language. 

## Reference docs

Whenever possible, automatically generate API and other reference documentation from source-code comments. However documentation is created, make sure it conforms to the [style conventions]({{< relref "../../style-guide/style-conventions/" >}}).
Pay particular attention to properly formatting the elements of code.

### Auto-generated documentation

The advantages of automating documentation programmatically are well-known, and include increased consistency and a reduction in human-made errors.
But behind every line of auto-generated content is a human author who is responsible for following the [style guide]({{< relref "../../style-guide/" >}}).

When writing documentation that will be used by an auto-generated program to create publishable content, keep the following things in mind:

An auto-generation tool can parse syntax, but when writing documentation that will be used to create automated content, it is up to you to add actionable insights. 
- What are the caveats, edge cases, and side effects?
- What is the bigger picture that is not self-evident in the code? 
- In short, what is everything that a developer needs to know to use the code?

> **Note:** It might be difficult to integrate auto-generated content with other documentation, such as relevant sections of _Get started_ guides, tutorials, or detailed code examples. You might need to ask a Docs team member to ensure that your content is properly cross-referenced.

### Elements of API references

Properly document the most common elements of an API reference, such as the title, parameters, return values, and so on. The following suggestions will help you to write more complete and consistent documentation:

| Element                                               | Description                                                           |
| ------------------------------------------------- | ----------------------------------------------------------------------- |
| Title and description       | The name of the element and a description of it in one or two sentences. Use back ticks (``) for API names, classes, methods, and so forth, so they display in a fixed-width font. |
| Syntax | The code signature that defines the element. If it is possible to use multiple programming languages, provide the syntax for each. Put the signature in `code font`. |
| Parameters | If the element has parameters, specify their descriptions, data types, and state whether they are optional or required. Put the parameters in _italic_. |
| Return values| If the element returns a value, describe the range of possibilities and the data type. |
| Error codes| Describe errors or exceptions and the conditions under which they occur. If possible, provide a way to resolve the issue. |
| Comments | Describe any important information that hasn't been previously included in the title, description, syntax, parameters, or return values. For example, you might explain non-obvious context, make a comparison to similar elements, or provide cautionary notes about potential gotchas. |

Additional tips:
- Write concisely. Don't say "The `AddUser` method adds a user." or "This method adds a user." when "Adds a user." will do.
- When the name of a code element is singular, don't pluralize it. Instead, follow it by plural noun. For example, write "`MyEvent` objects" instead of `MyEvents`.
- If the element performs an action, start the first sentence of the description with an action verb. 

## Code examples

Readers of documentation typically skim through it to find code samples they can copy and paste and run as is.
Because of this and whenever possible, provide production-ready examples.

However, it isn't necessary for every code example to be runnable in production.  Some code examples are written to illustrate a point so that the developer can learn how to do something similar on their own.

When providing an example, give a written description. You can put it either in the body of the document or as explanatory comments within the example code.

Remember the basic rule: explain _why_ your code does what it does, rather than describe _what_ it does.
For an in-depth, external resource about writing developer documentation, see [_Docs for Developers_](https://docsfordevelopers.com/).

### Formatting code examples

Here are guidelines to follow when formatting code examples.

- Use spaces, not tabs. 
- Follow Grafana's accepted [coding style guidelines](https://github.com/grafana/grafana/blob/main/contribute/style-guides/).
- Wrap lines at 80 characters.
- When omitting code, use three dots (...). Don't use the ellipsis character (…).

Introduce each code sample with a sentence or paragraph to establish its context. End the introduction with a colon if it immediately precedes the sample, or a period if it doesn't. 

### Paths, filenames, and URLs

There are many types of information that should be put in code font. Among these are filenames, paths, folders, and directories. However, don't format domain names or URLs as code if you intend the user to follow the link. 

While there are times when you may want to make a URL clickable, it is better to follow our [guidelines for references]({{< relref "../../writing-guide/references/" >}}).

## Command lines

Use the following conventions when you include commands on the command line in technical content.

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
```

If an argument is optional, enclose it in square brackets.

```console
ssh-rsa KEY_VALUE USERNAME [_FILENAME_]
```

Use descriptive words and phrases when including placeholders. Try to avoid using X or XXX. In Markdown, put a backtick followed by an asterisk in front and an asterisk followed by a backtick in front. For example, (*`A_PLACEHOLDER`*).

For more advice on formatting command lines, see the [Google developer style guide](https://developers.google.com/style/code-syntax).