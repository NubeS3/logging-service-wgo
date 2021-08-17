package logic

import (
	"encoding/json"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"github.com/nats-io/nats.go"
	"time"
)

func GetKeyPairLogQsub() *nats.Subscription {
	qsub, _ := js.QueueSubscribe("NUBES3."+keyPairSubj, "key-pair-log-qgroup", func(msg *nats.Msg) {
		go func() {
			var data common.KeyPairLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteKeyPairLog(data)
		}()
		msg.Ack()
	}, nats.Durable("NUBES3"))
	return qsub
}

func GetKeyPairLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(errSubj+"query", "key-pair-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.ReadKeyPairLogInDateRange(from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Type" {
			queryRes, _ := elasticsearchdb.ReadKeyPairLogByType(data.Data[0], data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		}
	})

	return qsub
}
