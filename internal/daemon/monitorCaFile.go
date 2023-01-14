package daemon

import (
	"dockernas/internal/config"
	"dockernas/internal/service"
	"dockernas/internal/utils"
)

var lastCerModTime int64 = 0
var lastKeyModTime int64 = 0

func MonitorCaFile() {
	if config.GetIsHttpGatewayEnabled() == false {
		return
	}

	cer, key, msg := service.GetCaFilePathOnHost(config.GetCaFileDir())
	if msg != "" {
		return
	}

	cerModTime, _ := utils.GetFileModTimestamp(cer)
	keyModTime, _ := utils.GetFileModTimestamp(key)

	if lastCerModTime == 0 || lastKeyModTime == 0 {
		lastCerModTime = cerModTime
		lastKeyModTime = keyModTime
	} else {
		lastCerModTime = cerModTime
		lastKeyModTime = keyModTime

		service.RestartHttpGateway()
	}
}
