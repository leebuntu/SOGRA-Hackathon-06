package user_controller

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	common_struct "SOGRA/common/structs"
	"SOGRA/db_controller"
)

func IsCorrectInfo(info common_struct.LoginRequest) bool {
	p := common_struct.User{}
	err := db_controller.UsersDBConn.Table("user").Where("id = ? AND password = ?", info.ID, info.Password).First(&p).Error

	return err == nil
}

func GetLoginManager(c *gin.Context) {
	info := common_struct.LoginRequest{}

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		// Handle error
	}

	json.Unmarshal([]byte(jsonData), &info)

	if IsCorrectInfo(info) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Incorrect ID or Password. Please try again.",
		})

	}
}
