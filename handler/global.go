package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	helloCount = 0
)

// Ping ping-pong
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Hello(c *gin.Context) {
	helloCount++
	c.JSON(200, gin.H{
		"message": "hello world",
		"count":   helloCount,
	})
}

// HealthCheck check
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("OK 122234522 %v", time.Now().Format(time.RFC3339)))
}
