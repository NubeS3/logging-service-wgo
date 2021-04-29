package common

import "time"

type FileLog struct {
	Event
	FId         string    `json:"file_id"`
	FileName    string    `json:"file_name"`
	Size        int64     `json:"size"`
	BucketId    string    `json:"bucket_id"`
	ContentType string    `json:"content_type"`
	UploadDate  time.Time `json:"upload_date"`
	Path        string    `json:"path"`
	IsHidden    bool      `json:"is_hidden"`
}
