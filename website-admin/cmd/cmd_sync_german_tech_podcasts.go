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
	germanTechPodcastsMaxImageSize     = 700
)

// syncGermanTechPodcastsCmd represents the german-tech-podcasts sync command
var syncGermanTechPodcastsCmd = &cobra.Command{
	Use:   "german-tech-podcasts",
	Short: "Sync the German Tech Podcasts directory from GitHub",
	Long: `Sync data from the German Tech Podcasts community repository.

The Engineering Kiosk website hosts a directory of German-language technology podcasts,
which is maintained as a separate community-driven GitHub repository. This command
synchronizes that external data into the website's content directory.

Source repository: https://github.com/EngineeringKiosk/GermanTechPodcasts

What this command does:
  1. Clones the GermanTechPodcasts repository to a temporary directory (shallow clone)
  2. Copies all JSON data files from generated/ to the local content directory
  3. Copies all podcast cover images from generated/images/
  4. Copies the OPML file for podcast app subscriptions
  5. Applies the following transformations:
     - Adjusts image paths in JSON files to be relative (e.g., "./podcast-name.jpg")
     - Resizes images larger than 700x700 pixels down to 700x700 for web optimization
  6. Cleans up the temporary clone

The JSON files contain structured podcast metadata including:
  - Podcast name, description, and website URL
  - RSS feed URL and language information
  - Tags/categories and cover image
  - Episode count and publish frequency

This command should be run periodically to pull in new podcast additions and updates
from the community repository.`,
	Example: `  # Sync German tech podcasts using default paths (run from project root)
  website-admin sync german-tech-podcasts

  # Specify a custom storage path for JSON and images
  website-admin sync german-tech-podcasts --storage-path ./custom/content/podcasts

  # Specify a custom path for the OPML file
  website-admin sync german-tech-podcasts --opml-path ./public/feeds/podcasts.opml

  # Use both custom paths
  website-admin sync german-tech-podcasts \
    --storage-path ./src/content/germantechpodcasts \
    --opml-path ./public/deutsche-tech-podcasts/podcasts.opml

  # Enable debug logging to see detailed file operations
  website-admin sync german-tech-podcasts --debug`,
	RunE:              RunSyncGermanTechPodcastsCmd,
	DisableAutoGenTag: true,
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
	var content map[string]any
	if err := json.Unmarshal(data, &content); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Transform image path to relative
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

// syncGermanTechPodcastsImageFiles copies and resizes image files from source to destination
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
			Msg("Copying image file")

		if err := copyFile(srcPath, dstPath); err != nil {
			logger.Warn().
				Err(err).
				Str("file", entry.Name()).
				Msg("Failed to copy image file")
			continue
		}

		// Check if image needs resizing
		width, height, err := utils.GetImageDimensions(dstPath)
		if err != nil {
			logger.Debug().
				Err(err).
				Str("file", entry.Name()).
				Msg("Could not get image dimensions, skipping resize")
		} else if width > germanTechPodcastsMaxImageSize && height > germanTechPodcastsMaxImageSize {
			logger.Info().
				Str("file", entry.Name()).
				Int("originalWidth", width).
				Int("originalHeight", height).
				Int("newSize", germanTechPodcastsMaxImageSize).
				Msg("Resizing image")

			if err := utils.ResizeImage(dstPath, germanTechPodcastsMaxImageSize, germanTechPodcastsMaxImageSize); err != nil {
				logger.Warn().
					Err(err).
					Str("file", entry.Name()).
					Msg("Failed to resize image")
			}
		}

		logger.Info().
			Str("file", entry.Name()).
			Msg("Copied image file")
	}

	return nil
}

// syncOPMLFile copies the OPML file from source to destination
func syncOPMLFile(logger zerolog.Logger, srcPath, dstPath string) error {
	logger.Info().
		Str("src", srcPath).
		Str("dst", dstPath).
		Msg("Copying OPML file")

	// Ensure destination directory exists
	destDir := filepath.Dir(dstPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	if err := copyFile(srcPath, dstPath); err != nil {
		return fmt.Errorf("failed to copy OPML file: %w", err)
	}

	logger.Info().
		Str("file", filepath.Base(dstPath)).
		Msg("Copied OPML file")

	return nil
}
