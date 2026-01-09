package cmd

import (
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Manage content tags and their SEO descriptions",
	Long: `Manage content tags and their SEO descriptions across the Engineering Kiosk website.

Tags are used throughout the website to categorize and organize content. Each tag should
have both a short and long SEO description to improve search engine visibility and provide
context to visitors browsing by tag.

This command group provides tools to:
  - Find tags that are missing SEO descriptions
  - Update tag description files with newly discovered tags
  - Track tag usage counts across different content types

Supported content collections:
  - website-content: Tags from podcast episodes and blog posts (src/data/tags.json)
  - german-tech-podcasts: Tags for the German tech podcasts directory (src/data/german-tech-podcasts-tags.json)
  - awesome-software-engineering-games: Genres for the games catalog (src/data/awesome-software-engineering-games-genres.json)

Each tag entry in the JSON files contains:
  - short_desc: A brief SEO-friendly description (for meta tags)
  - long_desc: A detailed description (for tag landing pages)
  - usage_count: How many content items use this tag`,
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
