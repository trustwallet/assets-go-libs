---
title: ValidateAsset — Contract and Invariants
category: architecture
tags: [validation, asset, info.json, god-node]
confidence: high
source: validation/info/asset.go, validation/info/fields_validators.go, validation/info/values.go
updated: 2026-07-22
---

# `ValidateAsset`

`ValidateAsset` is the central validation entry point for asset `info.json` files (fanOut=10 — it delegates to 10 distinct validators). Every asset PR through the Trust Wallet assets automation pipeline goes through this function.

## Signature

```go
func ValidateAsset(a AssetModel, chain coin.Coin, addr string) error
```

## Validation Sequence

1. **Required keys** (`ValidateAssetRequiredKeys`): all 9 fields must be non-nil and non-empty: `name`, `type`, `symbol`, `decimals`, `description`, `website`, `explorer`, `status`, `id`. Missing fields are listed by name.

2. **Type vs chain** (`ValidateAssetType`): `type` field must match the chain derived from the asset path (e.g. ERC20 on Ethereum, TRC20 on Tron). Must be ALL-CAPS.

3. **ID = address** (`ValidateAssetID`): `id` field must equal the token's on-chain address exactly (case-sensitive). A case mismatch returns a distinct error so callers can distinguish "wrong address" from "wrong case".

4. **Decimals range** (`ValidateDecimals`): 0–30 inclusive.

5. **BEP2 decimals** (`ValidateAssetDecimalsAccordingType`): BEP2 tokens must have exactly 8 decimals. This is the only type with a hardcoded decimals requirement.

6. **Status** (`ValidateStatus`): must be one of `active`, `spam`, `abandoned`.

7. **Description** (`ValidateDescription`): max 600 characters; no newlines or double spaces.

8. **Website requirement** (`ValidateDescriptionWebsite`): website must be non-empty unless description equals `"-"` (the sentinel for "no info available").

9. **Explorer URL** (`ValidateExplorer`): must match `coin.GetCoinExploreURL(chain, addr, type)` or a known alternative pattern. Alternatives for Ethereum: `etherscan.io/token/<name>`, `explorer.<name>.io`, `scan.<name>.io`. Chains with no explorer URL registered skip this check.

10. **Links** (`ValidateLinks`): if ≥ 2 links, each must have `name` + `url`; `name` must be a known key (github, twitter, telegram, medium, discord, reddit, etc.); URL must be HTTPS; some names enforce URL prefixes (e.g. github → `https://github.com/`, twitter → `https://x.com/`).

## Error Accumulation

All 10 validators run regardless of individual failures (composite error pattern). The caller receives a single `ErrComposite` listing every problem.

## What Calls It

`ValidateAsset` is called by the assets automation tooling (specifically by the assets repository's CI scripts — this library is the imported dependency). Within this library, `ValidateTokenList` indirectly calls it via `validateChainOrAssetInfo`.

## See Also
- [architecture/validation-domain.md](validation-domain.md)
- [architecture/data/models.md](data/models.md)
- [features/validation.md](../features/validation.md)
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
- [address and file validation](../security/address-and-file-validation.md) <!-- rel:strong -->
