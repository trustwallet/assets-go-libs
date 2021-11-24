package validation

import (
	"fmt"
	"io/fs"

	"github.com/trustwallet/assets-go-libs/pkg"
)

func ValidateHasFiles(files []fs.DirEntry, fileNames []string) error {
	if len(files) < len(fileNames) {
		return fmt.Errorf("%w, folders length shorter then needed", ErrMissingFile)
	}

	compErr := NewErrComposite()
OutLoop:
	for _, fName := range fileNames {
		for _, dirF := range files {
			if dirF.Name() == fName {
				continue OutLoop
			}
		}

		compErr.Append(fmt.Errorf("%w %s", ErrMissingFile, fName))
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func ValidateFilesNotInList(files []fs.DirEntry, fileList []string) error {
	compErr := NewErrComposite()

	for _, dir := range files {
		var found bool
		for _, f := range fileList {
			if dir.Name() == f {
				found = true
				break
			}
		}

		if !found {
			compErr.Append(fmt.Errorf("%w: %s", ErrMissingFile, dir.Name()))
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
		if !pkg.Contains(f.Name(), allowedFiles) {
			compErr.Append(fmt.Errorf("%w: %s", ErrNotAllowedFile, f.Name()))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}
