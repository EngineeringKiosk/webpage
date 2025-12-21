package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/EngineeringKiosk/website/website-admin/utils"
)

const (
	ModeWebsiteContent                  = "website-content"
	ModeGermanTechPodcasts              = "german-tech-podcasts"
	ModeAwesomeSoftwareEngineeringGames = "awesome-software-engineering-games"
)

// TagDescription represents a tag with its SEO descriptions
type TagDescription struct {
	ShortDesc  string `json:"short_desc"`
	LongDesc   string `json:"long_desc"`
	UsageCount int    `json:"usage_count"`
}

// tagsFindCmd represents the tags find command
var tagsFindCmd = &cobra.Command{
	Use:   "find [mode]",
	Short: "Find tags without descriptions",
	Long: `Find tags or genres without SEO descriptions in content files.

Modes:
  website-content              			Scan podcast episodes and blog posts for tags
  german-tech-podcasts         			Scan German tech podcasts for tags
  awesome-software-engineering-games  	Scan software engineering games for genres

Without --write-file flag, tags are printed to stdout and exits with code 1 if any are found.
With --write-file flag, missing tags are added to the respective JSON file with empty descriptions.`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{ModeWebsiteContent, ModeGermanTechPodcasts, ModeAwesomeSoftwareEngineeringGames},
	RunE:      RunTagsFindCmd,
}

func init() {
	tagsCmd.AddCommand(tagsFindCmd)

	// TODO Add the flags as the same structure as other commands (env.go)
	tagsFindCmd.Flags().BoolVarP(&flagTagsWriteFile, "write-file", "w", false, "Modify the local tag storage file and add missing tags")
	tagsFindCmd.Flags().StringVarP(&flagTagsContentDir, "content-dir", "c", "", "Content directory to scan (overrides default for mode)")
	tagsFindCmd.Flags().StringVarP(&flagTagsDescFile, "desc-file", "f", "", "Tag description file path (overrides default for mode)")
}

func RunTagsFindCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Str("mode", args[0]).
		Msg("starting")

	mode := args[0]

	// Determine paths and JSON key based on mode
	var contentPaths []string
	var tagFilePath string
	var jsonKey string

	switch mode {
	case ModeWebsiteContent:
		tagFilePath = "src/data/tags.json"
		contentPaths = []string{"src/content/podcast"}
		jsonKey = "tags"

	case ModeGermanTechPodcasts:
		tagFilePath = "src/data/german-tech-podcasts-tags.json"
		contentPaths = []string{"src/content/germantechpodcasts"}
		jsonKey = "tags"

	case ModeAwesomeSoftwareEngineeringGames:
		tagFilePath = "src/data/awesome-software-engineering-games-genres.json"
		contentPaths = []string{"src/content/awesome-software-engineering-games"}
		jsonKey = "german_content.genres"

	default:
		return fmt.Errorf("invalid mode: %s", mode)
	}

	// Allow overriding paths via flags
	if flagTagsContentDir != "" {
		contentPaths = []string{flagTagsContentDir}
	}
	if flagTagsDescFile != "" {
		tagFilePath = flagTagsDescFile
	}

	// Adjust paths
	tagFilePath = utils.BuildCorrectFilePath(tagFilePath)
	for i := range contentPaths {
		contentPaths[i] = utils.BuildCorrectFilePath(contentPaths[i])
	}

	// Read all tags from content files
	logger.Info().Msg("Reading tags from content files")
	tags, err := readAllTagsFromContentFiles(contentPaths, jsonKey, logger)
	if err != nil {
		return fmt.Errorf("failed to read tags from content files: %w", err)
	}
	lenContentTags := len(tags)

	// Read existing tag descriptions
	logger.Info().
		Str("file", tagFilePath).
		Msg("Reading existing tag descriptions")
	tagDescriptions, err := readTagDescriptions(tagFilePath)
	if err != nil {
		return fmt.Errorf("failed to read tag descriptions: %w", err)
	}

	// Determine tags without descriptions
	logger.Info().
		Int("uniqueTags", lenContentTags).
		Msg("Determining tags with missing descriptions")
	tagsWithoutDesc := getAllTagsWithoutDescription(tagDescriptions, tags)

	if len(tagsWithoutDesc) == 0 {
		logger.Info().
			Int("missingDescriptions", len(tagsWithoutDesc)).
			Int("uniqueTags", lenContentTags).
			Msg("All tags have descriptions. Nothing to do.")
		return nil
	}

	// Should we just print or write to file?
	if !flagTagsWriteFile {
		for tag := range tagsWithoutDesc {
			logger.Warn().
				Str("tag", tag).
				Msg("Tag missing description")
		}

		// Exit with error if there are missing descriptions
		// This is useful for CI/CD pipelines
		if len(tagsWithoutDesc) > 0 {
			logger.Error().
				Int("missingDescriptions", len(tagsWithoutDesc)).
				Msg("Found tags with missing descriptions")
			return fmt.Errorf("found %d tags with missing descriptions", len(tagsWithoutDesc))
		}

		return nil
	}

	// Write mode: update the tag file
	logger.Info().
		Str("file", tagFilePath).
		Msg("Writing missing tag structures to file")

	// Add missing tags
	for tag := range tagsWithoutDesc {
		if _, exists := tagDescriptions[tag]; !exists {
			tagDescriptions[tag] = &TagDescription{}
		}
	}

	// Update usage counts and remove unused tags
	for tag := range tagDescriptions {
		usageCount := 0
		if count, exists := tags[tag]; exists {
			usageCount = count
		}

		// Delete tags with no content
		if usageCount == 0 {
			delete(tagDescriptions, tag)
			continue
		}

		tagDescriptions[tag].UsageCount = usageCount
	}

	// Write the file
	data, err := json.MarshalIndent(tagDescriptions, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal tag descriptions: %w", err)
	}

	if err := os.WriteFile(tagFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write tag file: %w", err)
	}

	logger.Info().
		Str("file", tagFilePath).
		Msg("Writing missing tag structures to file ... done")

	return nil
}

