package server

import (
	"tinycloud/internal/api"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	apiv1 := router.Group("/api")
	{
		apiv1.GET("app", api.GetApps)
	}
}
