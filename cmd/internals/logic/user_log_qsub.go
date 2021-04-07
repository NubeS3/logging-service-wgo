package logic

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log-service-go/cmd/internals/models/eventstoredb"
	"time"
)

type UserLog struct {
	EventLog Event `json:"event_log"`

	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username"`
	Pass      string    `json:"password"`
	Email     string    `json:"email"`
	Dob       time.Time `json:"dob"`
	Company   string    `json:"company"`
	Gender    bool      `json:"gender"`
	IsActive  bool      `json:"is_active"`
	IsBanned  bool      `json:"is_banned"`
}

func GetUserLogQsub() stan.Subscription {
	qsub, _ := sc.QueueSubscribe(userSubj, "user-log-qgroup", func(msg *stan.Msg) {
		go func() {
			var data UserLog
			_ = json.Unmarshal(msg.Data, &data)
			_ = eventstoredb.UserLog(data)
		}()
	})
	return qsub
}
