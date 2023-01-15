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
		MaxSize:    128, // megabytes
		MaxBackups: 128,
		MaxAge:     30, //days
	})
}
