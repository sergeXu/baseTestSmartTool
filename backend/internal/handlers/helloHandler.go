package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func HelloHandler(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Hello from Hertz backend!",
		"data":    []string{"Hertz", "Go", "Vue3"},
	})
}
