package repo

import (
	"gorm.io/gorm"
)

const (
	AsyncTaskStatusInitiated = "Initiated"
	AsyncTaskStatusCreated   = "Created"
	AsyncTaskStatusRunning   = "Running"
	AsyncTaskStatusFinished  = "Finished"
)

type AsyncTask struct {
	gorm.Model
	TaskID        string `gorm:"unique;not null"`
	OwnerID       string `gorm:"not null"`
	Status        string `gorm:"not null"`
	StatusDetail  string
	TaskSignature []byte `gorm:"not null"`
	RetryCount    int32  `gorm:"not null;default:0"`
	State1        bool   `gorm:"not null;default:false;comment:任务1完成状态"`
	State2        bool   `gorm:"not null;default:false"`
	State3        bool   `gorm:"not null;default:false"`
	State4        bool   `gorm:"not null;default:false"`
}

// AutoMigrate gorm table auto migrate
func (e AsyncTask) AutoMigrate() error {
	return DB.AutoMigrate(&AsyncTask{})
}

func CreateAsyncTask(task *AsyncTask) error {
	return DB.Create(task).Error
}

func QueryAsyncTask(name string) (*AsyncTask, error) {
	var task AsyncTask
	err := DB.Where("name = ?", name).First(&task).Error
	return &task, err
}

func GetAllPendingTask() ([]*AsyncTask, error) {
	var tasks = make([]*AsyncTask, 0)
	err := DB.Not("status = ?", AsyncTaskStatusFinished).Find(&tasks).Error
	return tasks, err
}

func UpdateAsyncTaskStateByName(taskId string, statues string) error {
	return DB.Model(&AsyncTask{}).Where("task_id = ?", taskId).Update("status", statues).Error
}
