package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/credondocr/go-rest-api-best-practices/internal/service"
)

// HelloEndpoint godoc
// @Summary Hello example
// @Description Say hello
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func HelloEndpoint(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "world"
	}

	message, err := service.CallGRPCService(name)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error calling gRPC service"})
		return
	}

	c.JSON(200, gin.H{"message": message})
}
