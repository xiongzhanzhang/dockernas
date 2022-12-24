package server

import (
	"tinycloud/internal/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.New()
	router.Use(
		gin.Logger(),
		middleware.Recovery(),
	)

	router.Static("/static", "./frontend/dist")
	router.Static("/apps", "./apps")
	router.StaticFile("/", "./frontend/dist/index.html")
	// router.NoRoute(func(ctx *gin.Context) {
	// 	ctx.Request.URL.Path = "/"
	// 	router.HandleContext(ctx)
	// })

	registerRoutes(router)

	router.Run()
}
