package server

import (
	"tinycloud/internal/api"
	"tinycloud/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	apiv1 := router.Group("/api", middleware.Authentication(), middleware.PathCheck())
	{
		apiv1.GET("app", api.GetApps)
		apiv1.GET("app/:name", api.GetAppByName)

		apiv1.POST("instance", api.PostInstance)
		apiv1.GET("instance", api.GetInstance)
		apiv1.GET("instance/:name", api.GetInstanceByName)
		apiv1.PATCH("instance/:name", api.PatchInstance)
		apiv1.DELETE("instance/:name", api.DeleteInstance)

		apiv1.GET("instance/:name/log", api.GetInstanceLog)
		apiv1.GET("instance/:name/event", api.GetInstanceEvent)

		apiv1.GET("instancestats", api.GetAllInstanceStats)
		apiv1.GET("instancestats/:name", api.GetInstanceStatsByName)
		apiv1.GET("instancehttpport", api.GetInstanceHttpPorts)

		apiv1.GET("dfsdir", api.GetDfsDir)
		apiv1.GET("systemdir", api.GetSystemDir)
		apiv1.POST("basepath", api.SetBasePath)

		apiv1.GET("host", api.GetHostInfo)
		apiv1.GET("storage", api.GetStorageInfo)
		apiv1.GET("network", api.GetNetworkInfo)
		apiv1.GET("httpproxyconfig", api.GetHttpProxyConfig)
		apiv1.POST("httpproxyconfig", api.CreateHttpProxyConfig)
		apiv1.DELETE("httpproxyconfig/:hostname", api.DelHttpProxyConfig)
		apiv1.POST("domain", api.EditDomain)

		apiv1.POST("login", api.Login)
	}
}
