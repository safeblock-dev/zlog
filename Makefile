.PHONY: *
.DEFAULT_GOAL := help

# Go
GO_BIN ?= $(shell go env GOPATH)/bin

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

$(GO_BIN)/golangci-lint:
	@echo "==> Installing golangci-lint within "${GO_BIN}""
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: $(GO_BIN)/golangci-lint ## Run linting on Go files
	@echo "==> Linting Go source files"
	@golangci-lint run -v --fix -c .golangci.yml ./...

test: ## Run tests
	go test -race -v ./... -coverprofile ./coverage.txt

bench: ## Run benchmarks. See https://pkg.go.dev/cmd/go#hdr-Testing_flags
	go test ./... -bench . -benchtime 5s -timeout 0 -run=XXX -cpu 1 -benchmem

update: ## Update packages
	@find . -name "go.mod" -execdir sh -c 'echo "Updating $$(pwd)"; go get -u ./... && go mod tidy' \;

generate: generate-go## Generate Go files.
generate-go: $(GO_FILES) ## Generate Go files
	go generate ./...
