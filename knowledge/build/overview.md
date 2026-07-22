---
title: Build and Development
category: build
tags: [go, makefile, ci, lint, test]
confidence: high
source: Makefile, go.mod, .github/workflows/check.yml, .golangci.yml
updated: 2026-07-22
---

# Build and Development

## Local Development

```bash
make          # runs: fmt + lint + test
make test     # go test -cover -race -coverprofile=coverage.txt -v ./...
make fmt      # gofmt -w on all .go files
make lint     # golangci-lint run --timeout=2m
```

`golangci-lint` is pinned to v1.45.2 and installed to `./bin/` if not present.

## Module

Module path: `github.com/trustwallet/assets-go-libs` (Go 1.18+)

### Key Dependencies

| Dependency | Purpose |
|---|---|
| `github.com/trustwallet/go-primitives` | Coin/chain definitions, token types, address utilities |
| `github.com/trustwallet/go-libs` | HTTP client with JSON + caching support |
| `golang.org/x/image` | PNG decoding for image dimension validation |
| `gopkg.in/yaml.v2` | YAML parsing (transitive) |
| `github.com/prometheus/client_golang` | Metrics (transitive via go-libs) |

`go-primitives` is the authoritative source for blockchain identifiers (`coin.ETHEREUM`, `coin.TRON`, etc.) and token type values (`ERC20`, `BEP20`, `TRC20`, etc.). Any addition of a new chain or token type flows from there.

## CI

GitHub Actions runs on PRs and pushes to `master` (`.github/workflows/check.yml`):
1. Set up Go 1.17
2. `go get -v -t -d ./...`
3. `make test`
4. `make lint`

The workflow targets `master` (not `main`). The `knowledge-sync.yml` workflow (added by sdd-init) syncs knowledge to the SDD KB on every push to `main`.

## Linting

`.golangci.yml` configures the linters. `nolint:gochecknoglobals` is applied to package-level var blocks in `validation/info/values.go` — these are intentional global value sets (allowed status values, link keys) that are read-only and safe.

## See Also
- [validation domain](../architecture/validation-domain.md) <!-- rel:strong -->
- [address and file validation](../security/address-and-file-validation.md) <!-- rel:strong -->
- [client domain](../architecture/client-domain.md) <!-- rel:strong -->
- [validate asset explain](../architecture/validate-asset-explain.md) <!-- rel:strong -->
- [file domain](../architecture/file-domain.md) <!-- rel:related -->
