package core

import (
	"fmt"

	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/assets-go-libs/pkg/asset"
	"github.com/trustwallet/assets-go-libs/pkg/validation/info"
	"github.com/trustwallet/assets-go-libs/src/client/binance/dex"
	"github.com/trustwallet/assets-go-libs/src/client/binance/explorer"
	"github.com/trustwallet/assets-go-libs/src/config"
	assetlib "github.com/trustwallet/go-primitives/asset"
	"github.com/trustwallet/go-primitives/coin"
	"github.com/trustwallet/go-primitives/numbers"
	"github.com/trustwallet/go-primitives/types"
)

const (
	assetsPage       = 1
	assetsRows       = 1000
	marketPairsLimit = 1000
	tokensListLimit  = 10000
)

func (s *Service) UpdateBinanceTokens() error {
	explorerClient := explorer.NewClient(config.Default.ClientURLs.Binance.Explorer, nil)

	bep2AssetList, err := explorerClient.GetBep2Assets(assetsPage, assetsRows)
	if err != nil {
		return err
	}

	dexClient := dex.NewClient(config.Default.ClientURLs.Binance.Dex, nil)

	marketPairs, err := dexClient.GetMarketPairs(marketPairsLimit)
	if err != nil {
		return err
	}

	tokensList, err := dexClient.GetTokensList(tokensListLimit)
	if err != nil {
		return err
	}

	err = fetchMissingAssets(bep2AssetList.AssetInfoList)
	if err != nil {
		return err
	}

	return createTokenListJSON(marketPairs, tokensList)
}

func fetchMissingAssets(assets []explorer.Bep2Asset) error {
	chain, err := types.GetChainFromAssetType(string(types.BEP2))
	if err != nil {
		return err
	}

	for _, a := range assets {
		if a.AssetImg == "" || a.Decimals == 0 {
			continue
		}

		assetLogoPath := asset.GetAssetLogoPath(chain.Handle, a.Asset)
		if pkg.FileExists(assetLogoPath) {
			continue
		}

		if err = createLogo(assetLogoPath, a); err != nil {
			return err
		}

		if err = createInfoJSON(chain, a); err != nil {
			return err
		}
	}

	return nil
}

func createLogo(assetLogoPath string, a explorer.Bep2Asset) error {
	pkg.CreateDirPath(assetLogoPath)

	return pkg.CreatePNGFromURL(a.AssetImg, assetLogoPath)
}

func createInfoJSON(chain coin.Coin, a explorer.Bep2Asset) error {
	explorerURL, err := coin.GetCoinExploreURL(chain, a.Asset)
	if err != nil {
		return err
	}

	assetType := "BEP2"
	website := ""
	description := "-"
	status := "active"

	assetInfo := info.AssetModel{
		Name:        &a.Name,
		Type:        &assetType,
		Symbol:      &a.MappedAsset,
		Decimals:    &a.Decimals,
		Website:     &website,
		Description: &description,
		Explorer:    &explorerURL,
		Status:      &status,
		ID:          &a.Asset,
	}

	assetInfoPath := asset.GetAssetInfoPath(chain.Handle, a.Asset)

	return pkg.CreateJSONFile(assetInfoPath, &assetInfo)
}

func createTokenListJSON(marketPairs []dex.MarketPair, tokenList []dex.Token) error {
	tokens, err := generateTokenList(marketPairs, tokenList)
	if err != nil {
		return nil
	}

	return nil
}

func generateTokenList(marketPairs []dex.MarketPair, tokenList []dex.Token) ([]TokenItem, error) {
	if len(marketPairs) < 5 {
		return nil, fmt.Errorf("no markets info is returned from Binance DEX: %d", len(marketPairs))
	}

	if len(tokenList) < 5 {
		return nil, fmt.Errorf("no tokens info is returned from Binance DEX: %d", len(tokenList))
	}

	pairsMap := make(map[string][]Pair)
	pairsList := make(map[string]struct{})
	tokensMap := make(map[string]dex.Token)

	for _, token := range tokenList {
		tokensMap[token.Symbol] = token
	}

	for _, marketPair := range marketPairs {
		key := marketPair.QuoteAssetSymbol

		if val, exists := pairsMap[key]; exists {
			val = append(val, getPair(marketPair))
			pairsMap[key] = val
		} else {
			pairsMap[key] = []Pair{getPair(marketPair)}
		}

		pairsList[marketPair.BaseAssetSymbol] = struct{}{}
		pairsList[marketPair.QuoteAssetSymbol] = struct{}{}
	}

	var tokenItems []TokenItem

	for pair := range pairsList {
		token := tokensMap[pair]

		tokenItems = append(tokenItems, TokenItem{
			Asset:    getAssetIDSymbol(token.Symbol, coin.Coins[coin.BINANCE].Symbol, coin.BINANCE),
			Type:     getTokenType(token.Symbol, coin.Coins[coin.BINANCE].Symbol, string(types.BEP2)),
			Address:  token.Symbol,
			Name:     token.Name,
			Symbol:   token.OriginalSymbol,
			Decimals: coin.Coins[coin.BINANCE].Decimals,
			LogoURI:  getLogoURI(token.Symbol, coin.Coins[coin.BINANCE].Handle, coin.Coins[coin.BINANCE].Symbol),
			Pairs:    pairsMap[token.Symbol],
		})
	}

	return tokenItems, nil
}

func getPair(marketPair dex.MarketPair) Pair {
	return Pair{
		Base:     getAssetIDSymbol(marketPair.BaseAssetSymbol, coin.Coins[coin.BINANCE].Symbol, coin.BINANCE),
		LotSize:  numbers.ToSatoshi(marketPair.LotSize),
		TickSize: numbers.ToSatoshi(marketPair.TickSize),
	}
}

func getAssetIDSymbol(tokenID string, nativeCoinID string, coinType uint) string {
	if tokenID == nativeCoinID {
		return assetlib.BuildID(coinType, "")
	}

	return assetlib.BuildID(coinType, tokenID)
}

func getTokenType(symbol string, nativeCoinSymbol string, tokenType string) string {
	if symbol == nativeCoinSymbol {
		return "coin"
	}

	return tokenType
}

func getLogoURI(id, githubChainFolder, nativeCoinSymbol string) string {
	assetsURI := "https://assets.trustwalletapp.com"

	if id == nativeCoinSymbol {
		return fmt.Sprintf("%s/blockchains/%s/info/logo.png", assetsURI, githubChainFolder)
	}

	return fmt.Sprintf("%s/blockchains/%s/assets/%s/logo.png", assetsURI, githubChainFolder, id)
}
