---
title: Errata for doc-validator
menuTitle: Errata
description: Error codes and descriptions for doc-validator errata.
---

[//]: # "DO NOT EDIT. This file is generated from https://github.com/grafana/technical-documentation/blob/main/tools/cmd/doc-validator/errata.hcl"

# Errata for doc-validator

## anchor-does-not-exist

**Severity**: error

**Categories**: link

### Message

The anchor '%s' does not exist in the file '%s'.

### Guide

Replace the anchor with one of the available anchors.

Available anchors are: %q

## canonical-does-not-match-pretty-URL

**Severity**: error

**Categories**: front matter,canonical

### Message

The canonical '%s' in the front matter does not match the source file pretty URL path '%s'.

### Guide

Update the canonical URL to match the pretty URL for the source file path.

To determine the pretty URL for the source file path:

1. Start with the root documentation URL.
   For the Grafana project this would be "https://grafana.com/docs/grafana/latest/".

2. Append the source file path, ignoring the "docs/sources/" directory.
   For the path "docs/sources/administration/\_index.md", append the path "administration/\_index.md".
   The canonical URL is now "https://grafana.com/docs/grafana/latest/administration/_index.md".

   For the path "docs/sources/administration/users.md", append the path "administration/users.md".
   The canonical URL is now "https://grafana.com/docs/grafana/latest/administration/users.md".

3. Remove any "index.md", "\_index.md", or ".md" from the URL.
   For the path "docs/sources/administration/\_index.md", remove "\_index.md" from the canonical URL.
   The canonical URL is now "https://grafana.com/docs/grafana/latest/administration/"

   For the path "docs/sources/administration/users.md", remove ".md" from the canonical URL.
   The canonical URL is now "https://grafana.com/docs/grafana/latest/administration/users"

4. Append a trailing slash if there isn't one already.
   For the path "docs/sources/administration/\_index.md", do nothing.
   The canonical URL remains "https://grafana.com/docs/grafana/latest/administration/"

   For the path "docs/sources/administration/users.md", add a trailing slash.
   The canonical URL is now "https://grafana.com/docs/grafana/latest/administration/users/"

## canonical-is-not-valid-url

**Severity**: error

**Categories**: front matter,canonical

### Message

The canonical '%s' in the front matter is not a valid URL.

## canonical-validation-requires-additional-arguments

**Severity**: error

**Categories**: front matter,canonical

### Message

`doc-validator` cannot validate canonical without the [URL PATH PREFIX] argument being set.

### Guide

Invoke `doc-validator` with the [URL PATH PREFIX] argument set to the path prefix for the project documentation.
For grafana, this path prefix is '/docs/grafana/latest/'.

## external-link-invalid

**Severity**: error

**Categories**: link

### Message

The link to the external page '%s' is invalid. Links to external pages must be an absolute URI, with either the scheme https or http.

### Guide

For an explanation of absolute URIs, refer to https://www.rfc-editor.org/rfc/rfc3986#page-27.
For an explanation of schemes, refer to https://www.rfc-editor.org/rfc/rfc3986#page-17.

## grafana-com-link-invalid

**Severity**: error

**Categories**: link

### Message

The link target '%s' is invalid. Links to grafana.com pages must not include the hostname so that they can be resolved regardless of domain.

### Guide

Links to grafana.com pages must be one of the following:

- For linking to headings within the current page, use an anchor starting with a hash (#).
  For example, #heading.

- For linking to other pages in the current documentation set, use a Hugo relref with either an absolute path or relative path parameter.
  Using a relative path parameter, for example, `{{</* relref "./path/to/page" */>}}` or `{{</* relref "../other/path" */>}}`.
  Using an absolute path parameter, for example, `{{</* relref "/docs/grafana-cloud" */>}}`.

- For linking to any other page in the https://grafana.com site, use a partial URI consisting of an absolute path.
  For example, /blogs/.

For an explanation of partial URIs, refer to https://www.w3.org/Addressing/URL/4_3_Partial.html.

## h1-does-not-match-title

**Severity**: error

**Categories**: front matter,heading

### Message

The first heading '%s' does not match the title '%s'.

### Guide

Decide which of the first heading or title is most applicable and change the other to match, or update them both to a new title.

## image-does-not-exist

**Severity**: error

**Categories**: image

### Message

The image '%s' does not exist.

## image-located-outside-bundle

**Severity**: error

**Categories**: image

### Message

The image '%s' must be located in the same directory.

## image-not-linked-from-bundle

**Severity**: error

**Categories**: image

### Message

The image '%s' can only be linked from an index.md or \_index.md file.

### Guide

Move '%s' to a Hugo bundle and the image inside the bundle directory.

## invalid-alias

**Severity**: error

**Categories**: front matter

### Message

The front matter alias '%s' is from the Writers' Toolkit template and must be removed.

### Guide

Remove the alias from the file front matter.

If there are no other aliases in the front matter, remove the entire 'aliases' mapping entry.

## invalid-description

**Severity**: error

**Categories**: front matter

### Message

The front matter description is from the Writers' Toolkit template and must be replaced with a description of the content.

### Guide

Update the front matter description following the guidance in https://grafana.com/docs/writers-toolkit/writing-guide/front-matter/.

## link-does-not-exist

**Severity**: error

**Categories**: link

### Message

The link '%s' references a page in the project that does not exist.

### Guide

Check if the page has moved within the repository.

Verify the link by building the docs locally using the `make docs` webserver.

## parameter-must-be-present

**Severity**: error

**Categories**: front matter

### Message

The '%s' parameter in the front matter must be present.

## parameter-must-be-string

**Severity**: error

**Categories**: front matter

### Message

The '%s' parameter in the front matter must be a string.

## parameter-must-be-string-sequence

**Severity**: error

**Categories**: front matter

### Message

The '%s' parameter in the front matter must be a YAML sequence of strings.

## project-link-invalid

**Severity**: error

**Categories**: link

### Message

The link to the project page '%s' is invalid.
Links to project pages must use a Hugo relref shortcode like `{{</* relref "./path/to/page" */>}}`.

Links to project assets must be made from a bundle (index.md or \_index.md file).
The asset must be contained within the bundle directory, it must have an extension, and must not have the extension ".md".

## relref-can-resolve-ambiguously

**Severity**: error

**Categories**: link

### Message

The relref '%s' can resolve ambiguously because it is not a relative or absolute path.

### Guide

If a relref has ambiguous resolution, the link won't work and will instead point to the current page.

A relative path begins with either `./` or `../`.
An absolute path begins with a `/`.

## relref-has-trailing-slash

**Severity**: error

**Categories**: link

### Message

The relref '%s' has a trailing slash, which can break the resolution.

### Guide

You can reference an `index.md` file either by its path or by its containing folder without the ending `/`.
You can reference an `_index.md` file only by its containing folder.

Remove the trailing slash to make sure that changing the index type doesn't break the link.

## relref-is-malformed

**Severity**: error

**Categories**: link

### Message

The relref '%s' is malformed.
