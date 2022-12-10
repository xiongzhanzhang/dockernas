package models

type App struct {
	Name           string
	Category       []string
	Summary        string
	IconUrl        string
	DockerVersions []DockerTemplate
}

type ParamItem struct {
	Prompt string
	Key    string // resource in contariner
	Value  string // resource in host
	Reg    string // Regular for input
}

type DockerTemplate struct {
	ImageUrl    string
	Version     string
	PortParams  []ParamItem
	EnvParams   []ParamItem
	LocalVolume []ParamItem
	DfsVolume   []ParamItem
}
