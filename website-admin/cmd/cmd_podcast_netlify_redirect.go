package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"

	"github.com/EngineeringKiosk/website/website-admin/utils"
)

// episodeNumberRegex matches episode numbers at the start of filenames.
// It captures only digits (not dashes) to correctly handle filenames like "170-404-not-found.md"
// where we want to extract "170", not "170-404".
var episodeNumberRegex = regexp.MustCompile(`^(\d+)-`)

// RedirectExistence tracks which redirect types exist for an episode.
type RedirectExistence struct {
	EpisodesRedirect bool // /episodes/{n} redirect exists
	EpRedirect       bool // /ep{n} redirect exists
}

// BothExist returns true if both redirect types exist for the episode.
func (r RedirectExistence) BothExist() bool {
	return r.EpisodesRedirect && r.EpRedirect
}

// CheckBothRedirectsExist checks if both redirect types (/episodes/{n} and /ep{n})
// exist for a given episode number in the redirect existence map.
// Returns true only if both redirects are present.
func CheckBothRedirectsExist(redirectExistence map[string]RedirectExistence, episodeNumber string) bool {
	existence, found := redirectExistence[episodeNumber]
	if !found {
		return false
	}
	return existence.BothExist()
}

// ExtractEpisodeNumber extracts the episode number from a podcast episode filename.
// It returns the episode number as a string with leading zeros removed.
// If no valid episode number is found, it returns an empty string and false.
//
// Examples:
//   - "42-clean-code.md" -> "42", true
//   - "042-clean-code.md" -> "42", true
//   - "170-404-not-found.md" -> "170", true
//   - "-1-special-episode.md" -> "", false
//   - "invalid.md" -> "", false
func ExtractEpisodeNumber(filename string) (string, bool) {
	matches := episodeNumberRegex.FindStringSubmatch(filename)
	if len(matches) < 2 {
		return "", false
	}

	episodeNumber := matches[1]

	// Remove leading zeros
	episodeNumber = strings.TrimLeft(episodeNumber, "0")

	// Handle the case where the number was all zeros (e.g., "00-episode.md")
	if episodeNumber == "" {
		episodeNumber = "0"
	}

	return episodeNumber, true
}

// NetlifyConfig represents the netlify.toml structure
type NetlifyConfig struct {
	Redirects []Redirect `toml:"redirects"`
	Headers   []Header   `toml:"headers"`
	Build     Build      `toml:"build"`
}

type Redirect struct {
	From   string `toml:"from"`
	To     string `toml:"to"`
	Status int    `toml:"status,omitempty"`
	Force  bool   `toml:"force,omitempty"`
}

type Header struct {
	For    string       `toml:"for"`
	Values HeaderValues `toml:"headers"`
}

type HeaderValues struct {
	ContentType        string `toml:"Content-Type,omitempty"`
	ContentDisposition string `toml:"Content-Disposition,omitempty"`
}

type Build struct {
	Publish string `toml:"publish"`
	Command string `toml:"command"`
}

// podcastNetlifyRedirectCmd represents the netlify-redirect command
var podcastNetlifyRedirectCmd = &cobra.Command{
	Use:   "netlify-redirect",
	Short: "Generate short URL redirects for podcast episodes in netlify.toml",
	Long: `Generate short URL redirects for podcast episodes in netlify.toml.

This command scans the podcast episodes directory and creates user-friendly short URL
redirects for each episode. These short URLs make it easy to share episode links
verbally or in print.

For each episode, two redirect entries are created:
  /episodes/{n}  ->  /podcast/episode/{full-slug}?pkn=shortlink
  /ep{n}         ->  /podcast/episode/{full-slug}?pkn=shortlink

Where {n} is the episode number (e.g., 42) and {full-slug} is the episode filename
without the .md extension (e.g., "42-clean-code-prinzipien").

The redirects use HTTP 301 (permanent redirect) status codes with force=true to
ensure they work even if a file exists at that path.

Behavior:
  - Reads the existing netlify.toml file and preserves all other configuration
  - Skips episodes that already have redirects configured
  - Episode numbers are extracted from filenames (e.g., "042-title.md" -> 42)
  - Leading zeros are stripped from episode numbers
  - The ?pkn=shortlink query parameter helps track traffic from short links`,
	Example: `  # Generate redirects using default paths (run from project root)
  website-admin podcast netlify-redirect

  # Specify custom paths for all options
  website-admin podcast netlify-redirect \
    --toml-file ./netlify.toml \
    --episodes-dir ./src/content/podcast \
    --redirect-prefix /episodes/

  # Use environment variables instead of flags
  export WEBSITEADMIN_NETLIFY_REDIRECT_TOML_FILE=./netlify.toml
  export WEBSITEADMIN_NETLIFY_REDIRECT_EPISODES_DIR=./src/content/podcast
  website-admin podcast netlify-redirect

  # Enable debug logging to see each redirect being processed
  website-admin podcast netlify-redirect --debug`,
	RunE:              RunPodcastNetlifyRedirectCmd,
	DisableAutoGenTag: true,
}

