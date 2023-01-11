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

func GetAppByName(c *gin.Context) {
	name := c.Param("name")
	app := service.GetAppByName(name)
	if app == nil {
		panic("cant get app " + name)
	}
	c.JSON(200, *app)
}

func GetExtraAppByName(c *gin.Context) {
	name := c.Param("name")
	dir := c.Param("dir")
	app := service.GetAppByName(dir + "/" + name)
	if app == nil {
		panic("cant get app " + name)
	}
	c.JSON(200, *app)
}
