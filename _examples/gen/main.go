package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rookie-luochao/go-openapi-ui"

	_ "github.com/rookie-luochao/go-openapi-ui/_examples/gen/docs"

	fiberopenapiui "github.com/rookie-luochao/go-openapi-ui/fiber"
)

//go:generate swag init

// @title Fiber Example API
// @version 1.0
// @description Fiber example for openapi spec generation
// @host localhost:8000
// @BasePath /
func main() {
	doc := doc.Doc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/swagger.json",
		DocsPath:    "/docs",
	}

	r := fiber.New()
	r.Use(fiberredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Listen(":8000"))
}
