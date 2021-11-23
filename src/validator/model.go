package validator

import "github.com/trustwallet/assets-go-libs/pkg/file"

type Validator struct {
	Name     string
	FileType string

	Run func(f *file.AssetFile) error
}

type Fixer struct {
	Validator
}
