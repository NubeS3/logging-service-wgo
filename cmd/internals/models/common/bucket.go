package common

type BucketLog struct {
	Event
	UserId   string `json:"uid"`
	BucketId string `json:"bid"`
	Name     string `json:"name"`
	Region   string `json:"region"`
}
