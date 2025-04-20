package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/serge-xu/go-vue3-project/internal/handlers"
)

func RegisterRoutes(h *server.Hertz) {
	// 健康检查接口
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{"message": "pong"})
	})
	v1 := h.Group("/api/v1")
	{
		v1.GET("/hello", handlers.HelloHandler)
	}
}
