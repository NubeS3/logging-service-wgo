package logic

import (
	"encoding/json"
	"log-service-go/cmd/internals/models/eventstoredb"
	"time"

	"github.com/nats-io/stan.go"
)

type KeyPairLog struct {
	EventLog Event `json:"event_log"`

	Public       string    `json:"public"`
	Private      string    `json:"private"`
	BucketId     string    `json:"bucket_id"`
	GeneratorUid string    `json:"generator_uid"`
	ExpiredDate  time.Time `json:"expired_date"`
	Permissions  []string  `json:"permissions"`
}

func GetKeyPairLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(keyPairSubj, "key-pair-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data KeyPairLog
			_ = json.Unmarshal(msg.Data, &data)
			if data.EventLog.Type == "query" {
				eventstoredb.Query("keyPairStream", data)
			} else {
				_ = eventstoredb.KeyPairLog(data)
			}
		}()
	})
	return qsub
}
