package user_controller

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"SOGRA/common"
	common_struct "SOGRA/common/structs"
	"SOGRA/db_controller"
)

func IsCorrectInfo(info common_struct.LoginRequest) bool {
	c := common_struct.User{}
	return db_controller.IsValueExist("user", &c, "id = ? AND password = ?", info.ID, info.Password)
}

func GetLoginManager(c *gin.Context) {
	info := common_struct.LoginRequest{}

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"message": common.IncorrectJson})
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		// Handle error
	}

	json.Unmarshal([]byte(jsonData), &info)

	if IsCorrectInfo(info) {
		c.JSON(200, gin.H{
			"token": common.Sucess,
		})
	} else {
		c.JSON(400, gin.H{
			"message": common.IncorrectInfo,
		})

	}
}
