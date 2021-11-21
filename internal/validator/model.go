package validator

import "github.com/trustwallet/assets-go-libs/pkg/file"

type Validator struct {
	ValidationName string
	FileType       string

	Run func(f *file.AssetFile) error
}
