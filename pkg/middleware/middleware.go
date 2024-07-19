package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()

		log.Printf("Request: %s %s - %d - %v", c.Request.Method, c.Request.URL, status, latency)
	}
}
