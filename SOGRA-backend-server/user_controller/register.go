package user_controller

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	common_struct "SOGRA/common/structs"
	"SOGRA/db_controller"
)

func isDuplicateID(id string) bool {
	p := common_struct.User{}
	err := db_controller.UsersDBConn.Table("user").Where("id = ?", id).First(&p).Error

	return err == nil
}

func checkInputData(info common_struct.RegisterRequest) bool {
	if info.ID == "" || info.Password == "" || info.Email == "" {
		return false
	}

	return true
}

func inputAccountData(info common_struct.RegisterRequest) {
	p := common_struct.User{}
	p.Id = info.ID
	p.Password = info.Password
	p.Email = info.Email
	p.Profile_image = info.Profile
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
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		// Handle error
		return
	}

	json.Unmarshal([]byte(jsonData), &info)

	if registerAccount(info) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Duplicate ID. Please try again.",
		})
	}
}
