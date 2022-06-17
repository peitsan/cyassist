package login

import (
	"checkinWebsite/database"
	"github.com/gin-gonic/gin"
)

func Login(user, password string, c *gin.Context) (code int, message string, token string) {
	code, message = database.Login(user, password)
	token = ""
	if code == 2200 {
		var err error
		token, err = GetLoginToken(user)
		if err != nil {
			return 500, "加解密失败", ""
		}
	}
	return code, message, token
}

func GetLoginToken(user string) (string, error) {
	login, err := EncryptLogin(user)
	if err != nil {
		return "", err
	}
	return login, nil
}
