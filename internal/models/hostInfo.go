package models

type HostInfo struct {
	HostName         string `json:"hostname"`
	DockerNASVersion string `json:"dockerNASVersion"`
	Platform         string `json:"platform"`
	BootTime         uint64 `json:"bootTime"`
	ModelName        string `json:"modelName"`
	MemSize          uint64 `json:"memSize"`
}
