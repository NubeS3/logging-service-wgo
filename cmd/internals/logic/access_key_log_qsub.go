package logic

import (
	"encoding/json"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func GetAccessKeyLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(errSubj, "access-key-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.AccessKeyLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteAccessKeyLog(data)
		}()
	})
	return qsub
}

func GetAccessKeyLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(errSubj+"query", "access-key-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.ReadAccessKeyLogInDateRange(from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Type" {
			queryRes, _ := elasticsearchdb.ReadAccessKeyLogByType(data.Data[0], data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		}
	})

	return qsub
}
