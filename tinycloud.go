package main

import (
	"tinycloud/internal/config"
	"tinycloud/internal/daemon"
	"tinycloud/internal/server"
)

func main() {
	config.InitConfig()
	daemon.StartBackgroundTask()
	server.StartServer()
}
