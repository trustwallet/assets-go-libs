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
	assetType := "BEP2"

	chain, err := types.GetChainFromAssetType(assetType)
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

		log.WithField("path", assetLogoPath).Debug("Missing logo")

		pkg.CreateDirPath(assetLogoPath)
		err = pkg.CreatePNGFromURL(a.AssetImg, assetLogoPath)
		if err != nil {
			return err
		}

		assetInfoPath := asset.GetAssetInfoPath(chain.Handle, a.Asset)

		explorerURL, err := coin.GetCoinExploreURL(chain, a.Asset)
		if err != nil {
			return err
		}

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

		err = pkg.CreateJSONFile(assetInfoPath, &assetInfo)
		if err != nil {
			return err
		}
	}

	return nil
}
