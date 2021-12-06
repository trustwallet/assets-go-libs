package validation

import (
	"fmt"
	"image/png"
	"io"
	"os"
)

const (
	kbInByte    = 1000
	sizeLimitKB = 100
	maxW        = 512
	maxH        = 512
	minW        = 128
	minH        = 128
)

func ValidatePngImageDimension(file io.Reader) error {
	img, err := png.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	if img.Width > maxW || img.Height > maxH || img.Height < minH || img.Width < minW || img.Width != img.Height {
		return fmt.Errorf("%w: logo should be 256x256: given %dx%d", ErrInvalidImgDimension, img.Width, img.Height)
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
	logoSizeKB := imgBytesCount / kbInByte

	if logoSizeKB > sizeLimitKB {
		return fmt.Errorf("%w: logo should be less than %dKB, given %dKB", ErrInvalidFileSize, sizeLimitKB, logoSizeKB)
	}

	return nil
}