// readAllTagsFromContentFiles reads tags from markdown/mdx and JSON files
func readAllTagsFromContentFiles(paths []string, jsonKey string, logger zerolog.Logger) (map[string]int, error) {
	tags := make(map[string]int)

	for _, path := range paths {
		logger.Info().
			Str("path", path).
			Msg("Reading tags from files")

		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read directory %s: %w", path, err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			fileName := entry.Name()
			fullPath := filepath.Join(path, fileName)

			// Process Markdown/MDX files
			if strings.HasSuffix(fileName, ".md") || strings.HasSuffix(fileName, ".mdx") {
				logger.Debug().
					Str("file", fullPath).
					Msg("Processing markdown file")

				fileTags, err := readTagsFromMarkdownFile(fullPath, jsonKey)
				if err != nil {
					logger.Warn().
						Err(err).
						Str("file", fullPath).
						Msg("Failed to read tags from markdown file")
					continue
				}

				for _, tag := range fileTags {
					tags[tag]++
				}
			}

			// Process JSON files
			if strings.HasSuffix(fileName, ".json") {
				logger.Debug().
					Str("file", fullPath).
					Msg("Processing JSON file")

				fileTags, err := readTagsFromJSONFile(fullPath, jsonKey)
				if err != nil {
					logger.Warn().
						Err(err).
						Str("file", fullPath).
						Msg("Failed to read tags from JSON file")
					continue
				}

				for _, tag := range fileTags {
					tags[tag]++
				}
			}
		}
	}

	return tags, nil
}

// readTagsFromMarkdownFile extracts tags from a markdown file's frontmatter
func readTagsFromMarkdownFile(filePath, jsonKey string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matter struct {
		Tags interface{} `yaml:"tags"`
	}

	_, err = frontmatter.Parse(file, &matter)
	if err != nil {
		return nil, err
	}

	// Convert to string slice
	return convertToStringSlice(matter.Tags), nil
}

// readTagsFromJSONFile extracts tags from a JSON file
func readTagsFromJSONFile(filePath, jsonKey string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var content map[string]interface{}
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}

	// Navigate nested JSON keys (e.g., "german_content.genres")
	parts := strings.Split(jsonKey, ".")
	var current interface{} = content
	for _, part := range parts {
		if m, ok := current.(map[string]interface{}); ok {
			current = m[part]
		} else {
			return nil, nil // Key path not found
		}
	}

	return convertToStringSlice(current), nil
}

// convertToStringSlice converts interface{} to []string
func convertToStringSlice(data interface{}) []string {
	if data == nil {
		return nil
	}

	switch v := data.(type) {
	case []interface{}:
		result := make([]string, 0, len(v))
		for _, item := range v {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	case []string:
		return v
	default:
		return nil
	}
}

// readTagDescriptions reads the tag description JSON file
func readTagDescriptions(filePath string) (map[string]*TagDescription, error) {
	// If file doesn't exist, return empty map
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return make(map[string]*TagDescription), nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	descriptions := make(map[string]*TagDescription)
	if err := json.Unmarshal(data, &descriptions); err != nil {
		return nil, err
	}

	return descriptions, nil
}

// getAllTagsWithoutDescription returns tags that don't have complete descriptions
func getAllTagsWithoutDescription(tagDescriptions map[string]*TagDescription, tags map[string]int) map[string]bool {
	tagsWithoutDesc := make(map[string]bool)

	// Check existing tags in description file
	for tagName, desc := range tagDescriptions {
		// If keys are missing or empty
		if desc.ShortDesc == "" || desc.LongDesc == "" {
			tagsWithoutDesc[tagName] = true
		}
	}

	// Check for new tags not in description file
	for tagName := range tags {
		if _, exists := tagDescriptions[tagName]; !exists {
			tagsWithoutDesc[tagName] = true
		}
	}

	return tagsWithoutDesc
}
