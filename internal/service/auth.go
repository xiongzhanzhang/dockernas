package service

import (
	"log"
	"tinycloud/internal/config"

	"github.com/google/uuid"
)

var userToken string

func IsTokenValid(token string) bool {
	log.Println(userToken)
	return token != "" && userToken == token
}

func GenToken(user string, passwd string) string {
	realUserName, realPasswd := config.GetUserInfo()
	if realUserName != user || realPasswd != passwd {
		panic("user password error")
	}

	userToken = uuid.New().String()
	return userToken
}
