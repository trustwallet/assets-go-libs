package validators

import "github.com/trustwallet/assets-go-libs/pkg/assetfs"

type Validator struct {
	ValidationName string
	FileType       string

	Run func(f *assetfs.AssetFile) error
}
