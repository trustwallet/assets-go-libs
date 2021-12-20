package file

import (
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func CreateDirPath(path string) error {
	dirPath := filepath.Dir(path)

	return os.MkdirAll(dirPath, os.ModePerm)
}
