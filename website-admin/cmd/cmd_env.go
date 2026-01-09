package cmd

import (
	"os"

	"github.com/EngineeringKiosk/website/website-admin/utils"
	"github.com/spf13/cobra"
)

const (
	EnvVarEpisodesStorePath             = "WEBSITEADMIN_EPISODES_STORE_PATH"
	EnvVarRssFeedUrl                    = "WEBSITEADMIN_RSS_FEED_URL"
	EnvVarTranscriptPath                = "WEBSITEADMIN_TRANSCRIPT_PATH"
	EnvVarImagesPath                    = "WEBSITEADMIN_IMAGES_PATH"
	EnvVarNetlifyRedirectTomlFile       = "WEBSITEADMIN_NETLIFY_REDIRECT_TOML_FILE"
	EnvVarNetlifyRedirectEpisodesDir    = "WEBSITEADMIN_NETLIFY_REDIRECT_EPISODES_DIR"
	EnvVarNetlifyRedirectRedirectPrefix = "WEBSITEADMIN_NETLIFY_REDIRECT_REDIRECT_PREFIX"
)

var environmentVariables = map[string]string{
	EnvVarEpisodesStorePath:             "Path to the Episode YAML storage",
	EnvVarRssFeedUrl:                    "URL of the RSS Feed",
	EnvVarTranscriptPath:                "Path to store podcast transcripts",
	EnvVarImagesPath:                    "Path to store podcast images",
	EnvVarNetlifyRedirectTomlFile:       "Path to netlify.toml file for redirects",
	EnvVarNetlifyRedirectEpisodesDir:    "Directory containing episode markdown files",
	EnvVarNetlifyRedirectRedirectPrefix: "Prefix for redirect URLs",
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Display and validate environment variable configuration",
	Long: `Display and validate environment variable configuration for website-admin.

Several website-admin commands can be configured via environment variables instead of
(or in addition to) command-line flags. This command helps you verify that all the
environment variables you need are properly set before running other commands.

Checked environment variables:
  WEBSITEADMIN_EPISODES_STORE_PATH        Path to store episode Markdown files
  WEBSITEADMIN_RSS_FEED_URL               URL of the podcast RSS feed
  WEBSITEADMIN_TRANSCRIPT_PATH            Path to store podcast transcripts
  WEBSITEADMIN_IMAGES_PATH                Path to store podcast images
  WEBSITEADMIN_NETLIFY_REDIRECT_TOML_FILE Path to netlify.toml for redirects
  WEBSITEADMIN_NETLIFY_REDIRECT_EPISODES_DIR  Directory containing episode files
  WEBSITEADMIN_NETLIFY_REDIRECT_REDIRECT_PREFIX  Prefix for redirect URLs

Behavior:
  - Logs INFO for each environment variable that is set
  - Logs WARN for each environment variable that is not set
  - Always exits with code 0 (missing variables are warnings, not errors)
  - Does NOT validate whether the paths exist or URLs are valid

You can also use a .env file in the current directory. The tool automatically loads
it at startup if present.

This command is useful for:
  - Debugging configuration issues
  - Verifying CI/CD environment setup
  - Documenting which variables are available`,
	Example: `  # Check all environment variables
  website-admin env

  # Check with debug logging for more details
  website-admin env --debug

  # Typical workflow: check env, then run a command
  website-admin env && website-admin podcast sync-from-rss`,
	RunE:              RunEnvCmd,
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(envCmd)
}

func RunEnvCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Msg("starting")

	for envVar, envDescription := range environmentVariables {
		value := os.Getenv(envVar)
		if len(value) == 0 {
			logger.Warn().
				Str("env_var", envVar).
				Str("description", envDescription).
				Msg("environment variable is not set")
			continue
		}

		logger.Info().
			Str("env_var", envVar).
			Str("description", envDescription).
			Msg("environment variable set")
	}
	return nil
}
