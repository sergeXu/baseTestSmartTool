package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/serge-xu/go-vue3-project/utils"
)

var Logger *utils.Log

// Logger简易初始化
func InitLogger() {
	logger, err := utils.NewLog("debug", "console.log")
	Logger = logger
	if err != nil {
		fmt.Printf("BeanInit error %v\n", err)
		panic(err)
	}
	fmt.Printf("BeanInit is nil %v\n", Logger == nil)
}
func CloseLogger() {
	if Logger != nil {
		Logger.Close()
	}
}

// 访问日志中间件示例
func AccessLogMiddleware() app.HandlerFunc {
	Logger.Infof("AccessLogMiddleware start")
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
