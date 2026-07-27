---
title: Code Conventions
category: code-conventions
tags: [go, error-handling, naming, patterns]
confidence: high
source: validation/errors.go, validation/info/asset.go, file/types.go
updated: 2026-07-22
---

# Code Conventions

## Error Handling Pattern

All validators in this library follow the composite error pattern:

```go
compErr := validation.NewErrComposite()
if err := someCheck(); err != nil {
    compErr.Append(err)
}
// ... more checks ...
if compErr.Len() > 0 {
    return compErr
}
return nil
```

**Rule**: wrap failures with a sentinel error using `fmt.Errorf("%w: detail", validation.ErrSomething)`. Never invent new error variables in sub-packages — use the shared sentinels from `validation/errors.go`.

**Rule**: do NOT return on first error. Always run all checks so the caller gets the full picture.

## Required Fields Pattern

Required-field checks (e.g. `ValidateAssetRequiredKeys`) use pointer types: `*string` is nil when absent, non-nil but empty when provided as `""`. Both cases are treated as "missing". The private `isEmpty(field string) bool` helper checks for empty string after a nil guard.

## Type Constants

File type constants in `file/types.go` follow `TypeXxxYyyy` — noun + descriptor, singular for files/folders. Example: `TypeAssetInfoFile`, `TypeChainFolder`. These constants are keys in the regex map — adding a new file type requires:
1. A constant in `file/types.go`
2. A regex in `file/path.go`
3. An entry in `GetRegexMap()`

## Naming

- Exported validators: `ValidateXxx` (not `CheckXxx`)
- Path constructors: `GetXxxPath` for string path builders (in `path/assets.go`)
- File constructors: `NewXxx` for value objects
- Private helpers: uncapitalized, e.g. `isEmpty`, `validateLogoSize`

## Concurrency

`file.Service` is the only concurrently-accessed type. It uses `sync.RWMutex` explicitly. No goroutines are spawned by this library — callers own their concurrency model.

## Testing

Tests live alongside source files (`*_test.go`). Run with `make test` (covers `-race` flag). See `file/json_test.go` and `path/assets_test.go` for the testing style.

## See Also
- [validation domain](../architecture/validation-domain.md) <!-- rel:strong -->
- [file domain](../architecture/file-domain.md) <!-- rel:strong -->
- [validate asset explain](../architecture/validate-asset-explain.md) <!-- rel:strong -->
- [models](../architecture/data/models.md) <!-- rel:strong -->
- [client domain](../architecture/client-domain.md) <!-- rel:strong -->
