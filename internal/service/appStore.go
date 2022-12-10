package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"tinycloud/internal/models"
)

func GetApps() []models.App {
	return getAppsFromPath("./apps")
}

func getAppsFromPath(path string) []models.App {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Print("list dir error", err)
		return nil
	}

	apps := []models.App{}
	for _, fi := range dirs {
		if fi.IsDir() {
			var app models.App
			app.Name = fi.Name()

			introduction, error := ioutil.ReadFile(path + "/" + fi.Name() + "/introduction.json")
			if error == nil {
				err := json.Unmarshal([]byte(introduction), &app)
				if err != nil {
					log.Print("read json error", err)
				}
			}

			app.ImgUrl = strings.Replace(path, "./", "/", 1) + "/" + fi.Name() + "/icon.jpg"
			apps = append(apps, app)
		}
	}

	return apps
}
