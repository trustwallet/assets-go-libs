# Data Models & Schemas

<!-- sdd-knowledge-generated -->

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

## Link

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

