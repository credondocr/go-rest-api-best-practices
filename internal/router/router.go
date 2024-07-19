package router

import (
	"github.com/gin-gonic/gin"

	"github.com/credondocr/go-rest-api-best-practices/internal/handler"
	"github.com/credondocr/go-rest-api-best-practices/pkg/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RequestLogger())

	api := r.Group("/api/v1")
	{
		api.GET("/ping", handler.PingEndpoint)
		api.GET("/hello", handler.HelloEndpoint)
		api.GET("/products", handler.GetAllProductsEndpoint)
		api.GET("/products/search", handler.SearchProductsEndpoint)

	}

	return r
}
