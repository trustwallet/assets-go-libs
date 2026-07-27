---
category: architecture
confidence: low
documentType: explanation
scope: repo
contentHash: c5a3f3dd89f3
tags: [architecture, domain, service, dependency]
source: architecture/call-graph.md
verified: 2026-07-22
splitPartIndex: 2
splitPartTotal: 2
canonical: false
synthetic: split-part
---

## Call Graph (Part 2)

- Call edges: **52** across **140** symbols
- Reachable from entry points (exported symbols + routes): **138**
- Dependency cycles: **0**
- Cross-domain bridges: **0**
- Dead-code candidates (non-exported, zero callers, unreachable): **2**

### Most-coupled symbols (god-node ranking)

| Symbol | Fan-in | Fan-out | Degree |
|--------|--------|---------|--------|
| `ValidateAsset` | 0 | 10 | 10 |
| `ValidateCoin` | 0 | 8 | 8 |
| `GetTokenInfo` | 0 | 5 | 5 |
| `NewAssetFile` | 3 | 1 | 4 |
| `ValidateValidatorsAddress` | 0 | 4 | 4 |
| `ValidateLinks` | 2 | 2 | 4 |
| `ValidateTokenList` | 0 | 3 | 3 |
| `NewPath` | 1 | 1 | 2 |
| `ValidateAssetAddress` | 0 | 2 | 2 |
| `ValidateTronAddress` | 2 | 0 | 2 |
| `NewErrComposite` | 2 | 0 | 2 |
| `validateLogoSize` | 2 | 0 | 2 |
| `isEmpty` | 2 | 0 | 2 |
| `ValidateAssetRequiredKeys` | 1 | 1 | 2 |
| `ValidateCoinRequiredKeys` | 1 | 1 | 2 |

### Possible duplicate entities (name variants)

> Symbol names that normalize identically — likely the same entity spelled inconsistently. Unify or distinguish in the docs.

- `ValidateAssetID` / `validateAssetID`

### Dead-code candidates

> Non-exported symbols with no resolved callers, unreachable from any entry point. Static analysis cannot see dynamic dispatch — verify before removing.

- `defineFileType` (`file/path.go`)
- `getFile` (`file/service.go`)

## See Also
- [validation](../features/validation.md) <!-- rel:strong -->
- [file](../features/file.md) <!-- rel:weak -->
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:weak -->
