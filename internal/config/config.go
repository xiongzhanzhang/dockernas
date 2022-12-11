package config

import (
	"os"
	"tinycloud/internal/utils"
)

func GetBasePath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/data"
	return path
}

func GetDBFilePath() string {
	basePath := GetBasePath()
	basePath += "/meta"
	utils.CheckCreateDir(basePath)
	return basePath + "/data.db3"
}

func GetLocalVolumePath(instanceName string, volumeName string) string {
	basePath := GetBasePath()
	basePath += "/local/" + instanceName + "/" + volumeName
	utils.CheckCreateDir(basePath)
	return basePath
}
