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
	Short: "Checks if all required environment variables are set.",
	Long: `Checks if all required environment variables are set.
	
	If any are missing, a warning is logged.
	It is advised to run this command before using other commands to ensure all necessary environment variables are configured.
	This command does not validate the content of the environment variables.`,
	RunE: RunEnvCmd,
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
