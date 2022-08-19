package mysql

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

// Table general table defination
type Table interface {
	AutoMigrate() error
}

// Init init db
func Init(password, ip string, port int, name string) error {
	var err error
	once.Do(func() {
		dsn := buildDbUrl("root", password, ip, port)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		//	auto-migrate all the tables here.
		//	consider using api to expose table migrate feature
		println("syncing db table...")

		err := DB.AutoMigrate()
		if err != nil {
			println("create table error", err)
			return
		}
	})
	return err
}

func buildDbUrl(user, pwd, ip string, port int) string {
	s1 := strings.Join([]string{user, pwd}, ":")
	s2 := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	s3 := fmt.Sprintf("tcp(%v)", s2)
	s4 := s1 + "@" + s3
	s5 := strings.Join([]string{"charset=utf8mb4", "parseTime=True", "loc=Local"}, "&")
	return fmt.Sprintf("%v/%v?%v", s4, "ti_env_db", s5)
}
