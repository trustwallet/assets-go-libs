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

func (s *Service) ValidateRootFolder(file *file.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	err = validation.ValidateAllowedFiles(dirFiles, config.Default.ValidatorsSettings.RootFolder.AllowedFiles)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateChainFolder(file *file.AssetFile) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	var compErr = validation.NewErrComposite()

	err = validation.ValidateLowercase(fileInfo.Name())
	if err != nil {
		compErr.Append(err)
	}

	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	err = validation.ValidateAllowedFiles(dirFiles, config.Default.ValidatorsSettings.ChainFolder.AllowedFiles)
	if err != nil {
		compErr.Append(err)
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func (s *Service) ValidateImage(file *file.AssetFile) error {
	var compErr = validation.NewErrComposite()

	err := validation.ValidateLogoFileSize(file.File)
	if err != nil {
		compErr.Append(err)
	}

	err = validation.ValidatePngImageDimension(file)
	if err != nil {
		compErr.Append(err)
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func (s *Service) ValidateAssetFolder(file *file.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	var compErr = validation.NewErrComposite()

	err = validation.ValidateAllowedFiles(dirFiles, config.Default.ValidatorsSettings.AssetFolder.AllowedFiles)
	if err != nil {
		compErr.Append(err)
	}

	err = validation.ValidateAssetAddress(file.Info.Chain(), file.Info.Asset())
	if err != nil {
		compErr.Append(err)
	}

	errInfo := validation.ValidateHasFiles(dirFiles, []string{"info.json"})
	errLogo := validation.ValidateHasFiles(dirFiles, []string{"logo.png"})

	if errLogo != nil || errInfo != nil {
		infoFile, err := s.fileService.GetAssetFile(fmt.Sprintf("%s/info.json", file.Info.Path()))
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
			compErr.Append(fmt.Errorf("%w: logo.png for non-spam assest", validation.ErrMissingFile))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func (s *Service) ValidateDappsFolder(file *file.AssetFile) error {
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	var compErr = validation.NewErrComposite()
	for _, dirFile := range dirFiles {
		err = validation.ValidateExtension(dirFile.Name(), config.Default.ValidatorsSettings.DappsFolder.Ext)
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

func (s *Service) ValidateChainInfoFile(file *file.AssetFile) error {
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
		return fmt.Errorf("%w: failed to seek reader", validation.ErrInvalidJson)
	}

	var payload info.CoinModel
	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		return fmt.Errorf("%w: failed to decode", err)
	}

	tags := make([]string, len(config.Default.ValidatorsSettings.CoinInfoFile.Tags))
	for i, t := range config.Default.ValidatorsSettings.CoinInfoFile.Tags {
		tags[i] = t.ID
	}

	err = info.ValidateCoin(payload, file.Info.Chain(), file.Info.Asset(), tags)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateAssetInfoFile(file *file.AssetFile) error {
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
		return fmt.Errorf("%w: failed to seek reader", validation.ErrInvalidJson)
	}

	var payload info.AssetModel
	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		return fmt.Errorf("%w: failed to decode", err)
	}

	err = info.ValidateAsset(payload, file.Info.Chain(), file.Info.Asset())
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateValidatorsListFile(file *file.AssetFile) error {
	if !isStackingChain(file.Info.Chain()) {
		return nil
	}

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

	listIDs := make([]string, len(model))
	for i, listItem := range model {
		listIDs[i] = *listItem.ID
	}

	assetsPath := fmt.Sprintf("blockchains/%s/validators/assets", file.Info.Chain().Handle)
	assetFolder, err := s.fileService.GetAssetFile(assetsPath)
	if err != nil {
		return err
	}

	dirAssetFolderFiles, err := assetFolder.ReadDir(0)
	if err != nil {
		return err
	}

	err = validation.ValidateAllowedFiles(dirAssetFolderFiles, listIDs)
	if err != nil {
		return err
	}

	return nil
}

func isStackingChain(c coin.Coin) bool {
	for _, stackingChain := range config.StackingChains {
		if c.ID == stackingChain.ID {
			return true
		}
	}

	return false
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
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	compErr := validation.NewErrComposite()
	err = validation.ValidateValidatorsAddress(file.Info.Chain(), file.Info.Asset())
	if err != nil {
		compErr.Append(err)
	}

	err = validation.ValidateHasFiles(dirFiles, config.Default.ValidatorsSettings.ChainValidatorsAssetFolder.HasFiles)
	if err != nil {
		compErr.Append(err)
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}
