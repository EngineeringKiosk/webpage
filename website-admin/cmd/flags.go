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
	flagNetlifyRedirectEpisodesPath   string
	flagNetlifyRedirectRedirectPrefix string

	// Podcast
	// TODO Unify with flagEpisodesStorePath?
	flagPodcastEpisodesPath string

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
