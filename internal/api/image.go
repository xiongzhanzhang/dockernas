package api

import (
	"tinycloud/internal/models"
	"tinycloud/internal/service"

	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	images := service.GetImages()
	c.JSON(200, gin.H{
		"list": images,
	})
}

func DelImage(c *gin.Context) {
	var param models.ImageInfo
	c.BindJSON(&param)

	service.DelImage(param)

	c.JSON(200, gin.H{
		"msg": "ok",
	})
}
