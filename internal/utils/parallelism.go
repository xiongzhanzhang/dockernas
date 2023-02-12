package utils

import (
	"log"
	"time"
)

func RunTaskSafe(aFun func()) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("run background task error:", err)
		}
	}()
	aFun()
}

func RunBackgroundTaskSafe(aFun func(), delay time.Duration) {
	go func() {
		time.Sleep(delay)
		RunTaskSafe(aFun)
	}()
}
