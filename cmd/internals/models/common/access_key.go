package common

type AccessKeyLog struct {
	Event
	UserId   string `json:"uid"`
	BucketId string `json:"bucket_id"`
	Key      string `json:"key"`
}
