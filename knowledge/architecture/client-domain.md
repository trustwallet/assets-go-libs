---
title: Client Domain — Assets Manager HTTP API
category: architecture
tags: [client, http, assets-manager, api]
confidence: high
source: client/assets-manager/client.go, client/assets-manager/model.go
updated: 2026-07-22
---

# Client Domain — Assets Manager

The `client/assets-manager` package provides an HTTP client for the Trust Wallet assets-manager service. This service acts as an authoritative source for validation rules that cannot be hard-coded (e.g., the allowed tag values, which change as new categories are added).

## Client

```go
type Client struct{ req client.Request }

func InitClient(url string, errorHandler client.HttpErrorHandler) Client
```

`InitClient` wraps `go-libs/client.InitJSONClient`, which handles JSON serialization and error handling. The `Client` type is the sole entry point for this package.

## API Endpoints

### `POST /v1/validate/asset_info`
```go
func (c *Client) ValidateAssetInfo(req *AssetValidationReq) (AssetValidationResp, error)
```
Sends an `AssetValidationReq` (mirrors `info.AssetModel` field-for-field) to the assets-manager for server-side validation. Returns a list of errors if the asset info fails. This provides an alternative/complementary validation path to the local validators in `validation/info`.

### `GET /v1/values/tags`
```go
func (c *Client) GetTagValues() (TagValuesResp, error)
```
Fetches the allowed tag values. Responses are cached for 1 hour (via `go-libs/client.Request.GetWithCache`). The result is used by `info.ValidateTags` — callers fetch this list first, then pass it to `ValidateAsset`/`ValidateCoin`.

## Cross-Domain Integration

The client domain is the ONLY integration point between this library and an external HTTP service. There are no cross-domain bridges between `client` and `validation` within this library itself — the caller is responsible for fetching tag values and passing them to the validation functions. This is intentional: it keeps validation pure and testable without network access.

## Models

`AssetValidationReq` is structurally identical to `info.AssetModel` — both represent an asset's `info.json`. They exist as separate types because `AssetValidationReq` is a DTO for the HTTP API while `AssetModel` is the domain model for local validation. Any future divergence between the API contract and the local model should be handled here, not by merging the types.

## See Also
- [features/client.md](../features/client.md)
- [architecture/validation-domain.md](validation-domain.md)
- [overview](../build/overview.md) <!-- rel:strong -->
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
- [address and file validation](../security/address-and-file-validation.md) <!-- rel:strong -->
