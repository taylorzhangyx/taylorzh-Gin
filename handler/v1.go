package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"taylorzh.dev.com/toy-gin/biz/load_recorder"
)

type basePOSTReq struct {
}

type baseRsp struct {
	Code      ErrorCode `json:"code"`
	RequestID string    `json:"request_id"`
	Error     string    `json:"error"`
}

// Version get the version comment
func Version(c *gin.Context) {
	c.String(http.StatusOK, "TAYLORZH-GIN V1.0 - TAYLORZYX@HOTMAIL.COM")
}

func LoadCount(c *gin.Context) {
	t := time.Now()
	t.Unix()
	load_recorder.LoadChan <- t.Unix()

	c.JSON(http.StatusOK, gin.H{
		"message": "load counted",
	})
}

func LoadMetrics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total":   load_recorder.Loadtotal,
		"message": load_recorder.LoadMetrics,
	})
}

func LoadClear(c *gin.Context) {
	load_recorder.Reset()
	c.JSON(http.StatusOK, gin.H{
		"message": "load metrics cleared",
	})
}
