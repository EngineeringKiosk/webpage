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
	Short: "Validate that all podcast episodes have player URLs configured",
	Long: `Validate that all podcast episodes have their player URLs properly configured.

Podcast episodes need links to various streaming platforms so listeners can choose
their preferred player. This command scans all episode Markdown files and checks
that the following frontmatter fields contain non-empty URLs:

  - spotify:        Link to the episode on Spotify
  - apple_podcasts: Link to the episode on Apple Podcasts
  - amazon_music:   Link to the episode on Amazon Music
  - deezer:         Link to the episode on Deezer
  - youtube:        Link to the episode on YouTube

This is typically run after syncing new episodes from the RSS feed, as the RSS
feed doesn't include platform-specific URLs. These URLs usually need to be added
manually after each platform processes the new episode (which can take a few hours
to a day after publication).

Behavior:
  - Scans all .md files in the episodes directory
  - Reports each episode with missing URLs and which platforms are missing
  - Exits with code 0 if all episodes have complete URLs
  - Exits with code 1 if any URLs are missing (CI-friendly)

This command is useful for:
  - CI/CD pipelines to ensure episode completeness before deployment
  - Identifying episodes that need manual URL updates
  - Quality assurance after RSS feed syncs`,
	Example: `  # Check player URLs using default episodes directory
  website-admin podcast check-player-urls

  # Specify a custom episodes directory
  website-admin podcast check-player-urls --episodes-dir ./src/content/podcast

  # Enable debug logging to see each file being checked
  website-admin podcast check-player-urls --debug

  # Use in CI pipeline (exits with code 1 if any URLs missing)
  website-admin podcast check-player-urls || echo "Some episodes need player URLs!"`,
	RunE:              RunPodcastCheckPlayerURLsCmd,
	DisableAutoGenTag: true,
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
