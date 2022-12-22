package service

import (
	"io/ioutil"
	"log"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
)

func GetDfsDirInfo(path string) []models.DirInfo {
	basePath := config.GetFullDfsPath(path)
	dirInfoList := []models.DirInfo{}

	dirs, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Println("list dir error", err)
		panic(err)
	}

	relativePath := path + "/"
	if path == "/" {
		relativePath = "/"
	}

	for _, fi := range dirs {
		if fi.IsDir() {
			var dirInfo models.DirInfo
			dirInfo.Name = fi.Name()
			dirInfo.Label = relativePath + fi.Name()
			dirInfo.Value = dirInfo.Label
			dirInfoList = append(dirInfoList, dirInfo)
		}
	}

	return dirInfoList
}
