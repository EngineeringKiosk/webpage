.DEFAULT_GOAL := help

.PHONY: help
help: ## Outputs the help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Compiles the application into static content
	npm run build

.PHONY: run
run: ## Starts the development server
	npm run dev

.PHONY: clean
clean: ## Deletes the generated content and node_modules
	rm -rf ./dist
	rm -rf ./node_modules

.PHONY: update-content
update-content: ## Pulls the latest Podcast RSS feed and updates the content
	python ./scripts/podcast_feed_to_content.py sync

.PHONY: update-redirects
update-redirects: ## Writes all short url redirects for Podcast episodes to netlify.toml
	python ./scripts/podcast_feed_to_content.py redirect

.PHONY: episode-check
episode-check: ## Checks all Podcast Episodes if all player links (Spotify, etc.) are set
	python ./scripts/empty-player-urls.py

.PHONY: find-tags-that-need-descriptions-content-files
find-tags-that-need-descriptions-content-files: ## Checks all used tags in content files (blog posts and podcasts) that need SEO descriptions and output them on stdout
	python ./scripts/find-tags-that-need-descriptions.py website-content

.PHONY: find-tags-that-need-descriptions-german-tech-podcast-file
find-tags-that-need-descriptions-german-tech-podcast-file: ## Checks all used tags in the german tech podcasts that need SEO descriptions and output them on stdout
	python ./scripts/find-tags-that-need-descriptions.py german-tech-podcasts

.PHONY: find-tags-that-need-descriptions-content-files-dump
find-tags-that-need-descriptions-content-files-dump: ## Find all used tags in content files (blog posts and podcasts) that need SEO descriptions and dump it down to disk into tag-file
	python ./scripts/find-tags-that-need-descriptions.py -write-file website-content

.PHONY: find-tags-that-need-descriptions-german-tech-podcast-file-dump
find-tags-that-need-descriptions-german-tech-podcast-file-dump: ## Find all used tags in the german tech podcasts that need SEO descriptions and dump it down to disk into tag-file
	python ./scripts/find-tags-that-need-descriptions.py -write-file german-tech-podcasts

.PHONY: sync-german-tech-podcasts
sync-german-tech-podcasts: ## Syncs German Tech Podcasts data from https://github.com/EngineeringKiosk/GermanTechPodcasts
	python ./scripts/sync-german-tech-podcasts.py

.PHONY: init
init: ## Installs dependencies
	npm install

.PHONY: eslint
eslint: ## Statically analyzes of basic JavaScript scripts
	node_modules/.bin/eslint --config .eslintrc src/scripts