package logic

//
//import (
//	"encoding/json"
//	"log-service-go/cmd/internals/models/eventstoredb"
//	"time"
//
//	"github.com/nats-io/stan.go"
//)
//
//type FileLog struct {
//	EventLog    Event     `json:"event_log"`
//	Id          string    `json:"id"`
//	FId         string    `json:"f_id"`
//	Name        string    `json:"name"`
//	Size        int64     `json:"size"`
//	BucketId    string    `json:"bucket_id"`
//	ContentType string    `json:"content_type"`
//	UploadDate  time.Time `json:"upload_date"`
//	Path        string    `json:"path"`
//	IsHidden    bool      `json:"is_hidden"`
//}
//
//type StagingFileLog struct {
//	EventLog    Event     `json:"event_log"`
//	Name        string    `json:"name"`
//	Size        int64     `json:"size"`
//	BucketId    string    `json:"bucket_id"`
//	ContentType string    `json:"content_type"`
//	UploadDate  time.Time `json:"upload_date"`
//	Path        string    `json:"path"`
//	IsHidden    bool      `json:"is_hidden"`
//}
//
//type UploadSuccessFileLog struct {
//	EventLog    Event     `json:"event_log"`
//	Id          string    `json:"id"`
//	FId         string    `json:"f_id"`
//	Name        string    `json:"name"`
//	Size        int64     `json:"size"`
//	BucketId    string    `json:"bucket_id"`
//	ContentType string    `json:"content_type"`
//	UploadDate  time.Time `json:"upload_date"`
//	Path        string    `json:"path"`
//	IsHidden    bool      `json:"is_hidden"`
//}
//
//type FolderEvent struct {
//	EventLog Event `json:"event_log"`
//
//	Id       string `json:"-"`
//	OwnerId  string `json:"owner_id"`
//	Name     string `json:"name"`
//	Fullpath string `json:"fullpath"`
//}
//
//func GetFileDownloadedLogQsub() stan.Subscription {
//	qsub, _ := sc.QueueSubscribe(downloadFileSubj, "file-downloaded-log-qgroup", func(msg *stan.Msg) {
//		go func() {
//			var data FileLog
//			_ = json.Unmarshal(msg.Data, &data)
//			if data.EventLog.Type == "query" {
//				eventstoredb.Query("fileStream", data)
//			} else {
//				_ = eventstoredb.FileDownloadedLog(data)
//			}
//		}()
//	})
//	return qsub
//}
//
//func GetFileStagingLogQsub() stan.Subscription {
//	qsub, _ := sc.QueueSubscribe(stagingFileSubj, "file-staging-log-qgroup", func(msg *stan.Msg) {
//		go func() {
//			var data StagingFileLog
//			_ = json.Unmarshal(msg.Data, &data)
//			if data.EventLog.Type == "query" {
//				eventstoredb.Query("fileStream", data)
//			} else {
//				_ = eventstoredb.FileStagingLog(data)
//			}
//		}()
//	})
//	return qsub
//}
//
//func GetFileUploadedLogQsub() stan.Subscription {
//	qsub, _ := sc.QueueSubscribe(uploadFileSubj, "file-uploaded-log-qgroup", func(msg *stan.Msg) {
//		go func() {
//			var data FileLog
//			_ = json.Unmarshal(msg.Data, &data)
//			if data.EventLog.Type == "query" {
//				eventstoredb.Query("fileStream", data)
//			} else {
//				_ = eventstoredb.FileUploadedLog(data)
//			}
//		}()
//	})
//	return qsub
//}
//
//func GetFileUploadedSuccessLogQsub() stan.Subscription {
//	qsub, _ := sc.QueueSubscribe(uploadFileSuccessSubj, "file-uploaded-success-log-qgroup", func(msg *stan.Msg) {
//		go func() {
//			var data UploadSuccessFileLog
//			_ = json.Unmarshal(msg.Data, &data)
//			if data.EventLog.Type == "query" {
//				eventstoredb.Query("fileStream", data)
//			} else {
//				_ = eventstoredb.FileUploadedSuccessLog(data)
//			}
//		}()
//	})
//	return qsub
//}
//
//func GetFolderLogQsub() stan.Subscription {
//	qsub, _ := sc.QueueSubscribe(folderSubj, "folder-log-qgroup", func(msg *stan.Msg) {
//		go func() {
//			var data FolderEvent
//			_ = json.Unmarshal(msg.Data, &data)
//			if data.EventLog.Type == "query" {
//				eventstoredb.Query("folderStream", data)
//			} else {
//				_ = eventstoredb.FolderLog(data)
//			}
//		}()
//	})
//	return qsub
//}
