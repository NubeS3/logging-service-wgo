package logic

import (
	"encoding/json"
	"log-service-go/cmd/internals/models/eventstoredb"
	"time"

	"github.com/nats-io/stan.go"
)

type AccessKeyLog struct {
	EventLog Event `json:"event_log"`

	Key         string    `json:"key"`
	BucketId    string    `json:"bucket_id"`
	ExpiredDate time.Time `json:"expired_date"`
	Permissions []string  `json:"permissions"`
	Uid         string    `json:"uid"`
}

func GetAccessKeyLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(accessKeySubj, "access-key-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data AccessKeyLog
			_ = json.Unmarshal(msg.Data, &data)
			if data.EventLog.Type == "query" {
				eventstoredb.Query("keyStream", data)
			} else {
				_ = eventstoredb.AccessKeyLog(data)
			}
		}()
	})
	return qsub
}
