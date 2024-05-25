package account_controller

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"SOGRA/common"
	"SOGRA/db_controller"
)

func isDuplicateID(id string) bool {
	return false
}

func registerAccount(info common.RegisterFormat) bool {
	if isDuplicateID(info.ID) {
		return false
	} else {
		return db_controller.InsertAccountInfo(info)
	}
}

func GetRegisterManager(c *gin.Context) {
	info := common.RegisterFormat{}

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
