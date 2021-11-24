package validation

import (
	"fmt"
	"io/fs"

	"github.com/trustwallet/assets-go-libs/pkg"
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
		if !pkg.Contains(f.Name(), allowedFiles) {
			compErr.Append(fmt.Errorf("%w: %s", ErrNotAllowedFile, f.Name()))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}
