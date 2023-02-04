package models

type App struct {
	Name           string           `json:"name"`
	Category       []string         `json:"category"`
	Summary        string           `json:"summary"`
	Url            string           `json:"url"`
	Path           string           `json:"path"`
	UrlPrefix      string           `json:"urlPrefix"`
	IconUrl        string           `json:"iconUrl"`
	DockerVersions []DockerTemplate `json:"dockerVersions"`
}
