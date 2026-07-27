---
category: guides
subcategory: troubleshooting
confidence: low
documentType: how-to
scope: org
contentHash: 1b3a37ae63d1
tags: [anti-pattern, troubleshooting, faq]
source: (synthesized)
verified: 2026-07-22
synthetic: synthesized-faq
---

## Common Mistakes and Anti-Patterns

<!-- sdd-knowledge-synthesized -->

> Auto-generated from anti-patterns found across the knowledge base.

### Ci

**From [assets-go-libs](../../ci/assets-go-libs.md):**
After making your changes don't forget to run unit tests, go formatting and linter:

``` shell

### Code-conventions

**From [Code Conventions](../../code-conventions/code-conventions.md):**
**Rule**: wrap failures with a sentinel error using `fmt.Errorf("%w: detail", validation.ErrSomething)`. Never invent new error variables in sub-packages — use the shared sentinels from `validation/errors.go`.

**Rule**: do NOT return on first error. Always run all checks so the caller gets the full picture.

