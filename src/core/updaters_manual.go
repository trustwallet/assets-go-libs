package core

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/trustwallet/assets-go-libs/pkg"
)

var (
	// Force include list of token symbols, or symbol pairs (e.g. ["BAT", "YFI-WETH"]).
	UniswapForceInclude = []string{"TUSD", "STAKE", "YFI", "BAT", "MANA", "1INCH", "REP", "KP3R", "UNI", "WBTC", "HEX", "CREAM", "SLP", "REN", "XOR", "Link", "sUSD", "HEGIC", "RLC", "DAI", "SUSHI", "FYZ", "DYT", "AAVE", "LEND", "UBT", "DIA", "RSR", "SXP", "OCEAN", "MKR", "USDC", "CEL", "BAL", "BAND", "COMP", "SNX", "OMG", "AMPL", "USDT", "KNC", "ZRX", "AXS", "ENJ", "STMX", "DPX", "FTT", "DPI", "PAX"}
	// Force include list of token symbols, or symbol pairs (e.g. ["Cake", "DAI-WBNB"]).
	PolygonSwapForceInclude = []string{}
	// Force include list of token symbols, or symbol pairs (e.g. ["Cake", "DAI-WBNB"]).
	PancakeSwapForceInclude = []string{"Cake", "DAI", "ETH", "TWT", "VAI", "USDT", "BLINK", "BTCB", "ALPHA", "INJ", "CTK", "UNI", "XVS", "BUSD", "HARD", "BIFI", "FRONT"}
)

const (
	UniswapTradingPairsQuery = `
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
	`

	PolygonSwap_TradingPairsQuery = `
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
`

	PancakeSwap_TradingPairsQuery = `
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
	`

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

	return nil
}

func (s *Service) UpdatePolygonTokenlist() error {
	return nil
}

func (s *Service) UpdateSmartchainTokenlist() error {
	return nil
}

func retrieveUniswapPairs(url, query string, minLiquidity, minVol24, minTxCount24 int,
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
		pair.Token1 = tokens[0]
		if len(tokens) >= 2 {
			pair.Token2 = tokens[1]
		}

		result = append(result, pair)
	}

	return result
}

func getTradingPairs(url, query string) (*TradingPairs, error) {
	postData := fmt.Sprintf("{\"operationName\":\"pairs\", \"variables\":{}, \"query\":\" %s\"}", query)
	jsonValue, err := json.Marshal(postData)
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

	return true
}
