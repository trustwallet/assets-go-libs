package validators

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/trustwallet/assets-go-libs/internal/config"
	"github.com/trustwallet/assets-go-libs/pkg/assetfs"
	"github.com/trustwallet/assets-go-libs/pkg/assetfs/validation"
	"github.com/trustwallet/assets-go-libs/pkg/assetfs/validation/info"
	"github.com/trustwallet/assets-go-libs/pkg/assetfs/validation/list"
	"github.com/trustwallet/go-primitives/coin"
)

type Service struct {
	imgConf              config.ImageFile
	coinInfoConf         config.CoinInfoFile
	chainFolderConf      config.ChainFolder
	rootFolderConf       config.RootFolder
	assetFolderConf      config.AssetFolder
	infoFolderConf       config.ChainInfoFolder
	dappsFolderConfig    config.DaapsFolder
	binanceAssetsSymbols []string

	fileProvider *assetfs.FileProvider
}

func NewService(
	settings config.ValidatorsSettings,
	binanceAssetsSymbols []string,
	fileProvider *assetfs.FileProvider,
) *Service {
	return &Service{
		imgConf:           settings.ImageFile,
		coinInfoConf:      settings.CoinInfoFile,
		chainFolderConf:   settings.ChainFolder,
		rootFolderConf:    settings.RootFolder,
		assetFolderConf:   settings.AssetFolder,
		infoFolderConf:    settings.ChainInfoFolder,
		dappsFolderConfig: settings.DaapsFolder,

		binanceAssetsSymbols: binanceAssetsSymbols,

		fileProvider: fileProvider,
	}
}

func (s *Service) GetValidatorForFile(file *assetfs.AssetFile) *Validator {
	fileType := file.Info.Type()
	switch fileType {
	case assetfs.TypeAssetInfoFile:
	//	return &Validator{
	//		ValidationName: "Asset info (is valid json, fields)",
	//		Run:            s.ValidateAssetInfoFile,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeAssetLogoFile, assetfs.TypeChainLogoFile, assetfs.TypeValidatorsLogoFile, assetfs.TypeDappsLogoFile:
	//	return &Validator{
	//		ValidationName: "Logos (size, dimension)",
	//		Run:            s.ValidateImage,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeChainInfoFile:
	//	return &Validator{
	//		ValidationName: "Chain Info (is valid json, fields)",
	//		Run:            s.ValidateChainInfoFile,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeValidatorsListFile:
	//	if !isStackingChain(file.Info.Chain()) {
	//		return nil
	//	}
	//
	//	return &Validator{
	//		ValidationName: "Validators list file",
	//		Run:            s.ValidateValidatorsListFile,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeTokenListFile:
	//	return &Validator{
	//		ValidationName: "Token list (if assets from list present in chain)",
	//		Run:            s.ValidateTokenListFile,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeAssetFolder:
	//	return &Validator{
	//		ValidationName: "Each asset folder (valid asset address, contains logo/info)",
	//		Run:            s.ValidateAssetFolder,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeAssetsFolder:
	//	return &Validator{
	//		ValidationName: "Chain assets folder (chain specific validation also here)",
	//		Run:            s.ValidateChainAssetsFolder,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeChainFolder:
	//	return &Validator{
	//		ValidationName: "Chains folder (is files are lower cased, chain specific validations)",
	//		Run:            s.ValidateChainFolder,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeChainsFolder:
	//	return &Validator{
	//		ValidationName: "Each chain folders (lower case, contains files)",
	//		Run:            s.ValidateChainsFolder,
	//		FileType:       fileType,
	//	}
	//case assetfs.TypeDaapsFolder:
	//	return &Validator{
	//		ValidationName: "Daaps folder (allowed only png files, lowercase)",
	//		Run:            s.ValidateDaapsFolder,
	//		FileType:       fileType,
	//	}
	case assetfs.TypeRootFolder:
		return &Validator{
			ValidationName: "Root folder (contains only allowed files)",
			Run:            s.ValidateRootFolder,
			FileType:       fileType,
		}
		//case assetfs.TypeChainInfoFolder:
		//	return &Validator{
		//		ValidationName: "Chain Info Folder (has files)",
		//		Run:            s.ValidateInfoFolder,
		//		FileType:       fileType,
		//	}
		//case assetfs.TypeValidatorsFolder:
		//	return nil
		//case assetfs.TypeValidatorsAssetsFolder:
		//	return nil
		//case assetfs.TypeValidatorsAssetFolder:
		//	return &Validator{
		//		ValidationName: "Validators asset folder (has logo, valid asset address)",
		//		Run:            s.ValidateValidatorsAssetFolder,
		//		FileType:       fileType,
		//	}
	}

	return nil
}

