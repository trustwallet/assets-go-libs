---
title: Security and Address Validation
category: security
tags: [address-validation, eip55, blockchain, evm, tron]
confidence: high
source: validation/address.go, validation/image.go, validation/file.go
updated: 2026-07-22
---

# Security — Address Validation and File Safety

## Blockchain Address Validation

Address validation is chain-specific, dispatched in `ValidateAssetAddress`:

| Chain | Validation rule |
|---|---|
| EVM chains (Ethereum, BNB Chain, Polygon, etc.) | EIP-55 checksum — address must equal its own EIP-55 checksummed form |
| TRON (TRC20) | 34 chars, starts with `T`, mixed case |
| TRON (TRC10) | Numeric string only |
| Wanchain | EIP-55 checksum with reversed case (custom format) |
| Tezos | Must start with `tz` |
| Cosmos | `cosmosvaloper1` prefix, 52 chars, lowercase |
| Terra | `terravaloper1` prefix, 51 chars, lowercase |
| Kava | `kavavaloper1` prefix, 50 chars, lowercase |
| Waves | `3P` prefix, 35 chars, mixed case |

EIP-55 validation uses `go-primitives/address.EIP55Checksum` — it computes the correct checksum and errors if the provided address differs. This prevents case-mangled contract addresses from entering the assets repository.

## File Safety

`ValidateFileInPR` enforces that only safe file paths are accepted in pull requests to the assets repo:
- `dapps/**` — any file
- `blockchains/**` — only asset files (has "assets" in path), `allowlist.json`, or `validators/list.json`

This prevents arbitrary file injection into the repository via automated PRs.

## Image Safety

PNG logos are validated for:
- Dimensions: 128×128 to 512×512 pixels, square (width == height)
- File size: ≤ 100 KB
- Format: must be a valid PNG (decoded via `golang.org/x/image/png`)

`ValidatePngImageDimensionForCI` uses a relaxed 60px minimum for the CI linter, which handles legacy logos that predate the 128px requirement. This variant should be retired once all logos meet the current standard (see the TODO comment in `validation/image.go`).

## No Secrets in This Library

This library contains no authentication logic, no API keys, and no network credentials. The HTTP client (`client/assets-manager`) is initialized with a URL by the caller — credentials are the caller's responsibility. The `http.GetHTTPResponseBytes` utility in `http/http.go` uses the standard `http.DefaultClient` with no custom TLS or auth.

## See Also
- [validation domain](../architecture/validation-domain.md) <!-- rel:strong -->
- [overview](../build/overview.md) <!-- rel:related -->
- [client domain](../architecture/client-domain.md) <!-- rel:related -->
- [validate asset explain](../architecture/validate-asset-explain.md) <!-- rel:related -->
- [models](../architecture/data/models.md) <!-- rel:related -->
