package api

import (
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func GetDfsDir(c *gin.Context) {
	path := c.Query("path")
	dirInfoList := service.GetDfsDirInfo(path)
	c.JSON(200, gin.H{
		"list": dirInfoList,
	})
}
