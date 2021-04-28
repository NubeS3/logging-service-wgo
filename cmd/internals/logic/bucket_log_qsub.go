package logic

import (
	"encoding/json"
	"log-service-go/cmd/internals/models/eventstoredb"

	"github.com/nats-io/stan.go"
)

type BucketLog struct {
	EventLog Event  `json:"event_log"`
	Id       string `json:"id"`
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Region   string `json:"region"`
}

func GetBucketLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(bucketSubj, "bucket-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data BucketLog
			_ = json.Unmarshal(msg.Data, &data)
			if data.EventLog.Type == "query" {
				eventstoredb.Query("bucketStream", data)
			} else {
				_ = eventstoredb.BucketLog(data)
			}
		}()
	})
	return qsub
}
