package validator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/assets-go-libs/pkg/validation/info"
	"github.com/trustwallet/assets-go-libs/pkg/validation/list"
	"github.com/trustwallet/assets-go-libs/src/config"
	"github.com/trustwallet/go-primitives/coin"
)

type Service struct {
	binanceAssetsSymbols []string
	fileProvider         *file.FileProvider
}

func NewService(fileProvider *file.FileProvider) (*Service, error) {
	var binanceAssetsSymbols []string

	// binancedexClient := binancedex.InitBinanceDexClient(config.Default.ClientURLs.Binancedex, nil)

	// bep8, err := binancedexClient.GetBep8Assets(1000)
	// if err != nil {
	// 	return nil, err
	// }
	// <-time.After(time.Second * 1) // TODO: binance dex blocked too often requests (request timout)
	// bep2, err := binancedexClient.GetBep2Assets(1000)
	// if err != nil {
	// 	return nil, err
	// }

	// for _, a := range bep8 {
	// 	binanceAssetsSymbols = append(binanceAssetsSymbols, a.Symbol)
	// }
	// for _, a := range bep2 {
	// 	binanceAssetsSymbols = append(binanceAssetsSymbols, a.Symbol)
	// }

	return &Service{
		binanceAssetsSymbols: binanceAssetsSymbols,
		fileProvider:         fileProvider,
	}, nil
}

func (s *Service) GetValidatorForFilesAndFolders(f *file.AssetFile) *Validator {
	fileType := f.Info.Type()
	switch fileType {
	case file.TypeRootFolder:
		return &Validator{
			ValidationName: "Root folder contains only allowed files",
			Run:            s.ValidateRootFolder,
			FileType:       fileType,
		}
	case file.TypeChainFolder:
		return &Validator{
			ValidationName: "Chain folders are lowercase and contains only allowed files",
			Run:            s.ValidateChainFolder,
			FileType:       fileType,
		}
	case file.TypeChainLogoFile, file.TypeAssetLogoFile, file.TypeValidatorsLogoFile, file.TypeDappsLogoFile:
		return &Validator{
			ValidationName: "Logos (size, dimension)",
			Run:            s.ValidateImage,
			FileType:       fileType,
		}
	case file.TypeAssetFolder:
		return &Validator{
			ValidationName: "Each asset folder has valid asset address and contains logo/info",
			Run:            s.ValidateAssetFolder,
			FileType:       fileType,
		}
	case file.TypeDappsFolder:
		return &Validator{
			ValidationName: "Dapps folder (allowed only png files, lowercase)",
			Run:            s.ValidateDappsFolder,
			FileType:       fileType,
		}
		// case file.TypeAssetInfoFile:
		// 	return &Validator{
		// 		ValidationName: "Asset info (is valid json, fields)",
		// 		Run:            s.ValidateAssetInfoFile,
		// 		FileType:       fileType,
		// 	}
		// case file.TypeChainInfoFile:
		// 	return &Validator{
		// 		ValidationName: "Chain Info (is valid json, fields)",
		// 		Run:            s.ValidateChainInfoFile,
		// 		FileType:       fileType,
		// 	}
		// case file.TypeValidatorsListFile:
		// 	if !isStackingChain(f.Info.Chain()) {
		// 		return nil
		// 	}

		// 	return &Validator{
		// 		ValidationName: "Validators list file",
		// 		Run:            s.ValidateValidatorsListFile,
		// 		FileType:       fileType,
		// 	}
		// case file.TypeTokenListFile:
		// 	return &Validator{
		// 		ValidationName: "Token list (if assets from list present in chain)",
		// 		Run:            s.ValidateTokenListFile,
		// 		FileType:       fileType,
		// 	}
		// case file.TypeAssetsFolder:
		// 	return &Validator{
		// 		ValidationName: "Chain assets folder (chain specific validation also here)",
		// 		Run:            s.ValidateChainAssetsFolder,
		// 		FileType:       fileType,
		// 	}
		// case file.TypeChainInfoFolder:
		// 	return &Validator{
		// 		ValidationName: "Chain Info Folder (has files)",
		// 		Run:            s.ValidateInfoFolder,
		// 		FileType:       fileType,
		// 	}
		// case file.TypeValidatorsFolder:
		// 	return nil
		// case file.TypeValidatorsAssetsFolder:
		// 	return nil
		// case file.TypeValidatorsAssetFolder:
		// 	return &Validator{
		// 		ValidationName: "Validators asset folder (has logo, valid asset address)",
		// 		Run:            s.ValidateValidatorsAssetFolder,
		// 		FileType:       fileType,
		// 	}
	}

	return nil
}

