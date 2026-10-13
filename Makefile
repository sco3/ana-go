APP_NAME := pal-go
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

.DEFAULT_GOAL := help

.PHONY: help build build-prod run test clean

help: ## Show available commands
	@echo "Available commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
	@echo ""

build: ## Build for development
	go build -o bin/$(APP_NAME) .

build-prod: ## Production build (like cargo build --release)
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.Version=$(VERSION)" -o bin/$(APP_NAME) .

run: ## Run the project
	go run .

test: ## Run tests
	go test ./...

clean: ## Clean build artifacts
	rm -rf bin/