package server

import (
	"dockernas/internal/api"
	"dockernas/internal/middleware"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func initStaticFileRouter(router *gin.Engine) {
	router.StaticFile("/", "./frontend/dist/index.html")
	router.Static("/assets", "./frontend/dist/assets")
	router.Static("/apps", "./apps")

	dir1, err1 := ioutil.ReadDir("./frontend/dist")
	if err1 != nil {
		log.Println("list dir error", err1)
	} else {
		for _, fi1 := range dir1 {
			if !fi1.IsDir() {
				router.StaticFile("/"+fi1.Name(), "./frontend/dist/"+fi1.Name())
			}
		}
	}
}

func registerRoutes(router *gin.Engine) {
	router.GET("/terminal", api.InstanceWebTerminal)

	apiv1 := router.Group("/api", middleware.Authentication(), middleware.PathCheck())
	{
		apiv1.GET("app", api.GetApps)
		apiv1.GET("app/:name", api.GetAppByName)
		apiv1.GET("extra/app/:dir/:name", api.GetExtraAppByName)

		apiv1.GET("image", api.GetImages)
		apiv1.DELETE("image", api.DelImage)

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
		apiv1.GET("network/ipv4", api.GetHostIpv4)
		apiv1.GET("network/ipv6", api.GetHostIpv6)
		apiv1.GET("httpproxyconfig", api.GetHttpProxyConfig)
		apiv1.POST("httpproxyconfig", api.CreateHttpProxyConfig)
		apiv1.DELETE("httpproxyconfig/:hostname", api.DelHttpProxyConfig)
		apiv1.POST("domain", api.EditDomain)
		apiv1.POST("httpgateway", api.EnableHttpGateway)
		apiv1.PATCH("httpgateway", api.PatchHttpGateway)
		apiv1.POST("httpgateway/https", api.EnableHttps)
		apiv1.DELETE("httpgateway/https", api.DisableHttps)
		apiv1.POST("httpgateway/capath", api.SetCaFileDir)

		apiv1.GET("subscribe", api.GetSubscribes)
		apiv1.PATCH("subscribe", api.UpdateSubscribe)
		apiv1.POST("subscribe", api.AddSubscribe)
		apiv1.DELETE("subscribe/:name", api.DelSubscribe)

		apiv1.POST("login", api.Login)
	}
}
