package api

import (
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

type userInfo struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

func Login(c *gin.Context) {
	var param userInfo
	c.BindJSON(&param)

	c.JSON(200, gin.H{
		"token": service.GenToken(param.User, param.Passwd),
	})
}
