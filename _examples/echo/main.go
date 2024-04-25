package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rookie-luochao/go-openapi-ui/pkg/doc"
	middleware_echo "github.com/rookie-luochao/go-openapi-ui/pkg/middleware/echo"
	"os"
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
	r.Use(middleware_echo.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Start(":8000"))
}
