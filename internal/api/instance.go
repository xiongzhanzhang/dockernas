package api

import (
	"tinycloud/internal/models"

	"github.com/gin-gonic/gin"
)

func PostInstance(c *gin.Context) {
	var param models.InstanceParam

	c.BindJSON(&param)

	c.JSON(200, gin.H{
		"list": "ok",
	})
}
