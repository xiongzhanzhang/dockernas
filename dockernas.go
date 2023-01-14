package main

import (
	"dockernas/internal/config"
	"dockernas/internal/daemon"
	"dockernas/internal/server"
)

func main() {
	config.InitConfig()
	daemon.StartBackgroundTask()
	server.StartServer()
}
