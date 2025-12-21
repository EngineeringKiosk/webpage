package cmd

import (
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Tag management commands",
	Long: `Tag management commands for finding and managing SEO descriptions
for tags used in podcast episodes, blog posts, German tech podcasts, and games.`,
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
