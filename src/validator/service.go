package validator

import (
	"github.com/trustwallet/assets-go-libs/pkg/file"
)

type Service struct {
	fileProvider *file.FileProvider
}

func NewService(fileProvider *file.FileProvider) (*Service, error) {
	return &Service{
		fileProvider: fileProvider,
	}, nil
}

// nolint:funlen
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
	case file.TypeAssetInfoFile:
		return &Validator{
			ValidationName: "Asset info (is valid json, fields)",
			Run:            s.ValidateAssetInfoFile,
			FileType:       fileType,
		}
	case file.TypeChainInfoFile:
		return &Validator{
			ValidationName: "Chain Info (is valid json, fields)",
			Run:            s.ValidateChainInfoFile,
			FileType:       fileType,
		}
	case file.TypeValidatorsListFile:
		return &Validator{
			ValidationName: "Validators list file",
			Run:            s.ValidateValidatorsListFile,
			FileType:       fileType,
		}
	case file.TypeTokenListFile:
		return &Validator{
			ValidationName: "Token list (if assets from list present in chain)",
			Run:            s.ValidateTokenListFile,
			FileType:       fileType,
		}
	case file.TypeChainInfoFolder:
		return &Validator{
			ValidationName: "Chain Info Folder (has files)",
			Run:            s.ValidateInfoFolder,
			FileType:       fileType,
		}
	case file.TypeValidatorsAssetFolder:
		return &Validator{
			ValidationName: "Validators asset folder (has logo, valid asset address)",
			Run:            s.ValidateValidatorsAssetFolder,
			FileType:       fileType,
		}
	}

	return nil
}
