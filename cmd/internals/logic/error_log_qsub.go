package logic

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"logging-service-wgo/cmd/internals/models/common"
	"logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"time"
)

func GetErrLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(errSubj, "error-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.ErrLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteErrLog(data)
		}()
	})
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
		}
	})

	return qsub
}
