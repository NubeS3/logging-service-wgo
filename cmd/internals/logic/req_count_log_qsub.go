package logic

import (
	"encoding/json"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"strconv"
	"time"
)

func GetUnauthCountLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(reqSubj+"unauth", "unauth-count-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.UnauthReqLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteUnauthReqCountLog(data)
		}()
	})
	return qsub
}

func GetAuthCountLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(reqSubj+"auth", "auth-count-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.AuthReqLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteAuthReqCountLog(data)
		}()
	})
	return qsub
}

func GetAccessKeyCountLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(reqSubj+"access-key", "access-key-count-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.AccessKeyReqLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteAccessKeyReqCountLog(data)
		}()
	})
	return qsub
}

func GetSignedCountLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(reqSubj+"signed", "signed-count-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.SignedReqLog
			_ = json.Unmarshal(msg.Data, &data)
			elasticsearchdb.WriteSignedReqCountLog(data)
		}()
	})
	return qsub
}

func GetUnauthCountLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(reqSubj+"unauth"+"query", "unauth-req-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.ReadUnauthReqCountInDateRange(from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "All" {
			queryRes, _ := elasticsearchdb.ReadUnauthReqCountLog(data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Date-count" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			queryRes, _ := elasticsearchdb.CountUnauthReqCountInDateRange(from, to, data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		} else if data.Type == "All-count" {
			queryRes, _ := elasticsearchdb.CountUnauthReqCountLog(data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		}
	})

	return qsub
}

func GetAuthCountLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(reqSubj+"auth"+"query", "auth-req-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			uid := data.Data[2]
			queryRes, _ := elasticsearchdb.ReadAuthReqCountInDateRange(uid, from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Date-count" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			uid := data.Data[2]
			queryRes, _ := elasticsearchdb.CountAuthReqCountInDateRange(uid, from, to, data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		}
	})

	return qsub
}

func GetAccessKeyCountLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(reqSubj+"access-key"+"query", "access-key-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			key := data.Data[2]
			queryRes, _ := elasticsearchdb.ReadAccessKeyReqCountInDateRange(key, from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "All" {
			queryRes, _ := elasticsearchdb.ReadAccessKeyReqCount(data.Data[0], data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Date-count" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			key := data.Data[2]
			queryRes, _ := elasticsearchdb.CountAccessKeyReqCountInDateRange(key, from, to, data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		} else if data.Type == "All-count" {
			queryRes, _ := elasticsearchdb.CountAccessKeyReqCount(data.Data[0], data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		}
	})

	return qsub
}

func GetSignedCountLogQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(reqSubj+"signed"+"query", "signed-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			key := data.Data[2]
			queryRes, _ := elasticsearchdb.ReadSignedReqCountInDateRange(key, from, to, data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "All" {
			queryRes, _ := elasticsearchdb.ReadSignedReqCount(data.Data[0], data.Limit, data.Offset)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "Date-count" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			key := data.Data[2]
			queryRes, _ := elasticsearchdb.CountSignedReqCountInDateRange(key, from, to, data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		} else if data.Type == "All-count" {
			queryRes, _ := elasticsearchdb.CountSignedReqCount(data.Data[0], data.Limit, data.Offset)
			_ = msg.Respond([]byte(strconv.FormatInt(queryRes, 10)))
		}
	})

	return qsub
}

func GetReportQSubMsgHandler() *nats.Subscription {
	qsub, _ := nc.QueueSubscribe(reqSubj+"report"+"query", "report-log-query", func(msg *nats.Msg) {
		var data common.Req
		_ = json.Unmarshal(msg.Data, &data)
		if data.Type == "Date" {
			from, _ := time.Parse(time.RFC3339, data.Data[0])
			to, _ := time.Parse(time.RFC3339, data.Data[1])
			uid := data.Data[2]
			queryRes, _ := elasticsearchdb.CountReqByClassUsingUidInDateRAnge(uid, from, to)
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		} else if data.Type == "All" {
			queryRes, _ := elasticsearchdb.CountReqByClassUsingUidInDateRAnge(data.Data[0], time.Time{}, time.Now())
			jsonData, _ := json.Marshal(queryRes)
			_ = msg.Respond(jsonData)
		}
	})

	return qsub
}
