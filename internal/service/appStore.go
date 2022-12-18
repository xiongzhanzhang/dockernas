package service

import (
	"io/ioutil"
	"log"
	"strings"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func GetApps() []models.App {
	models.GetDb()
	return getAppsFromPath("./apps")
}

func GetAppByName(name string) models.App {
	var path = "./apps"
	var app models.App
	app.Name = name
	app.IconUrl = strings.Replace(path, "./", "/", 1) + "/" + name + "/icon.jpg"
	utils.GetObjFromJsonFile(path+"/"+name+"/introduction.json", &app)
	app.DockerVersions = getDockerTemplates(path + "/" + name + "/docker")

	return app
}

func getAppsFromPath(path string) []models.App {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("list dir error", err)
		return nil
	}

	apps := []models.App{}
	for _, fi := range dirs {
		if fi.IsDir() {
			app := GetAppByName(fi.Name())
			apps = append(apps, app)
		}
	}

	return apps
}

func getDockerTemplates(path string) []models.DockerTemplate {
	var dockerTemplates []models.DockerTemplate

	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("list dir error", err)
		return dockerTemplates
	}

	for _, fi := range dirs {
		if fi.IsDir() {
			var dockerTemplate models.DockerTemplate
			dockerTemplate.Version = fi.Name()
			if utils.GetObjFromJsonFile(path+"/"+fi.Name()+"/template.json", &dockerTemplate) != nil {
				dockerTemplates = append(dockerTemplates, dockerTemplate)
			}
		}
	}

	return dockerTemplates
}
