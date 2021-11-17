package validation

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif" // Added gif and jpeg configs for image because repo has images not only in .png extension, gif and jpeg too, but name contains .png
	_ "image/jpeg"
	_ "image/png"
	"io"

	_ "golang.org/x/image/webp"
)

func ValidateImageDimension(file io.Reader, maxW int, maxH int, minW int, minH int) error {
	img, name, err := image.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("failed to decode image, err %w", err)
	}

	if img.Width > maxW || img.Height > maxH || img.Height < minH || img.Width < minW {
		return fmt.Errorf(
			"%w %s , max - %dpx x %dpx, min - %dpx x %dpx; given %dpx %d",
			ErrInvalidImgDimension,
			name,
			maxW,
			maxH,
			minW,
			minH,
			img.Width,
			img.Height,
		)
	}

	return nil
}

func ValidateSize(b []byte, needleSize int) error {
	if len(b) > kbToBytes(needleSize) {
		return fmt.Errorf("%w, should be less than %dKB", ErrInvalidFileSize, needleSize)
	}

	return nil
}

func ValidateJson(b []byte) error {
	if !json.Valid(b) {
		return ErrInvalidJson
	}

	return nil
}

func kbToBytes(b int) int {
	return b * 1024
}
