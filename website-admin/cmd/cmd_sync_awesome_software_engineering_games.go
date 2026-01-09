package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/EngineeringKiosk/website/website-admin/utils"
)

const (
	awesomeGamesGitRepo          = "https://github.com/EngineeringKiosk/awesome-software-engineering-games.git"
	awesomeGamesRepoName         = "awesome-software-engineering-games"
	awesomeGamesJSONPathInRepo   = "generated"
	awesomeGamesImagesPathInRepo = "generated/images"
	awesomeGamesDefaultStorage   = "src/content/awesome-software-engineering-games"
)

// syncAwesomeSoftwareEngineeringGamesCmd represents the awesome-software-engineering-games sync command
var syncAwesomeSoftwareEngineeringGamesCmd = &cobra.Command{
	Use:   "awesome-software-engineering-games",
	Short: "Sync the Awesome Software Engineering Games catalog from GitHub",
	Long: `Sync data from the Awesome Software Engineering Games community repository.

The Engineering Kiosk website features a catalog of video games that teach programming,
software engineering, and computer science concepts. This curated list is maintained
as a separate community-driven GitHub repository, and this command synchronizes that
data into the website's content directory.

Source repository: https://github.com/EngineeringKiosk/awesome-software-engineering-games

What this command does:
  1. Clones the awesome-software-engineering-games repository to a temporary directory (shallow clone)
  2. Copies all JSON data files from generated/ to the local content directory
  3. Copies all game cover images from generated/images/
  4. Applies the following transformations:
     - Adjusts image paths in JSON files to be relative (e.g., "./game-name.jpg")
     - Normalizes genre names for consistency:
       * "MMO" -> "Massively Multiplayer"
       * "Simulationen" -> "Simulation"
  5. Cleans up the temporary clone

The JSON files contain structured game metadata including:
  - Game name, description, and official website
  - Available platforms (PC, Mac, Linux, consoles)
  - Genres and educational topics covered
  - German-specific content including localized descriptions

This command should be run periodically to pull in new game additions and updates
from the community repository.`,
	Example: `  # Sync games catalog using default paths (run from project root)
  website-admin sync awesome-software-engineering-games

  # Specify a custom storage path for JSON and images
  website-admin sync awesome-software-engineering-games --storage-path ./custom/content/games

  # Enable debug logging to see detailed file operations
  website-admin sync awesome-software-engineering-games --debug

  # Disable all logging for quiet operation (useful in scripts)
  website-admin sync awesome-software-engineering-games --disable-logging`,
	RunE:              RunSyncAwesomeSoftwareEngineeringGamesCmd,
	DisableAutoGenTag: true,
}

func init() {
	syncCmd.AddCommand(syncAwesomeSoftwareEngineeringGamesCmd)

	syncAwesomeSoftwareEngineeringGamesCmd.Flags().StringVarP(&flagSyncAwesomeGamesStoragePath, "storage-path", "s", "", "Path to store JSON and image files (default: src/content/awesome-software-engineering-games)")
}

func RunSyncAwesomeSoftwareEngineeringGamesCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Msg("starting")

	// Set default storage path if not provided
	storagePath := flagSyncAwesomeGamesStoragePath
	if storagePath == "" {
		storagePath = awesomeGamesDefaultStorage
	}
	storagePath = utils.BuildCorrectFilePath(storagePath)

	// Create temp directory for git clone
	tmpDir, err := os.MkdirTemp("", awesomeGamesRepoName+"-*")
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
	if err := cloneRepository(logger, awesomeGamesGitRepo, tmpDir); err != nil {
		return err
	}

	// Copy and transform JSON files
	jsonSourceDir := filepath.Join(tmpDir, awesomeGamesJSONPathInRepo)
	if err := syncJSONFiles(logger, jsonSourceDir, storagePath); err != nil {
		return err
	}

	// Copy image files
	imagesSourceDir := filepath.Join(tmpDir, awesomeGamesImagesPathInRepo)
	if err := syncImageFiles(logger, imagesSourceDir, storagePath); err != nil {
		return err
	}

	logger.Info().Msg("Sync completed successfully")
	return nil
}

// cloneRepository clones a git repository to the specified directory
func cloneRepository(logger zerolog.Logger, repoURL, destDir string) error {
	logger.Info().
		Str("repo", repoURL).
		Str("dest", destDir).
		Msg("Cloning repository")

	cmd := exec.Command("git", "clone", "--depth", "1", repoURL, destDir)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	logger.Info().
		Str("repo", repoURL).
		Msg("Repository cloned successfully")

	return nil
}

// syncJSONFiles copies and transforms JSON files from source to destination
func syncJSONFiles(logger zerolog.Logger, sourceDir, destDir string) error {
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

		if err := copyAndTransformJSONFile(srcPath, dstPath); err != nil {
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

// copyAndTransformJSONFile copies a JSON file and applies transformations
func copyAndTransformJSONFile(srcPath, dstPath string) error {
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

	// Transform image path to relative
	if imagePath, ok := content["image"].(string); ok {
		content["image"] = "./" + filepath.Base(imagePath)
	}

	// Normalize genres in german_content
	if germanContent, ok := content["german_content"].(map[string]interface{}); ok {
		if genres, ok := germanContent["genres"].([]interface{}); ok {
			germanContent["genres"] = normalizeGenres(genres)
		}
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

// normalizeGenres normalizes genre names to avoid duplicates
func normalizeGenres(genres []interface{}) []interface{} {
	result := make([]interface{}, 0, len(genres))
	hasMMO := false
	hasSimulationen := false
	hasMassivelyMultiplayer := false
	hasSimulation := false

	// First pass: detect what we have
	for _, g := range genres {
		genre, ok := g.(string)
		if !ok {
			result = append(result, g)
			continue
		}

		switch genre {
		case "MMO":
			hasMMO = true
		case "Massively Multiplayer":
			hasMassivelyMultiplayer = true
			result = append(result, genre)
		case "Simulationen":
			hasSimulationen = true
		case "Simulation":
			hasSimulation = true
			result = append(result, genre)
		default:
			result = append(result, genre)
		}
	}

	// Second pass: add normalized versions if needed
	if hasMMO && !hasMassivelyMultiplayer {
		result = append(result, "Massively Multiplayer")
	}
	if hasSimulationen && !hasSimulation {
		result = append(result, "Simulation")
	}

	return result
}

// syncImageFiles copies image files from source to destination
func syncImageFiles(logger zerolog.Logger, sourceDir, destDir string) error {
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

		logger.Info().
			Str("file", entry.Name()).
			Msg("Copied image file")
	}

	return nil
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := sourceFile.Close(); err != nil {
			fmt.Println("error when closing source file:", err)
		}
	}()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if err := destFile.Close(); err != nil {
			fmt.Println("error when closing dest file:", err)
		}
	}()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
