package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/EngineeringKiosk/website/website-admin/episode"
	"github.com/EngineeringKiosk/website/website-admin/utils"
)

// podcastCheckPlayerURLsCmd represents the check-player-urls command
var podcastCheckPlayerURLsCmd = &cobra.Command{
	Use:   "check-player-urls",
	Short: "Check if all podcast episodes have player URLs set",
	Long: `Check if all podcast episodes have their player URLs properly set.

This command iterates through all podcast episode files and verifies that
the following player links are filled:
  - spotify
  - apple_podcasts
  - amazon_music
  - deezer
  - youtube

If any player URL is missing, the episode and missing fields are reported.
The command exits with code 1 if any URLs are missing (useful for CI/CD).`,
	RunE: RunPodcastCheckPlayerURLsCmd,
}

func init() {
	podcastCmd.AddCommand(podcastCheckPlayerURLsCmd)

	podcastCheckPlayerURLsCmd.Flags().StringVarP(&flagPodcastEpisodesDir, "episodes-dir", "e", os.Getenv(EnvVarNetlifyRedirectEpisodesDir), environmentVariables[EnvVarNetlifyRedirectEpisodesDir])
}

func RunPodcastCheckPlayerURLsCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Msg("starting")

	// Set default if not provided
	episodesDir := flagPodcastEpisodesDir
	if episodesDir == "" {
		episodesDir = "src/content/podcast"
	}

	// Adjust path based on working directory
	episodesDir = utils.BuildCorrectFilePath(episodesDir)

	logger.Info().
		Str("episodesDir", episodesDir).
		Msg("Checking player URLs for episodes")

	// Read all episode files
	entries, err := os.ReadDir(episodesDir)
	if err != nil {
		return fmt.Errorf("failed to read episodes directory: %w", err)
	}

	// Track missing URLs
	episodesWithMissingURLs := 0
	totalEpisodes := 0

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		totalEpisodes++
		episodePath := filepath.Join(episodesDir, entry.Name())

		logger.Debug().
			Str("file", episodePath).
			Msg("Processing episode")

		// Load episode
		ep, err := episode.NewEpisode(episodePath)
		if err != nil {
			logger.Warn().
				Err(err).
				Str("file", episodePath).
				Msg("Failed to load episode")
			continue
		}

		fm := ep.GetFrontmatter()

		// Check for missing player URLs
		missing := checkMissingPlayerURLs(fm)

		if len(missing) > 0 {
			episodesWithMissingURLs++
			missingURLs := strings.Join(missing, ", ")
			logger.Error().
				Str("file", entry.Name()).
				Str("missing", missingURLs).
				Msg("Episode is missing player URLs")
		}
	}

	logger.Info().
		Int("totalEpisodes", totalEpisodes).
		Int("episodesWithMissingURLs", episodesWithMissingURLs).
		Msg("Player URL check completed")

	if episodesWithMissingURLs > 0 {
		return fmt.Errorf("found %d episodes with missing player URLs", episodesWithMissingURLs)
	}

	logger.Info().Msg("All episodes have their player URLs set")
	return nil
}

// checkMissingPlayerURLs checks which player URLs are missing from the episode
func checkMissingPlayerURLs(fm episode.EpisodeFrontmatter) []string {
	var missing []string

	playerURLs := map[string]string{
		"spotify":        fm.Spotify,
		"apple_podcasts": fm.ApplePodcasts,
		"amazon_music":   fm.AmazonMusic,
		"deezer":         fm.Deezer,
		"youtube":        fm.YouTube,
	}

	for name, url := range playerURLs {
		if url == "" {
			missing = append(missing, name)
		}
	}

	return missing
}
