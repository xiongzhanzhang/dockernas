package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenToken() string {
	return uuid.New().String()
}

func GenPasswd() string {
	token := GenToken()
	return strings.Replace(token, "-", "", -1)[0:10]
}
