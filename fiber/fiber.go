package fiberopenapiui

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/rookie-luochao/go-openapi-ui"
)

func New(doc doc.Doc) fiber.Handler {
	return adaptor.HTTPHandlerFunc(doc.Handler())
}
