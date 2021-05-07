package common

type BandwidthLog struct {
	Event
	Size     int64  `json:"size"`
	BucketId string `json:"bucket_id"`
	Uid      string `json:"uid"`
	From     string `json:"from"`
}
