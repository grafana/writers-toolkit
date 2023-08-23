---
headless: true
description: Shared file for example Hugo error output.
labels:
  products:
    - oss
---

[//]: # "This file documents an example Hugo error output for relref and links."
[//]: # "This shared file is included in these locations:"
[//]: # "/writers-toolkit/review/run-a-local-webserver"
[//]: # "/writers-toolkit/write/references/index.md"
[//]: # "/tempo/docs/sources/tempo/traceql/query_editor.md"
[//]: #
[//]: # "If you make changes to this file, verify that the meaning and content are not changed in any place where the file is included."
[//]: # "Any links should be fully qualified and not relative: /docs/grafana/ instead of ../grafana/."

Hugo emits `REF_NOT_FOUND` warnings indicating the filename and location of such references when building the docs, for example with `make docs` in `grafana/grafana` or `make server-quick` in `grafana/website`:

```
WARN 2022/08/04 21:35:37 [en] REF_NOT_FOUND: Ref "../../enterprise": "/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md:14:47": page not found
```

In this example,

- `Ref "../../enterprise"` is the destination of the reference that Hugo can't resolve
- `\/hugo/content/docs/grafana/next/administration/roles-and-permissions/access-control/assign-rbac-roles.md` is the document containing the reference, where the path after `/next/` is relative to the documentation root of the component repository
- `:14` represents the line number containing the unresolved reference
- `:47` represents the character in that line where the unresolved reference begins

If the reference's destination appears to be invalid, for example due to a typo in the reference or the depth of the`relref` directory, then you should be able to resolve this by correcting the reference target.