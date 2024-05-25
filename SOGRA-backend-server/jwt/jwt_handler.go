package jwt

import (
	"SOGRA/common"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
