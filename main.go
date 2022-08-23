package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"taylorzh.dev.com/toy-gin/api"
	"taylorzh.dev.com/toy-gin/biz/async_task"
	"taylorzh.dev.com/toy-gin/biz/load_recorder"
	"taylorzh.dev.com/toy-gin/repo"
)

func main() {
	dbPw := flag.String("dP", "*", "database pa")
	dbN := flag.String("dn", "*", "database name")
	dbPort := flag.Int("dp", 8080, "database port")
	dbIp := flag.String("di", "*", "database ip")

	flag.Parse()

	println("starting taylorzh Gin Server...")
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.ForceConsoleColor()

	s := gin.Default()
	s.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	InitRepo(*dbPw, *dbIp, *dbPort, *dbN)
	InitBiz(repo.LocalConfig.AsyncTaskRunner)

	s = api.SetupRouter(s)

	s.Run(":8080")
}

func InitBiz(c *repo.AsyncTaskConfig) {
	load_recorder.Init()
	if err := async_task.Init(c.RedisIp, c.RedisPort); err == nil {
		fmt.Println("async task is started and running...")
	}
}

func InitRepo(pw, ip string, port int, schema string) {

	err := repo.Init(pw, ip, port, schema)
	if err != nil {
		log.Fatalln("failed to init mysql", err.Error())
	}
}
