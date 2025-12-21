.DEFAULT_GOAL := help

VENV_NAME?=scripts/venv
VENV_ACTIVATE=$(VENV_NAME)/bin/activate

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
clean: ## Deletes all generated items (like node_modules, build output, caches)
	rm -rf dist
	rm -rf node_modules
	rm -rf .ruff_cache
	rm -rf .astro
	rm -rf scripts/__pycache__
	rm -rf scripts/.ruff_cache
	rm -rf scripts/venv
	rm -rf scripts/.cache

.PHONY: init
init: init-javascript init-python ## Installs all dependencies (JavaScript and python)

.PHONY: init-javascript
init-javascript: ## Installs JavaScript dependencies
	npm install

.PHONY: prettier
prettier: ## Run code formatter prettier (for JavaScript)
	node_modules/.bin/prettier -w .

# target is only run if the target or any of the prerequisites have changed (touchfile)
scripts/venv/touchfile: scripts/requirements.txt
	test -d $(VENV_NAME) || python3 -m venv $(VENV_NAME)
	. $(VENV_ACTIVATE) && cd scripts && pip install --require-virtualenv -Ur requirements.txt
	touch $(VENV_NAME)/touchfile

.PHONY: init-python
init-python: scripts/venv/touchfile ## Installs python dependencies and creates a virtualenv

# Python scripts
.PHONY: update-episode-content
update-episode-content: ## Pulls the latest Podcast RSS feed from RedCircle and updates the content
	. $(VENV_ACTIVATE); python ./scripts/podcast_feed_to_content.py sync

.PHONY: update-episode-content-no-api
update-episode-content-no-api: ## Pulls the latest Podcast RSS feed from RedCircle and updates the content (without doing external API calls)
	. $(VENV_ACTIVATE); python ./scripts/podcast_feed_to_content.py sync --no-api-calls

.PHONY: update-episode-redirects
update-episode-redirects: ## Writes all short url redirects for Podcast episodes to netlify.toml
	. $(VENV_ACTIVATE); python ./scripts/podcast_feed_to_content.py redirect

.PHONY: update-german-tech-podcasts
update-german-tech-podcasts: ## Updates the German Tech Podcasts data from https://github.com/EngineeringKiosk/GermanTechPodcasts
	. $(VENV_ACTIVATE); python ./scripts/sync_german_tech_podcasts.py
