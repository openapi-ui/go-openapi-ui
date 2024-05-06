.PHONY: all deps golangci-lint

DOC_PATH=pkg/doc/assets/openapi-ui.umd.js
DOC_URL=https://unpkg.com/openapi-ui-dist@latest/lib/openapi-ui.umd.js

all: download lint test

download:
	curl -sL -o $(DOC_PATH) $(DOC_URL)

lint:
	go fmt ./...
	go vet ./...

test:
	go test -race ./...

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	
golangci-lint:
	golangci-lint run ./...

run-echo:
	cd _examples/echo && go run .

run-fiber:
	cd _examples/fiber && go run .

run-gin:
	cd _examples/gin && go run .

run-gorilla:
	cd _examples/gorilla && go run .

run-http:
	cd _examples/http && go run .
