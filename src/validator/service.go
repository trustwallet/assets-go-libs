package validator

import (
	"github.com/trustwallet/assets-go-libs/pkg/file"
)

type Service struct {
	fileService *file.FileService
}

func NewService(fileProvider *file.FileService) (*Service, error) {
	return &Service{
		fileService: fileProvider,
	}, nil
}

// nolint:funlen
func (s *Service) GetValidator(f *file.AssetFile) *Validator {
	fileType := f.Info.Type()

	switch fileType {
	case file.TypeRootFolder:
		return &Validator{
			Name:     "Root folder contains only allowed files",
			Run:      s.ValidateRootFolder,
			FileType: fileType,
		}
	case file.TypeChainFolder:
		return &Validator{
			Name:     "Chain folders are lowercase and contains only allowed files",
			Run:      s.ValidateChainFolder,
			FileType: fileType,
		}
	case file.TypeChainLogoFile, file.TypeAssetLogoFile, file.TypeValidatorsLogoFile, file.TypeDappsLogoFile:
		return &Validator{
			Name:     "Logos (size, dimension)",
			Run:      s.ValidateImage,
			FileType: fileType,
		}
	case file.TypeAssetFolder:
		return &Validator{
			Name:     "Each asset folder has valid asset address and contains logo/info",
			Run:      s.ValidateAssetFolder,
			FileType: fileType,
		}
	case file.TypeDappsFolder:
		return &Validator{
			Name:     "Dapps folder (allowed only png files, lowercase)",
			Run:      s.ValidateDappsFolder,
			FileType: fileType,
		}
	case file.TypeAssetInfoFile:
		return &Validator{
			Name:     "Asset info (is valid json, fields)",
			Run:      s.ValidateAssetInfoFile,
			FileType: fileType,
		}
	case file.TypeChainInfoFile:
		return &Validator{
			Name:     "Chain Info (is valid json, fields)",
			Run:      s.ValidateChainInfoFile,
			FileType: fileType,
		}
	case file.TypeValidatorsListFile:
		return &Validator{
			Name:     "Validators list file",
			Run:      s.ValidateValidatorsListFile,
			FileType: fileType,
		}
	case file.TypeTokenListFile:
		return &Validator{
			Name:     "Token list (if assets from list present in chain)",
			Run:      s.ValidateTokenListFile,
			FileType: fileType,
		}
	case file.TypeChainInfoFolder:
		return &Validator{
			Name:     "Chain Info Folder (has files)",
			Run:      s.ValidateInfoFolder,
			FileType: fileType,
		}
	case file.TypeValidatorsAssetFolder:
		return &Validator{
			Name:     "Validators asset folder (has logo, valid asset address)",
			Run:      s.ValidateValidatorsAssetFolder,
			FileType: fileType,
		}
	}

	return nil
}

func (s *Service) GetFixer(f *file.AssetFile) *Validator {
	fileType := f.Info.Type()

	switch fileType {

	}

	return nil
}
