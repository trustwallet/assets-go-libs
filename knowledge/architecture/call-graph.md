---
category: architecture
confidence: low
documentType: explanation
scope: repo
contentHash: 93083d083b10
tags: [architecture]
source: architecture/call-graph.md
verified: 2026-07-22
splitPartIndex: 1
splitPartTotal: 2
canonical: true
---

## Call Graph
<!-- sdd-knowledge-generated -->

> Deterministic call graph extracted via tree-sitter (no LLM). Direct calls are resolved by lexical scope + imports; **dynamic dispatch and ambiguous name matches are withheld** rather than guessed. The `## Calls` table is `EXTRACTED` (a single resolved target). Member calls (`obj.method()`) whose method name resolves to exactly one definition repo-wide are recovered separately under `## Inferred calls` (`INFERRED` — receiver type unverified, but a single plausible target).

## Calls

| Caller | Callee | Example args | Sites |
|--------|--------|--------------|-------|
| `NewAssetFile` | `NewPath(path: string): *Path` | `(path)` | 1 |
| `NewPath` | `GetRegexMap(): map[string]*regexp.Regexp` | `()` | 1 |
| `getFile` | `NewAssetFile(path: string): *AssetFile` | `(path)` | 1 |
| `NewService` | `NewAssetFile(path: string): *AssetFile` | `(path)` | 1 |
| `UpdateFile` | `NewAssetFile(path: string): *AssetFile` | `(newPath)` | 1 |
| `GetHTTPResponse` | `GetHTTPResponseBytes(url: string): ([]byte, error)` | `(url)` | 1 |
| `GetPNGImageDimensions` | `GetPNGImageDimensionsFromReader(r: io.Reader): (width, height int, err error)` | `(file)` | 1 |
| `ValidateAssetAddress` | `ValidateETHForkAddress(chain: coin.Coin, addr: string): error` | `(chain, address)` | 1 |
| `ValidateAssetAddress` | `ValidateTronAddress(addr: string): error` | `(address)` | 1 |
| `ValidateValidatorsAddress` | `ValidateAddress(address: string, length: int): error` | `(address, "cosmosvaloper1", 52)` | 3 |
| `ValidateValidatorsAddress` | `ValidateTezosAddress(addr: string): error` | `(address)` | 1 |
| `ValidateValidatorsAddress` | `ValidateTronAddress(addr: string): error` | `(address)` | 1 |
| `ValidateValidatorsAddress` | `ValidateWavesAddress(addr: string): error` | `(address)` | 1 |
| `ValidateJSON` | `checkDuplicateKey(d: *json.Decoder, path: []string): error` | `(…, nil)` | 1 |
| `ValidateAllowedFiles` | `NewErrComposite(): *ErrComposite` | `()` | 1 |
| `ValidateHasFiles` | `NewErrComposite(): *ErrComposite` | `()` | 1 |
| `ValidateLogoFileSize` | `validateLogoSize(imgBytesCount: int): error` | `(…)` | 1 |
| `ValidateLogoStreamSize` | `validateLogoSize(imgBytesCount: int): error` | `(…)` | 1 |
| `ValidatePngImageDimension` | `ValidateImageDimension(width: int): error` | `(width, height)` | 1 |
| `ValidateAsset` | `ValidateAssetDecimalsAccordingType(assetType: string, decimals: int): error` | `(…, …)` | 1 |
| `ValidateAsset` | `ValidateAssetID(id: string): error` | `(…, addr)` | 1 |
| `ValidateAsset` | `ValidateAssetRequiredKeys(a: AssetModel): error` | `(a)` | 1 |
| `ValidateAsset` | `ValidateAssetType(assetType: string, chain: coin.Coin): error` | `(…, chain)` | 1 |
| `ValidateAsset` | `ValidateDecimals(decimals: int): error` | `(…)` | 1 |
| `ValidateAsset` | `ValidateDescription(description: string): error` | `(…)` | 1 |
| `ValidateAsset` | `ValidateDescriptionWebsite(description: string): error` | `(…, …)` | 1 |
| `ValidateAsset` | `ValidateExplorer(explorer: string, chain: coin.Coin, addr: string): error` | `(…, …, chain, addr, …)` | 1 |
| `ValidateAsset` | `ValidateLinks(links: []Link): error` | `(a.Links)` | 1 |
| `ValidateAsset` | `ValidateStatus(status: string): error` | `(…)` | 1 |
| `ValidateCoin` | `ValidateCoinRequiredKeys(c: CoinModel): error` | `(c)` | 1 |
| `ValidateCoin` | `ValidateCoinType(assetType: string): error` | `(…)` | 1 |
| `ValidateCoin` | `ValidateDecimals(decimals: int): error` | `(…)` | 1 |
| `ValidateCoin` | `ValidateDescription(description: string): error` | `(…)` | 1 |
| `ValidateCoin` | `ValidateDescriptionWebsite(description: string): error` | `(…, …)` | 1 |
| `ValidateCoin` | `ValidateLinks(links: []Link): error` | `(c.Links)` | 1 |
| `ValidateCoin` | `ValidateStatus(status: string): error` | `(…)` | 1 |
| `ValidateCoin` | `ValidateTags(tags: []string): error` | `(c.Tags, allowedTags)` | 1 |
| `GetTokenInfo` | `GetTokenInfoForERC20(tokenID: string): (*TokenInfo, error)` | `(tokenID)` | 1 |
| `GetTokenInfo` | `GetTokenInfoByScraping(url: string): (*TokenInfo, error)` | `(…)` | 4 |
| `GetTokenInfo` | `GetTokenInfoForSPL(tokenID: string): (*TokenInfo, error)` | `(tokenID)` | 1 |
| `GetTokenInfo` | `GetTokenInfoForTRC10(tokenID: string): (*TokenInfo, error)` | `(tokenID)` | 1 |
| `GetTokenInfo` | `GetTokenInfoForTRC20(tokenID: string): (*TokenInfo, error)` | `(tokenID)` | 1 |
| `ValidateAssetRequiredKeys` | `isEmpty(field: string): bool` | `(…)` | 7 |
| `ValidateCoinRequiredKeys` | `isEmpty(field: string): bool` | `(…)` | 7 |
| `ValidateExplorer` | `explorerURLAlternatives(chain: string): []string` | `(chain.Handle, name)` | 1 |
| `ValidateLinks` | `linkNameAllowed(str: string): bool` | `(…)` | 1 |
| `ValidateLinks` | `supportedLinkNames(): []string` | `()` | 1 |
| `ValidateList` | `validateRequiredFields(model: Model): error` | `(validator)` | 1 |
| `validateTokenAddress` | `validateAssetID(chain: coin.Coin, id: string): error` | `(chain, token.Asset)` | 2 |
| `ValidateTokenList` | `validateChainOrAssetInfo(token: Token, chain: coin.Coin, tokenListPath: string): error` | `(token, chain, tokenListPath)` | 1 |
| `ValidateTokenList` | `validateTokenAddress(chain: coin.Coin, token: Token): error` | `(chain, token)` | 1 |
| `ValidateTokenList` | `validateTokenListPairs(model: Model): error` | `(model)` | 1 |

