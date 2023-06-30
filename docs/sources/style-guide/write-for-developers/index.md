---
title: Write for developers
description: Learn to write documentation for software developers and engineers.
weight: 500
keywords:
  - Grafana
  - Docs for developers
---

# Write for developers

The guidelines that follow provide suggestions for writing documentation for software developers and engineers.
Follow these tips to write useful API documentation, code examples, and other technical material.
To learn how to communicate effectively with the developers who enhance and work with the code of Grafana Labs projects or products, read these guidelines in the context of the [Style guide]({{< relref "../../style-guide" >}}).

## Developer documentation basics

Writing technical content for software developers is similar to writing content for users or administrators of software products, so the same general guidelines apply.
However, developer-facing documentation tends to be more technical, and relies on certain important conventions.
Because this type of documentation contains details about code, it's important to know how to structure, format, and identify common problems that might arise.

When your readers are developers, you can assume that they are familiar with general programming concepts.
There is no need to explain elementary ideas. Instead, introduce those concepts and features that are specific to Grafana Labs products.
For example, instead of covering the fundamentals of UI design _in general_, explain how Grafana Labs software or APIs _interpret_ those principles.

## Code comments

The foundation of strong documentation is well-written comments in code that are concise, relevant, and current.

For do's and don'ts of writing comments, refer to the [Guidelines for code comments in grafana-\* packages](https://github.com/grafana/grafana/blob/main/contribute/style-guides/code-comments.md).

For more general advice, consult one of the reputable [Google Style Guides](https://google.github.io/styleguide/) for your favorite programming language.

## Reference docs

Whenever possible, automatically generate API and other reference documentation from source-code comments. However documentation is created, make sure it conforms to the [style conventions]({{< relref "../../style-guide/style-conventions" >}}).
Pay particular attention to properly formatting the elements of code.

### Auto-generated documentation

The advantages of automating documentation programmatically are well-known, and include increased consistency and a reduction in human-made errors.
But behind every line of auto-generated content is a human author who is responsible for following the [style guide]({{< relref "../../style-guide" >}}).

When writing documentation that will be used by an auto-generated program to create publishable content, keep the following things in mind:

An auto-generation tool can parse syntax, but when writing documentation that will be used to create automated content, it is up to you to add actionable insights.

- What are the caveats, edge cases, and side effects?
- What is the bigger picture that is not self-evident in the code?
- In short, what is everything that a developer needs to know to use the code?

> **Note:** It might be difficult to integrate auto-generated content with other documentation, such as relevant sections of _Get started_ guides, tutorials, or detailed code examples. You might need to ask a Docs team member to ensure that your content is properly cross-referenced.

### Elements of API references

Properly document the most common elements of an API reference, such as the title, parameters, return values, and so on. The following suggestions will help you to write more complete and consistent documentation:

| Element               | Description                                                                                                                                                                                                                                                                              |
| --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Title and description | The name of the element and a description of it in one or two sentences. Use backticks (``) for API names, classes, methods, and so forth, so they display in a fixed-width font.                                                                                                        |
| Syntax                | The code signature that defines the element. If it is possible to use multiple programming languages, provide the syntax for each. Put the signature in `code font`.                                                                                                                     |
| Parameters            | If the element has parameters, specify their descriptions, data types, and state whether they are optional or required. Put the parameters in _italic_.                                                                                                                                  |
| Return values         | If the element returns a value, describe the range of possibilities and the data type.                                                                                                                                                                                                   |
| Error codes           | Describe errors or exceptions and the conditions under which they occur. If possible, provide a way to resolve the issue.                                                                                                                                                                |
| Comments              | Describe any important information that hasn't been previously included in the title, description, syntax, parameters, or return values. For example, you might explain non-obvious context, make a comparison to similar elements, or provide cautionary notes about potential gotchas. |

Additional tips:

- Remember to write concisely. Don't say "This method adds a user." when "Adds a user." will do. If your linter requires the description to begin with the element name, you may say "The AddUser method adds a user." to avoid an error message.
- When the names of code elements are singular, don't make them plural. Instead, add a plural noun to describe them. For example, don't change `MyEvent` to `MyEvents`; refer to the `MyEvent` objects.
- If the element does some sort of action, start the first sentence of the description with an action verb.

## Code examples

Readers of documentation typically skim through it to find code samples they can copy and paste and run as is.
Because of this and whenever possible, provide production-ready examples.

However, it isn't necessary for every code example to be runnable in production. Some code examples are written to illustrate a point so that the developer can learn how to do something similar on their own.

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

Many types of information belong in fixed-width font. Among these are path names, file names, directories, or folders. However, don't format domain names or URLs as code if you intend the user to follow the link.

Although there are times when you might want to make a URL clickable, it is better to follow the guidelines for [References]({{< relref "../../write/references" >}}).

## Command lines

Use the following conventions when you include commands on the command line in technical content.

- Do not assume everyone is using Linux. Make sure instructions include enough information for Windows and Mac users to successfully complete procedures.

- Do not add `$` before commands. Make it easy for users to copy and paste commands.

  - **Right:** `sudo yum install grafana`
  - **Wrong:** `$ sudo yum install grafana`

- Include `sudo` before commands that require `sudo` to work.

For terminal examples and configurations, use a `bash` code block. In raw markdown:

````
```bash
sudo yum install grafana
```
````

It produces:

```bash
sudo yum install grafana
```

If your command-line instructions include a combination of input and output lines, use separate code blocks for input and output, and use a `console` code block for the output.

The input, in raw markdown:

````
```bash
cat ~/.ssh/my-ssh-key.pub
```
````

It produces the following output:

```console
cat ~/.ssh/my-ssh-key.pub
```

For HTTP request/response, use an `http` code block in raw markdown:

````
```http
GET /api/dashboards/id/1/permissions HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```
````

It produces:

```http
GET /api/dashboards/id/1/permissions HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer  eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

If an argument is optional, enclose it in square brackets. For example, prefer the following raw markdown:

````
```typescript
ssh-rsa KEY_VALUE USERNAME [_FILENAME_]
```
````

It produces:

```typescript
ssh-rsa KEY_VALUE USERNAME [_FILENAME_]
```

Use descriptive words and phrases when including placeholders, and avoid using X or XXX.
In Markdown, in front of a placeholder, use an asterisk followed by a backtick and an open angle bracket. Use a dash instead of an underscore in the placeholder name if necessary. At the end of the placeholder, use an open angle bracket followed by a backtick and an asterisk. For example, refer to the following raw markdown:

````
```
The following text is a placeholder: *`<A-PLACEHOLDER>`*.
```
````

It produces:

```
The following text is a placeholder: *`<A-PLACEHOLDER>*`.
```

For more information about formatting command lines, see [Document command-line syntax](https://developers.google.com/style/code-syntax) from Google.
