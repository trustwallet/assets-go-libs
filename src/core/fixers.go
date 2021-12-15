package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/assets-go-libs/pkg/validation/info"
	"github.com/trustwallet/go-primitives/address"
	"github.com/trustwallet/go-primitives/coin"
	"github.com/trustwallet/go-primitives/types"

	log "github.com/sirupsen/logrus"
)

func (s *Service) FixJSON(file *file.AssetFile) error {
	return pkg.FormatJSONFile(file.Info.Path())
}

func (s *Service) FixETHAddressChecksum(file *file.AssetFile) error {
	if !coin.IsEVM(file.Info.Chain().ID) {
		return nil
	}

	assetDir := filepath.Base(file.Info.Path())

	err := validation.ValidateETHForkAddress(file.Info.Chain(), assetDir)
	if err != nil {
		checksum, e := address.EIP55Checksum(assetDir)
		if e != nil {
			return fmt.Errorf("failed to get checksum: %s", e)
		}

		newName := fmt.Sprintf("blockchains/%s/assets/%s", file.Info.Chain().Handle, checksum)

		if e = os.Rename(file.Info.Path(), newName); e != nil {
			return fmt.Errorf("failed to rename dir: %s", e)
		}

		s.fileService.UpdateFile(file, checksum)

		log.WithField("from", assetDir).
			WithField("to", checksum).
			Debug("Renamed asset")
	}

	return nil
}

func (s *Service) FixLogo(file *file.AssetFile) error {
	width, height, err := pkg.GetPNGImageDimensions(file.Info.Path())
	if err != nil {
		return err
	}

	var isLogoTooLarge bool
	if width > validation.MaxW || height > validation.MaxH {
		isLogoTooLarge = true
	}

	if isLogoTooLarge {
		log.WithField("path", file.Info.Path()).Debug("Fixing too large image")

		targetW, targetH := calculateTargetDimension(width, height)

		err = pkg.ResizePNG(file.Info.Path(), targetW, targetH)
		if err != nil {
			return err
		}
	}

	err = validation.ValidateLogoFileSize(file.Info.Path())
	if err != nil { // nolint:staticcheck
		// TODO: Compress images.
	}

	return nil
}

func calculateTargetDimension(width, height int) (targetW, targetH int) {
	widthFloat := float32(width)
	heightFloat := float32(height)

	maxEdge := widthFloat
	if heightFloat > widthFloat {
		maxEdge = heightFloat
	}

	ratio := validation.MaxW / maxEdge

	targetW = int(widthFloat * ratio)
	targetH = int(heightFloat * ratio)

	return targetW, targetH
}

func (s *Service) FixChainInfoJSON(file *file.AssetFile) error {
	chainInfo := info.CoinModel{}

	err := pkg.ReadJSONFile(file.Info.Path(), &chainInfo)
	if err != nil {
		return err
	}

	expectedType := string(types.Coin)
	if chainInfo.Type == nil || *chainInfo.Type != expectedType {
		chainInfo.Type = &expectedType
	}

	return pkg.CreateJSONFile(file.Info.Path(), &chainInfo)
}

func (s *Service) FixAssetInfoJSON(file *file.AssetFile) error {

	return nil
}
