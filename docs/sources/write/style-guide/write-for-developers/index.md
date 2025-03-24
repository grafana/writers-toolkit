---
aliases:
  - /docs/writers-toolkit/style-guide/write-for-developers
  - /docs/writers-toolkit/write/style-guide/write-for-developers
date: "2023-03-27T13:09:54-07:00"
description: Learn to write documentation for software developers and engineers using Grafana Labs products.
keywords:
  - Grafana
  - Docs for developers
review_date: "2024-06-27"
title: Write for developers
weight: 500
---

# Write for developers

The guidelines that follow provide suggestions for writing documentation for software developers and engineers.
Follow these tips to write useful API documentation, code examples, and other technical material.
To learn how to communicate effectively with the developers who enhance and work with the code of Grafana Labs projects or products, read these guidelines in the context of the [Style guide](https://grafana.com/docs/writers-toolkit/write/style-guide/).

## Developer documentation basics

Writing technical content for software developers is similar to writing content for users or administrators of software products, so the same general guidelines apply.
However, developer-facing documentation tends to be more technical, and relies on certain important conventions.
Because this type of documentation contains details about code, it's important to know how to structure, format, and identify common problems that might arise.

When your readers are developers, you can assume that they're familiar with general programming concepts.
There is no need to explain elementary ideas.
Instead, introduce those concepts and features that are specific to Grafana Labs products.
For example, instead of covering the fundamentals of UI design _in general_, explain how Grafana Labs software or APIs _interpret_ those principles.

## Code comments

The foundation of strong documentation is well-written comments in code that are concise, relevant, and current.

For dos and don'ts of writing comments, refer to the [Guidelines for code comments in `grafana-\*` packages](https://github.com/grafana/grafana/blob/main/contribute/style-guides/code-comments.md).

For more general advice, consult one of the reputable [Google Style Guides](https://google.github.io/styleguide/) for your programming language.

## Reference documentation

Whenever possible, automatically generate API and other reference documentation from source-code comments.
However documentation is created, make sure it conforms to the [style conventions](https://grafana.com/docs/writers-toolkit/write/style-guide/style-conventions/).
Pay particular attention to properly formatting the elements of code.

### Auto-generated documentation

The advantages of automating documentation programmatically are well-known, and include increased consistency and a reduction in human-made errors.
But behind every line of auto-generated content is a human author who is responsible for following the style guide.

When writing documentation used by an auto-generated program to create publishable content, keep the following things in mind:

An auto-generation tool can parse syntax, but when writing documentation used to create automated content, it's up to you to add actionable insights.

- What are the caveats, edge cases, and side effects?
- What's the bigger picture that's not self-evident in the code?
- In short, what's everything that a developer needs to know to use the code?

{{< admonition type="note" >}}
It might be difficult to integrate auto-generated content with other documentation, such as relevant sections of _Get started_ guides, tutorials, or detailed code examples.
You might need to ask a member of the Grafana Labs documentation team to ensure that your content is properly cross-referenced.
{{< /admonition >}}

### Elements of API references

Properly document the most common elements of an API reference, such as the title, parameters, return values, and so on.
The following suggestions help you to write more complete and consistent documentation:

| Element               | Description                                                                                                                                                                                                                                                                              |
| --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Title and description | The name of the element and a description of it in one or two sentences. Use backticks (``) for API names, classes, methods, and so forth, so they display in a fixed-width font.                                                                                                        |
| Syntax                | The code signature that defines the element. If it's possible to use multiple programming languages, provide the syntax for each. Put the signature in `code font`.                                                                                                                      |
| Parameters            | If the element has parameters, specify their descriptions, data types, and state whether they're optional or required. If the parameters are optional, add the prefix "Optional:". Put the parameters in _italic_.                                                                       |
| Return values         | If the element returns a value, describe the range of possibilities and the data type.                                                                                                                                                                                                   |
| Error codes           | Describe errors or exceptions and the conditions under which they occur. If possible, provide a way to resolve the issue.                                                                                                                                                                |
| Comments              | Describe any important information that hasn't been previously included in the title, description, syntax, parameters, or return values. For example, you might explain non-obvious context, make a comparison to similar elements, or provide cautionary notes about potential gotchas. |

Additional tips:

- Remember to write concisely.
  Don't say "This method adds a user." when you could "Adds a user.".
  If your linter requires the description to begin with the element name, you may say "The AddUser method adds a user." to avoid an error message.
- When the names of code elements are singular, don't make them plural.
  Instead, add a plural noun to describe them.
  For example, don't change `MyEvent` to `MyEvents`; refer to the `MyEvent` objects.
- If the element does some sort of action, start the first sentence of the description with an action verb.

### Reference structure

Organize references alphabetically, organizing by frequency of use is too subjective.

If the reference has required and optional components, document them separately.
Document the required components first, then the optional components.
Organize each set of components alphabetically.

## Code examples

Readers of documentation typically skim through it to find code samples they can copy and paste and run as is.
Because of this and whenever possible, provide production-ready examples.

However, it isn't necessary for every code example to be runnable in production.
Some code examples are written to illustrate a point so that the developer can learn how to do something similar on their own.
Such examples should be clearly marked as partial.

When providing an example, give a written description.
You can put it either in the body of the document or as explanatory comments within the example code.
However, don't put comments in command line examples.

Remember the basic rule: explain _why_ your code does what it does, rather than describe _what_ it does.
For an in-depth, external resource about writing developer documentation, refer to [_Docs for Developers_](https://docsfordevelopers.com/).

### Full code examples

To format code examples, use the following guidelines.

#### Highlight syntax

In Markdown, an info string after the first three backticks describes the language contained within.
The website uses this information to apply syntax highlighting to code examples.

The following Markdown sets the info string to `json`:

````markdown
```json
{ "key": "value" }
```
````

Some common languages and their info strings are:

- Bash: `bash`
- Console: `console`
- Go: `go`
- JSON: `json`
- PromQL: `promql`
- River: `river`
- Shell: `shell`
- YAML: `yaml`

#### Use appropriate whitespace for indentation

Spaces (` `) are a good default but note that some languages use alternative indentation.
For example, Makefiles and Go source code uses tabs (`	`) for indentation.

Use 2 spaces unless there is a clearly established alternative convention.
For example, Python generally uses 4 spaces for indentation.

Above all else, be consistent with existing documentation in your area of the documentation or in your project.

#### Introduce each code sample with a sentence or paragraph to establish its context

End the introduction with a colon (`:`) if it immediately precedes the sample, or a period (`.`) if it doesn't.

#### Use the `code` shortcode for tabbed examples

If you have the same example in multiple languages, use the [`code` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#code).
The website presents the snippets in tabs and remembers the user's preferred choice.

### Partial code examples

Partial code examples are shorter and focus the reader on a specific area of the code.
However, they require the user to integrate the partial code example with their existing configuration or other examples.

In addition to the general guidance for formatting code examples, when working with partial code examples, use the following guidelines.

#### Wherever possible, ensure the partial example is copy-pasteable

In JSON or YAML, the partial example should be valid on its own.

In programming language code examples, if the example won't compile on its own, make sure the snippet is a recognizable unit of a source file.

<!-- vale Grafana.Parentheses = NO -->

Don't use an ellipsis (`…`) or three periods (`...`) to omit information.

It breaks the readers ability to copy-paste the example and provides no additional context for the omission.
Instead use the preceding sentences or valid code comments to clearly explain the scope of the example and what's omitted.

<!-- vale Grafana.Parentheses = YES -->

For example:

> The following YAML example demonstrates the configuration of a single port numbered `80` using the `TCP` protocol.
>
> It's part of a Kubernetes Service specification.
> It's not a complete Service specification and you must incorporate it with the rest of a Service specification.
>
> ```yaml
> ports:
>   - port: 80
>     protocol: TCP
> ```

#### Explain how to integrate the example

With YAML examples especially, it's important to explain how to integrate the example within a wide configuration.

Explain which key the example is the value for.
The following example extends the one in the previous section:

> The following YAML example demonstrates the configuration of a single port numbered `80` using the `TCP` protocol.
>
> It's part of a Kubernetes Service specification.
> It's not a complete Service specification and you must incorporate it with the rest of a Service specification.
>
> To incorporate the example, you must include it as the value to the Service `spec` mapping.
> If there is an existing `ports` value, you must choose to replace it or merge the two.
>
> ```yaml
> ports:
>   - port: 80
>     protocol: TCP
> ```

#### Refer to nested fields

With configuration in JSON or YAML, you may need to refer to deeply nested fields.
Using natural language is laborious to write and read.

Instead use _dot notation_ to separate nested fields.
For example, `spec.template.metadata` to refer to the `metadata` field within the `template` field, which is itself within the `spec` field.

If the field name has dots, surround the name with square brackets instead.
For example, `spec.selector[app.kubernetes.io/name]`.

To refer to any member of an array, use `[*]`.

To refer to a specific index in an array, use `[<INDEX>]`.
For example, the following dot notation refers to the first container within `spec`, and the first port within that container: `spec.containers[0].ports[0].containerPort`.

The use of _dot notation_ for nested fields is common in other engineering documentation such as [Kubernetes documentation](https://kubernetes.io/docs/).

### Paths, filenames, and URLs

Many types of information belong in fixed-width font.
Among these are paths, filenames, directories, and folders.
However, don't format domain names or URLs as code if you intend the user to follow the link.

## Command lines

Use the following conventions when you include commands on the command line in technical content.

- Don't assume everyone is using Linux. Make sure instructions include enough information for Windows and Mac users to successfully complete procedures.

- Don't add `$` before commands so users can copy and paste them.

  - **Right:** `sudo yum install grafana`
  - **Wrong:** `$ sudo yum install grafana`

- Include `sudo` before commands that require `sudo` to work.

For terminal examples and configurations, use a `bash` code block. In raw Markdown:

````markdown
```bash
sudo yum install grafana
```
````

It produces:

```bash
sudo yum install grafana
```

If your command-line instructions include a combination of input and output lines, use separate code blocks for input and output, and use a `console` code block for the output.

The input, in raw Markdown:

````markdown
```bash
cat ~/.ssh/my-ssh-key.pub
```
````

It produces the following output:

```console
cat ~/.ssh/my-ssh-key.pub
```

For HTTP request/response, use an `http` code block in raw Markdown:

````markdown
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
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

If an argument is optional, enclose it in square brackets.
For required arguments, refer to [Placeholder variables](#placeholder-variables).
For example, prefer the following raw Markdown:

````markdown
```bash
ssh-rsa <KEY_VALUE> <USERNAME> [FILENAME]
```
````

It produces:

```bash
ssh-rsa <KEY_VALUE> <USERNAME> [FILENAME]
```

## Placeholder variables

Use descriptive words and phrases when including placeholders, and avoid using X or XXX.
In Markdown, in front of a placeholder, use an underscore (`_`) followed by a backtick (`` ` ``) and a less-than sign (`<`).
At the end of the placeholder, use a greater-than sign (`>`) followed by a backtick (`` ` ``) and an underscore (`_`).

For example, refer to the following raw Markdown:

```Markdown
The following text is a placeholder: _`<PLACEHOLDER>`_.
```

It produces:

The following text is a placeholder: _`<PLACEHOLDER>`_.

For more information about formatting command lines, refer to [Document command-line syntax](https://developers.google.com/style/code-syntax) from Google.
