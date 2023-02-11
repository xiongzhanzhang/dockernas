package server

import (
	"dockernas/internal/config"
	"dockernas/internal/middleware"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	gin.DisableConsoleColor()
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()

	router := gin.New()
	router.Use(
		gin.Logger(),
		middleware.Recovery(),
	)

	initStaticFileRouter(router)

	// if config.IsBasePathSet() {
	// 	router.Static("/extra/apps", config.GetExtraAppPath())
	// }
	
	router.NoRoute(func(ctx *gin.Context) {
		if strings.Index(ctx.Request.URL.Path, "/index/") == 0 ||
			ctx.Request.URL.Path == "/login" ||
			ctx.Request.URL.Path == "/basepath" {
			ctx.Request.URL.Path = "/"
			router.HandleContext(ctx)
		}
	})

	registerRoutes(router)
	router.Run(config.GetBindAddr())
}
