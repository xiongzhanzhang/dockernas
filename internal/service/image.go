package service

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"tinycloud/internal/backend/docker"
	"tinycloud/internal/models"
)

var mutex sync.Mutex
var pullingImageStateMap map[string]*models.ImageInfo = make(map[string]*models.ImageInfo)

func getImageName(imageUrl string) string {
	if strings.Contains(imageUrl, ":") == false {
		return imageUrl + ":latest"
	}
	return imageUrl
}

func ProcessImagePullMsg(imageUrl string, msg string) {
	log.Println(msg)

	var msgObj models.ImagePullMsg
	err := json.Unmarshal([]byte(msg), &msgObj)
	if err != nil || msgObj.ProgressDetail.Total == 0 {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	fullName := getImageName(imageUrl)
	if strings.Contains(msgObj.Status, "Downloading") {
		if _, ok := pullingImageStateMap[fullName]; !ok {
			pullingImageStateMap[fullName] = &models.ImageInfo{
				Name: fullName,
				Size: msgObj.ProgressDetail.Total,
			}
		}

		pullingImageStateMap[fullName].State =
			strconv.FormatInt(msgObj.ProgressDetail.Current*100/msgObj.ProgressDetail.Total, 10) + "%"
	}
	if strings.Contains(msgObj.Status, "Download complete") {
		delete(pullingImageStateMap, fullName)
	}
}

func ReportImagePullTimeout(imageUrl string) {
	mutex.Lock()
	defer mutex.Unlock()
	fullName := getImageName(imageUrl)
	delete(pullingImageStateMap, fullName)
}

func GetImages() []models.ImageInfo {
	mutex.Lock()
	defer mutex.Unlock()
	infos := docker.ListImage()
	for _, value := range pullingImageStateMap {
		infos = append(infos, *value)
	}
	return infos
}

func DelImage(info models.ImageInfo) {
	if info.Id == "" {
		ReportImagePullTimeout(info.Name)
	} else {
		docker.DelImage(info.Id)
	}
}
