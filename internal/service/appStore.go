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

func getAppsFromPath(path string) []models.App {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("list dir error", err)
		return nil
	}

	apps := []models.App{}
	for _, fi := range dirs {
		if fi.IsDir() {
			var app models.App
			app.Name = fi.Name()
			app.IconUrl = strings.Replace(path, "./", "/", 1) + "/" + fi.Name() + "/icon.jpg"
			utils.GetObjFromJsonFile(path+"/"+fi.Name()+"/introduction.json", &app)
			app.DockerVersions = getDockerTemplates(path + "/" + fi.Name() + "/docker")
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
