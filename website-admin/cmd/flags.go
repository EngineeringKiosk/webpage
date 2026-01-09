package cmd

var (
	// RSS Feed
	flagRSSFeedURL string

	// Episodes
	flagEpisodesStorePath string
	flagImagesPath        string
	flagTranscriptPath    string

	// Netlify Redirects
	flagNetlifyRedirectTomlFile       string
	flagNetlifyRedirectRedirectPrefix string

	// Tags
	flagTagsWriteFile   bool
	flagTagsContentPath string
	flagTagsDescFile    string

	// Sync
	flagSyncAwesomeGamesStoragePath       string
	flagSyncGermanTechPodcastsStoragePath string
	flagSyncGermanTechPodcastsOPMLPath    string

	// Logging
	flagDebugLogging   bool
	flagDisableLogging bool
)
