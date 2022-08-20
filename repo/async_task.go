package repo

import (
	"gorm.io/gorm"
)

type AsyncTask struct {
	gorm.Model
	Name         string `gorm:"unique;not null"`
	OwnerID      string `gorm:"not null"`
	Status       string `gorm:"not null"`
	StatusDetail string
	state1       bool `gorm:"not null;default:false;comment:任务1完成状态"`
	state2       bool `gorm:"not null;default:false"`
	state3       bool `gorm:"not null;default:false"`
	state4       bool `gorm:"not null;default:false"`
}

// AutoMigrate gorm table auto migrate
func (e AsyncTask) AutoMigrate() error {
	return DB.AutoMigrate(&AsyncTask{})
}
