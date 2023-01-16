package config

import (
	"dockernas/internal/utils"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	utils.CheckCreateDir("./logs")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/server.log",
		MaxSize:    64, // megabytes
		MaxBackups: 32,
		MaxAge:     30, //days
	})
}
