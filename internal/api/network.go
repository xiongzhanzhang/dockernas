package api

import (
	"tinycloud/internal/models"
	"tinycloud/internal/service"
	"tinycloud/internal/utils"

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
	config := models.GetHttpProxyConfigByHostName(hostName)

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

func EditDomain(c *gin.Context) {
	postMap := map[string]string{}
	c.BindJSON(&postMap)
	domain, ok := postMap["domain"]
	if ok == false {
		panic("EditDomain without a param")
	}

	service.SetDomain(domain)

	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func EnableHttpGateway(c *gin.Context) {
	service.EnableHttpGateway()
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func PatchHttpGateway(c *gin.Context) {
	patchMap := map[string]string{}
	c.BindJSON(&patchMap)

	op, _ := patchMap["op"]
	if op == "restart" {
		service.RestartHttpGateway()
	} else if op == "stop" {
		service.StopHttpGateway()
	} else {
		panic("unkown patch request")
	}

	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func EnableHttps(c *gin.Context) {
	service.EnableHttps()
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func DisableHttps(c *gin.Context) {
	service.DisableHttps()
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func SetCaFileDir(c *gin.Context) {
	postMap := map[string]string{}
	c.BindJSON(&postMap)
	path, ok := postMap["path"]
	if ok == false {
		panic("SetCaFilePath without a param")
	}

	service.SetCaFileDir(path)

	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func GetHostIpv4(c *gin.Context) {
	c.String(200, utils.GetLocalAddress())
}

func GetHostIpv6(c *gin.Context) {
	c.String(200, utils.GetLocalAddressIpv6())
}
