package api

import (
	"tinycloud/internal/models"
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func PostInstance(c *gin.Context) {
	var param models.InstanceParam
	c.BindJSON(&param)

	service.CreateInstance(param)

	c.JSON(200, gin.H{
		"list": "ok",
	})
}
