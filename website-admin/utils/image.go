package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
)

// ResizeImage resizes an image to the specified width and height
func ResizeImage(inputPath string, width, height int) error {
	// Open the input file
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Create a new image with the target dimensions
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// Resize using bilinear interpolation
	draw.BiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	// Open output file
	outFile, err := os.Create(inputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()

	// Encode based on original format
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(outFile, dst, &jpeg.Options{Quality: 90})
	case "png":
		err = png.Encode(outFile, dst)
	default:
		// Default to JPEG
		err = jpeg.Encode(outFile, dst, &jpeg.Options{Quality: 90})
	}

	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}

// GetImageDimensions returns the width and height of an image file
func GetImageDimensions(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to open image: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to decode image config: %w", err)
	}

	return config.Width, config.Height, nil
}

// ImageExists checks if an image file exists with any common extension
func ImageExists(baseFilename string) (string, bool) {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}

	for _, ext := range extensions {
		fullPath := baseFilename + ext
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}

	// Also check if the baseFilename already has an extension
	if filepath.Ext(baseFilename) != "" {
		if _, err := os.Stat(baseFilename); err == nil {
			return baseFilename, true
		}
	}

	return "", false
}
