package utils

import (
	"os"
	"path/filepath"
)

// BuildCorrectFilePath adjusts paths based on current working directory
func BuildCorrectFilePath(filePath string) string {
	// Determine if the script is called from root or from a subdirectory
	directory, err := os.Getwd()
	if err != nil {
		return filePath
	}

	folderName := filepath.Base(directory)
	if folderName == "website-admin" {
		return filepath.Join("..", filePath)
	}

	return filePath
}
