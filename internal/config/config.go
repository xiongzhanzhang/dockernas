package config

import (
	"dockernas/internal/utils"
)

const configFilePath = "./config.json"

var configMap map[string]string = map[string]string{}

func InitConfig() {
	if utils.IsFileExist(configFilePath) {
		if utils.GetObjFromJsonFile(configFilePath, &configMap) == nil {
			panic("read config file error")
		}
	} else {
		SetConfig("user", "admin")
		SetConfig("passwd", utils.GenPasswd())
		SetConfig("bindAddr", "0.0.0.0:8080")
		SaveConfig()
	}
}

func GetConfig(key string, defualt string) string {
	value, ok := configMap[key]
	if ok {
		return value
	}

	return defualt
}

func SetConfig(key string, value string) {
	configMap[key] = value

}

func SaveConfig() {
	utils.WriteFile(configFilePath, utils.GetJsonFromObj(configMap))
}
