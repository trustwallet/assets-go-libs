package file

import (
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func CreateDirPath(path string) error {
	dirPath := filepath.Dir(path)

	return os.MkdirAll(dirPath, os.ModePerm)
}

func CreateFileWithPath(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), os.ModePerm); err != nil {
		return nil, err
	}

	return os.Create(p)
}
