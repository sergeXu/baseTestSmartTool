package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	// 创建Hertz实例
	h := server.Default(
		server.WithHostPorts(":7001"),
		server.WithHandleMethodNotAllowed(true),
	)

	// 注册中间件
	h.Use(accessLogMiddleware())

	// 注册路由
	registerRoutes(h)

	// 启动服务
	h.Spin()
}

func registerRoutes(h *server.Hertz) {
	v1 := h.Group("/api/v1")
	{
		v1.GET("/hello", helloHandler)
	}
}

func helloHandler(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Hello from Hertz backend!",
		"data":    []string{"Hertz", "Go", "Vue3"},
	})
}

// 访问日志中间件示例
func accessLogMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 前置处理
		ctx.Next(c) // 执行后续处理
		// 后置处理
		fmt.Printf("[%s] %s %d\n",
			ctx.Request.Header.Method(),
			ctx.Request.URI().Path(),
			ctx.Response.StatusCode())
	}
}
