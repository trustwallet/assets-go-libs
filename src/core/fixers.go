package core

import (
	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/assets-go-libs/pkg/file"
)

func (s *Service) FixJSON(file *file.AssetFile) error {
	return pkg.FormatJSONFile(file.Info.Path())
}
