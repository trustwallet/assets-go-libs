package pkg

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func GetPNGImageDimensions(path string) (width, height int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	img, err := png.DecodeConfig(file)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to decode image: %w", err)
	}

	return img.Width, img.Height, nil
}

func ResizePNG(path string, targetWidth, targetHeight int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image (from PNG to image.Image).
	src, err := png.Decode(file)
	if err != nil {
		return err
	}

	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	// Set the expected size that you want.
	dst := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	// Resize.
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	err = png.Encode(output, dst)
	if err != nil {
		return err
	}

	return nil
}
