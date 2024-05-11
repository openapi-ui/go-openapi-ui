package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
	fiberopenapiui "github.com/openapi-ui/go-openapi-ui/pkg/middleware/fiber"
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
	r.Use(fiberopenapiui.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Listen(":8000"))
}
