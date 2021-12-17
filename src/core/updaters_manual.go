package core

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/assets-go-libs/pkg/asset"
	"github.com/trustwallet/assets-go-libs/src/config"
	"github.com/trustwallet/go-libs/client/api/backend"
	"github.com/trustwallet/go-primitives/address"
	"github.com/trustwallet/go-primitives/coin"
	"github.com/trustwallet/go-primitives/types"
)

var (
	UniswapForceInclude = []string{"TUSD", "STAKE", "YFI", "BAT", "MANA", "1INCH", "REP", "KP3R", "UNI", "WBTC", "HEX", "CREAM", "SLP", "REN", "XOR", "Link", "sUSD", "HEGIC", "RLC", "DAI", "SUSHI", "FYZ", "DYT", "AAVE", "LEND", "UBT", "DIA", "RSR", "SXP", "OCEAN", "MKR", "USDC", "CEL", "BAL", "BAND", "COMP", "SNX", "OMG", "AMPL", "USDT", "KNC", "ZRX", "AXS", "ENJ", "STMX", "DPX", "FTT", "DPI", "PAX"}
	UniswapForceExclude = []string{"STARL", "UFO"}

	PolygonSwapForceInclude = []string{}
	PolygonSwapForceExclude = []string{}

	PancakeSwapForceInclude = []string{"Cake", "DAI", "ETH", "TWT", "VAI", "USDT", "BLINK", "BTCB", "ALPHA", "INJ", "CTK", "UNI", "XVS", "BUSD", "HARD", "BIFI", "FRONT"}
	PancakeSwapForceExclude = []string{}
)

var (
	UniswapTradingPairsQuery = map[string]string{
		"operationName": "pairs",
		"query": `
			query pairs {
				pairs(first: 800, orderBy: reserveUSD, orderDirection: desc) {
					id reserveUSD trackedReserveETH volumeUSD txCount untrackedVolumeUSD __typename
					token0 {
						id symbol name decimals __typename
					}
					token1 {
						id symbol name decimals __typename
					}
				}
			}
		`,
	}

	PolygonSwap_TradingPairsQuery = map[string]string{
		"operationName": "pairs",
		"query": `
			{
				ethereum(network: matic) {
					dexTrades(date: {is: "$DATE$"}) {
						sellCurrency {address symbol name decimals}
						buyCurrency {address symbol name decimals}
						trade: count
						tradeAmount(in: USD)
					}
				}
			}
		`,
	}

	PancakeSwap_TradingPairsQuery = map[string]string{
		"operationName": "pairs",
		"query": `
			query pairs {
				pairs(first: 300, orderBy: reserveUSD, orderDirection: desc) {
					id reserveUSD volumeUSD txCount __typename
					token0 {
						id symbol name decimals __typename
					}
					token1 {
						id symbol name decimals __typename
					}
				}
			}
		`,
	}

	UniswapTradingPairsUrl     = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
	PolygonSwapTradingPairsUrl = "https://graphql.bitquery.io"
	PancakeSwapradingPairsUrl  = "https://api.bscgraph.org/subgraphs/name/cakeswap"
)

const (
	UniswapMinLiquidity = 2000000
	UniswapMinVol24     = 1000000
	UniswapMinTxCount24 = 480

	PolygonSwapMinVol24     = 500000
	PolygonSwapMinTxCount24 = 288

	PancakeSwapMinLiquidity = 1000000
	PancakeSwapMinVol24     = 500000
	PancakeSwapMinTxCount24 = 288
)

var (
	PrimaryTokensETH = []string{"WETH", "ETH"}
)

func (s *Service) UpdateEthereumTokenlist() error {
	tradingPairs, err := retrieveUniswapPairs(UniswapTradingPairsUrl, UniswapTradingPairsQuery,
		UniswapMinLiquidity, UniswapMinVol24, UniswapMinTxCount24, PrimaryTokensETH)
	if err != nil {
		return err
	}

	pairs := make([][]TokenItem, 0)

	for _, tradingPair := range tradingPairs {
		tokenItem0, err := getTokenInfoFromSubgraphToken(tradingPair.Token0)
		if err != nil {
			return err
		}

		tokenItem1, err := getTokenInfoFromSubgraphToken(tradingPair.Token1)
		if err != nil {
			return err
		}

		if !isTokenPrimary(tradingPair.Token0, PrimaryTokensETH) {
			tokenItem0, tokenItem1 = tokenItem1, tokenItem0
		}

		pairs = append(pairs, []TokenItem{*tokenItem0, *tokenItem1})
	}

	return rebuildTokenList(coin.Coins[coin.ETHEREUM], pairs, UniswapForceExclude)
}

func (s *Service) UpdatePolygonTokenlist() error {
	return nil
}

