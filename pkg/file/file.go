package file

import (
	"os"

	"github.com/trustwallet/go-primitives/coin"
)

type AssetFile struct {
	*os.File

	Info *AssetInfo
}

func NewAssetFile(path string) *AssetFile {
	info := AssetInfo{
		path: NewPath(path),
	}

	return &AssetFile{
		Info: &info,
	}
}

func (f *AssetFile) Open() error {
	file, err := os.Open(f.Info.Path())
	if err != nil {
		return err
	}

	f.File = file

	return nil
}

type AssetInfo struct {
	path *Path
}

func (i *AssetInfo) Path() string {
	return i.path.String()
}

func (i *AssetInfo) Type() string {
	return i.path.fileType
}

func (i *AssetInfo) Chain() coin.Coin {
	return i.path.chain
}

func (i *AssetInfo) Asset() string {
	return i.path.asset
}
