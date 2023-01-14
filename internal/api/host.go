package api

import (
	"dockernas/internal/service"

	"github.com/gin-gonic/gin"
)

func GetHostInfo(c *gin.Context) {
	hostData := service.GetHostInfo()
	c.JSON(200, hostData)
}

func GetStorageInfo(c *gin.Context) {
	storageInfo := service.GetStorageInfo()
	c.JSON(200, storageInfo)
}
