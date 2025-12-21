SHELL := /bin/bash
.DEFAULT_GOAL := help

.PHONY: help
help: ## help message, list all command
	@echo -e "$$(grep -hE '^\S+.*:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

.PHONY: run
run: build ## Starts Pod Admin CLI
	./website-admin $(filter-out $@,$(MAKECMDGOALS)) || true

.PHONY: test
test: ## Runs all unit tests
	go test -v -race ./...

.PHONY: build
build: ## Compiles the application
	go build -race -ldflags "-X main.version=`git rev-parse --abbrev-ref HEAD`" -o website-admin

.PHONY: staticcheck
staticcheck: ## Runs static code analyzer staticcheck
	go install honnef.co/go/tools/cmd/staticcheck@2025.1.1
	staticcheck ./...

.PHONY: vet
vet: ## Runs go vet
	go vet ./...
