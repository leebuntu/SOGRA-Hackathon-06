package user_controller

import (
	common_struct "SOGRA/common/structs"
	"SOGRA/db_controller"

	"github.com/gin-gonic/gin"
)

func GetSave(c *gin.Context) {
	u := common_struct.User{}
	s := common_struct.User_saved_event{}

	err := db_controller.UsersDBConn.Table("user").Where("user_id = ?", c.Param("id")).First(&u).Error
	if err != nil {
		panic(err)
	}

	err = db_controller.UsersDBConn.Table("user_saved_event").Where("user_id = ?", u.Saved_event_list).First(&s).Error
	if err != nil {
		panic(err)
	}
}

func AddSave(c *gin.Context) {

}

func DeleteSave(c *gin.Context) {

}

func GetProfileImage(c *gin.Context) {

}

func GetNickname(c *gin.Context) {

}
