package common

import "time"

type FileLog struct {
	Event
	Id         string    `json:"id"`
	FId        string    `json:"file_id"`
	FileName   string    `json:"file_name"`
	Size       int64     `json:"size"`
	BucketId   string    `json:"bucket_id"`
	UploadDate time.Time `json:"upload_date"`
	Uid        string    `json:"uid"`
}

type SumAggregateRes struct {
	Value float64 `json:"value"`
}
