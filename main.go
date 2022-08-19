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
	"taylorzh.dev.com/toy-gin/biz/load_recorder"
	"taylorzh.dev.com/toy-gin/repo/mysql"
)

var dbPw = flag.String("dP", "*", "database pa")
var dbN = flag.String("dn", "*", "database name")
var dbPort = flag.Int("dp", 8080, "database port")
var dbIp = flag.String("di", "*", "database ip")

func main() {
	println("starting taylorzh Gin Server...")
	f, _ := os.Create("ti-env-manager-go.log")
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

	InitBiz()
	InitRepo()

	s = api.SetupRouter(s)

	s.Run(":8080")
}

func InitBiz() {
	load_recorder.Init()
}

func InitRepo() {

	err := mysql.Init(*dbPw, *dbIp, *dbPort, *dbN)
	if err != nil {
		log.Fatalln("failed to init mysql", err.Error())
	}
}
