package echoopenapiui

import (
	"github.com/labstack/echo/v4"
	"github.com/rookie-luochao/go-openapi-ui"
)

func New(doc doc.Doc) echo.MiddlewareFunc {
	handle := doc.Handler()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			handle(ctx.Response(), ctx.Request())

			if ctx.Response().Committed {
				return nil
			}

			return next(ctx)
		}
	}
}
