package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"taylorzh.dev.com/toy-gin/biz/async_task"
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

// *****************************************
// 		load counter related handlers
// *****************************************

func LoadCount(c *gin.Context) {

	load_recorder.Count()
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

// *****************************************
// 		async task related handlers
// *****************************************

type AsyncTaskAddParam struct {
	Nums []int `form:"nums" json:"nums" xml:"nums"  binding:"required"`
}

func AsyncTaskAdd(c *gin.Context) {
	var params AsyncTaskAddParam
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := async_task.ScheduleTaskAdd(c, params.Nums...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"res": fmt.Sprintln(params.Nums)})
}

func DescribeTaskResult(c *gin.Context) {
	
}