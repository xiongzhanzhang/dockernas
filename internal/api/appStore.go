package api

import (
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
	app := service.GetAppByName(name, true)
	if app == nil {
		panic("cant get app " + name)
	}
	c.JSON(200, *app)
}

func GetAppImage(c *gin.Context) {
	path1 := c.Param("path1")
	path2 := c.Param("path2")

	c.File(service.GetIconPath(path1, path2))
}

func GetExtraAppByName(c *gin.Context) {
	name := c.Param("name")
	dir := c.Param("dir")
	app := service.GetAppByName(dir+"/"+name, true)
	if app == nil {
		panic("cant get app " + name)
	}
	c.JSON(200, *app)
}
