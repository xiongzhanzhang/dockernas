package service

import (
	"dockernas/internal/backend/docker"
	"dockernas/internal/models"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
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

		pullingImageStateMap[fullName].Size = msgObj.ProgressDetail.Total
		pullingImageStateMap[fullName].State =
			strconv.FormatInt(msgObj.ProgressDetail.Current*100/msgObj.ProgressDetail.Total, 10) + "%"
	}
	if strings.Contains(msgObj.Status, "Download complete") {
		delete(pullingImageStateMap, fullName)
	}
}

func ReportImagePullStoped(imageUrl string) {
	mutex.Lock()
	defer mutex.Unlock()
	fullName := getImageName(imageUrl)
	delete(pullingImageStateMap, fullName)
}

func GetImagePullState(imageUrl string) string {
	mutex.Lock()
	defer mutex.Unlock()
	fullName := getImageName(imageUrl)
	if _, ok := pullingImageStateMap[fullName]; ok {
		return pullingImageStateMap[fullName].State
	}
	return ""
}

func GetImages() []models.ImageInfo {
	mutex.Lock()
	defer mutex.Unlock()
	var infos []models.ImageInfo
	for _, value := range pullingImageStateMap {
		infos = append(infos, *value)
	}
	infos = append(infos, docker.ListImage()...)
	return infos
}

func DelImage(info models.ImageInfo) {
	if info.Id == "" {
		ReportImagePullStoped(info.Name)
	} else {
		docker.DelImage(info.Id)
	}
}
