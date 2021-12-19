package processor

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
	UniswapForceInclude = []string{
		"TUSD", "STAKE", "YFI", "BAT", "MANA", "1INCH", "REP", "KP3R", "UNI", "WBTC", "HEX", "CREAM", "SLP",
		"REN", "XOR", "Link", "sUSD", "HEGIC", "RLC", "DAI", "SUSHI", "FYZ", "DYT", "AAVE", "LEND", "UBT",
		"DIA", "RSR", "SXP", "OCEAN", "MKR", "USDC", "CEL", "BAL", "BAND", "COMP", "SNX", "OMG", "AMPL",
		"USDT", "KNC", "ZRX", "AXS", "ENJ", "STMX", "DPX", "FTT", "DPI", "PAX",
	}
	UniswapForceExclude = []string{"STARL", "UFO"}

	PolygonSwapForceInclude = []string{}
	PolygonSwapForceExclude = []string{}

	PancakeSwapForceInclude = []string{
		"Cake", "DAI", "ETH", "TWT", "VAI", "USDT", "BLINK", "BTCB", "ALPHA", "INJ", "CTK", "UNI", "XVS",
		"BUSD", "HARD", "BIFI", "FRONT",
	}
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
	log.WithFields(log.Fields{
		"limit_liquidity": UniswapMinLiquidity,
		"volume":          UniswapMinVol24,
		"tx_count":        UniswapMinTxCount24,
	}).Debug("Retrieving pairs from Uniswap")

	tradingPairs, err := retrieveUniswapPairs(UniswapTradingPairsUrl, UniswapTradingPairsQuery,
		UniswapMinLiquidity, UniswapMinVol24, UniswapMinTxCount24, UniswapForceInclude, PrimaryTokensETH)
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
	forceIncludeList []string, primaryTokens []string) ([]TradingPair, error) {
	includeList := parseForceList(forceIncludeList)

	pairs, err := fetchTradingPairs(url, query)
	if err != nil {
		return nil, err
	}

	filtered := make([]TradingPair, 0)

	for _, pair := range pairs.Data.Pairs {
		ok, err := checkTradingPairOK(pair, minLiquidity, minVol24, minTxCount24, primaryTokens, includeList)
		if err != nil {
			log.Debug(err)
		}
		if ok {
			filtered = append(filtered, pair)
		}
	}

	return filtered, nil
}

func parseForceList(forceList []string) []ForceListPair {
	result := make([]ForceListPair, 0, len(forceList))

	for _, item := range forceList {
		tokens := strings.Split(item, "-")
		pair := ForceListPair{
			Token0: tokens[0],
		}

		if len(tokens) >= 2 {
			pair.Token1 = tokens[1]
		}

		result = append(result, pair)
	}

	return result
}

func fetchTradingPairs(url string, query map[string]string) (*TradingPairs, error) {
	jsonValue, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	log.WithField("url", url).Debug("Retrieving trading pair infos")

	var result TradingPairs
	err = pkg.PostHTTPResponse(url, jsonValue, &result)
	if err != nil {
		return nil, err
	}

	log.Debugf("Retrieved %d trading pair infos", len(result.Data.Pairs))

	return &result, nil
}

