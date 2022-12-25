package models

type StorageInfo struct {
	BaseDir         string           `json:"baseDir"`
	Device          string           `json:"device"`
	Fstype          string           `json:"fstype"`
	Capacity        int64            `json:"capacity"`
	FreeSize        int64            `json:"freeSize"`
	LocalSize       int64            `json:"localSize"`
	DfsSize         int64            `json:"dfsSize"`
	OtherSize       int64            `json:"otherSize"`
	InstanceSizeMap map[string]int64 `json:"instanceSizeMap"`
}