func (s *Service) ValidateInfoFolder(file *file.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	err = validation.ValidateHasFiles(dirFiles, config.Default.ValidatorsSettings.ChainInfoFolder.HasFiles)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateValidatorsAssetFolder(file *file.AssetFile) error {
	assetInfo := file.Info
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	compErr := validation.NewErrComposite()
	err = validation.ValidateValidatorsAddress(assetInfo.Chain(), assetInfo.Asset())
	if err != nil {
		compErr.Append(err)
	}

	err = validation.ValidateHasFiles(dirFiles, []string{"logo.png"})
	if err != nil {
		compErr.Append(err)
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func (s *Service) ValidateChainAssetsFolder(file *file.AssetFile) error {
	if file.Info.Chain().ID == coin.BINANCE {
		dirFiles, err := file.ReadDir(0)
		if err != nil {
			return err
		}

		err = validation.ValidateAllowedFiles(dirFiles, s.binanceAssetsSymbols)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) ValidateTokenListFile(file *file.AssetFile) error {
	buf := bytes.NewBuffer(nil)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return err
	}

	err = validation.ValidateJson(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateValidatorsListFile(file *file.AssetFile) error {
	buf := bytes.NewBuffer(nil)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return err
	}

	err = validation.ValidateJson(buf.Bytes())
	if err != nil {
		return err
	}

	var model []list.Model
	err = json.Unmarshal(buf.Bytes(), &model)
	if err != nil {
		return err
	}

	err = list.ValidateList(model)
	if err != nil {
		return err
	}

	var listIDs []string
	for _, listItem := range model {
		listIDs = append(listIDs, *listItem.ID)
	}

	assetsPath := getValidatorsAssetsPath(file.Info.Chain())
	assetFolder, err := s.fileProvider.GetAssetFile(assetsPath)
	if err != nil {
		return err
	}

	dirAssetFolderFiles, err := assetFolder.ReadDir(0)
	if err != nil {
		return err
	}

	err = validation.ValidateHasFiles(dirAssetFolderFiles, listIDs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateChainInfoFile(file *file.AssetFile) error {
	fileInfo := file.Info
	buf := bytes.NewBuffer(nil)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return err
	}

	err = validation.ValidateJson(buf.Bytes())
	if err != nil {
		return err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("%w, failed to seek reader", validation.ErrInvalidJson)
	}

	var payload info.CoinModel
	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		return fmt.Errorf("%w, failed to decode", err)
	}

	var tags []string
	for _, t := range config.Default.ValidatorsSettings.CoinInfoFile.Tags {
		tags = append(tags, t.ID)
	}

	err = info.ValidateCoin(payload, fileInfo.Chain(), fileInfo.Asset(), tags)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateAssetInfoFile(file *file.AssetFile) error {
	fileInfo := file.Info
	buf := bytes.NewBuffer(nil)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return err
	}

	err = validation.ValidateJson(buf.Bytes())
	if err != nil {
		return err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("%w, failed to seek reader", validation.ErrInvalidJson)
	}

	var payload info.AssetModel
	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		return fmt.Errorf("%w, failed to decode", err)
	}

	err = info.ValidateAsset(payload, fileInfo.Chain(), fileInfo.Asset())
	if err != nil {
		return err
	}

	return nil
}

// TODO: figure out how to do it other way...
func getValidatorsAssetsPath(chain coin.Coin) string {
	return fmt.Sprintf("blockchains/%s/validators/assets", chain.Handle)
}

// func isStackingChain(c coin.Coin) bool {
// 	for _, stackingChain := range config.StackingChains {
// 		if c.ID == stackingChain.ID {
// 			return true
// 		}
// 	}

// 	return false
// }
