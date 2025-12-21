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
	Short: "Generate URL redirects for podcast episodes in netlify.toml",
	Long: `Generate URL redirects for podcast episodes in netlify.toml.

This command scans the podcast episodes directory and creates short URL redirects
for each episode in the netlify.toml configuration file. Two redirects are created
per episode:
  - /episodes/{n} -> /podcast/episode/{slug}
  - /ep{n} -> /podcast/episode/{slug}`,
	RunE: RunPodcastNetlifyRedirectCmd,
}

func init() {
	podcastCmd.AddCommand(podcastNetlifyRedirectCmd)

	// TODO Unify parameters
	podcastNetlifyRedirectCmd.Flags().StringVarP(&flagNetlifyRedirectTomlFile, "toml-file", "t", os.Getenv(EnvVarNetlifyRedirectTomlFile), environmentVariables[EnvVarNetlifyRedirectTomlFile])
	podcastNetlifyRedirectCmd.Flags().StringVarP(&flagNetlifyRedirectEpisodesDir, "episodes-dir", "e", os.Getenv(EnvVarNetlifyRedirectEpisodesDir), environmentVariables[EnvVarNetlifyRedirectEpisodesDir])
	podcastNetlifyRedirectCmd.Flags().StringVarP(&flagNetlifyRedirectRedirectPrefix, "redirect-prefix", "r", os.Getenv(EnvVarNetlifyRedirectRedirectPrefix), environmentVariables[EnvVarNetlifyRedirectRedirectPrefix])
}

func RunPodcastNetlifyRedirectCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Msg("starting")

	logger.Info().
		Str("tomlFile", flagNetlifyRedirectTomlFile).
		Str("episodesDir", flagNetlifyRedirectEpisodesDir).
		Msg("Generating podcast netlify redirects")

	// Adjust paths
	fileToParse := utils.BuildCorrectFilePath(flagNetlifyRedirectTomlFile)
	flagNetlifyRedirectEpisodesDir = utils.BuildCorrectFilePath(flagNetlifyRedirectEpisodesDir)

	// Read existing TOML file
	var config NetlifyConfig
	tomlData, err := os.ReadFile(fileToParse)
	if err != nil {
		return fmt.Errorf("failed to read TOML file: %w", err)
	}
	if err := toml.Unmarshal(tomlData, &config); err != nil {
		return fmt.Errorf("failed to parse TOML file: %w", err)
	}

	// Build redirect map for easy lookup
	redirectMap := make(map[string]Redirect)
	redirectEpisodeNumberRegex := regexp.MustCompile(fmt.Sprintf(`^%s([-\d]*)$`, regexp.QuoteMeta(flagNetlifyRedirectRedirectPrefix)))

	for _, redirect := range config.Redirects {
		matches := redirectEpisodeNumberRegex.FindStringSubmatch(redirect.From)
		if len(matches) > 1 {
			episodeNumber := matches[1]
			redirectMap[episodeNumber] = redirect
		}
	}

	// Get existing podcast episodes
	entries, err := os.ReadDir(flagNetlifyRedirectEpisodesDir)
	if err != nil {
		return fmt.Errorf("failed to read episodes directory: %w", err)
	}

	episodeNumberRegex := regexp.MustCompile(`^([-\d]*)-`)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		episodeName := entry.Name()

		// Extract episode number
		// TODO Fix bug, there is an episode with number episodeNumber=170-404
		// I would expect only to get 170, not 170-404
		matches := episodeNumberRegex.FindStringSubmatch(episodeName)
		if len(matches) < 2 {
			continue
		}

		episodeNumber := matches[1]

		// Remove leading zero if present
		episodeNumber = strings.TrimPrefix(episodeNumber, "0")

		// Check if redirect already exists
		// TODO Fix bug: This check currently only works for /episodes/{n} redirects, not /ep{n} (or vice versa)
		// It only checks if one of those exists, not both
		if _, exists := redirectMap[episodeNumber]; exists {
			logger.Debug().
				Str("episodeNumber", episodeNumber).
				Msg("Skipping redirect processing for episode: Redirect exists already")
			continue
		}

		logger.Info().
			Str("episodeNumber", episodeNumber).
			Msg("Processing redirect for episode")

		// Build episode file path (without .md extension)
		episodeFile := strings.TrimSuffix(episodeName, ".md")

		// Add /episodes/{n} redirect
		newRedirectShortlink := Redirect{
			From:   fmt.Sprintf("%s%s", flagNetlifyRedirectRedirectPrefix, episodeNumber),
			To:     fmt.Sprintf("/podcast/episode/%s?pkn=shortlink", episodeFile),
			Status: 301,
			Force:  true,
		}
		config.Redirects = append(config.Redirects, newRedirectShortlink)

		// Add /ep{n} redirect
		newRedirectEpisodeShortlink := Redirect{
			From:   fmt.Sprintf("/ep%s", episodeNumber),
			To:     fmt.Sprintf("/podcast/episode/%s?pkn=shortlink", episodeFile),
			Status: 301,
			Force:  true,
		}
		config.Redirects = append(config.Redirects, newRedirectEpisodeShortlink)
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
