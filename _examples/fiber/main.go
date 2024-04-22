package main

import (
	"github.com/gofiber/fiber/v2"
	fiberredoc "github.com/rookie-luochao/go-openapi-ui/fiber"
	"github.com/rookie-luochao/go-openapi-ui/pkg/doc"
)

func main() {
	doc := doc.Doc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := fiber.New()
	r.Use(fiberredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Listen(":8000"))
}
