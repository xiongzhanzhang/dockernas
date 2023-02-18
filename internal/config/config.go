package config

import (
	"dockernas/internal/utils"
	"os"
)

var configMap map[string]string = map[string]string{}

func IsRunInConainer() bool {
	return os.Getenv("DOCKERNAS_RUN_IN_CONTAINER") == "true"
}

func getConfigPath() string {
	if IsRunInConainer() {
		utils.CheckCreateDir("./data")
		return "./data/config.json"
	}
	return "./config.json"
}

func InitConfig() {
	if utils.IsFileExist(getConfigPath()) {
		if utils.GetObjFromJsonFile(getConfigPath(), &configMap) == nil {
			panic("read config file error")
		}
	} else {
		SetConfig("user", "admin")
		SetConfig("passwd", utils.GenPasswd())
		SetConfig("bindAddr", "0.0.0.0:8080")
		if IsRunInConainer() {
			SetConfig("basePath", "/home/dockernas/data")
		}
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
	utils.WriteFile(getConfigPath(), utils.GetJsonFromObj(configMap))
}

func GetDockerNASVersion() string {
	return "0.3.0"
}
