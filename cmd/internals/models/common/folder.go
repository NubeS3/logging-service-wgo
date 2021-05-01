package common

type FolderLog struct {
	Event
	Id       string `json:"id"`
	UserId   string `json:"uid"`
	Name     string `json:"name"`
	FullPath string `json:"full_path"`
}
