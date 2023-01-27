package config

import (
	"dockernas/internal/utils"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	logPath := "./logs"
	if IsBasePathSet() {
		logPath = GetBasePath() + "/logs"
	}
	utils.CheckCreateDir(logPath)
	log.SetOutput(&lumberjack.Logger{
		Filename:   logPath + "/server.log",
		MaxSize:    64, // megabytes
		MaxBackups: 32,
		MaxAge:     30, //days
	})
}
