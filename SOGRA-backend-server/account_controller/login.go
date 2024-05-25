package account_controller

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"SOGRA/common"
)

func isCorrectInfo(info common.LoginFormat) bool {
	return false
}

func createToken(userID string) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", common.JWTSecretKey)
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetLoginManager(c *gin.Context) {
	info := common.LoginFormat{}

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		// Handle error
	}

	json.Unmarshal([]byte(jsonData), &info)

	if isCorrectInfo(info) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "Incorrect ID or Password. Please try again.",
		})

	}
}
