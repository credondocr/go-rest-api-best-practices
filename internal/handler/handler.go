package handler

import (
	"github.com/gin-gonic/gin"
)

// Controller example
type Handler struct {
}

// NewHandler example
func NewHandler() *Handler {
	return &Handler{}
}

// Ping godoc
// @Summary Ping example
// @Description Do ping
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func (Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
