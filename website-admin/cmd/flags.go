package cmd

var (
	// RSS Feed
	flagRSSFeedURL string

	// Episodes
	flagEpisodesStorePath string
	flagImagesDir         string
	flagTranscriptDir     string

	// TODO Unify "Dir" vs "Path" naming

	// Netlify Redirects
	flagNetlifyRedirectTomlFile       string
	flagNetlifyRedirectEpisodesDir    string
	flagNetlifyRedirectRedirectPrefix string

	// Podcast
	// TODO Unify with flagEpisodesStorePath?
	flagPodcastEpisodesDir string

	// Tags
	flagTagsWriteFile  bool
	flagTagsContentDir string
	flagTagsDescFile   string

	// Sync
	flagSyncAwesomeGamesStoragePath string

	// Logging
	flagDebugLogging   bool
	flagDisableLogging bool
)
