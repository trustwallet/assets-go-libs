package image

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/draw"

	"github.com/trustwallet/assets-go-libs/http"
)

func GetPNGImageDimensions(path string) (width, height int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return GetPNGImageDimensionsFromReader(file)
}

func GetPNGImageDimensionsFromReader(r io.Reader) (width, height int, err error) {
	img, err := png.DecodeConfig(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to decode config: %w", err)
	}

	return img.Width, img.Height, nil
}

func ResizePNG(path string, targetWidth, targetHeight int) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Decode the image (from PNG to image.Image).
	src, err := png.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image bytes: %w", err)
	}

	output, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer output.Close()

	// Set the expected size.
	dst := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	// Resize.
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	err = png.Encode(output, dst)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}

func CreatePNGFromURL(logoURL, logoPath string) error {
	imgBytes, err := http.GetHTTPResponseBytes(logoURL)
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return fmt.Errorf("failed to decode image bytes: %w", err)
	}

	out, err := os.Create(logoPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}
