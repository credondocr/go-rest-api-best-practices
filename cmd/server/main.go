package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/credondocr/go-rest-api-best-practices/internal/docs"
	"github.com/credondocr/go-rest-api-best-practices/internal/graphql"
	"github.com/credondocr/go-rest-api-best-practices/internal/router"
	"github.com/credondocr/go-rest-api-best-practices/internal/service"
	"github.com/credondocr/go-rest-api-best-practices/pkg/config"
)

//	@title			Go Rest API Best Practices
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	cfg := config.LoadConfig()
	r := router.SetupRouter()
	service.InitHandler("grpc-server:50051")
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
	prometheus.MustRegister(httpRequestsTotal)

	r.Use(func(c *gin.Context) {
		path := c.FullPath()
		httpRequestsTotal.WithLabelValues(path).Inc()
		c.Next()
	})

	// Configurar servidor GraphQL
	graphqlHandler := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.POST("/graphql", gin.WrapH(graphqlHandler))
	r.GET("/graphql", gin.WrapH(graphqlHandler))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	r.Run(cfg.Server.Port)
}
