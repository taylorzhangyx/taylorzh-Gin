package api

import (
	"github.com/gin-gonic/gin"
	"taylorzh.dev.com/toy-gin/handler"
)

// SetupRouter init
func SetupRouter(r *gin.Engine) *gin.Engine {

	// global
	r.GET("/hello", handler.Hello)
	r.POST("/ping", handler.Ping)
	r.GET("/healthcheck", handler.HealthCheck)

	// v1
	v1 := r.Group("/v1")
	{
		v1.GET("/version", handler.Version)
	}

	return r
}
