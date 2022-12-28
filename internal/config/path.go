package config

import (
	"os"
	"path/filepath"
	"tinycloud/internal/utils"
)

// func GetBasePath() string {
// 	basePath, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}
// 	basePath = filepath.Join(basePath, "data")
// 	return basePath
// }

func GetBasePath() string {
	basePath := GetConfig("basePath", "")
	if basePath == "" {
		panic("base data path is not set")
	}
	return basePath
}

func IsBasePathSet() bool {
	return GetConfig("basePath", "") != ""
}

func SetBasePath(path string) {
	if IsBasePathSet() {
		panic("base data path has set")
	}
	SetConfig("basePath", path)
	SaveConfig()
}

func GetFullDfsPath(path string) string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "dfs", path)
	return basePath
}

func GetDBFilePath() string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "meta")
	utils.CheckCreateDir(basePath)
	return filepath.Join(basePath, "data.db3")
}

func GetAppLocalPath(instanceName string) string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "local", instanceName)
	return basePath
}

func GetAppLocalFilePath(instanceName string, fileName string) string {
	return filepath.Join(GetAppLocalPath(instanceName), fileName)
}

func GetLocalVolumePath(instanceName string, volumeName string) string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "local", instanceName, volumeName)
	utils.CheckCreateDir(basePath)
	return basePath
}

func GetAppMountFilePath(appName string, version string, fileName string) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(pwd, "apps", appName, "docker", version, fileName)
}
