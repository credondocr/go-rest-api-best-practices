package handler

import (
	"github.com/gin-gonic/gin"
)

// PingEndpoint godoc
// @Summary Ping example
// @Description Do ping
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
