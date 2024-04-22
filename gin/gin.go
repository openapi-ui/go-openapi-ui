package ginopenapiui

import (
	"github.com/gin-gonic/gin"
	doc "github.com/rookie-luochao/go-openapi-ui"
)

func New(doc doc.Doc) gin.HandlerFunc {
	handle := doc.Handler()

	return func(ctx *gin.Context) {
		handle(ctx.Writer, ctx.Request)
		ctx.Next()
	}
}
