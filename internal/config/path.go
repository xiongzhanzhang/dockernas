package config

import (
	"dockernas/internal/utils"
	"path/filepath"
	"strings"
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
	InitLogger()
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

func GetExtraAppPath() string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "apps")
	utils.CheckCreateDir(basePath)
	return basePath
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

func GetRelativePath(path string) string {
	if path[0] == '.' {
		return path
	}
	basePath := filepath.Join(GetBasePath(), "")
	if strings.Index(path, basePath) == 0 {
		return path[len(basePath):]
	}

	panic("can't change to relative path: " + path)
}

func GetAbsolutePath(path string) string {
	if path[0] == '.' {
		return path
	}
	return filepath.Join(GetBasePath(), path)
}

func GetAppMountFilePath(path string, fileName string) string {
	return filepath.Join(GetAbsolutePath(path), fileName)
}