## Callers (reverse)

| Symbol | Called by |
|--------|-----------|
| `checkDuplicateKey` | `ValidateJSON` |
| `explorerURLAlternatives` | `ValidateExplorer` |
| `GetHTTPResponseBytes` | `GetHTTPResponse` |
| `GetPNGImageDimensionsFromReader` | `GetPNGImageDimensions` |
| `GetRegexMap` | `NewPath` |
| `GetTokenInfoByScraping` | `GetTokenInfo` |
| `GetTokenInfoForERC20` | `GetTokenInfo` |
| `GetTokenInfoForSPL` | `GetTokenInfo` |
| `GetTokenInfoForTRC10` | `GetTokenInfo` |
| `GetTokenInfoForTRC20` | `GetTokenInfo` |
| `isEmpty` | `ValidateAssetRequiredKeys`, `ValidateCoinRequiredKeys` |
| `linkNameAllowed` | `ValidateLinks` |
| `NewAssetFile` | `NewService`, `UpdateFile`, `getFile` |
| `NewErrComposite` | `ValidateAllowedFiles`, `ValidateHasFiles` |
| `NewPath` | `NewAssetFile` |
| `supportedLinkNames` | `ValidateLinks` |
| `ValidateAddress` | `ValidateValidatorsAddress` |
| `ValidateAssetDecimalsAccordingType` | `ValidateAsset` |
| `validateAssetID` | `validateTokenAddress` |
| `ValidateAssetID` | `ValidateAsset` |
| `ValidateAssetRequiredKeys` | `ValidateAsset` |
| `ValidateAssetType` | `ValidateAsset` |
| `validateChainOrAssetInfo` | `ValidateTokenList` |
| `ValidateCoinRequiredKeys` | `ValidateCoin` |
| `ValidateCoinType` | `ValidateCoin` |
| `ValidateDecimals` | `ValidateAsset`, `ValidateCoin` |
| `ValidateDescription` | `ValidateAsset`, `ValidateCoin` |
| `ValidateDescriptionWebsite` | `ValidateAsset`, `ValidateCoin` |
| `ValidateETHForkAddress` | `ValidateAssetAddress` |
| `ValidateExplorer` | `ValidateAsset` |
| `ValidateImageDimension` | `ValidatePngImageDimension` |
| `ValidateLinks` | `ValidateAsset`, `ValidateCoin` |
| `validateLogoSize` | `ValidateLogoFileSize`, `ValidateLogoStreamSize` |
| `validateRequiredFields` | `ValidateList` |
| `ValidateStatus` | `ValidateAsset`, `ValidateCoin` |
| `ValidateTags` | `ValidateCoin` |
| `ValidateTezosAddress` | `ValidateValidatorsAddress` |
| `validateTokenAddress` | `ValidateTokenList` |
| `validateTokenListPairs` | `ValidateTokenList` |
| `ValidateTronAddress` | `ValidateAssetAddress`, `ValidateValidatorsAddress` |
| `ValidateWavesAddress` | `ValidateValidatorsAddress` |

