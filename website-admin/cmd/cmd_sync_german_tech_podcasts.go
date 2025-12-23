package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/EngineeringKiosk/website/website-admin/utils"
)

const (
	germanTechPodcastsGitRepo          = "https://github.com/EngineeringKiosk/GermanTechPodcasts.git"
	germanTechPodcastsRepoName         = "GermanTechPodcasts"
	germanTechPodcastsJSONPathInRepo   = "generated"
	germanTechPodcastsImagesPathInRepo = "generated/images"
	germanTechPodcastsOPMLPathInRepo   = "podcasts.opml"
	germanTechPodcastsDefaultStorage   = "src/content/germantechpodcasts"
	germanTechPodcastsDefaultOPMLPath  = "public/deutsche-tech-podcasts/podcasts.opml"
	germanTechPodcastsImageMaxSize     = 700
)

// syncGermanTechPodcastsCmd represents the german-tech-podcasts sync command
var syncGermanTechPodcastsCmd = &cobra.Command{
	Use:   "german-tech-podcasts",
	Short: "Sync German Tech Podcasts data",
	Long: `Sync data from the German Tech Podcasts repository.

This command clones the repository from GitHub, copies the JSON data files,
images, and OPML file to the local content directory, and performs necessary transformations:
  - Adjusts image paths to be relative
  - Resizes images larger than 700x700 to 700x700

Source: https://github.com/EngineeringKiosk/GermanTechPodcasts`,
	RunE: RunSyncGermanTechPodcastsCmd,
}

func init() {
	syncCmd.AddCommand(syncGermanTechPodcastsCmd)

	syncGermanTechPodcastsCmd.Flags().StringVarP(&flagSyncGermanTechPodcastsStoragePath, "storage-path", "s", "", "Path to store JSON and image files (default: src/content/germantechpodcasts)")
	syncGermanTechPodcastsCmd.Flags().StringVarP(&flagSyncGermanTechPodcastsOPMLPath, "opml-path", "o", "", "Path to store OPML file (default: public/deutsche-tech-podcasts/podcasts.opml)")
}

func RunSyncGermanTechPodcastsCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Msg("starting")

	// Set default storage path if not provided
	storagePath := flagSyncGermanTechPodcastsStoragePath
	if storagePath == "" {
		storagePath = germanTechPodcastsDefaultStorage
	}
	storagePath = utils.BuildCorrectFilePath(storagePath)

	// Set default OPML path if not provided
	opmlPath := flagSyncGermanTechPodcastsOPMLPath
	if opmlPath == "" {
		opmlPath = germanTechPodcastsDefaultOPMLPath
	}
	opmlPath = utils.BuildCorrectFilePath(opmlPath)

	// Create temp directory for git clone
	tmpDir, err := os.MkdirTemp("", germanTechPodcastsRepoName+"-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer func() {
		logger.Info().
			Str("path", tmpDir).
			Msg("Cleaning up temp directory")
		if err := os.RemoveAll(tmpDir); err != nil {
			logger.Warn().Err(err).Msg("Failed to remove temp directory")
		}
	}()

	// Clone the repository
	if err := cloneRepository(logger, germanTechPodcastsGitRepo, tmpDir); err != nil {
		return err
	}

	// Copy and transform JSON files
	jsonSourceDir := filepath.Join(tmpDir, germanTechPodcastsJSONPathInRepo)
	if err := syncGermanTechPodcastsJSONFiles(logger, jsonSourceDir, storagePath); err != nil {
		return err
	}

	// Copy and resize image files
	imagesSourceDir := filepath.Join(tmpDir, germanTechPodcastsImagesPathInRepo)
	if err := syncGermanTechPodcastsImageFiles(logger, imagesSourceDir, storagePath); err != nil {
		return err
	}

	// Copy OPML file
	opmlSourcePath := filepath.Join(tmpDir, germanTechPodcastsOPMLPathInRepo)
	if err := syncOPMLFile(logger, opmlSourcePath, opmlPath); err != nil {
		return err
	}

	logger.Info().Msg("Sync completed successfully")
	return nil
}

