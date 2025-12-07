---
date: "2025-12-04T14:00:01Z"
description: Guidelines for how to write security-concious documentation.
keywords:
  - Grafana
  - security
review_date: "2025-12-04"
title: Security
weight: 0
---

# Security

While there is no expectation that writers or engineers are security experts, we want to ensure that our documentation takes security into mind.

## Invalid tokens

It is rarely necessary for documentation to include tokens that are or look valid.
We prefer to mimic valid token formats, where we replace the majority of the tokens with text that makes them obviously invalid.

Examples of obviously invalid tokens would be:

- Grafana Labs tokens: `glsa_iNValIdinValiDinvalidinvalidinva_5b582697`.
- GitHub tokens: `github_pat_XXXXXXXXXXXXXXXX`.
