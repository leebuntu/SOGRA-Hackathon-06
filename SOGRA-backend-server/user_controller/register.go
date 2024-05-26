package user_controller

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"SOGRA/common"
	common_struct "SOGRA/common/structs"
	"SOGRA/db_controller"
)

func isDuplicateID(id string) bool {
	c := common_struct.User{}
	return db_controller.IsValueExist("user", &c, "id = ?", id)
}

func checkInputData(info common_struct.RegisterRequest) bool {
	if info.ID == "" || info.Password == "" || info.Email == "" {
		return false
	}

	return true
}

func inputAccountData(info common_struct.RegisterRequest) {
	saved_event := common_struct.User_saved_event{}
	db_controller.InsertValue("user_saved_event", saved_event)

	saved_course := common_struct.User_saved_course{}
	db_controller.InsertValue("user_saved_course", saved_course)

	p := common_struct.User{}
	p.User_id = 0
	p.Id = info.ID
	p.Password = info.Password
	p.Email = info.Email
	p.Profile_image = info.Profile
	p.Saved_event_list = saved_event.User_id
	p.Saved_course_list = saved_course.User_id

	db_controller.UsersDBConn.Table("user").Create(&p)
}

func registerAccount(info common_struct.RegisterRequest) bool {
	if checkInputData(info) {
		if isDuplicateID(info.ID) {
			return false
		} else {
			inputAccountData(info)
			return true
		}
	} else {
		return false
	}
}

func GetRegisterManager(c *gin.Context) {
	info := common_struct.RegisterRequest{}

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"message": common.IncorrectJson})
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		// Handle error
		return
	}

	json.Unmarshal([]byte(jsonData), &info)

	info.Password = common.SHA256Hash(info.Password)

	if registerAccount(info) {
		c.JSON(200, gin.H{
			"message": common.Sucess,
		})
	} else {
		c.JSON(400, gin.H{
			"message": common.DuplicateId,
		})
	}
}
