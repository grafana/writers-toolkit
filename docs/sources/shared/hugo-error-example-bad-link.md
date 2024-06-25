---
date: "2023-09-04T07:43:24-04:00"
description: Understand the REF_NOT_FOUND error Hugo emits for broken relref links.
review_date: "2024-05-28"
title: REF_NOT_FOUND Hugo output
---

[//]: # "This file documents an example Hugo error output for relref and links."
[//]: # "This shared file is included in these locations:"
[//]: # "- Page: [Test documentation changes](https://grafana.com/docs/writers-toolkit/review/test-documentation-changes/#example-rebuild-failed-due-to-missing-shortcode)"
[//]: # "  Source: [test-documentation-changes/index.md](https://github.com/grafana/writers-toolkit/blob/main/docs/sources/review/test-documentation-changes/index.md?plain=1#L99)"
[//]: #
[//]: # "If you make changes to this file, verify that the meaning and content are not changed in any place where the file is included."

Hugo emits `REF_NOT_FOUND` warnings indicating the filename and location of such references when building the docs, for example with `make docs` in `grafana/grafana` or `make server-quick` in `grafana/website`:

```
WARN 2022/08/04 21:35:37 [en] REF_NOT_FOUND: Ref "../../enterprise": "/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md:14:47": page not found
```

In this example,

- `Ref "../../enterprise"` is the destination of the reference that Hugo can't resolve
- `\/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md` is the document containing the reference, where the path after `/next/` is relative to the documentation root of the component repository
- `:14` represents the line number containing the unresolved reference
- `:47` represents the character in that line where the unresolved reference begins

If you see this error, then the reference's destination is invalid.
This may be due to a typo in the reference or having the incorrect path to the destination directory.
Fix the error by correcting the reference target.
