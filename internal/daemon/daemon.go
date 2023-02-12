package daemon

import (
	"dockernas/internal/utils"
	"time"
)

func runTaskEveryMinute(aFun func()) {
	for {
		sleepSecond := 61 - time.Now().Unix()%60
		time.Sleep(time.Duration(sleepSecond) * time.Second)
		utils.RunTaskSafe(aFun)
	}
}

func StartBackgroundTask() {
	go runTaskEveryMinute(monitorContainer)
	go runTaskEveryMinute(MonitorCaFile)
	// time.Sleep(1000000 * time.Second)
}
