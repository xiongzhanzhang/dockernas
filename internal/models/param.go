package models

type ParamItem struct {
	Prompt    string `json:"prompt"`
	Name      string `json:"name"`      // used for resource identification like local volume
	Key       string `json:"key"`       // resource in contariner
	Value     string `json:"value"`     // resource in host
	Reg       string `json:"reg"`       // Regular for input
	Hide      bool   `json:"hide"`      // is show on frontend
	Protocol  string `json:"protocol"`  // network protocol, used by port param
	MountFile bool   `json:"mountFile"` // mount file to container
	OtherType string `json:"otherType"` // other param type
	Passwd    bool   `json:"passwd"`    // is a password config
}

type DockerTemplate struct {
	ImageUrl    string      `json:"imageUrl"`
	Version     string      `json:"version"`
	PortParams  []ParamItem `json:"portParams"`
	EnvParams   []ParamItem `json:"envParams"`
	LocalVolume []ParamItem `json:"localVolume"`
	DfsVolume   []ParamItem `json:"dfsVolume"`
	OtherParams []ParamItem `json:"otherParams"`
	Privileged  bool        `json:"privileged"`
	Cmd         string      `json:"cmd"`
	OSList      string      `json:"osList"`
}

type InstanceParam struct {
	Name     string `json:"name"`
	AppName  string `json:"appName"`
	Summary  string `json:"summary"`
	IconUrl  string `json:"iconUrl"`
	HostOnly bool   `json:"hostOnly"`  // can only access by 127.0.0.1
	DockerTemplate
}

const (
	PLACEHOLDER_PARAM = "placeholder"
)
