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

	cer, key, msg := utils.GetCaFilePathOnHost(config.GetFullDfsPath(config.GetCaFileDir()), config.GetDomain())
	if msg != "" {
		return
	}

	cerModTime, _ := utils.GetFileModTimestamp(cer)
	keyModTime, _ := utils.GetFileModTimestamp(key)

	if lastCerModTime == 0 || lastKeyModTime == 0 {
		lastCerModTime = cerModTime
		lastKeyModTime = keyModTime
	} else if lastCerModTime < cerModTime || lastKeyModTime < keyModTime {
		lastCerModTime = cerModTime
		lastKeyModTime = keyModTime
		service.RestartHttpGateway()
	}
}
