package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rookie-luochao/go-openapi-ui"
	echoredoc "github.com/rookie-luochao/go-openapi-ui/echo"
)

func main() {
	doc := doc.Doc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := echo.New()
	r.Use(echoredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Start(":8000"))
}
