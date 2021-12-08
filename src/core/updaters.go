package core

import (
	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/assets-go-libs/pkg/asset"
	"github.com/trustwallet/assets-go-libs/pkg/validation/info"
	"github.com/trustwallet/assets-go-libs/src/client/binancedex"
	"github.com/trustwallet/assets-go-libs/src/config"
	"github.com/trustwallet/go-primitives/coin"
	"github.com/trustwallet/go-primitives/types"

	log "github.com/sirupsen/logrus"
)

const (
	assetsRows = 1000
	assetsPage = 1
)

func (s *Service) UpdateBinanceTokens() error {
	c := binancedex.NewClient(config.Default.ClientURLs.Binancedex, nil)

	bep2AssetList, err := c.GetBep2Assets(assetsPage, assetsRows)
	if err != nil {
		return err
	}

	err = fetchMissingAssets(bep2AssetList.AssetInfoList)
	if err != nil {
		return err
	}

	return nil
}

func fetchMissingAssets(assets []binancedex.Bep2Asset) error {
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

func createLogo(assetLogoPath string, a binancedex.Bep2Asset) error {
	log.WithField("path", assetLogoPath).Debug("Missing logo")

	pkg.CreateDirPath(assetLogoPath)

	return pkg.CreatePNGFromURL(a.AssetImg, assetLogoPath)
}

func createInfoJSON(chain coin.Coin, a binancedex.Bep2Asset) error {
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
