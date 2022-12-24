package main

import (
	"tinycloud/internal/config"
	"tinycloud/internal/server"
)

func main() {
	config.InitConfig()
	server.StartServer()
}
