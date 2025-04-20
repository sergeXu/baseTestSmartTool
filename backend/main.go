package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/serge-xu/go-vue3-project/config"
	"github.com/serge-xu/go-vue3-project/router"
	"github.com/serge-xu/go-vue3-project/test"
)

func main() {
	config.InitLogger()
	defer config.CloseLogger()

	test.TestStart()
	// 创建Hertz实例
	h := server.Default(
		server.WithHostPorts(":7001"),
		server.WithHandleMethodNotAllowed(true),
	)
	// 注册中间件
	h.Use(config.AccessLogMiddleware())

	// 注册路由
	router.RegisterRoutes(h)

	// 启动服务
	h.Spin()
}
