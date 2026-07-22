---
title: Validation Domain — Architecture and Intent
category: architecture
tags: [validation, asset, coin, tokenlist, errors]
confidence: high
source: validation/*.go, validation/info/*.go, validation/tokenlist/*.go, validation/list/*.go
updated: 2026-07-22
---

# Validation Domain

The `validation` package family is the core of `assets-go-libs`. It implements the business rules that guard the Trust Wallet [assets repository](https://github.com/trustwallet/assets) — every asset addition or change must pass these checks before it is accepted.

## Sub-packages

| Package | Responsibility |
|---|---|
| `validation` | Shared error types, composite error accumulator, address validators, file/image validators |
| `validation/info` | Asset info (`info.json`) and coin info (`info.json`) validators — models + field validators |
| `validation/tokenlist` | Token list (`tokenlist.json`) validators — cross-referencing info files |
| `validation/list` | Validator list validators |

## The Composite Error Pattern

Every validator in this library collects all failures before returning, using `ErrComposite` (`validation/errors.go`):

```go
compErr := validation.NewErrComposite()
if err := SomeCheck(); err != nil {
    compErr.Append(err)
}
if compErr.Len() > 0 {
    return compErr
}
return nil
```

This ensures that a single `ValidateAsset` call surfaces all problems at once (not just the first), so callers can present the full list to a contributor. `ErrComposite` implements `error`.

## Sentinel Errors

All validation failures are expressed as wrapped sentinel errors from `validation/errors.go`:

| Sentinel | Usage |
|---|---|
| `ErrMissingFile` | Required file absent from folder |
| `ErrMissingField` | Required JSON field empty/nil |
| `ErrInvalidField` | Field present but fails a constraint |
| `ErrInvalidAddress` | Blockchain address fails chain-specific format check |
| `ErrInvalidJSON` | File is not valid JSON |
| `ErrInvalidImgDimension` | Logo PNG outside allowed dimensions |
| `ErrInvalidFileNameCase` | File/folder name case violation |
| `ErrInvalidFileExt` | Wrong file extension |
| `ErrInvalidFileSize` | Logo exceeds 100 KB limit |
| `ErrInvalidFileNameLength` | Address/folder name wrong length |

Callers can use `errors.Is` to test for a specific sentinel.

## Validation Entry Points

### `ValidateAsset(a AssetModel, chain coin.Coin, addr string) error`
The primary validator for asset `info.json` files. Validates: required keys presence → type vs chain matching → ID equals address → decimals range → BEP2 decimals exactly 8 → status in {active, spam, abandoned} → description ≤ 600 chars / no whitespace → website required unless description is "-" → explorer URL → links (≥2 required, HTTPS, known names).

### `ValidateCoin(c CoinModel, allowedTags []string) error`
Validates coin `info.json`. Same structure as `ValidateAsset` but without ID validation; validates coin type = "coin" and tags against the allowed list fetched from assets-manager.

### `ValidateTokenList(model Model, chain coin.Coin, tokenListPath string) error`
Cross-validates `tokenlist.json`: for each token, reads the corresponding `info.json` from disk and checks type/symbol/decimals/name consistency plus active status. Also validates that all pair references point to tokens that exist in the same token list.

### `ValidateAssetAddress(chain coin.Coin, address string) error`
Dispatches to chain-specific address validators. EVM chains use EIP-55 checksum validation; TRON uses prefix + length + mixed-case checks.

## Image Constraints

Logo files must be PNG, square, between 128×128 and 512×512 pixels, and ≤ 100 KB. The CI-specific `ValidatePngImageDimensionForCI` uses a relaxed 60px minimum to accommodate legacy logos that predate the current standard.

## See Also
- [architecture/call-graph.md](call-graph.md)
- [features/validation.md](../features/validation.md)
- [architecture/data/models.md](data/models.md)
- [address and file validation](../security/address-and-file-validation.md) <!-- rel:strong -->
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:related -->
