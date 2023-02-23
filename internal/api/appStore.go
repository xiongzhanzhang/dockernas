package api

import (
	"dockernas/internal/config"
	"dockernas/internal/service"

	"github.com/gin-gonic/gin"
)

func GetApps(c *gin.Context) {
	apps := service.GetApps()
	c.JSON(200, gin.H{
		"list": apps,
	})
}

func GetAppByName(c *gin.Context) {
	name := c.Param("name")
	if name[0] == '/' {
		name = name[1:]
	}
	app := service.GetAppByName(name, true)
	if app == nil {
		panic("cant get app " + name)
	}
	c.JSON(200, *app)
}

func GetIcon(c *gin.Context) {
	iconPath := c.Query("path")
	c.File(config.GetAbsolutePath(iconPath))
}
