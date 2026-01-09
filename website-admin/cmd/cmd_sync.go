package cmd

import (
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize content from external GitHub repositories",
	Long: `Synchronize content from external GitHub repositories into the website.

The Engineering Kiosk website aggregates content from several community-maintained
GitHub repositories. This command group handles the synchronization of that external
data into the local content directories.

Supported external sources:
  - German Tech Podcasts: A curated directory of German-language technology podcasts
    Source: https://github.com/EngineeringKiosk/GermanTechPodcasts

  - Awesome Software Engineering Games: A catalog of games that teach programming
    and software engineering concepts
    Source: https://github.com/EngineeringKiosk/awesome-software-engineering-games

The sync process:
  1. Clones the external repository to a temporary directory
  2. Copies JSON data files and images to the local content directory
  3. Applies necessary transformations (path adjustments, image resizing)
  4. Cleans up the temporary clone

This keeps the website content up-to-date with the latest community contributions
without requiring manual file copying.`,
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
