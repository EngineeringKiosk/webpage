package cmd

import (
	"github.com/spf13/cobra"
)

// podcastCmd represents the podcast command
var podcastCmd = &cobra.Command{
	Use:   "podcast",
	Short: "Podcast management commands",
	Long: `Podcast management commands for handling various podcast-related tasks
such as creating redirects and other podcast operations.`,
}

func init() {
	rootCmd.AddCommand(podcastCmd)
}
