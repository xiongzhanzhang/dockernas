package api

import (
	"tinycloud/internal/models"
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func GetNetworkInfo(c *gin.Context) {
	networkInfo := service.GetNetworkInfo()
	c.JSON(200, networkInfo)
}

func GetInstanceHttpPorts(c *gin.Context) {
	ports := service.GetInstanceHttpPorts()
	c.JSON(200, gin.H{
		"list": ports,
	})
}

func CreateHttpProxyConfig(c *gin.Context) {
	var param models.HttpProxyConfig
	c.BindJSON(&param)

	service.CreateHttpProxyConfig(param)

	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func DelHttpProxyConfig(c *gin.Context) {
	hostName := c.Param("hostname")
	config :=models.GetHttpProxyConfigByHostName(hostName)

	service.DelHttpProxyConfig(config)

	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func GetHttpProxyConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"list": models.GetHttpProxyConfig(),
	})
}
