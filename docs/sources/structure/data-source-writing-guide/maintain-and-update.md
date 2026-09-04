---
date: "2025-09-10T00:00:00+01:00"
description: How to keep data source documentation accurate and up-to-date by verifying it against the source code.
keywords:
  - data source
  - plugin
  - maintenance
menuTitle: Maintain and update
review_date: "2027-09-04"
title: Maintain and update data source documentation
weight: 400
---

# Maintain and update data source documentation

Documentation drifts from the implementation over time.
When you update data source documentation, verify accuracy against the source code rather than trusting the existing text.

## Verify accuracy against the source code

Check these files to confirm the documentation is accurate:

| File                                   | What it tells you                                                                                             |
| -------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `plugin.json`                          | Supported features (`alerting`, `annotations`), minimum Grafana version (`grafanaDependency`), and plugin ID. |
| `spec.ts` or similar                   | API endpoints, query parameters, field names, and descriptions.                                               |
| `*-editor.spec.ts` or end-to-end tests | Actual UI field names, error messages, and user flows.                                                        |
| Backend test files                     | Provisioning keys and configuration structure.                                                                |
| `CHANGELOG.md`                         | Version history, feature additions, and breaking changes.                                                     |

## Watch for common drift

These are the areas where documentation most often becomes outdated:

| What drifts                             | How to verify                                                               |
| --------------------------------------- | --------------------------------------------------------------------------- |
| Feature support (alerting, annotations) | Check `plugin.json` for `"alerting": true` and `"annotations": true`.       |
| UI field names                          | Check end-to-end tests for the actual label text.                           |
| Provisioning keys                       | Check test files for `secureJsonData` and `jsonData` keys.                  |
| Required permissions                    | Map API endpoints in `spec.ts` to the external platform's permission model. |
| Query parameters                        | Check `spec.ts` for all available fields, including optional ones.          |

## Keep documentation current

- **Match the version.** Ensure the plugin version referenced in the documentation matches the current version in `CHANGELOG.md`.
- **Update the review date.** Set the `review_date` front matter field to the date of your review.
- **Add aliases when you restructure.** Whenever you rename or move a page, add aliases that point to the old URLs. Refer to [File structure and front matter](https://grafana.com/docs/writers-toolkit/structure/data-source-writing-guide/file-structure-and-front-matter/).
- **Document all query parameters.** Include optional parameters, not just the common ones.
- **Verify the connection message.** Confirm the success message on the Configure page matches the plugin's health check implementation.
