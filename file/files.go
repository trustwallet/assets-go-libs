package file

import (
	"fmt"
	"io/fs"
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

func ReadDir(path string) ([]fs.DirEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	dirFiles, err := file.ReadDir(0)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}

	return dirFiles, nil
}
