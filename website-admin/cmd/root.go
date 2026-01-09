package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "website-admin",
	Short: "Administrative CLI tool for the Engineering Kiosk website",
	Long: `website-admin is a command-line tool for managing the Engineering Kiosk website.

This tool provides various administrative commands for:
  - Managing podcast episodes (syncing from RSS, generating redirects, checking player URLs)
  - Handling content tags and SEO descriptions
  - Synchronizing external data sources (German Tech Podcasts, Awesome Software Engineering Games)
  - Validating environment configuration

The Engineering Kiosk website (engineeringkiosk.dev) is a German-language software
engineering podcast website built with Astro. This admin tool automates many of the
content management and deployment tasks required to keep the site up-to-date.

Use "website-admin [command] --help" for more information about a command.`,
	DisableAutoGenTag: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&flagDebugLogging, "debug", "d", false, "Enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&flagDisableLogging, "disable-logging", "D", false, "Disable all logging output")

	// Load .env file if it exists - this must happen before flag initialization
	// Only log if the file exists but fails to load (don't log if file doesn't exist)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
	}
}
