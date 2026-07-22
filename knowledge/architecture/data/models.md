---
category: architecture
subcategory: data
confidence: low
documentType: explanation
scope: repo
contentHash: d2fcd33367bb
tags: [architecture, domain]
source: architecture/data/models.md
verified: 2026-07-22
splitPartIndex: 1
splitPartTotal: 3
canonical: true
---

## Data Models & Schemas
<!-- sdd-knowledge-generated -->

## Model Semantics

The models in this library represent the Trust Wallet assets repository's JSON file formats. All pointer fields in `AssetModel` and `CoinModel` are optional in JSON but **validated as required** by `ValidateAssetRequiredKeys` / `ValidateCoinRequiredKeys` — the pointer type allows distinguishing "not provided" from "empty string" during validation.

**`AssetModel` vs `AssetValidationReq`**: structurally identical but semantically distinct. `AssetModel` is the domain model for local validation; `AssetValidationReq` is the DTO for the remote assets-manager API. Keep them separate — any future divergence (e.g. server adds a field the local validator ignores) stays isolated.

**`tokenlist.Token.Asset`**: composite ID `<chain_handle>/<token_address>`, parsed by `go-primitives/asset.ParseID`. On EVM chains, the parsed address must pass EIP-55 checksum.

**`validation/list.Model`**: represents a staking validator entry (`validators/list.json`). `Staking.FreeSpace` is the number of open delegation slots; `Payout.PayoutDelay` is in blocks.

**`ValidateAssetID` vs `tokenlist.validateAssetID`**: distinct functions in different packages — see [validation-domain.md](../validation-domain.md).

> Field-level shape of data models extracted via tree-sitter: TS interfaces / type aliases, Zod `z.object` schemas, Go/Rust/Swift structs, Kotlin data classes, and Python dataclasses. Scoped to domain data — UI views/props, view-models, design tokens (theme/style/colors), and constant/identifier namespaces are excluded. Deterministic, no LLM.

## AssetFile

_struct · `file/asset_file.go`:7_

| Field | Type | Optional |
|-------|------|----------|
| `path` | `*Path` | no |

## AssetModel

_struct · `validation/info/model.go`:18_

| Field | Type | Optional |
|-------|------|----------|
| `Name` | `*string` | no |
| `Symbol` | `*string` | no |
| `Type` | `*string` | no |
| `Decimals` | `*int` | no |
| `Description` | `*string` | no |
| `Website` | `*string` | no |
| `Explorer` | `*string` | no |
| `Research` | `string` | no |
| `Status` | `*string` | no |
| `ID` | `*string` | no |
| `Links` | `[]Link` | no |
| `ShortDesc` | `*string` | no |
| `Audit` | `*string` | no |
| `AuditReport` | `*string` | no |
| `Tags` | `[]string` | no |
| `Code` | `*string` | no |
| `Ticker` | `*string` | no |
| `ExplorerEth` | `*string` | no |
| `Address` | `*string` | no |
| `Twitter` | `*string` | no |
| `CoinMarketcap` | `*string` | no |
| `DataSource` | `*string` | no |

## AssetValidationReq

_struct · `client/assets-manager/model.go`:4_

| Field | Type | Optional |
|-------|------|----------|
| `Name` | `*string` | no |
| `Symbol` | `*string` | no |
| `Type` | `*string` | no |
| `Decimals` | `*int` | no |
| `Description` | `*string` | no |
| `Website` | `*string` | no |
| `Explorer` | `*string` | no |
| `Research` | `string` | no |
| `Status` | `*string` | no |
| `ID` | `*string` | no |
| `Links` | `[]Link` | no |
| `ShortDesc` | `*string` | no |
| `Audit` | `*string` | no |
| `AuditReport` | `*string` | no |
| `Tags` | `[]string` | no |
| `Code` | `*string` | no |
| `Ticker` | `*string` | no |
| `ExplorerEth` | `*string` | no |
| `Address` | `*string` | no |
| `Twitter` | `*string` | no |
| `CoinMarketcap` | `*string` | no |
| `DataSource` | `*string` | no |

## AssetValidationResp

_struct · `client/assets-manager/model.go`:29_

| Field | Type | Optional |
|-------|------|----------|
| `Status` | `string` | no |
| `Errors` | `[]Error` | no |

## Client

_struct · `client/assets-manager/client.go`:9_

| Field | Type | Optional |
|-------|------|----------|
| `req` | `client.Request` | no |

## CoinModel

_struct · `validation/info/model.go`:4_

| Field | Type | Optional |
|-------|------|----------|
| `Name` | `*string` | no |
| `Website` | `*string` | no |
| `Description` | `*string` | no |
| `Explorer` | `*string` | no |
| `Research` | `string` | no |
| `Symbol` | `*string` | no |
| `Type` | `*string` | no |
| `Decimals` | `*int` | no |
| `Status` | `*string` | no |
| `Tags` | `[]string` | no |
| `Links` | `[]Link` | no |

## Data

_struct · `validation/info/external/spl.go`:16_

| Field | Type | Optional |
|-------|------|----------|
| `Decimals` | `int` | no |

## ErrComposite

_struct · `validation/errors.go`:26_

| Field | Type | Optional |
|-------|------|----------|
| `errors` | `[]error` | no |

## Error

_struct · `client/assets-manager/model.go`:39_

| Field | Type | Optional |
|-------|------|----------|
| `Message` | `string` | no |

## Link

_struct · `client/assets-manager/model.go`:34_

| Field | Type | Optional |
|-------|------|----------|
| `Name` | `*string` | no |
| `URL` | `*string` | no |
