---
title: File Domain — AssetFile, Path, and Service
category: architecture
tags: [file, path, asset-file, service, type-system]
confidence: high
source: file/path.go, file/asset_file.go, file/service.go, file/types.go, file/files.go
updated: 2026-07-22
---

# File Domain

The `file` package provides the type system and utilities for navigating the Trust Wallet assets repository's directory structure. Every path in the assets repo is classified into a known type via regex matching.

## Directory Structure Model

The assets repository follows a fixed layout:

```
<root>/
  blockchains/
    <chain>/                  → TypeChainFolder
      info/                   → TypeChainInfoFolder
        info.json             → TypeChainInfoFile
        logo.png              → TypeChainLogoFile
      assets/                 → TypeAssetsFolder
        <token_address>/      → TypeAssetFolder
          info.json           → TypeAssetInfoFile
          logo.png            → TypeAssetLogoFile
      tokenlist.json          → TypeTokenListFile
      tokenlist-extended.json → TypeTokenListExtendedFile
      validators/             → TypeValidatorsFolder
        list.json             → TypeValidatorsListFile
        assets/               → TypeValidatorsAssetsFolder
          <address>/          → TypeValidatorsAssetFolder
            logo.png          → TypeValidatorsLogoFile
  dapps/                      → TypeDappsFolder
    <name>.png                → TypeDappsLogoFile
```

The 21 type constants in `file/types.go` map 1:1 to positions in this hierarchy.

## `Path` — the classification engine

`NewPath(path string) *Path` is the central factory. It tests the path against all 21 regex patterns (`GetRegexMap`) and, on match, extracts:
- `chain` — the `coin.Coin` for the first capture group (chain handle)
- `asset` — the token address from the second capture group (when present)
- `fileType` — one of the 21 type constants

Callers use `AssetFile.Type()`, `.Chain()`, `.Asset()` to branch on the result without re-parsing strings. The AST-inferred `defineFileType` is private; the public surface is `NewPath`.

## `AssetFile` — the value object

```go
type AssetFile struct{ path *Path }
```

`NewAssetFile(path string)` wraps a path in a `Path` and exposes `.Path()`, `.Type()`, `.Chain()`, `.Asset()`. It is the canonical way to represent a file in the assets repo throughout the automation code — it is constructed in `file.Service` and in test code.

`NewAssetFile` is the most-called constructor in the library (fanIn=3, called by `Service.NewService`, `Service.getFile`, and `Service.UpdateFile`).

## `Service` — thread-safe file cache

`Service` is a `sync.RWMutex`-protected cache of `*AssetFile` values keyed by path. It lazily constructs `AssetFile` entries on first access and provides `UpdateFile` to rename a file across all cached paths (used when a file is moved during an asset PR review).

### Dead-code note

`Service.getFile` is identified as dead code by the static analysis (zero external fan-in within this repo). It is private and used only via `GetAssetFile` — this is a false positive; the private method is the implementation, not an unused symbol.

## See Also
- [features/file.md](../features/file.md)
- [features/path.md](../features/path.md)
- [architecture/project-structure.md](project-structure.md)
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
- [overview](../build/overview.md) <!-- rel:related -->
