package common

type BucketLog struct {
	Event
	UserId   string `json:"uid"`
	BucketId string `json:"id"`
	Name     string `json:"name"`
	Region   string `json:"region"`
}
