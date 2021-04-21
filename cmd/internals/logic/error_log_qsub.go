package logic

import (
	"encoding/json"
	"fmt"
	"log-service-go/cmd/internals/models/eventstoredb"

	"github.com/nats-io/stan.go"
)

type ErrLog struct {
	EventLog Event  `json:"event_log"`
	Error    string `json:"content"`
}

func GetErrLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(errSubj, "error-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data ErrLog
			fmt.Println(msg.String())
			_ = json.Unmarshal(msg.Data, &data)
			if data.EventLog.Type == "query" {
				eventstoredb.Query("errorStream", data.Error)
			}
			_ = eventstoredb.SaveErrorLog(data)
		}()
	})
	return qsub
}