func (s *Service) UpdateSmartchainTokenlist() error {
	return nil
}

func retrieveUniswapPairs(url string, query map[string]string, minLiquidity, minVol24, minTxCount24 int,
	primaryTokens []string) ([]TradingPair, error) {
	includeList := parseForceList(UniswapForceInclude)

	pairs, err := getTradingPairs(url, query)
	if err != nil {
		return nil, err
	}

	filtered := make([]TradingPair, 0)
	for _, pair := range pairs.Data.Pairs {
		if checkTradingPairOK(pair, minLiquidity, minVol24, minTxCount24, primaryTokens, includeList) {
			filtered = append(filtered, pair)
		}
	}

	return filtered, nil
}

func parseForceList(forceList []string) []ForceListPair {
	result := make([]ForceListPair, 0, len(forceList))

	for _, item := range forceList {
		pair := ForceListPair{}

		tokens := strings.Split(item, "-")
		pair.Token0 = tokens[0]
		if len(tokens) >= 2 {
			pair.Token1 = tokens[1]
		}

		result = append(result, pair)
	}

	return result
}

func getTradingPairs(url string, query map[string]string) (*TradingPairs, error) {
	jsonValue, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	var result TradingPairs
	err = pkg.PostHTTPResponse(url, jsonValue, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func checkTradingPairOK(pair TradingPair, minLiquidity, minVol24, minTxCount24 int, primaryTokens []string,
	forceIncludeList []ForceListPair) bool {
	if pair.ID == "" || pair.ReserveUSD == "" || pair.VolumeUSD == "" || pair.TxCount == "" ||
		pair.Token0 == nil || pair.Token1 == nil {
		return false
	}

	if !(isTokenPrimary(pair.Token0, primaryTokens) || isTokenPrimary(pair.Token1, primaryTokens)) {
		return false
	}

	if isPairMatchedToForceList(getTokenItemFromInfo(pair.Token0), getTokenItemFromInfo(pair.Token1), forceIncludeList) {
		return true
	}

	reserveUSD, err := strconv.Atoi(pair.ReserveUSD)
	if err != nil {
		return false
	}

	volumeUSD, err := strconv.Atoi(pair.VolumeUSD)
	if err != nil {
		return false
	}

	txCount, err := strconv.Atoi(pair.TxCount)
	if err != nil {
		return false
	}

	if reserveUSD < minLiquidity || volumeUSD < minVol24 || txCount < minTxCount24 {
		return false
	}

	return true
}

func getTokenItemFromInfo(tokenInfo *TokenInfo) *TokenItem {
	decimals, err := strconv.Atoi(tokenInfo.Decimals)
	if err != nil {
		return nil
	}

	return &TokenItem{
		Asset:    tokenInfo.ID,
		Address:  tokenInfo.ID,
		Name:     tokenInfo.Name,
		Symbol:   tokenInfo.Symbol,
		Decimals: uint(decimals),
	}
}

func getTokenInfoFromSubgraphToken(token *TokenInfo) (*TokenItem, error) {
	checksum, err := address.EIP55Checksum(token.ID)
	if err != nil {
		return nil, err
	}

	decimals, err := strconv.Atoi(token.Decimals)
	if err != nil {
		return nil, err
	}

	return &TokenItem{
		Asset:    getAssetIDSymbol(checksum, coin.Coins[coin.ETHEREUM].Symbol, coin.ETHEREUM),
		Type:     string(types.ERC20),
		Address:  checksum,
		Name:     token.Name,
		Symbol:   token.Symbol,
		Decimals: uint(decimals),
		LogoURI:  getLogoURI(token.Symbol, coin.Coins[coin.ETHEREUM].Handle, coin.Coins[coin.ETHEREUM].Symbol),
		Pairs:    make([]Pair, 0),
	}, nil
}

func isTokenPrimary(token *TokenInfo, primaryTokens []string) bool {
	if token == nil {
		return false
	}

	for _, primaryToken := range primaryTokens {
		if strings.EqualFold(primaryToken, token.Symbol) {
			return true
		}
	}

	return false
}

func isPairMatchedToForceList(token0, token1 *TokenItem, forceIncludeList []ForceListPair) bool {
	var matched bool

	for _, forcePair := range forceIncludeList {
		if matchPairToForceListEntry(token0, token1, forcePair) {
			matched = true
		}
	}

	return matched
}

func matchPairToForceListEntry(token0, token1 *TokenItem, forceListEntry ForceListPair) bool {
	if forceListEntry.Token1 == "" {
		// entry is token only
		if matchTokenToForceListEntry(token0, forceListEntry.Token0) ||
			(token1 != nil && matchTokenToForceListEntry(token1, forceListEntry.Token0)) {
			return true
		}

		return false
	}

	if token1 == nil {
		return false
	}

	if matchTokenToForceListEntry(token0, forceListEntry.Token0) && matchTokenToForceListEntry(token0, forceListEntry.Token1) {
		return true
	}

	if matchTokenToForceListEntry(token0, forceListEntry.Token1) && matchTokenToForceListEntry(token1, forceListEntry.Token0) {
		return true
	}
	return false
}

func matchTokenToForceListEntry(token *TokenItem, forceListEntry string) bool {
	if strings.EqualFold(forceListEntry, token.Symbol) ||
		strings.EqualFold(forceListEntry, token.Asset) ||
		strings.EqualFold(forceListEntry, token.Name) {
		return true
	}

	return false
}

func matchPairToForceList(token0, token1 *TokenItem, forceList []ForceListPair) bool {
	var matched bool
	for _, forcePair := range forceList {
		if matchPairToForceListEntry(token0, token1, forcePair) {
			matched = true
		}
	}

	return matched
}

func rebuildTokenList(chain coin.Coin, pairs [][]TokenItem, forceExcludeList []string) error {
	if pairs == nil || len(pairs) < 5 {
		return nil
	}

	excludeList := parseForceList(forceExcludeList)

	pairs2 := make([][]TokenItem, 0)

	for _, pair := range pairs {
		if !checkTokenExists(chain.Handle, pair[0].Address) {
			return fmt.Errorf("pair with unsupported 1st coin: %s-%s", pair[0].Symbol, pair[1].Symbol)
		}

		if !checkTokenExists(chain.Handle, pair[1].Address) {
			return fmt.Errorf("pair with unsupported 2nd coin: %s-%s", pair[0].Symbol, pair[1].Symbol)
		}

		if matchPairToForceList(&pair[0], &pair[1], excludeList) {
			return fmt.Errorf("pair excluded due to FORCE EXCLUDE: %s-%s", pair[0].Symbol, pair[1].Symbol)
		}

		pairs2 = append(pairs2, pair)
	}

	filteredCount := len(pairs) - len(pairs2)
	log.Debugf("%d unsupported tokens filtered out, %d pairs", filteredCount, len(pairs2))

	tokenListPath := fmt.Sprintf("blockchains/%s/tokenlist.json", chain.Handle)

	var oldTokenList TokenList
	err := pkg.ReadJSONFile(tokenListPath, &oldTokenList)
	if err != nil {
		return nil
	}

	removeAllPairs(&oldTokenList)

	for _, pair := range pairs2 {
		err = addPairIfNeeded(&pair[0], &pair[1], &oldTokenList)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkTokenExists(chain, tokenID string) bool {
	logoPath := asset.GetAssetLogoPath(chain, tokenID)

	return pkg.FileExists(logoPath)
}

func removeAllPairs(list *TokenList) {
	for _, token := range list.Tokens {
		token.Pairs = make([]Pair, 0)
	}
}

func addPairIfNeeded(token0, token1 *TokenItem, list *TokenList) error {
	err := addTokenIfNeeded(token0, list)
	if err != nil {
		return err
	}

	err = addTokenIfNeeded(token1, list)
	if err != nil {
		return err
	}

	addPairToToken(token1, token0, list)

	return nil
}

func addTokenIfNeeded(token *TokenItem, list *TokenList) error {
	for _, token := range list.Tokens {
		if strings.EqualFold(token.Address, token.Address) {
			return nil
		}
	}

	err := updateTokenInfo(token)
	if err != nil {
		return err
	}

	list.Tokens = append(list.Tokens, *token)

	return nil
}

func updateTokenInfo(token *TokenItem) error {
	backendClient := backend.InitClient(config.Default.ClientURLs.BackendAPI, nil)
	assetInfo, err := backendClient.GetAssetInfo(token.Asset)
	if err != nil {
		return err
	}

	if token.Name != assetInfo.Name {
		token.Name = assetInfo.Name
	}

	if token.Symbol != assetInfo.Symbol {
		token.Symbol = assetInfo.Symbol
	}

	if token.Decimals != uint(assetInfo.Decimals) {
		token.Decimals = uint(assetInfo.Decimals)
	}

	return nil
}

func addPairToToken(pairToken, token *TokenItem, list *TokenList) {
	var tokenInList *TokenItem
	for _, t := range list.Tokens {
		if t.Address == token.Address {
			tokenInList = &t
			break
		}
	}

	if tokenInList != nil {
		tokenInList.Pairs = make([]Pair, 0)
	}

	for _, pair := range tokenInList.Pairs {
		if pair.Base == pairToken.Asset {
			return
		}
	}

	tokenInList.Pairs = append(tokenInList.Pairs, Pair{
		Base: pairToken.Asset,
	})
}
