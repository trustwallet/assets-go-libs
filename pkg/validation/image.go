package validation

import (
	"fmt"
	"os"

	"github.com/trustwallet/assets-go-libs/pkg"
)

const (
	bytesInKB   = 1024
	sizeLimitKB = 100

	MaxW = 512
	MaxH = 512
	MinW = 128
	MinH = 128
)

// TODO: Fix all logos in "assets" and then we can use ValidatePngImageDimension in CI.
// Old logo's have dimensions like 195x163, 60x60 and etc. This method is used only in CI.
// ValidatePngImageDimensionForCI should be removed after logo's fixing.
func ValidatePngImageDimensionForCI(path string) error {
	imgWidth, imgHeight, err := pkg.GetPNGImageDimensions(path)
	if err != nil {
		return err
	}

	// TODO: If we fix all incorrect logos in assets repo, we could use "|| img.Width != img.Height" in addition.
	if imgWidth > MaxW || imgHeight > MaxH || imgWidth < 60 || imgHeight < 60 {
		return fmt.Errorf("%w: max - %dx%d, min - %dx%d; given %dx%d",
			ErrInvalidImgDimension, MaxW, MaxH, MinW, MinH, imgWidth, imgHeight)
	}

	return nil
}

func ValidatePngImageDimension(path string) error {
	imgWidth, imgHeight, err := pkg.GetPNGImageDimensions(path)
	if err != nil {
		return err
	}

	if imgWidth > MaxW || imgHeight > MaxH || imgHeight < MinH || imgWidth < MinW || imgWidth != imgHeight {
		return fmt.Errorf("%w: max - %dx%d, min - %dx%d; given %dx%d",
			ErrInvalidImgDimension, MaxW, MaxH, MinW, MinH, imgWidth, imgHeight)
	}

	return nil
}

func ValidateLogoFileSize(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	return validateLogoSize(int(fileInfo.Size()))
}

func ValidateLogoStreamSize(imgBytes []byte) error {
	return validateLogoSize(len(imgBytes))
}

func validateLogoSize(imgBytesCount int) error {
	logoSizeKB := imgBytesCount / bytesInKB

	if logoSizeKB > sizeLimitKB {
		return fmt.Errorf("%w: logo should be less than %dKB, given %dKB", ErrInvalidFileSize, sizeLimitKB, logoSizeKB)
	}

	return nil
}
