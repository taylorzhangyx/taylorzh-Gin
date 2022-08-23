package async_task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/google/uuid"
	"taylorzh.dev.com/toy-gin/repo"
)

func ScheduleTaskAdd(ctx context.Context, nums ...int) error {

	var args []tasks.Arg
	for _, n := range nums {
		args = append(args, tasks.Arg{Type: "int64", Value: n})
	}

	s := tasks.Signature{
		Name: TaskNameAdd,
		Args: args,
		UUID: fmt.Sprintf("task_%v", uuid.New().String()),
	}

	err := createTaskInDb(s.UUID, s)
	if err != nil {
		return err
	}

	return nil
}

func createTaskInDb(taskId string, s tasks.Signature) error {
	sJson, err := json.Marshal(s)
	if err != nil {
		return err
	}
	// insert into db before issue to async server
	t := &repo.AsyncTask{
		TaskID:        taskId,
		Status:        repo.AsyncTaskStatusInitiated,
		StatusDetail:  "Task is saved in db but not running yet",
		TaskSignature: sJson,
	}
	createTaskErr := repo.CreateAsyncTask(t)
	if createTaskErr != nil {
		return createTaskErr
	}
	return nil
}

func sendTaskToBroker(t *repo.AsyncTask) {
	s := &tasks.Signature{}
	err := json.Unmarshal(t.TaskSignature, s)
	if err != nil {
		log.ERROR.Printf("Failed to unmarshal task signature: %v", err)
		return
	}
	_, sendErr := AsyncServer.SendTaskWithContext(context.Background(), s)
	if sendErr != nil {
		fmt.Println(sendErr)
		return
	}
	_ = repo.UpdateAsyncTaskStateByName(t.TaskID, repo.AsyncTaskStatusCreated)

	fmt.Println("task send", t.TaskID)
}
