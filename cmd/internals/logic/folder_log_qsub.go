package logic

import (
	"encoding/json"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"time"

	"github.com/nats-io/nats.go"
)

func GetFolderLogQsub() *nats.Subscription {
	qsub, _ := js.QueueSubscribe("NUBES3."+fileSubj, "folder-uploaded-success-log-qgroup", func(msg *nats.Msg) {
		go func() {
			var data common.FolderLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteFolderLog(data)
		}()
		msg.Ack()
	})
	return qsub
}

func GetFolderLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(fileSubj+"query", "folder-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.ReadFolderLogInDateRange(from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "All" {
			queryRes, _ := elasticsearchdb.ReadFolderLogByType(data.Data[0], data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		}
	})

	return qsub
}
