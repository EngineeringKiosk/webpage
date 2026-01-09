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
clean: ## Deletes all generated items (like node_modules, build output, caches)
	rm -rf dist
	rm -rf node_modules
	rm -rf .ruff_cache
	rm -rf .astro

.PHONY: init
init: init-javascript ## Installs all dependencies (JavaScript)

.PHONY: init-javascript
init-javascript: ## Installs JavaScript dependencies
	npm install

.PHONY: prettier
prettier: ## Run code formatter prettier (for JavaScript)
	node_modules/.bin/prettier -w .

# Go sync scripts (website-admin)
.PHONY: update-german-tech-podcasts
update-german-tech-podcasts: ## Updates the German Tech Podcasts data from https://github.com/EngineeringKiosk/GermanTechPodcasts
	./website-admin/website-admin sync german-tech-podcasts
