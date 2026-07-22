---
title: scip-go unavailable in bootstrap pod (Go not in shell PATH)
date: 2026-07-22
pr: TBD
area: [tooling, scip, kb-bootstrap]
files: [.github/workflows/knowledge-sync.yml]
symptom: "Step 2.5 SCIP indexing silently skipped — no index.scip.json produced, no error except 'go: command not found'"
tags: [scip, go, kb-bootstrap, tooling, pod]
summary: "Go binary is not in the login shell's PATH in the KB bootstrap pod, so scip-go cannot be installed or run; the bootstrap falls through to the AST path as intended."
---

## The Pattern

The KB bootstrap template (step 2.5) attempts to run SCIP indexing via a `bash -l` login shell so that environment setup scripts (`.bashrc`, `.bash_profile`) install `go` on demand. In this pod, `go` is not available via the login shell despite the container having Go-based tools installed at other paths. The `bash -lc` invocation returned `go: command not found`, causing the `scip-go` install and run to silently skip.

## The Rule

SCIP indexing is **best-effort and always optional** for Go repos in this environment. When `go` is unavailable in the bootstrap pod's login shell, the `sdd-knowledge --full` AST path runs unchanged and the KB is still fully buildable. Do not attempt to work around the missing `go` by hardcoding paths — the bootstrap is designed to fall through cleanly.

The AST path correctly extracts 140 symbols, 122 edges, and 14 docs for a 34-file Go repo. The SCIP path would have added resolved member/dynamic-dispatch edges that the AST withholds — a quality improvement, not a correctness requirement.

## Detection

```bash
bash -lc 'which go || echo "go not in login shell"'
```

If this prints "go not in login shell", the SCIP step will silently skip. The bootstrap log will show "No SCIP artifacts produced" from the step 2.5 echo.

## Related

- KB Bootstrap step 2.5 template in `templates/building-blocks.md`
- `sdd-knowledge --full` AST fallback path
