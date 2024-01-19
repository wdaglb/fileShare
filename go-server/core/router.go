package core

import (
	"github.com/gin-gonic/gin"
	"ke.file.share/core/context"
)

type Handler = func(c *context.Context)

func convertGinHandler(ctx *gin.Context, handler Handler) {
	newContext := &context.Context{
		Context: ctx,
		UserId:  ctx.GetString("__userId"),
	}
	handler(newContext)
}

func (t *AppServer) Use(handler ...Handler) {
	for _, h := range handler {
		t.handle.Use(func(context *gin.Context) {
			convertGinHandler(context, h)
		})
	}
}

func (t *AppServer) GET(path string, handler Handler) {
	t.handle.GET(path, func(context *gin.Context) {
		convertGinHandler(context, handler)
	})
}

func (t *AppServer) POST(path string, handler Handler) {
	t.handle.POST(path, func(context *gin.Context) {
		convertGinHandler(context, handler)
	})
}
