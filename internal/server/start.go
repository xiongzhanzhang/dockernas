package server

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()
	registerRoutes(router)
	router.Run()
}
