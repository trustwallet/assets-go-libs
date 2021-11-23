package file

import (
	"sync"
)

type FileService struct {
	mu    *sync.RWMutex
	cache map[string]*AssetFile
}

func NewFileService(filePaths ...string) *FileService {
	var filesMap = make(map[string]*AssetFile)

	for _, path := range filePaths {
		assetF := newAssetFile(path)
		filesMap[path] = assetF
	}

	return &FileService{
		mu:    &sync.RWMutex{},
		cache: filesMap,
	}
}

func (f *FileService) GetAssetFile(path string) (*AssetFile, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return f.getFile(path)
}

func (f *FileService) getFile(path string) (*AssetFile, error) {
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