func (s *Service) ValidateValidatorsAssetFolder(file *assetfs.AssetFile) error {
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

func (s *Service) ValidateDaapsFolder(file *assetfs.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	var compErr = validation.NewErrComposite()
	for _, dirFile := range dirFiles {
		err = validation.ValidateExtension(dirFile.Name(), s.dappsFolderConfig.Ext)
		if err != nil {
			compErr.Append(err)
		}

		err = validation.ValidateLowercase(dirFile.Name())
		if err != nil {
			compErr.Append(err)
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func (s *Service) ValidateAssetFolder(file *assetfs.AssetFile) error {
	assetInfo := file.Info
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	var compErr = validation.NewErrComposite()
	err = validation.ValidateAllowedFiles(dirFiles, s.assetFolderConf.AllowedFiles)
	if err != nil {
		compErr.Append(err)
	}

	err = validation.ValidateAssetAddress(assetInfo.Chain(), assetInfo.Asset())
	if err != nil {
		compErr.Append(err)
	}

	errInfo := validation.ValidateHasFiles(dirFiles, []string{"info.json"})
	errLogo := validation.ValidateHasFiles(dirFiles, []string{"logo.png"})
	if errLogo != nil || errInfo != nil {
		infoFile, err := s.fileProvider.GetAssetFile(fmt.Sprintf(fmt.Sprintf("%s/info.json", assetInfo.Path())))
		if err != nil {
			return err
		}

		_, err = infoFile.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

		b, err := io.ReadAll(infoFile)
		if err != nil {
			return err
		}

		var infoJson info.AssetModel
		err = json.Unmarshal(b, &infoJson)
		if err != nil {
			return err
		}

		if infoJson.GetStatus() != "spam" && infoJson.GetStatus() != "abandoned" {
			compErr.Append(fmt.Errorf(
				"%w logo.png for non-spam assest",
				validation.ErrMissingFile,
			))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func (s *Service) ValidateInfoFolder(file *assetfs.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	err = validation.ValidateHasFiles(dirFiles, s.infoFolderConf.HasFiles)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateRootFolder(file *assetfs.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	err = validation.ValidateAllowedFiles(dirFiles, s.rootFolderConf.AllowedFiles)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateChainsFolder(file *assetfs.AssetFile) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	err = validation.ValidateLowercase(fileInfo.Name())
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateChainFolder(file *assetfs.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}
	err = validation.ValidateAllowedFiles(dirFiles, s.chainFolderConf.AllowedFiles)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateChainAssetsFolder(file *assetfs.AssetFile) error {
	assetInfo := file.Info
	if assetInfo.Chain().ID == coin.BINANCE {
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

func (s *Service) ValidateTokenListFile(file *assetfs.AssetFile) error {
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

func (s *Service) ValidateValidatorsListFile(file *assetfs.AssetFile) error {
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

func (s *Service) ValidateChainInfoFile(file *assetfs.AssetFile) error {
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
	for _, t := range s.coinInfoConf.Tags {
		tags = append(tags, t.ID)
	}

	err = info.ValidateCoin(payload, fileInfo.Chain(), fileInfo.Asset(), tags)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateAssetInfoFile(file *assetfs.AssetFile) error {
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

func (s *Service) ValidateImage(file *assetfs.AssetFile) error {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = validation.ValidateSize(b, s.imgConf.Size)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	err = validation.ValidateImageDimension(file, s.imgConf.MaxW, s.imgConf.MaxH, s.imgConf.MinW, s.imgConf.MinH)
	if err != nil {
		return err
	}

	return nil
}

//TODO figure out how to do it other way...
func getValidatorsAssetsPath(chain coin.Coin) string {
	return fmt.Sprintf("../assets_ts/blockchains/%s/validators/assets", chain.Handle)
}

func isStackingChain(c coin.Coin) bool {
	for _, stackingChain := range config.StackingChains {
		if c.ID == stackingChain.ID {
			return true
		}
	}

	return false
}