## Inferred calls (member, unique name)

> `INFERRED` (confidence 0.9): `obj.method()` where `method` resolves to exactly one definition repo-wide. Receiver type is not resolved; common method names (init/update/onCreate) remain withheld as ambiguous. Use SCIP (`--scip`) for type-resolved member calls at full certainty.

| Caller | Callee | Example args | Sites |
|--------|--------|--------------|-------|
| `Path` | `String(): string` | `()` | 1 |
| `NewPath` | `defineFileType(path: string): (string, *regexp.Regexp)` | `(path)` | 1 |
| `ReadLocalFileStructure` | `Contains(str: string, entries: []string): bool` | `(path, filesToSkip)` | 1 |
| `GetAssetFile` | `getFile(path: string): *AssetFile` | `(path)` | 1 |
| `UpdateFile` | `Contains(str: string, entries: []string): bool` | `(path, oldFileBaseName)` | 1 |
| `CreatePNGFromURL` | `GetHTTPResponseBytes(url: string): ([]byte, error)` | `(logoURL)` | 1 |
| `ValidateAddress` | `IsLowerCase(str: string): bool` | `(address)` | 1 |
| `ValidateETHForkAddress` | `ReverseCase(str: string): string` | `(checksum)` | 1 |
| `ValidateTronAddress` | `IsLowerCase(str: string): bool` | `(addr)` | 1 |
| `ValidateTronAddress` | `IsUpperCase(str: string): bool` | `(addr)` | 1 |
| `ValidateWavesAddress` | `IsLowerCase(str: string): bool` | `(addr)` | 1 |
| `ValidateWavesAddress` | `IsUpperCase(str: string): bool` | `(addr)` | 1 |
| `checkDuplicateKey` | `Token` | `()` | 4 |
| `ValidateAllowedFiles` | `Contains(str: string, entries: []string): bool` | `(…, allowedFiles)` | 1 |
| `ValidateAllowedFiles` | `Append(err: error)` | `(…)` | 1 |
| `ValidateAllowedFiles` | `Len(): int` | `()` | 1 |
| `ValidateHasFiles` | `Append(err: error)` | `(…)` | 2 |
| `ValidateHasFiles` | `Len(): int` | `()` | 1 |
| `ValidateLowercase` | `IsLowerCase(str: string): bool` | `(name)` | 1 |
| `ValidatePngImageDimension` | `GetPNGImageDimensions(path: string): (width, height int, err error)` | `(path)` | 1 |
| `ValidatePngImageDimensionForCI` | `GetPNGImageDimensions(path: string): (width, height int, err error)` | `(path)` | 1 |
| `ValidateAsset` | `Append(err: error)` | `(err)` | 9 |
| `ValidateAsset` | `Len(): int` | `()` | 1 |
| `ValidateAsset` | `NewErrComposite(): *ErrComposite` | `()` | 1 |
| `ValidateCoin` | `Append(err: error)` | `(err)` | 7 |
| `ValidateCoin` | `Len(): int` | `()` | 1 |
| `ValidateCoin` | `NewErrComposite(): *ErrComposite` | `()` | 1 |
| `GetTokenInfoForERC20` | `GetHTTPResponse(url: string, result: interface{}): error` | `(url, …)` | 1 |
| `GetTokenInfoByScraping` | `GetHTTPResponseBytes(url: string): ([]byte, error)` | `(url)` | 1 |
| `GetTokenInfoForSPL` | `GetHTTPResponse(url: string, result: interface{}): error` | `(url, …)` | 1 |
| `GetTokenInfoForTRC10` | `GetHTTPResponse(url: string, result: interface{}): error` | `(url, …)` | 1 |
| `GetTokenInfoForTRC20` | `GetHTTPResponse(url: string, result: interface{}): error` | `(url, …)` | 1 |
| `ValidateAssetRequiredKeys` | `Difference(a: []string): []string` | `(requiredAssetFields, fields)` | 1 |
| `ValidateCoinRequiredKeys` | `Difference(a: []string): []string` | `(requiredCoinFields, fields)` | 1 |
| `ValidateDescription` | `Contains(str: string, entries: []string): bool` | `(description, ch)` | 1 |
| `ValidateLinks` | `Contains(str: string, entries: []string): bool` | `(…, "medium.com")` | 1 |
| `ValidateTags` | `Contains(str: string, entries: []string): bool` | `(t, allowedTags)` | 1 |
| `ValidateList` | `Append(err: error)` | `(err)` | 1 |
| `ValidateList` | `Len(): int` | `()` | 1 |
| `ValidateList` | `NewErrComposite(): *ErrComposite` | `()` | 1 |
| `validateAssetID` | `ValidateETHForkAddress(chain: coin.Coin, addr: string): error` | `(chain, addr)` | 1 |
| `validateChainOrAssetInfo` | `GetAssetInfoPath(chain: string): string` | `(chain.Handle, token.Address)` | 1 |
| `validateChainOrAssetInfo` | `GetChainInfoPath(chain: string): string` | `(chain.Handle)` | 1 |
| `validateChainOrAssetInfo` | `GetStatus(): string` | `()` | 1 |
| `validateTokenAddress` | `ValidateETHForkAddress(chain: coin.Coin, addr: string): error` | `(chain, token.Address)` | 1 |
| `ValidateTokenList` | `Append(err: error)` | `(err)` | 3 |
| `ValidateTokenList` | `Len(): int` | `()` | 1 |
| `ValidateTokenList` | `NewErrComposite(): *ErrComposite` | `()` | 1 |
| `validateTokenListPairs` | `Append(err: error)` | `(…)` | 1 |
| `validateTokenListPairs` | `Len(): int` | `()` | 1 |
| `validateTokenListPairs` | `NewErrComposite(): *ErrComposite` | `()` | 1 |
