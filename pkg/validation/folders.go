package validation

import (
	"fmt"
	"io/fs"

	log "github.com/sirupsen/logrus"

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

func ValidateAllowedFiles(files []fs.DirEntry, allowedFiles []string) error {
	compErr := NewErrComposite()
	for _, f := range files {
		log.WithField("allowed_file", f.Name()).WithField("files", files).Debug("Allowed files validation")
		if !pkg.Contains(f.Name(), allowedFiles) {
			compErr.Append(fmt.Errorf("%w %s", ErrNotAllowedFile, f.Name()))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}
