package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
	echoopenapiui "github.com/openapi-ui/go-openapi-ui/pkg/middleware/echo"
)

func main() {
	path, _ := os.Getwd()

	fmt.Println("a:", path)

	doc := doc.Doc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := echo.New()
	r.Use(echoopenapiui.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Start(":8000"))
}
