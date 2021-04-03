package internals

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log-service-go/cmd/internals/logic"
	"log-service-go/cmd/internals/models/eventstoredb"
	"os"
	"os/signal"
	"syscall"
)

func Run() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	eventstoredb.Initialize()

	subs := []stan.Subscription{}
	subs = append(subs, logic.GetErrLogQsub())

	sigs := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		for len(subs) > 0 {
			_ = subs[0].Unsubscribe()
			subs = subs[1:]
		}
		cleanupDone <- true
	}()

	fmt.Println("Listening for log events")
	<-cleanupDone
	return nil
}
