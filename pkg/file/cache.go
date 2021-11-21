package file

import (
	"sync"
)

type FileProvider struct {
	mu    *sync.RWMutex
	cache map[string]*AssetFile
}

func NewFileProvider(filePaths ...string) *FileProvider {
	var filesMap = make(map[string]*AssetFile)

	for _, path := range filePaths {
		assetF := newAssetFile(path)
		filesMap[path] = assetF
	}

	return &FileProvider{
		mu:    &sync.RWMutex{},
		cache: filesMap,
	}
}

func (f *FileProvider) GetAssetFile(path string) (*AssetFile, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return f.getFile(path)
}

func (f *FileProvider) getFile(path string) (*AssetFile, error) {
	if file, exists := f.cache[path]; exists {
		err := file.open()
		if err != nil {
			return nil, err
		}

		return file, nil
	}

	assetF := newAssetFile(path)
	f.cache[path] = assetF

	err := assetF.open()
	if err != nil {
		return nil, err
	}

	return assetF, nil
}
