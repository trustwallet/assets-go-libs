package core

import (
	"fmt"

	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/go-primitives/address"
	"github.com/trustwallet/go-primitives/coin"
)

func (s *Service) FixJSON(file *file.AssetFile) error {
	return pkg.FormatJSONFile(file.Info.Path())
}

func (s *Service) FixETHAddressChecksum(file *file.AssetFile) error {
	if !coin.IsEVM(file.Info.Chain().ID) {
		return nil
	}

	assetsDirs, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	for _, assetDir := range assetsDirs {
		err = validation.ValidateETHForkAddress(file.Info.Chain(), assetDir.Name())
		if err != nil {
			checksum, e := address.EIP55Checksum(assetDir.Name())
			if e != nil {
				return err
			}

			if _, e = pkg.RunCmd(getGitMoveCommand(file.Info.Path(), assetDir.Name(), checksum), false); err != nil {
				return e
			}
		}
	}

	return nil
}

func getGitMoveCommand(dirPath, oldFileName, newFileName string) string {
	oldFullName := fmt.Sprintf("%s/%s", dirPath, oldFileName)
	newFullName := fmt.Sprintf("%s/%s", dirPath, newFileName)

	return fmt.Sprintf("git mv %s %s-temp && git mv %s-temp %s", oldFullName, newFullName, newFullName, newFullName)
}
