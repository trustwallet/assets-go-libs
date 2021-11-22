package validator

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/assets-go-libs/pkg/validation/info"
	"github.com/trustwallet/assets-go-libs/src/config"
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

	err := validation.ValidateSize(file, config.Default.ValidatorsSettings.ImageFile.Size)
	if err != nil {
		compErr.Append(err)
	}

	err = validation.ValidateImageDimension(file,
		config.Default.ValidatorsSettings.ImageFile.MaxW,
		config.Default.ValidatorsSettings.ImageFile.MaxH,
		config.Default.ValidatorsSettings.ImageFile.MinW,
		config.Default.ValidatorsSettings.ImageFile.MinH,
	)
	if err != nil {
		compErr.Append(err)
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func (s *Service) ValidateAssetFolder(file *file.AssetFile) error {
	assetInfo := file.Info
	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	var compErr = validation.NewErrComposite()

	err = validation.ValidateAllowedFiles(dirFiles, config.Default.ValidatorsSettings.AssetFolder.AllowedFiles)
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
		infoFile, err := s.fileProvider.GetAssetFile(fmt.Sprintf("%s/info.json", assetInfo.Path()))
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
