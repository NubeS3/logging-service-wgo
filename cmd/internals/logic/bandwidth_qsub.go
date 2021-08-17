package logic

import (
	"encoding/json"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"github.com/nats-io/nats.go"
	"strconv"
	"time"
)

func GetBandwidthQsub() *nats.Subscription {
	qsub, _ := js.QueueSubscribe("NUBES3."+bandwidthSubj, "bandwidth-log-qgroup", func(msg *nats.Msg) {
		go func() {
			var data common.BandwidthLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteBandwidthLog(data)
		}()
		msg.Ack()
	})
	return qsub
}

func GetBandwidthMsgHandler() *nats.Subscription {
	msgHandler, _ := nc.QueueSubscribe(bandwidthSubj+"query", "bandwidth-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		from, _ := time.Parse(time.RFC3339, data.Data[0])
		to, _ := time.Parse(time.RFC3339, data.Data[1])
		if data.Type == "Uid" {
			uid := data.Data[2]
			queryRes, _ := elasticsearchdb.GetTotalBandwidthInDateByUid(from, to, uid)
			_ = msg.Respond([]byte(strconv.FormatFloat(queryRes, 'f', 6, 64)))
		} else if data.Type == "Bid" {
			bid := data.Data[2]
			queryRes, _ := elasticsearchdb.GetTotalBandwidthInDateByBucketId(from, to, bid)
			_ = msg.Respond([]byte(strconv.FormatFloat(queryRes, 'f', 6, 64)))
		} else if data.Type == "From" {
			fromSrc := data.Data[2]
			queryRes, _ := elasticsearchdb.GetTotalBandwidthInDateByFrom(from, to, fromSrc)
			_ = msg.Respond([]byte(strconv.FormatFloat(queryRes, 'f', 6, 64)))
		}
	})

	return msgHandler
}
