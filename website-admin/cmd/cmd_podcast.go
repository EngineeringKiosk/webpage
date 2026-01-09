package cmd

import (
	"github.com/spf13/cobra"
)

// podcastCmd represents the podcast command
var podcastCmd = &cobra.Command{
	Use:   "podcast",
	Short: "Manage podcast episodes and related configuration",
	Long: `Manage podcast episodes and related configuration for the Engineering Kiosk podcast.

The Engineering Kiosk is a German-language software engineering podcast. This command
group provides tools for automating various podcast management tasks:

  - Syncing episode data from the RedCircle RSS feed
  - Generating Netlify URL redirects for short episode links
  - Validating that all episodes have player URLs configured

Episode files are stored as Markdown with YAML frontmatter in src/content/podcast/
and contain metadata such as:
  - Audio file URLs
  - Chapter markers with timestamps
  - Player links (Spotify, Apple Podcasts, Amazon Music, Deezer, YouTube)
  - Episode descriptions and transcripts
  - Tags and speaker information

The podcast is hosted on RedCircle and distributed across major podcast platforms.`,
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(podcastCmd)
}
