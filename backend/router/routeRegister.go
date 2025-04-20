package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/serge-xu/go-vue3-project/internal/handlers"
)

func RegisterRoutes(h *server.Hertz) {
	v1 := h.Group("/api/v1")
	{
		v1.GET("/hello", handlers.HelloHandler)
	}
}
