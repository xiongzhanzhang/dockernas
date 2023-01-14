package api

import (
	"dockernas/internal/models"
	"dockernas/internal/service"

	"github.com/gin-gonic/gin"
)

func AddSubscribe(c *gin.Context) {
	var param models.Subscribe
	c.BindJSON(&param)
	service.AddSubscribe(param)
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func GetSubscribes(c *gin.Context) {
	c.JSON(200, gin.H{
		"list": models.GetSubscribe(),
	})
}

func DelSubscribe(c *gin.Context) {
	name := c.Param("name")
	service.DelSubscribe(name)
	c.JSON(200, gin.H{
		"state": "ok",
	})
}

func UpdateSubscribe(c *gin.Context) {
	service.UpdateSubscribe()
	c.JSON(200, gin.H{
		"state": "ok",
	})
}
