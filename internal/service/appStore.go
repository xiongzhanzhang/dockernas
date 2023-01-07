package service

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"tinycloud/internal/config"
	"tinycloud/internal/models"
	"tinycloud/internal/utils"
)

func GetApps() []models.App {
	apps := []models.App{}

	dirs, err := ioutil.ReadDir("./apps")
	if err != nil {
		log.Println("list dir error", err)
	} else {
		for _, fi := range dirs {
			if fi.IsDir() {
				app := GetAppByName(fi.Name())
				apps = append(apps, app)
			}
		}
	}

	dir1, err1 := ioutil.ReadDir(config.GetExtraAppPath())
	if err1 != nil {
		log.Println("list dir error", err1)
	} else {
		for _, fi1 := range dir1 {
			if fi1.IsDir() {
				dir2, err2 := ioutil.ReadDir(filepath.Join(config.GetExtraAppPath(), fi1.Name()))
				if err2 != nil {
					log.Println("list dir error", err2)
				} else {
					for _, fi2 := range dir2 {
						if fi2.IsDir() {
							app := GetAppByName(fi1.Name() + "/" + fi2.Name())
							apps = append(apps, app)
						}
					}
				}
			}
		}
	}

	return apps
}

func GetAppByName(name string) models.App {
	var app *models.App
	if !strings.Contains(name, "/") {
		app = GetAppByNameAndPath(name, "./apps", "/apps")
	} else {
		app = GetAppByNameAndPath(name, config.GetExtraAppPath(), "/extra/apps")
	}

	if app == nil {
		panic("can't get app " + name)
	}
	return *app
}

func GetAppByNameAndPath(name string, path string, urlPrefix string) *models.App {
	var app models.App
	app.IconUrl = urlPrefix + "/" + name + "/icon.jpg"
	app.DockerVersions = getDockerTemplates(path + "/" + name + "/docker")
	if utils.GetObjFromJsonFile(path+"/"+name+"/introduction.json", &app) == nil {
		return nil
	}
	app.Name = name

	return &app
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
