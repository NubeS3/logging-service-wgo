package logic

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log-service-go/cmd/internals/models/eventstoredb"
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
			_ = eventstoredb.BucketLog(data)
		}()
	})
	return qsub
}
