package server

import "github.com/gin-gonic/gin"

func registerRoutes(router *gin.Engine){
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}