package common

type AccessKeyLog struct {
	Event
	UserId   string `json:"uid"`
	BucketId string `json:"bid"`
	Key      string `json:"id"`
}
