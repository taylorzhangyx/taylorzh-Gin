package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	
	s := gin.Default()
	count := 0
	s.GET("/", func(c *gin.Context) {
		count++
		c.JSON(200, gin.H{
			"message": "hello world",
			"count":   count,
		})
	})
	s.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.POST("/SayHello", func(c *gin.Context) {

		res := "Version3 Hello "
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	s.POST("/whoami/*key", func(c *gin.Context) {
		keypath := c.Param("key")
		key := strings.Trim(keypath, "/")
		if key == "" {
			key = "whoami"
		}
		whoami := os.Getenv(key)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("key=%v, val=%v", key, whoami),
		})
	})
	s.POST("/loop/*iteration", func(c *gin.Context) {
		iteratepath := c.Param("iteration")
		iterate := strings.Trim(iteratepath, "/")
		println("iterate=", iterate)
		i, err := strconv.ParseInt(iterate, 10, 32)
		if err != nil {
			println(err.Error())
			i = 10
		}

		go func(ite int64) {
			counter := 0
			for n := int64(0); n < ite; n++ {
				counter++
			}
		}(i)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("for loop starting, iteration=%v", i),
		})
	})
	s.POST("/memwaste/*len", func(c *gin.Context) {
		length := strings.Trim(c.Param("len"), "/")
		println("length=", length)
		len, err := strconv.ParseInt(length, 10, 32)
		if err != nil {
			println(err.Error())
			len = 10
		}

		go func(ite int64) {
			l := make([]int64, 0)
			for n := int64(0); n < ite; n++ {
				l = append(l, n)
			}
		}(len)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("for loop starting, iteration=%v", len),
		})
	})

	var loadtotal int64
	loadMetrics := make(map[int64]int64)
	loadchan := make(chan int64, 10000)

	go func() {
		for {
			select {
			case t := <-loadchan:
				loadMetrics[t]++
				loadtotal++
			}
		}
	}()

	s.POST("/load/count", func(c *gin.Context) {
		t := time.Now()
		t.Unix()
		loadchan <- t.Unix()

		c.JSON(http.StatusOK, gin.H{
			"message": "load counted",
		})
	})

	s.GET("/load/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"total":   loadtotal,
			"message": loadMetrics,
		})
	})

	s.POST("/load/clear", func(c *gin.Context) {
		loadMetrics = make(map[int64]int64)
		loadtotal = 0
		c.JSON(http.StatusOK, gin.H{
			"message": "load metrics cleared",
		})
	})

	s.Run(":8080")

}
