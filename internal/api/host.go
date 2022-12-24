package api

import (
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func GetHostInfo(c *gin.Context) {
	hostData := service.GetHostInfo()
	c.JSON(200, hostData)
}
