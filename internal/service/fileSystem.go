package service

import (
	"dockernas/internal/config"
	"dockernas/internal/models"
	"io/ioutil"
	"log"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

func getDirInfo(fullPath string, relativePath string) []models.DirInfo {
	dirInfoList := []models.DirInfo{}

	dirs, err := ioutil.ReadDir(fullPath)
	if err != nil {
		log.Println("list dir error", err)
		panic(err)
	}

	for _, fi := range dirs {
		if fi.IsDir() {
			var dirInfo models.DirInfo
			dirInfo.Name = fi.Name()
			if strings.HasSuffix(relativePath, "/") {
				dirInfo.Label = relativePath + fi.Name()
			} else {
				dirInfo.Label = relativePath + "/" + fi.Name()
			}
			dirInfo.Value = dirInfo.Label
			dirInfoList = append(dirInfoList, dirInfo)
		}
	}

	return dirInfoList
}

func GetDfsDirInfo(path string) []models.DirInfo {
	basePath := config.GetFullDfsPath(path)
	return getDirInfo(basePath, path)
}

func GetSystemDirInfo(path string) []models.DirInfo {
	if path == "" {
		dirInfoList := []models.DirInfo{}
		infos, err := disk.Partitions(false)
		if err != nil {
			panic(err)
		}
		for _, info := range infos {
			var dirInfo models.DirInfo
			dirInfo.Name = info.Mountpoint
			dirInfo.Label = info.Mountpoint + "/"
			dirInfo.Value = info.Mountpoint + "/"
			dirInfoList = append(dirInfoList, dirInfo)
		}
		return dirInfoList
	} else {
		return getDirInfo(path, path)
	}

}

func SetBasePath(path string) {
	config.SetBasePath(path)
}
