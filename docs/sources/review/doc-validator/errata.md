---
aliases:
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/validate-technical-documentation/errata/
  - /docs/writers-toolkit/review/doc-validator/errata/
review_date: 2024-05-23
description: A reference of error codes and descriptions reported by doc-validator when linting Grafana Labs technical documentation.
title: Errata for doc-validator
---

<!-- DO NOT EDIT. This file is generated from <https://github.com/grafana/technical-documentation/blob/main/tools/cmd/doc-validator/errata.hcl> -->

# Errata for doc-validator

## `anchor-does-not-exist`

**Severity**: error

**Categories**: link

### Message

The anchor '%s' doesn't exist in the file '%s'.

### Guide

Replace the anchor with one of the available anchors.

Available anchors are: %q

## `canonical-does-not-match-pretty-URL`

**Severity**: error

**Categories**: front matter,canonical

### Message

The canonical '%s' in the front matter doesn't match the source file pretty URL path '%s'.

### Guide

Update the canonical URL to match the pretty URL for the source path.

To determine the pretty URL for the source path:

1. Start with the root documentation URL.
   For the Grafana project this would be `https://grafana.com/docs/grafana/latest/`.

2. Append the source path, ignoring the `docs/sources/` directory.
   For the path `docs/sources/administration/_index.md`, append the path `administration/_index.md`.
   The canonical URL becomes `https://grafana.com/docs/grafana/latest/administration/_index.md`.

   For the path `docs/sources/administration/users.md`, append the path `administration/users.md`.
   The canonical URL becomes `https://grafana.com/docs/grafana/latest/administration/users.md`.

3. Remove any `index.md`, `_index.md`, or `.md` from the URL.
   For the path `docs/sources/administration/_index.md`, remove `_index.md` from the canonical URL.
   The canonical URL becomes `https://grafana.com/docs/grafana/latest/administration/`

   For the path `docs/sources/administration/users.md`, remove `.md` from the canonical URL.
   The canonical URL becomes `https://grafana.com/docs/grafana/latest/administration/users`

4. Append a trailing slash if there isn't one already.
   For the path `docs/sources/administration/_index.md`, do nothing.
   The canonical URL remains `https://grafana.com/docs/grafana/latest/administration/`

   For the path `docs/sources/administration/users.md`, add a trailing slash.
   The canonical URL becomes `https://grafana.com/docs/grafana/latest/administration/users/`

## `canonical-is-not-valid-url`

**Severity**: error

**Categories**: front matter,canonical

### Message

The canonical '%s' in the front matter isn't a valid URL.

## `canonical-validation-requires-additional-arguments`

**Severity**: error

**Categories**: front matter,canonical

### Message

`doc-validator` can't validate canonical without the `URL PATH PREFIX` argument being set.

### Guide

Invoke `doc-validator` with the `URL PATH PREFIX` argument set to the path prefix for the project documentation.
For the Grafana project, this path prefix is `/docs/grafana/latest/`.

## `h1-does-not-match-title`

**Severity**: error

**Categories**: front matter,heading

### Message

The first heading '%s' doesn't match the title '%s'.

### Guide

Decide which of the first heading or title is most applicable and change the other to match, or update them both.

## `image-does-not-exist`

**Severity**: error

**Categories**: image

### Message

The image '%s' doesn't exist.

## `image-located-outside-bundle`

**Severity**: error

**Categories**: image

### Message

The image '%s' must be located in the same directory.

## `image-not-linked-from-bundle`

**Severity**: error

**Categories**: image

### Message

The image '%s' can only be linked from an index.md or \_index.md file.

### Guide

Move '%s' to a Hugo bundle and the image inside the bundle directory.

## `invalid-alias`

**Severity**: error

**Categories**: front matter

### Message

The front matter alias '%s' is from the Writers' Toolkit template and must be removed.

### Guide

Remove the alias from the file front matter.

If there are no other aliases in the front matter, remove the entire 'aliases' mapping entry.

## `invalid-description`

**Severity**: error

**Categories**: front matter

### Message

The front matter description is from the Writers' Toolkit template and must be replaced with a description of the content.

### Guide

Update the front matter description following the guidance in https://grafana.com/docs/writers-toolkit/write/front-matter/.

## `link-does-not-exist`

**Severity**: error

**Categories**: link

### Message

The link '%s' references a page in the project that doesn't exist.

### Guide

Check if the page has moved within the repository.

Verify the link by building the docs locally using the `make docs` webserver.

## `link-invalid`

**Severity**: error

**Categories**: link

### Message

The link '%s' is invalid. Links must be an absolute URI, with either the scheme `https` or `http`.

### Guide

For an explanation of absolute URIs, refer to https://www.rfc-editor.org/rfc/rfc3986#page-27.
For an explanation of schemes, refer to https://www.rfc-editor.org/rfc/rfc3986#page-17.

## `parameter-must-be-present`

**Severity**: error

**Categories**: front matter

### Message

The '%s' parameter in the front matter must be present.

## `parameter-must-be-string`

**Severity**: error

**Categories**: front matter

### Message

The '%s' parameter in the front matter must be a string.

## `parameter-must-be-string-sequence`

**Severity**: error

**Categories**: front matter

### Message

The '%s' parameter in the front matter must be a YAML sequence of strings.

## `project-link-invalid`

**Severity**: error

**Categories**: link

### Message

The link to the project page '%s' is invalid.
Links to project pages must use a Hugo `relref` shortcode like `{{</* relref "./path/to/page" */>}}`.

Links to project assets must be made from a bundle (index.md or \_index.md file).
The asset must be contained within the bundle directory, it must have an extension, and must not have the extension `.md`.

## `relref-can-resolve-ambiguously`

**Severity**: error

**Categories**: link

### Message

The `relref` shortcode argument '%s' can resolve ambiguously because it isn't a relative or absolute path.

### Guide

If a `relref` shortcode has ambiguous resolution, the link destination is the current page and not the intended page.

A relative path begins with either `./` or `../`.
An absolute path begins with a `/`.

## `relref-has-trailing-slash`

**Severity**: error

**Categories**: link

### Message

The `relref` shortcode argument '%s' has a trailing slash, which can break the resolution.

### Guide

You can reference an `index.md` file either by its path or by its containing folder without the ending `/`.
You can reference an `_index.md` file only by its containing folder.

Remove the trailing slash to make sure that changing the index type doesn't break the link.

## `relref-is-malformed`

**Severity**: error

**Categories**: link

### Message

The `relref` shortcode '%s' is malformed.
