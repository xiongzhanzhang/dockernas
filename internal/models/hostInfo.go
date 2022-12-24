package models

type HostInfo struct {
	HostName  string `json:"hostname"`
	Platform  string `json:"platform"`
	BootTime  uint64 `json:"bootTime"`
	ModelName string `json:"modelName"`
	MemSize   uint64 `json:"memSize"`
}
