---
category: architecture
subcategory: data
confidence: low
documentType: explanation
scope: repo
contentHash: c81c649a29d1
tags: [architecture, service]
source: architecture/data/models.md
verified: 2026-07-22
splitPartIndex: 2
splitPartTotal: 3
canonical: false
synthetic: split-part
---

## Data Models & Schemas (Part 2)

_struct · `validation/info/model.go`:43_

| Field | Type | Optional |
|-------|------|----------|
| `Name` | `*string` | no |
| `URL` | `*string` | no |

## Model

_struct · `validation/list/model.go`:4_

| Field | Type | Optional |
|-------|------|----------|
| `ID` | `*string` | no |
| `Name` | `*string` | no |
| `Description` | `*string` | no |
| `Website` | `*string` | no |
| `Staking` | `Staking` | no |
| `Payout` | `Payout` | no |
| `Status` | `Status` | no |

## Model

_struct · `validation/tokenlist/model.go`:6_

| Field | Type | Optional |
|-------|------|----------|
| `Name` | `string` | no |
| `LogoURI` | `string` | no |
| `Timestamp` | `string` | no |
| `Tokens` | `[]Token` | no |
| `Version` | `Version` | no |

## Pair

_struct · `validation/tokenlist/model.go`:25_

| Field | Type | Optional |
|-------|------|----------|
| `Base` | `string` | no |
| `LotSize` | `string` | no |
| `TickSize` | `string` | no |

## Path

_struct · `file/path.go`:78_

| Field | Type | Optional |
|-------|------|----------|
| `path` | `string` | no |
| `chain` | `coin.Coin` | no |
| `asset` | `string` | no |
| `fileType` | `string` | no |
| `regexpMap` | `map[string]*regexp.Regexp` | no |

## Payout

_struct · `validation/list/model.go`:20_

| Field | Type | Optional |
|-------|------|----------|
| `Commission` | `float64` | no |
| `PayoutDelay` | `int` | no |
| `PayoutPeriod` | `int` | no |

## Service

_struct · `file/service.go`:9_

| Field | Type | Optional |
|-------|------|----------|
| `mu` | `*sync.RWMutex` | no |
| `cache` | `map[string]*AssetFile` | no |

## Staking

_struct · `validation/list/model.go`:14_

| Field | Type | Optional |
|-------|------|----------|
| `FreeSpace` | `int` | no |
| `MinDelegation` | `int` | no |
| `OpenForDelegation` | `bool` | no |

## Status

_struct · `validation/list/model.go`:26_

| Field | Type | Optional |
|-------|------|----------|
| `Disabled` | `bool` | no |
| `Note` | `string` | no |

## Tag

_struct · `client/assets-manager/model.go`:49_

| Field | Type | Optional |
|-------|------|----------|
| `ID` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## TagValuesResp

_struct · `client/assets-manager/model.go`:45_

| Field | Type | Optional |
|-------|------|----------|
| `Tags` | `[]Tag` | no |

## Token

_struct · `validation/tokenlist/model.go`:14_

| Field | Type | Optional |
|-------|------|----------|
| `Asset` | `string` | no |
| `Type` | `types.TokenType` | no |
| `Address` | `string` | no |
| `Name` | `string` | no |
| `Symbol` | `string` | no |
| `Decimals` | `uint` | no |
| `LogoURI` | `string` | no |
| `Pairs` | `[]Pair` | no |

## TokenInfo

_struct · `validation/info/external/external.go`:18_

| Field | Type | Optional |
|-------|------|----------|
| `Symbol` | `string` | no |
| `Decimals` | `int` | no |
| `HoldersCount` | `int` | no |

## TokenInfoERC20

_struct · `validation/info/external/erc20.go`:12_

| Field | Type | Optional |
|-------|------|----------|
| `Decimals` | `string` | no |
| `HoldersCount` | `int` | no |

## TokenInfoSPL

_struct · `validation/info/external/spl.go`:11_
