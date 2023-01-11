package models

type ImageInfo struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Size  int64  `json:"size"`
	State string `json:"state"`
}

type ImagePullProgressDetail struct {
	Current int64 `json:"current"`
	Total   int64 `json:"total"`
}

type ImagePullMsg struct {
	Status         string                  `json:"status"`
	ProgressDetail ImagePullProgressDetail `json:"progressDetail"`
}
