package config

import "tinycloud/internal/utils"

const configFilePath = "./config.json"

var configMap map[string]string = map[string]string{}

func InitConfig() {
	if utils.IsFileExist(configFilePath) {
		utils.GetObjFromJsonFile(configFilePath, &configMap)
	} else {
		SetConfig("user", "admin")
		SetConfig("passwd", utils.GenPasswd())
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
