package validation

import (
	"fmt"
	"os"

	"github.com/trustwallet/assets-go-libs/image"
)

const (
	bytesInKB   = 1024
	sizeLimitKB = 100

	MaxW = 512
	MaxH = 512
	MinW = 128
	MinH = 128
)

// TODO: Fix all incorrect logos in "assets" and then we can use ValidatePngImageDimension instead.
// Old logo's have bad dimensions like 195x163, 60x60 and etc. So this method is used only in CI scripts.
// ValidatePngImageDimensionForCI should be removed after logo's fixing.
func ValidatePngImageDimensionForCI(path string) error {
	imgWidth, imgHeight, err := image.GetPNGImageDimensions(path)
	if err != nil {
		return err
	}

	if imgWidth > MaxW || imgHeight > MaxH || imgWidth < 60 || imgHeight < 60 {
		return fmt.Errorf("%w: max - %dx%d, min - %dx%d; given %dx%d",
			ErrInvalidImgDimension, MaxW, MaxH, MinW, MinH, imgWidth, imgHeight)
	}

	return nil
}

func ValidatePngImageDimension(path string) error {
	width, height, err := image.GetPNGImageDimensions(path)
	if err != nil {
		return fmt.Errorf("failed to get png dimensions: %w", err)
	}

	return ValidateImageDimension(width, height)
}

func ValidateImageDimension(width, height int) error {
	if width > MaxW || height > MaxH || width < MinW || height < MinH || width != height {
		return fmt.Errorf("%w: max - %dx%d, min - %dx%d; given %dx%d",
			ErrInvalidImgDimension, MaxW, MaxH, MinW, MinH, width, height)
	}

	return nil
}

func ValidateLogoFileSize(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
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
