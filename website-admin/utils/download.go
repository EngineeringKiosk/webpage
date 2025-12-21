package utils

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadFile downloads a file from a URL and saves it with the appropriate extension
func DownloadFile(url, localFilename string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Determine file extension from content type
	contentType := resp.Header.Get("Content-Type")
	var extensions []string
	switch contentType {
	// Force .jpg extension, as in mime.ExtensionsByType
	// the first entry is "".jpe"
	case "image/jpeg", "image/jpg", "image/pjpeg":
		extensions = []string{".jpg"}
	default:
		extensions, err = mime.ExtensionsByType(contentType)
		if err != nil || len(extensions) == 0 {
			// Try to get extension from URL
			ext := filepath.Ext(url)
			if ext == "" {
				ext = ".jpg" // Default extension
			}
			extensions = []string{ext}
		}
	}
	fileName := localFilename + extensions[0]

	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return fileName, nil
}
