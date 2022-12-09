package api

import (
	"tinycloud/internal/models"

	"github.com/gin-gonic/gin"
)

func GetApps(c *gin.Context) {
	var app models.App
	app.Category = "数据库"
	app.ImgUrl = "https://img-blog.csdnimg.cn/img_convert/1db846beadb17f6e4f919d259ee05a1b.png"
	app.Name = "mysql"

	c.JSON(200, gin.H{
		"list": []models.App{app},
	})
}
