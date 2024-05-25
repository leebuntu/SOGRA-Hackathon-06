package db_controller

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"SOGRA/common"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", common.DbID, common.DbPassword, common.DbAddr)
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
