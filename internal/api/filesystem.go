package api

import (
	"dockernas/internal/service"

	"github.com/gin-gonic/gin"
)

func GetDfsDir(c *gin.Context) {
	path := c.Query("path")
	dirInfoList := service.GetDfsDirInfo(path)
	c.JSON(200, gin.H{
		"list": dirInfoList,
	})
}

func GetSystemDir(c *gin.Context) {
	path := c.Query("path")
	dirInfoList := service.GetSystemDirInfo(path)
	c.JSON(200, gin.H{
		"list": dirInfoList,
	})
}

func SetBasePath(c *gin.Context) {
	postMap := map[string]string{}
	c.BindJSON(&postMap)
	path, ok := postMap["path"]
	if ok == false {
		panic("SetBasePath must has a path param")
	}

	service.SetBasePath(path)
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}
