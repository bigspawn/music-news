#!make
.PHONY: lint test run_parser run_notifier run clean build help

APP_NAME := music-news
MAIN_PATH := cmd/main.go
BUILD_DIR := build

help: ## Display available commands
	@echo "Available commands:"
	@echo
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo

lint: ## Run code linters
	@echo "Running linters..."
	@go vet ./...
	@echo "Linting completed"

test: ## Run tests with race condition detection
	@echo "Running tests..."
	CGO_ENABLED=0 go test -v -count 1 -race ./...

build: ## Build the application binary
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Build completed: $(BUILD_DIR)/$(APP_NAME)"

run_parser: ## Run the parser service
	@echo "Running parser..."
	go run $(MAIN_PATH)

run_notifier: ## Run the notifier service
	@echo "Running notifier..."
	ONLY_NOTIFIER=true go run $(MAIN_PATH)

run: ## Run service with environment variables from .env file
	@echo "Starting service with .env configuration..."
	@if [ -f .env.local ]; then \
		export $$(grep -v '^#' .env.local | xargs) && go run $(MAIN_PATH); \
	else \
		echo ".env.local file not found. Starting without environment variables from .env"; \
		go run $(MAIN_PATH); \
	fi

clean: ## Clean build artifacts and tidy dependencies
	@echo "Cleaning..."
	@go mod tidy -v
	@go clean ./...
	@rm -rf $(BUILD_DIR)
	@echo "Cleaning completed"

.DEFAULT_GOAL := help
