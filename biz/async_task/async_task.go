package async_task

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/google/uuid"
	"taylorzh.dev.com/toy-gin/repo"
)

var AsyncServer *machinery.Server
var workerTag string

const (
	TaskQueue    = "machinery_tasks"
	TaskNameAdd  = "add"
	TaskNameEcho = "echo"
)

func Init(redisIp string, redisPort int) error {
	conf := &config.Config{
		DefaultQueue:    TaskQueue,
		ResultsExpireIn: 3600,
		Broker:          fmt.Sprintf("redis://%s:%v", redisIp, redisPort),
		ResultBackend:   fmt.Sprintf("redis://%s:%v", redisIp, redisPort),
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}
	s, serverErr := machinery.NewServer(conf)
	if serverErr != nil {
		return serverErr
	}
	// save the machinery server as variable
	AsyncServer = s

	// task name - handler map
	taskConfig := map[string]interface{}{
		TaskNameEcho: echo,
		TaskNameAdd:  add,
	}
	registerErr := s.RegisterTasks(taskConfig)
	if registerErr != nil {
		return registerErr
	}

	workerTag = uuid.New().String()
	w := s.NewWorker(workerTag, 0)

	w.SetPostTaskHandler(posttaskHandler)
	w.SetErrorHandler(errorHandler)
	w.SetPreTaskHandler(pretaskHandler)

	go func() {
		// start a new goroutine to run worker
		fmt.Println("starting async task worker", workerTag)
		e := w.Launch()
		fmt.Println(e.Error())
	}()

	go taskDispatcherLoop()

	return nil
}

func errorHandler(err error) {
	log.ERROR.Println("I am an error handler:", err)
}

func pretaskHandler(signature *tasks.Signature) {
	log.INFO.Println("I am a start of task handler for:", signature.Name)
}

func posttaskHandler(signature *tasks.Signature) {
	log.INFO.Println("I am an end of task handler for:", signature.Name)
	s, _ := json.Marshal(signature)
	fmt.Println("signature:", signature, string(s))
}

func taskDispatcherLoop() {
	// create a token bucket with size of 1
	tokenBucket := make(chan struct{}, 1)
	tokenBucket <- struct{}{}

	for {
		select {
		case <-tokenBucket:
			fmt.Println("start scanning tasks...")
			// 	do db scan to schedule task
			taskDispatcher()

			// after task finished, send token back to the bucket
			time.AfterFunc(time.Duration(1000)*time.Millisecond, func() {
				tokenBucket <- struct{}{}
			})
		}
	}
}

func taskDispatcher() {
	dbTasks, err := repo.GetAllPendingTask()
	if err != nil {
		fmt.Println("read database error", err)
		return
	}

	// check if the task is processed by backend
	for _, t := range dbTasks {
		state, getErr := AsyncServer.GetBackend().GetState(t.TaskID)
		if getErr != nil && getErr.Error() == "redigo: nil returned" {
			fmt.Println("task not scheduled", t.TaskID, "start schedule process...")
			sendTaskToBroker(t)
			continue
		}
		if getErr != nil {
			fmt.Println("get state error", getErr, t.TaskID, getErr.Error() == "redigo: nil returned")
			continue
		}
		if state.IsSuccess() {
			err = repo.UpdateAsyncTaskStateByName(t.TaskID, repo.AsyncTaskStatusFinished)
			fmt.Println("task finished", t.TaskID, err)
		}
	}
}
