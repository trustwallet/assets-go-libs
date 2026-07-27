---
category: architecture
subcategory: data
confidence: low
documentType: explanation
scope: org
contentHash: db266fcfa937
tags: [architecture]
source: architecture/data/models.md
verified: 2026-07-22
splitPartIndex: 3
splitPartTotal: 3
canonical: false
synthetic: split-part
---

## Data Models & Schemas (Part 3)

| Field | Type | Optional |
|-------|------|----------|
| `Data` | `[]Data` | no |
| `HoldersCount` | `int` | no |

## TRC10TokensResponse

_struct · `validation/info/external/trc10.go`:12_

| Field | Type | Optional |
|-------|------|----------|
| `Data` | `[]struct { Symbol string `json:"abbr"` Decimals int `json:"precision"` HoldersCount int `json:"nrOfTokenHolders"` }` | no |

## TRC20TokensResponse

_struct · `validation/info/external/trc20.go`:12_

| Field | Type | Optional |
|-------|------|----------|
| `TRC20Tokens` | `[]struct { Symbol string `json:"symbol"` Decimals int `json:"decimals"` HoldersCount int `json:"holders_count"` }` | no |

## Version

_struct · `validation/tokenlist/model.go`:31_

| Field | Type | Optional |
|-------|------|----------|
| `Major` | `int` | no |
| `Minor` | `int` | no |
| `Patch` | `int` | no |

## See Also
- [validation](../../features/validation.md) <!-- rel:strong -->
- [go conventions](../../code-conventions/go-conventions.md) <!-- rel:strong -->
- [address and file validation](../../security/address-and-file-validation.md) <!-- rel:strong -->
- [client](../../features/client.md) <!-- rel:related -->
- [overview](../../build/overview.md) <!-- rel:related -->
