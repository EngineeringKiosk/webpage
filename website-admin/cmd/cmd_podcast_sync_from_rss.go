package cmd

import (
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/EngineeringKiosk/website/website-admin/episode"
	"github.com/EngineeringKiosk/website/website-admin/utils"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

// RSS/iTunes namespace structures
type RSSFeed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	PubDate     string      `xml:"pubDate"`
	Enclosure   Enclosure   `xml:"enclosure"`
	Image       ITunesImage `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
}

type Enclosure struct {
	URL string `xml:"url,attr"`
}

type ITunesImage struct {
	Href string `xml:"href,attr"`
}

// podcastSyncFromRSSCmd represents the podcast command
var podcastSyncFromRSSCmd = &cobra.Command{
	Use:   "sync-from-rss",
	Short: "Sync podcast episodes from RSS feed",
	Long:  `Syncs podcast episodes from an RSS feed to local markdown files with frontmatter.`,
	RunE:  RunPodcastSyncFromRSSCmd,
}

func init() {
	podcastCmd.AddCommand(podcastSyncFromRSSCmd)

	// Add flags with defaults from Python constants
	podcastSyncFromRSSCmd.Flags().StringVarP(&flagRSSFeedURL, "rss-feed-url", "r", os.Getenv(EnvVarRssFeedUrl), environmentVariables[EnvVarRssFeedUrl])
	podcastSyncFromRSSCmd.Flags().StringVarP(&flagEpisodesStorePath, "episodes-dir", "s", os.Getenv(EnvVarEpisodesStorePath), environmentVariables[EnvVarEpisodesStorePath])
	podcastSyncFromRSSCmd.Flags().StringVarP(&flagImagesDir, "images-dir", "i", os.Getenv(EnvVarImagesPath), environmentVariables[EnvVarImagesPath])
	podcastSyncFromRSSCmd.Flags().StringVarP(&flagTranscriptDir, "transcript-dir", "t", os.Getenv(EnvVarTranscriptPath), environmentVariables[EnvVarTranscriptPath])
}

func RunPodcastSyncFromRSSCmd(cmd *cobra.Command, args []string) error {
	logger := utils.GetLogger(flagDisableLogging, flagDebugLogging)
	logger.Info().
		Str("cmd", cmd.Use).
		Msg("starting")

	logger.Info().
		Str("rssFeed", flagRSSFeedURL).
		Msg("Reading Podcast RSS Feed")

	// Adjust paths if needed
	flagEpisodesStorePath = utils.BuildCorrectFilePath(flagEpisodesStorePath)
	flagImagesDir = utils.BuildCorrectFilePath(flagImagesDir)

	// Fetch RSS feed
	resp, err := http.Get(flagRSSFeedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status from RSS feed: %s", resp.Status)
	}

	logger.Info().
		Str("rssFeed", flagRSSFeedURL).
		Msg("Reading Podcast RSS Feed ... Successful")

	// Parse XML
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read RSS feed: %w", err)
	}

	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	logger.Info().
		Int("itemCount", len(feed.Channel.Items)).
		Msg("Number of podcast items found")

	// Default speakers
	defaultSpeaker := []episode.Speaker{
		{Name: "Andy Grunwald"},
		{Name: "Wolfi Gassler"},
	}

	// Process each episode
	for _, item := range feed.Channel.Items {
		title := item.Title
		description := item.Description

		// Parse headlines and add IDs to h3 tags
		htmlContent, headlines := parseHeadlinesFromHTML(description)
		descriptionHTML := htmlContent

		// Clean up HTML content
		htmlContent = strings.TrimSpace(descriptionHTML)
		htmlContent = utils.RemoveRelNofollowFromInternalLinks(htmlContent)

		// Extract chapters
		chapters := getChapterFromDescription(description)

		// Convert chapters to episode.Chapter format
		episodeChapters := make([]episode.Chapter, len(chapters))
		for i, ch := range chapters {
			episodeChapters[i] = episode.Chapter{
				Start: ch["start"],
				Title: ch["title"],
			}
		}

		// Parse headlines into string format
		headlineInfoParts := make([]string, 0, len(headlines))
		for slug, headline := range headlines {
			headlineInfoParts = append(headlineInfoParts, fmt.Sprintf("%s::%s", slug, headline))
		}
		headlineInfo := strings.Join(headlineInfoParts, "||")

		// Get description as text only
		descriptionTextOnly := utils.RemoveHTMLTags(description)

		// Cut text after intro
		descriptionShort := descriptionTextOnly
		if strings.Contains(descriptionTextOnly, "Feedback an stehtisch") {
			parts := strings.Split(descriptionTextOnly, "Feedback an stehtisch@engineeringkiosk.dev")
			descriptionShort = parts[0]
		} else if strings.Contains(descriptionTextOnly, "Feedback (gerne") {
			parts := strings.Split(descriptionTextOnly, "Feedback (gerne")
			descriptionShort = parts[0]
		}

		descriptionShort = strings.TrimSpace(descriptionShort)
		descriptionShort = html.UnescapeString(descriptionShort)

		// Parse date
		// Date format: Tue, 05 Apr 2022 04:25:00 +0000
		pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			logger.Warn().
				Err(err).
				Str("episode", title).
				Str("pubDate", item.PubDate).
				Msg("failed to parse publication date for episode")
			// Use current time as fallback
			pubDate = time.Now()
		}

		// Get MP3 link
		mp3Link := item.Enclosure.URL

		// Handle episode image
		var imageFilename string
		if item.Image.Href != "" {
			imageURL := item.Image.Href
			sluggedTitle := utils.Slugify(title)
			imageBaseFilename := filepath.Join(flagImagesDir, sluggedTitle)

			// Check if image already exists
			if existingImage, exists := utils.ImageExists(imageBaseFilename); exists {
				imageFilename = fmt.Sprintf("./%s", filepath.Base(existingImage))
			} else {
				// Download image
				downloadedImage, err := utils.DownloadFile(imageURL, imageBaseFilename)
				if err != nil {
					logger.Warn().
						Err(err).
						Str("episode", title).
						Str("imageURL", imageURL).
						Msg("failed to download episode image")
				} else {
					// Resize image to 700x700
					if err := utils.ResizeImage(downloadedImage, 700, 700); err != nil {
						logger.Warn().
							Err(err).
							Str("episode", title).
							Str("imageFile", downloadedImage).
							Msg("failed to resize episode image")
					}
					imageFilename = fmt.Sprintf("./%s", filepath.Base(downloadedImage))
				}
			}
		}

		// Build filename
		filename := utils.Slugify(title) + ".md"
		episodeNumber, err := episode.GetEpisodeNumberFromFilename(filename, true)
		if err != nil {
			logger.Warn().
				Err(err).
				Str("filename", filename).
				Msg("failed to get episode number from filename")
			episodeNumber = "00"
		}

		// Get transcript path
		transcriptPath := episode.GetTranscriptSlimPathByEpisodeNumber(episodeNumber, flagTranscriptDir)

		// Build episode data
		data := episode.EpisodeFrontmatter{
			Advertiser:     "",
			AmazonMusic:    "",
			ApplePodcasts:  "",
			Audio:          mp3Link,
			Chapter:        episodeChapters,
			Description:    descriptionShort,
			Headlines:      headlineInfo,
			Image:          imageFilename,
			LengthSecond:   0,
			PubDate:        episode.FrontmatterPubDateTime{Time: pubDate},
			Speaker:        defaultSpeaker,
			Tags:           []string{},
			Title:          title,
			TranscriptSlim: transcriptPath,
		}

		fullFilePath := filepath.Join(flagEpisodesStorePath, filename)

		// If file exists, preserve certain fields
		if _, err := os.Stat(fullFilePath); err == nil {
			existingEpisode, err := episode.NewEpisode(fullFilePath)
			if err == nil {
				existingFM := existingEpisode.GetFrontmatter()

				// Preserve manually-added fields
				if existingFM.Advertiser != "" {
					data.Advertiser = existingFM.Advertiser
				}
				if existingFM.AmazonMusic != "" {
					data.AmazonMusic = existingFM.AmazonMusic
				}
				if existingFM.ApplePodcasts != "" {
					data.ApplePodcasts = existingFM.ApplePodcasts
				}
				if existingFM.Deezer != "" {
					data.Deezer = existingFM.Deezer
				}
				if existingFM.LengthSecond > 0 {
					data.LengthSecond = existingFM.LengthSecond
				}
				if len(existingFM.Speaker) > 0 {
					data.Speaker = existingFM.Speaker
				}
				if existingFM.Spotify != "" {
					data.Spotify = existingFM.Spotify
				}
				if len(existingFM.Tags) > 0 {
					data.Tags = existingFM.Tags
				}
				if existingFM.YouTube != "" {
					data.YouTube = existingFM.YouTube
				}
			}
		}

		// Write file
		if err := writeEpisodeFile(fullFilePath, data, htmlContent); err != nil {
			logger.Warn().
				Err(err).
				Str("episode", title).
				Str("filePath", fullFilePath).
				Msg("failed to write episode file")
			continue
		}
	}

	logger.Info().
		Msg("Processing Podcast Episode items ... Successful")
	return nil
}

func writeEpisodeFile(filePath string, frontmatter episode.EpisodeFrontmatter, content string) error {
	// Convert frontmatter to YAML
	yamlData, err := yaml.Marshal(frontmatter)
	if err != nil {
		return fmt.Errorf("failed to marshal frontmatter: %w", err)
	}

	// Build complete file content
	fileContent := fmt.Sprintf("---\n%s---\n%s", string(yamlData), content)

	// Ensure directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write file
	if err := os.WriteFile(filePath, []byte(fileContent), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// parseHeadlinesFromHTML extracts h3 headlines from HTML and adds IDs
func parseHeadlinesFromHTML(rawHTML string) (string, map[string]string) {
	htmlContent := rawHTML
	headlineSlugs := make(map[string]string)

	// Find all h3 headlines with optional span tags
	headlineRegex := regexp.MustCompile(`<h3>(<span>)?(.*?)(</span>)?</h3>`)
	foundHeadlines := headlineRegex.FindAllStringSubmatch(htmlContent, -1)

	for _, h := range foundHeadlines {
		line := h[2]

		// Skip if line contains HTML tags (like <br>)
		if strings.Contains(line, "<") && strings.Contains(line, ">") {
			continue
		}

		slug := utils.Slugify(line)

		// Replace the h3 tag with one that includes an id
		oldH3 := fmt.Sprintf("<h3>%s%s%s", h[1], h[2], h[3])
		newH3 := fmt.Sprintf(`<h3 id="%s">%s`, slug, line)
		htmlContent = strings.Replace(htmlContent, oldH3, newH3, 1)

		headlineSlugs[slug] = line
	}

	return htmlContent, headlineSlugs
}

// getChapterFromDescription extracts chapter marks from description HTML
func getChapterFromDescription(description string) []map[string]string {
	var chapters []map[string]string

	// Try first pattern: <p><span>(00:00:00) Title</span></p>
	chapterRegex1 := regexp.MustCompile(`<p><span>\(([0-9:]*)\) ([^<]*)</span></p>`)
	foundChapters := chapterRegex1.FindAllStringSubmatch(description, -1)

	// If no matches, try second pattern: <p>(00:00:00) Title</p>
	if len(foundChapters) == 0 {
		chapterRegex2 := regexp.MustCompile(`<p>\(([0-9:]*)\) ([^<]*)</p>`)
		foundChapters = chapterRegex2.FindAllStringSubmatch(description, -1)
	}

	for _, c := range foundChapters {
		startTime := c[1]

		// Ensure time format is HH:MM:SS
		if strings.Count(startTime, ":") == 1 {
			startTime = "00:" + startTime
		}

		chapter := map[string]string{
			"start": startTime,
			"title": html.UnescapeString(c[2]),
		}
		chapters = append(chapters, chapter)
	}

	return chapters
}
