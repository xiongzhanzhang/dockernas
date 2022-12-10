package models

type App struct {
	Name           string           `json:"name"`
	Category       []string         `json:"category"`
	Summary        string           `json:"summary"`
	IconUrl        string           `json:"iconUrl"`
	DockerVersions []DockerTemplate `json:"dockerVersions"`
}

type ParamItem struct {
	Prompt string `json:"prompt"`
	Key    string `json:"key"`   // resource in contariner
	Value  string `json:"value"` // resource in host
	Reg    string `json:"reg"`   // Regular for input
}

type DockerTemplate struct {
	ImageUrl    string      `json:"imageUrl"`
	Version     string      `json:"c"`
	PortParams  []ParamItem `json:"portParams"`
	EnvParams   []ParamItem `json:"envParams"`
	LocalVolume []ParamItem `json:"localVolume"`
	DfsVolume   []ParamItem `json:"dfsVolume"`
}
