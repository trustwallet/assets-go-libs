package validators

import "github.com/trustwallet/assets-backend/pkg/assetfs"

type Validator struct {
	ValidationName string
	FileType       string

	Run func(f *assetfs.AssetFile) error
}
