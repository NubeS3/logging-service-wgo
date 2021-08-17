package logic

import (
	"encoding/json"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"github.com/nats-io/nats.go"
	"time"
)

func GetErrLogQsub() *nats.Subscription {
	qsub, _ := js.QueueSubscribe("NUBES3."+errSubj, "error-log-qgroup", func(msg *nats.Msg) {
		go func() {
			var data common.ErrLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteErrLog(data)
		}()
		msg.Ack()
	}, nats.Durable("NUBES3"))
	return qsub
}

func GetErrLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(errSubj+"query", "error-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.ReadErrLogInDateRange(from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Type" {
			queryRes, _ := elasticsearchdb.ReadErrLogByType(data.Data[0], data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "All" {
			queryRes, _ := elasticsearchdb.ReadErrLog(data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		}
	})

	return qsub
}
