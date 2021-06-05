package logic

import (
	"encoding/json"
	"fmt"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func GetFileLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(fileSubj, "file-uploaded-success-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.FileLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteFileLog(data)
		}()
	})
	return qsub
}

func GetFileLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(fileSubj+"query", "file-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.ReadFileLogInDateRange(from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "All" {
			queryRes, _ := elasticsearchdb.ReadFileLog(data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "AvgSize" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			uid := data.Data[2]
			queryRes, _ := elasticsearchdb.AvgSizeStoredByUidInDateRange(uid, from, to, data.Limit, data.Offset)
			_ = msg.Respond([]byte(fmt.Sprintf("%f", queryRes)))
		} else if data.Type == "AvgCount" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			uid := data.Data[2]
			queryRes, _ := elasticsearchdb.CountAvgObjectByUidInDateRange(uid, from, to, data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		}
	})

	return qsub
}
