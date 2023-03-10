package service

import (
	"dockernas/internal/backend/docker"
	"dockernas/internal/config"
	"dockernas/internal/models"
	"dockernas/internal/utils"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var appMap = map[string]models.App{}

func GetApps() []models.App {
	apps := []models.App{}

	dir1, err1 := ioutil.ReadDir("./apps")
	if err1 != nil {
		log.Println("list dir error", err1)
	} else {
		for _, fi1 := range dir1 {
			if fi1.IsDir() {
				dir2, err2 := ioutil.ReadDir(filepath.Join("./apps", fi1.Name()))
				if err2 != nil {
					log.Println("list dir error", err2)
				} else {
					for _, fi2 := range dir2 {
						if fi2.IsDir() {
							app := GetAppByNameAndPath(fi2.Name(), "./apps/"+fi1.Name()+"/"+fi2.Name())
							if app != nil {
								apps = append(apps, *app)
							}
						}
					}
				}
			}
		}
	}

	dir1, err1 = ioutil.ReadDir(config.GetExtraAppPath())
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
							app := GetAppByNameAndPath(fi1.Name()+"/"+fi2.Name(), config.GetExtraAppPath()+"/"+fi1.Name()+"/"+fi2.Name())
							if app != nil {
								apps = append(apps, *app)
							}
						}
					}
				}
			}
		}
	}

	for k := range appMap {
		delete(appMap, k)
	}
	for _, app := range apps {
		appMap[app.Name] = app
	}

	return apps
}

func GetAppByName(name string, flush bool) *models.App {
	app, ok := appMap[name]
	if ok {
		return GetAppByNameAndPath(app.Name, config.GetAbsolutePath(app.Path)) //get lastest data on disk
	}
	if !flush {
		return nil
	}

	GetApps()
	return GetAppByName(name, false)
}

func GetAppByNameAndPath(name string, path string) *models.App {
	var app models.App
	app.IconUrl = "/api/icon?path=" + config.GetRelativePath(path) + "/icon.jpg"
	app.DockerVersions = getDockerTemplates(path + "/docker")
	if len(app.DockerVersions) == 0 {
		return nil
	}
	if utils.GetObjFromJsonFile(path+"/introduction.json", &app) == nil {
		return nil
	}
	app.Name = name
	app.Path = config.GetRelativePath(path)

	return &app
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
				if dockerTemplate.OSList != "" &&
					strings.Contains(dockerTemplate.OSList, docker.DetectRealSystem()) == false {
					continue
				}
				dockerTemplate.Path = config.GetRelativePath(path) + "/" + fi.Name()
				dockerTemplates = append(dockerTemplates, dockerTemplate)
			} else {
				log.Println("load template error for " + fi.Name() + " under " + path)
			}
		}
	}

	return dockerTemplates
}