func checkTradingPairOK(pair TradingPair, minLiquidity, minVol24, minTxCount24 int, primaryTokens []string,
	forceIncludeList []ForceListPair) (bool, error) {
	if pair.ID == "" || pair.ReserveUSD == "" || pair.VolumeUSD == "" || pair.TxCount == "" ||
		pair.Token0 == nil || pair.Token1 == nil {

		return false, nil
	}

	if !(isTokenPrimary(pair.Token0, primaryTokens) || isTokenPrimary(pair.Token1, primaryTokens)) {
		log.Debugf("pair with no primary coin: %s -- %s", pair.Token0.Symbol, pair.Token1.Symbol)
		return false, nil
	}

	if isPairMatchedToForceList(getTokenItemFromInfo(pair.Token0), getTokenItemFromInfo(pair.Token1), forceIncludeList) {
		log.Debugf("pair included due to FORCE INCLUDE: %s -- %s", pair.Token0.Symbol, pair.Token1.Symbol)
		return true, nil
	}

	reserveUSD, err := strconv.ParseFloat(pair.ReserveUSD, 64)
	if err != nil {
		return false, err
	}
	if int(reserveUSD) < minLiquidity {
		log.Debugf("pair with low liquidity: %s -- %s", pair.Token0.Symbol, pair.Token1.Symbol)
		return false, nil
	}

	volumeUSD, err := strconv.ParseFloat(pair.VolumeUSD, 64)
	if err != nil {
		return false, err
	}
	if int(volumeUSD) < minVol24 {
		log.Debugf("pair with low volume: %s -- %s", pair.Token0.Symbol, pair.Token1.Symbol)
		return false, nil
	}

	txCount, err := strconv.ParseFloat(pair.TxCount, 64)
	if err != nil {
		return false, err
	}
	if int(txCount) < minTxCount24 {
		log.Debugf("pair with low tx count: %s -- %s", pair.Token0.Symbol, pair.Token1.Symbol)
		return false, nil
	}

	return true, nil
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
		if matchTokenToForceListEntry(token0, forceListEntry.Token0) ||
			(token1 != nil && matchTokenToForceListEntry(token1, forceListEntry.Token0)) {

			return true
		}

		return false
	}

	if token1 == nil {
		return false
	}

	if matchTokenToForceListEntry(token0, forceListEntry.Token0) &&
		matchTokenToForceListEntry(token0, forceListEntry.Token1) {

		return true
	}

	if matchTokenToForceListEntry(token0, forceListEntry.Token1) &&
		matchTokenToForceListEntry(token1, forceListEntry.Token0) {

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
			log.Debugf("pair with unsupported 1st coin: %s-%s", pair[0].Symbol, pair[1].Symbol)
			continue
		}

		if !checkTokenExists(chain.Handle, pair[1].Address) {
			log.Debugf("pair with unsupported 2nd coin: %s-%s", pair[0].Symbol, pair[1].Symbol)
			continue
		}

		if matchPairToForceList(&pair[0], &pair[1], excludeList) {
			log.Debugf("pair excluded due to FORCE EXCLUDE: %s-%s", pair[0].Symbol, pair[1].Symbol)
			continue
		}

		pairs2 = append(pairs2, pair)
	}

	filteredCount := len(pairs) - len(pairs2)

	log.Debugf("%d unsupported tokens filtered out, %d pairs", filteredCount, len(pairs2))

	tokenListPath := fmt.Sprintf("blockchains/%s/tokenlist.json", chain.Handle)

	var list TokenList
	err := pkg.ReadJSONFile(tokenListPath, &list)
	if err != nil {
		return nil
	}

	removeAllPairs(&list)

	for _, pair := range pairs2 {
		err = addPairIfNeeded(&pair[0], &pair[1], &list)
		if err != nil {
			return err
		}
	}

	log.Debugf("Tokenlist updated: %d tokens", len(list.Tokens))

	var totalPairs int

	for _, item := range list.Tokens {
		totalPairs += len(item.Pairs)
	}

	log.Debugf("Tokenlist: list with %d tokens and %d pairs written to %s.",
		len(list.Tokens), totalPairs, tokenListPath)

	return pkg.CreateJSONFile(tokenListPath, list)
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
	for _, t := range list.Tokens {
		if strings.EqualFold(t.Address, token.Address) {
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
	var tokenInListIndex = -1

	for i, t := range list.Tokens {
		if t.Address == token.Address {
			tokenInListIndex = i
			break
		}
	}

	if tokenInListIndex == -1 {
		return
	}

	if list.Tokens[tokenInListIndex].Pairs == nil {
		list.Tokens[tokenInListIndex].Pairs = make([]Pair, 0)
	}

	for _, pair := range list.Tokens[tokenInListIndex].Pairs {
		if pair.Base == pairToken.Asset {
			return
		}
	}

	list.Tokens[tokenInListIndex].Pairs = append(list.Tokens[tokenInListIndex].Pairs, Pair{Base: pairToken.Asset})
}
