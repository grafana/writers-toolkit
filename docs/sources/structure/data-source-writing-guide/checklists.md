---
date: "2025-09-10T00:00:00+01:00"
description: Checklists for creating and reviewing Grafana data source documentation.
keywords:
  - data source
  - plugin
  - checklist
menuTitle: Checklists
review_date: "2027-09-04"
title: Data source documentation checklists
weight: 500
---

# Data source documentation checklists

Use these checklists when you create or review data source documentation.

## Required pages

- [ ] Overview `_index.md` with an introduction and navigation.
- [ ] Install and upgrade page, for Enterprise data sources.
- [ ] Configure page with authentication methods.
- [ ] Query editor page with examples.
- [ ] Template variables page, if supported.
- [ ] Annotations page, if supported.
- [ ] Alerting page, if supported.
- [ ] Troubleshooting page, for complex data sources.

## Front matter and structure

- [ ] All front matter fields populated, including `review_date`.
- [ ] Aliases added for any renamed or moved pages.
- [ ] Plugin version in the documentation matches the current version in `CHANGELOG.md`.
- [ ] Links use the fully qualified form. No `refs:` section, relative links, or `.md` links.
- [ ] No gerunds in headings.
- [ ] Intro content between all headings.
- [ ] Key concepts table on the Configure page for platform-specific authentication terminology.
- [ ] Key concepts table on the Query editor page for platform-specific query terminology.
- [ ] Enterprise plugin note included for Enterprise data sources, specifying "Pro or Advanced" and not the free tier.
- [ ] Plugin updates section included on the overview page.

## Content

- [ ] Verify the connection section includes the exact success message users see when **Save & test** succeeds.
- [ ] Provisioning YAML example included.
- [ ] Terraform example included, for major data sources.
- [ ] Screenshots added with the figure shortcode and captions.
- [ ] Query examples with realistic use cases.
- [ ] Use cases section for real-world scenarios.
- [ ] Macros documented, if applicable.

## Accuracy verification

- [ ] Feature support verified against `plugin.json`.
- [ ] UI field names verified against end-to-end tests.
- [ ] Provisioning keys verified against test files.
- [ ] All query parameters documented, including optional ones.
- [ ] Permissions documented using the external platform's exact terminology.
- [ ] No positional language, such as "above" or "below".

<!-- vale Grafana.Gerunds = NO -->

## Troubleshooting page

<!-- vale Grafana.Gerunds = YES -->

- [ ] Organized by issue category, such as authentication, connection, and query.
- [ ] Version and upgrade guidance section included for Enterprise plugins, using the `#version-and-upgrade-guidance` heading.
- [ ] Error messages quoted exactly as headings.
- [ ] Symptoms section describes what users see.
- [ ] Tables used for multiple cause and solution pairs.
- [ ] Numbered lists for sequential solutions.
- [ ] Code examples for complex fixes, such as IAM policies or configuration changes.
- [ ] Debug logging instructions included.
- [ ] "Get additional help" section with a community forum link, a GitHub issues link, an external documentation link, and what to include when reporting issues.
