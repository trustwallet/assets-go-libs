package validation

import (
	"fmt"
	"path/filepath"

	"github.com/trustwallet/assets-go-libs/pkg"
)

func ValidateLowercase(name string) error {
	if !pkg.IsLowerCase(name) {
		return fmt.Errorf("%w: it should be in lowercase", ErrInvalidFileCase)
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
