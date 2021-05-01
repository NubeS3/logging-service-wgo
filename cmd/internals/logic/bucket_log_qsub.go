package logic

import (
	"encoding/json"
	"log-service-go/cmd/internals/models/common"
	"log-service-go/cmd/internals/models/elasticsearchdb"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func GetBucketLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(errSubj, "bucket-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.BucketLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteBucketLog(data)
		}()
	})
	return qsub
}

func GetBucketLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(errSubj+"query", "bucket-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.ReadBucketLogInDateRange(from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Type" {
			queryRes, _ := elasticsearchdb.ReadBucketLogByType(data.Data[0], data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		}
	})

	return qsub
}
