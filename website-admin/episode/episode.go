package episode

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/adrg/frontmatter"
)

// Speaker represents a podcast host or guest
type Speaker struct {
	Name             string `yaml:"name" json:"name"`
	TranscriptLetter string `yaml:"transcriptLetter" json:"transcriptLetter"`
}

// Chapter represents a chapter in a podcast episode
type Chapter struct {
	Start string `yaml:"start" json:"start"`
	Title string `yaml:"title" json:"title"`
}

// Episode represents a podcast episode with its metadata
type Episode struct {
	fileName           string
	episodeNumber      string
	episodeFrontmatter EpisodeFrontmatter
	content            string
}

// EpisodeFrontmatter represents the frontmatter of an episode
type EpisodeFrontmatter struct {
	Advertiser     string    `yaml:"advertiser" json:"advertiser"`
	AmazonMusic    string    `yaml:"amazon_music" json:"amazon_music"`
	ApplePodcasts  string    `yaml:"apple_podcasts" json:"apple_podcasts"`
	Audio          string    `yaml:"audio" json:"audio"`
	Chapter        []Chapter `yaml:"chapter" json:"chapter"`
	Deezer         string    `yaml:"deezer" json:"deezer"`
	Description    string    `yaml:"description" json:"description"`
	Headlines      string    `yaml:"headlines" json:"headlines"`
	Image          string    `yaml:"image" json:"image"`
	LengthSecond   int       `yaml:"length_second" json:"length_second"`
	PubDate        string    `yaml:"pubDate" json:"pubDate"`
	Speaker        []Speaker `yaml:"speaker" json:"speaker"`
	Spotify        string    `yaml:"spotify" json:"spotify"`
	Tags           []string  `yaml:"tags" json:"tags"`
	Title          string    `yaml:"title" json:"title"`
	TranscriptSlim string    `yaml:"transcript_slim" json:"transcript_slim"`
	YouTube        string    `yaml:"youtube" json:"youtube"`
}

// NewEpisode creates a new Episode instance from a file path
func NewEpisode(fileName string) (*Episode, error) {
	episode := &Episode{
		fileName: fileName,
	}

	episodeNumber, err := episode.getNumber(true)
	if err != nil {
		return nil, err
	}
	episode.episodeNumber = episodeNumber

	if err := episode.load(); err != nil {
		return nil, err
	}

	return episode, nil
}

// load reads and parses the episode file with frontmatter
func (e *Episode) load() error {
	file, err := os.Open(e.fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error when closing:", err)
		}
	}()

	var fm EpisodeFrontmatter
	rest, err := frontmatter.Parse(file, &fm)
	if err != nil {
		return err
	}

	e.episodeFrontmatter = fm
	e.content = string(rest)
	return nil
}

// GetFrontmatter returns the complete frontmatter
func (e *Episode) GetFrontmatter() EpisodeFrontmatter {
	return e.episodeFrontmatter
}

// getNumber extracts the episode number from the filename
// leadingZero determines whether to return with leading zeros (e.g., "04" vs "4")
//
// TODO Cleanup this function together with GetEpisodeNumberFromFilename (it is nearly the same)
func (e *Episode) getNumber(leadingZero bool) (string, error) {
	filename := filepath.Base(e.fileName)
	index := strings.Index(filename, "-")

	var episodeNumber string

	// Handle episodes starting with '-1' (negative episode numbers)
	if index == 0 {
		remainingFilename := filename[1:]
		nextIndex := strings.Index(remainingFilename, "-")
		if nextIndex == -1 {
			return "", fmt.Errorf("invalid filename format: %s", filename)
		}
		episodeNumber = filename[0 : nextIndex+1]
	} else if index == -1 {
		return "", fmt.Errorf("invalid filename format: %s", filename)
	} else {
		episodeNumber = filename[0:index]
	}

	if !leadingZero {
		episodeNumber = trimEpisodeNumber(episodeNumber)
	} else {
		// Pad with zeros for positive numbers
		if !strings.HasPrefix(episodeNumber, "-") {
			if num, err := strconv.Atoi(episodeNumber); err == nil && num < 10 {
				episodeNumber = fmt.Sprintf("%02d", num)
			}
		}
	}

	return episodeNumber, nil
}

// trimEpisodeNumber removes leading zeros from episode number
func trimEpisodeNumber(episodeNumber string) string {
	if strings.HasPrefix(episodeNumber, "0") {
		return strings.TrimLeft(episodeNumber, "0")
	}
	return episodeNumber
}
