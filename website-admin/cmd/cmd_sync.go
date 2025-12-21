package cmd

import (
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync external data sources",
	Long: `Sync external data sources into the website content.

This command group handles synchronization of data from external repositories
such as the German Tech Podcasts and Awesome Software Engineering Games.`,
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
