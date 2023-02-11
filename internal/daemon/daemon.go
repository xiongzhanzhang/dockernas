package daemon

import (
	"log"
	"time"
)

func runTaskSafe(aFun func()) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("run background task error:", err)
		}
	}()
	aFun()
}

func runTaskEveryMinute(aFun func()) {
	for {
		sleepSecond := 61 - time.Now().Unix()%60
		time.Sleep(time.Duration(sleepSecond) * time.Second)
		runTaskSafe(aFun)
	}
}

func StartBackgroundTask() {
	go runTaskEveryMinute(monitorContainer)
	// go runTaskEveryMinute(MonitorCaFile)
	// time.Sleep(1000000 * time.Second)
}
