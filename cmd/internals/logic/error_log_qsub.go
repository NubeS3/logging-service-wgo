package logic

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log-service-go/cmd/internals/models/eventstoredb"
	"time"
)

type ErrLog struct {
	Content string    `json:"content"`
	Type    string    `json:"type"`
	At      time.Time `json:"at"`
}

func GetErrLogQsub() stan.Subscription {
	sc, err := stan.Connect(viper.GetString("Cluster_id"), viper.GetString("Client_id"), stan.NatsURL("nats://"+viper.GetString("Nats_url")))
	if err != nil {
		panic(fmt.Errorf("Fatal error connecting nats stream: %s \n", err))
	}

	qsub, _ := sc.QueueSubscribe(viper.GetString("ErrLog_subject"), "error-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data ErrLog
			fmt.Println(msg.String())
			_ = json.Unmarshal(msg.Data, &data)
			_ = eventstoredb.SaveErrorLog(data)
		}()
	})
	return qsub
}
