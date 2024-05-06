package ginopenapiui

import (
	"github.com/gin-gonic/gin"
	"github.com/rookie-luochao/go-openapi-ui/pkg/doc"
)

func New(doc doc.Doc) gin.HandlerFunc {
	handle := doc.Handler()

	return func(ctx *gin.Context) {
		handle(ctx.Writer, ctx.Request)
		ctx.Next()
	}
}
