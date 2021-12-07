package validation

import (
	"fmt"
	"image/png"
	"io"
	"os"
)

const (
	bytesInKB   = 1024
	sizeLimitKB = 100

	maxW = 512
	maxH = 512
	minW = 128
	minH = 128
)

// TODO: Fix all logos in "assets" and then we can use ValidatePngImageDimension in CI.
// Old logo's have dimensions like 195x163, 60x60 and etc. This method is used only in CI.
// ValidatePngImageDimensionForCI should be removed after logo's fixing.
func ValidatePngImageDimensionForCI(file io.Reader) error {
	img, err := png.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// TODO: If we fix all incorrect logos in assets repo, we could use "|| img.Width != img.Height" in addition.
	if img.Width > maxW || img.Height > maxH || img.Height < minH {
		return fmt.Errorf("%w: max - %dx%d, min - %dx%d; given %dx%d",
			ErrInvalidImgDimension, maxW, maxH, minW, minH, img.Width, img.Height)
	}

	return nil
}

func ValidatePngImageDimension(file io.Reader) error {
	img, err := png.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	if img.Width > maxW || img.Height > maxH || img.Height < minH || img.Width < minW || img.Width != img.Height {
		return fmt.Errorf("%w: max - %dx%d, min - %dx%d; given %dx%d",
			ErrInvalidImgDimension, maxW, maxH, minW, minH, img.Width, img.Height)
	}

	return nil
}

func ValidateLogoFileSize(file *os.File) error {
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