func init() {
	podcastCmd.AddCommand(podcastNetlifyRedirectCmd)

	podcastNetlifyRedirectCmd.Flags().StringVarP(&flagNetlifyRedirectTomlFile, "toml-file", "t", os.Getenv(EnvVarNetlifyRedirectTomlFile), environmentVariables[EnvVarNetlifyRedirectTomlFile])
	podcastNetlifyRedirectCmd.Flags().StringVarP(&flagEpisodesStorePath, "episodes-dir", "e", os.Getenv(EnvVarEpisodesStorePath), environmentVariables[EnvVarEpisodesStorePath])
	podcastNetlifyRedirectCmd.Flags().StringVarP(&flagNetlifyRedirectRedirectPrefix, "redirect-prefix", "r", os.Getenv(EnvVarNetlifyRedirectRedirectPrefix), environmentVariables[EnvVarNetlifyRedirectRedirectPrefix])
}

func RunPodcastNetlifyRedirectCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Msg("starting")

	logger.Info().
		Str("tomlFile", flagNetlifyRedirectTomlFile).
		Str("episodesDir", flagEpisodesStorePath).
		Msg("Generating podcast netlify redirects")

	// Adjust paths
	fileToParse := utils.BuildCorrectFilePath(flagNetlifyRedirectTomlFile)
	flagEpisodesStorePath = utils.BuildCorrectFilePath(flagEpisodesStorePath)

	// Read existing TOML file
	var config NetlifyConfig
	tomlData, err := os.ReadFile(fileToParse)
	if err != nil {
		return fmt.Errorf("failed to read TOML file: %w", err)
	}
	if err := toml.Unmarshal(tomlData, &config); err != nil {
		return fmt.Errorf("failed to parse TOML file: %w", err)
	}

	// Build redirect existence map for easy lookup
	// Track both /episodes/{n} and /ep{n} redirects separately
	redirectExistence := make(map[string]RedirectExistence)
	episodesRedirectRegex := regexp.MustCompile(fmt.Sprintf(`^%s(\d+)$`, regexp.QuoteMeta(flagNetlifyRedirectRedirectPrefix)))
	epRedirectRegex := regexp.MustCompile(`^/ep(\d+)$`)

	for _, redirect := range config.Redirects {
		// Check for /episodes/{n} pattern
		if matches := episodesRedirectRegex.FindStringSubmatch(redirect.From); len(matches) > 1 {
			episodeNumber := matches[1]
			existence := redirectExistence[episodeNumber]
			existence.EpisodesRedirect = true
			redirectExistence[episodeNumber] = existence
		}
		// Check for /ep{n} pattern
		if matches := epRedirectRegex.FindStringSubmatch(redirect.From); len(matches) > 1 {
			episodeNumber := matches[1]
			existence := redirectExistence[episodeNumber]
			existence.EpRedirect = true
			redirectExistence[episodeNumber] = existence
		}
	}

	// Get existing podcast episodes
	entries, err := os.ReadDir(flagEpisodesStorePath)
	if err != nil {
		return fmt.Errorf("failed to read episodes directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		episodeName := entry.Name()

		// Extract episode number from filename
		episodeNumber, ok := ExtractEpisodeNumber(episodeName)
		if !ok {
			continue
		}

		// Check if both redirects already exist
		if CheckBothRedirectsExist(redirectExistence, episodeNumber) {
			logger.Debug().
				Str("episodeNumber", episodeNumber).
				Msg("Skipping redirect processing for episode: Both redirects exist already")
			continue
		}

		// Build episode file path (without .md extension)
		episodeFile := strings.TrimSuffix(episodeName, ".md")
		existence := redirectExistence[episodeNumber]

		// Add /episodes/{n} redirect if it doesn't exist
		if !existence.EpisodesRedirect {
			logger.Info().
				Str("episodeNumber", episodeNumber).
				Str("redirectType", "episodes").
				Msg("Adding redirect for episode")

			newRedirectShortlink := Redirect{
				From:   fmt.Sprintf("%s%s", flagNetlifyRedirectRedirectPrefix, episodeNumber),
				To:     fmt.Sprintf("/podcast/episode/%s?pkn=shortlink", episodeFile),
				Status: 301,
				Force:  true,
			}
			config.Redirects = append(config.Redirects, newRedirectShortlink)
		}

		// Add /ep{n} redirect if it doesn't exist
		if !existence.EpRedirect {
			logger.Info().
				Str("episodeNumber", episodeNumber).
				Str("redirectType", "ep").
				Msg("Adding redirect for episode")

			newRedirectEpisodeShortlink := Redirect{
				From:   fmt.Sprintf("/ep%s", episodeNumber),
				To:     fmt.Sprintf("/podcast/episode/%s?pkn=shortlink", episodeFile),
				Status: 301,
				Force:  true,
			}
			config.Redirects = append(config.Redirects, newRedirectEpisodeShortlink)
		}
	}

	// Write updated TOML file
	outFile, err := os.Create(fileToParse)
	if err != nil {
		return fmt.Errorf("failed to create TOML file: %w", err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()

	encoder := toml.NewEncoder(outFile)
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to write TOML file: %w", err)
	}

	logger.Info().
		Msg("Generating podcast netlify redirects ... Successful")
	return nil
}
