package server

import (
	"tinycloud/internal/api"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	apiv1 := router.Group("/api")
	{
		apiv1.GET("app", api.GetApps)
		apiv1.GET("app/:name", api.GetAppByName)

		apiv1.POST("instance", api.PostInstance)
		apiv1.GET("instance", api.GetInstance)
		apiv1.GET("instance/:name", api.GetInstanceByName)
		apiv1.PATCH("instance/:name", api.PatchInstance)
		apiv1.DELETE("instance/:name", api.DeleteInstance)
	}
}
