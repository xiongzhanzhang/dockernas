package main

import (
	"dockernas/internal/config"
	"dockernas/internal/daemon"
	"dockernas/internal/server"
)

func main() {
	config.InitConfig()
	config.InitLogger()
	daemon.StartBackgroundTask()
	server.StartServer()
}
