package main

import (
	"github.com/gin-gonic/gin"
	doc "github.com/rookie-luochao/go-openapi-ui"
	ginopenapiui "github.com/rookie-luochao/go-openapi-ui/gin"
)

func main() {
	doc := doc.Doc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := gin.New()
	r.Use(ginopenapiui.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Run(":8000"))
}
