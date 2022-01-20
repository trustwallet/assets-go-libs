package validation

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	str "github.com/trustwallet/assets-go-libs/strings"
)

func ValidateHasFiles(files []fs.DirEntry, fileNames []string) error {
	compErr := NewErrComposite()

	if len(files) < len(fileNames) {
		compErr.Append(fmt.Errorf("%w: this folder must have more files", ErrMissingFile))

		return compErr
	}

	for _, fName := range fileNames {
		var found bool

		for _, dirF := range files {
			if dirF.Name() == fName {
				found = true

				break
			}
		}

		if !found {
			compErr.Append(fmt.Errorf("%w: %s", ErrMissingFile, fName))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func ValidateAllowedFiles(files []fs.DirEntry, allowedFiles []string) error {
	compErr := NewErrComposite()

	for _, f := range files {
		if !str.Contains(f.Name(), allowedFiles) {
			compErr.Append(fmt.Errorf("%w: %s", ErrNotAllowedFile, f.Name()))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func ValidateFileInPR(path string) error {
	if strings.HasPrefix(path, "dapps") {
		return nil
	}

	if strings.HasPrefix(path, "blockchains") {
		if strings.Index(path, "assets") > 0 ||
			strings.HasSuffix(path, "allowlist.json") ||
			strings.HasSuffix(path, "validators/list.json") {
			return nil
		}
	}

	return fmt.Errorf("%w: %s", ErrNotAllowedFile, path)
}

func ValidateLowercase(name string) error {
	if !str.IsLowerCase(name) {
		return fmt.Errorf("%w: should be in lowercase", ErrInvalidFileNameCase)
	}

	return nil
}

func ValidateExtension(name, ext string) error {
	fileExtension := filepath.Ext(name)
	if fileExtension != ext {
		return fmt.Errorf("%w: only %s allowed in folder", ErrInvalidFileExt, ext)
	}

	return nil
}
