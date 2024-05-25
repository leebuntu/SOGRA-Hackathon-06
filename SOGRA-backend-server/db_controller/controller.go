package db_controller

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"SOGRA/common"
)

var (
	UsersDBConn  *gorm.DB
	EventsDBConn *gorm.DB
)

func userDBInit() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/users", common.DbID, common.DbPassword, common.DbAddr)
	UsersDBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func eventDBInit() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/events", common.DbID, common.DbPassword, common.DbAddr)
	EventsDBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func InitDB() {
	userDBInit()
	eventDBInit()
}
