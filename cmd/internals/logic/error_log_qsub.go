package logic

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"time"
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
			//_ = eventstoredb.SaveErrorLog(data)
		}()
	})
	return qsub
}
