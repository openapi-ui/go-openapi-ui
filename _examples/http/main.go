package main

import (
	"net/http"

	"github.com/rookie-luochao/go-openapi-ui"
)

func main() {
	doc := doc.Doc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(http.ListenAndServe(":8000", doc.Handler()))
}
