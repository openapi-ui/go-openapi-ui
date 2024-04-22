.PHONY: all lint test deps

DOC_PATH=assets/openapi-ui.umd.js
DOC_URL=https://cdn.jsdelivr.net/npm/openapi-ui-dist@latest/lib/openapi-ui.umd.js

all: $(DOC_PATH) lint test

lint:
	go fmt ./...
	go vet ./...
	golangci-lint run ./...

test:
	go test -race ./...

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(DOC_PATH):
	curl -sL -o $(DOC_PATH) $(DOC_URL)
