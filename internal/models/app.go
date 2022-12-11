package models

type App struct {
	Name           string           `json:"name"`
	Category       []string         `json:"category"`
	Summary        string           `json:"summary"`
	IconUrl        string           `json:"iconUrl"`
	DockerVersions []DockerTemplate `json:"dockerVersions"`
}


