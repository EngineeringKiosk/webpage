package episode

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/EngineeringKiosk/website/website-admin/utils"
)

// GetEpisodeNumberFromFilename extracts the episode number from a filename
// leadingZero determines whether to return with leading zeros
func GetEpisodeNumberFromFilename(filename string, leadingZero bool) (string, error) {
	filename = filepath.Base(filename)
	index := strings.Index(filename, "-")

	var episodeNumber string

	// Handle episodes starting with '-1' (negative episode numbers)
	if index == 0 {
		remainingFilename := filename[1:]
		nextIndex := strings.Index(remainingFilename, "-")
		if nextIndex == -1 {
			return "", fmt.Errorf("invalid filename format: %s", filename)
		}
		episodeNumber = filename[0 : nextIndex+1]
	} else if index == -1 {
		return "", fmt.Errorf("invalid filename format: %s", filename)
	} else {
		episodeNumber = filename[0:index]
	}

	if !leadingZero {
		episodeNumber = trimEpisodeNumber(episodeNumber)
	} else {
		// Pad with zeros for positive numbers if needed
		if !strings.HasPrefix(episodeNumber, "-") {
			// Parse and reformat
			var num int
			_, err := fmt.Sscanf(episodeNumber, "%d", &num)
			if err != nil {
				return "", fmt.Errorf("invalid episode number: %s", episodeNumber)
			}
			if num < 10 && num >= 0 {
				episodeNumber = fmt.Sprintf("%02d", num)
			}
		}
	}

	return episodeNumber, nil
}

// GetTranscriptSlimPathByEpisodeNumber returns the transcript path for an episode
func GetTranscriptSlimPathByEpisodeNumber(episodeNumber, transcriptDir string) string {
	transcriptFile := fmt.Sprintf("%s-transcript-slim.json", episodeNumber)
	filePath := filepath.Join(transcriptDir, transcriptFile)
	filePath = utils.BuildCorrectFilePath(filePath)

	if _, err := os.Stat(filePath); err == nil {
		// Adjust path based on working directory if needed
		filePath = strings.TrimPrefix(filePath, "../")
		return filePath
	}

	return ""
}
