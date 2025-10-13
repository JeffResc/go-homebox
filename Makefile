.PHONY: generate clean test help

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

generate: ## Generate client from OpenAPI spec
	@./scripts/generate.sh

clean: ## Clean generated files and temp directory
	@rm -rf tmp/
	@rm -rf client/
	@echo "Cleaned generated files"

test: ## Run tests
	@go test ./...

deps: ## Install/update dependencies
	@go mod tidy
	@go mod download

.DEFAULT_GOAL := help
