package service

import (
	"dockernas/internal/config"
	"time"

	"github.com/google/uuid"
)

var tokenMap = make(map[string]int64)

func IsTokenValid(token string) bool {
	// log.Println(userToken)
	_, ok := tokenMap[token]
	return ok
}

func GenToken(user string, passwd string) string {
	realUserName, realPasswd := config.GetUserInfo()
	if realUserName != user || realPasswd != passwd {
		panic("user password error")
	}

	userToken := uuid.New().String()
	tokenMap[userToken] = time.Now().UnixMilli()
	return userToken
}
