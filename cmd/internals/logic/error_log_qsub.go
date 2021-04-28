package logic

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log-service-go/cmd/internals/models/common"
)

func GetErrLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(errSubj, "error-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data common.ErrLog
			fmt.Println(msg.String())
			_ = json.Unmarshal(msg.Data, &data)
			//_ = eventstoredb.SaveErrorLog(data)
		}()
	})
	return qsub
}
