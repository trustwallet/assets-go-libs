package file

import (
	"os"

	"github.com/trustwallet/go-primitives/coin"
)

type AssetFile struct {
	Info *AssetInfo
	*os.File
}

func Open(path string) (*AssetFile, error) {
	file := newAssetFile(path)
	err := file.open()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func newAssetFile(path string) *AssetFile {
	p := NewPath(path)

	info := AssetInfo{
		path: p,
	}

	return &AssetFile{
		Info: &info,
	}
}

func (f *AssetFile) open() error {
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
	return i.path.type_
}
func (i *AssetInfo) Chain() coin.Coin {
	return i.path.chain
}

func (i *AssetInfo) Asset() string {
	return i.path.asset
}
