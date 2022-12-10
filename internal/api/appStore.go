package api

import (
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func GetApps(c *gin.Context) {
	apps := service.GetApps()
	c.JSON(200, gin.H{
		"list": apps,
	})
}