// syncGermanTechPodcastsJSONFiles copies and transforms JSON files from source to destination
func syncGermanTechPodcastsJSONFiles(logger zerolog.Logger, sourceDir, destDir string) error {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read JSON directory: %w", err)
	}

	jsonCount := 0
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		jsonCount++
	}

	logger.Info().
		Int("count", jsonCount).
		Str("dir", sourceDir).
		Msg("Found JSON files")

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		srcPath := filepath.Join(sourceDir, entry.Name())
		dstPath := filepath.Join(destDir, entry.Name())

		logger.Debug().
			Str("src", srcPath).
			Str("dst", dstPath).
			Msg("Processing JSON file")

		if err := copyAndTransformGermanTechPodcastsJSONFile(srcPath, dstPath); err != nil {
			logger.Warn().
				Err(err).
				Str("file", entry.Name()).
				Msg("Failed to process JSON file")
			continue
		}

		logger.Info().
			Str("file", entry.Name()).
			Msg("Copied and transformed JSON file")
	}

	return nil
}

// copyAndTransformGermanTechPodcastsJSONFile copies a JSON file and applies transformations
func copyAndTransformGermanTechPodcastsJSONFile(srcPath, dstPath string) error {
	// Read source file
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON
	var content map[string]interface{}
	if err := json.Unmarshal(data, &content); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Transform image path to relative (e.g., "./image.png")
	if imagePath, ok := content["image"].(string); ok {
		content["image"] = "./" + filepath.Base(imagePath)
	}

	// Write transformed JSON
	output, err := json.MarshalIndent(content, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(dstPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// syncGermanTechPodcastsImageFiles copies image files and resizes if necessary
func syncGermanTechPodcastsImageFiles(logger zerolog.Logger, sourceDir, destDir string) error {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read images directory: %w", err)
	}

	imageCount := 0
	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		imageCount++
	}

	logger.Info().
		Int("count", imageCount).
		Str("dir", sourceDir).
		Msg("Found image files")

	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		srcPath := filepath.Join(sourceDir, entry.Name())
		dstPath := filepath.Join(destDir, entry.Name())

		logger.Debug().
			Str("src", srcPath).
			Str("dst", dstPath).
			Msg("Processing image file")

		// Copy the file first
		if err := copyFile(srcPath, dstPath); err != nil {
			logger.Warn().
				Err(err).
				Str("file", entry.Name()).
				Msg("Failed to copy image file")
			continue
		}

		// Check if resizing is needed
		width, height, err := utils.GetImageDimensions(dstPath)
		if err != nil {
			logger.Debug().
				Err(err).
				Str("file", entry.Name()).
				Msg("Could not get image dimensions, skipping resize")
		} else if width > germanTechPodcastsImageMaxSize && height > germanTechPodcastsImageMaxSize {
			logger.Info().
				Str("file", entry.Name()).
				Int("originalWidth", width).
				Int("originalHeight", height).
				Int("targetSize", germanTechPodcastsImageMaxSize).
				Msg("Resizing image")

			if err := utils.ResizeImage(dstPath, germanTechPodcastsImageMaxSize, germanTechPodcastsImageMaxSize); err != nil {
				logger.Warn().
					Err(err).
					Str("file", entry.Name()).
					Msg("Failed to resize image")
			}
		}

		logger.Info().
			Str("file", entry.Name()).
			Msg("Processed image file")
	}

	return nil
}

// syncOPMLFile copies the OPML file to the destination
func syncOPMLFile(logger zerolog.Logger, srcPath, dstPath string) error {
	logger.Info().
		Str("src", srcPath).
		Str("dst", dstPath).
		Msg("Copying OPML file")

	if err := copyFile(srcPath, dstPath); err != nil {
		return fmt.Errorf("failed to copy OPML file: %w", err)
	}

	logger.Info().Msg("OPML file copied successfully")
	return nil
}
