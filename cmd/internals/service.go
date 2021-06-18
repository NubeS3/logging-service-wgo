package internals

import (
	"fmt"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/logic"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
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

	println("Logger Service v1.0.6-hotfix")
	//eventstoredb.Initialize()
	println("Init elasticDB")
	elasticsearchdb.Initialize()
	println("Finish init elasticDB")
	println("Connecting to NATS")
	logic.Initialize()
	println("Connected to NATS")

	println("Init msg handlers")
	stanSubs := []stan.Subscription{}
	natsSubs := []*nats.Subscription{}
	stanSubs = append(stanSubs, logic.GetErrLogQsub())
	natsSubs = append(natsSubs, logic.GetErrLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetFileLogQsub())
	natsSubs = append(natsSubs, logic.GetFileLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetUnauthCountLogQsub())
	natsSubs = append(natsSubs, logic.GetUnauthCountLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetAuthCountLogQsub())
	natsSubs = append(natsSubs, logic.GetAuthCountLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetAccessKeyCountLogQsub())
	natsSubs = append(natsSubs, logic.GetAccessKeyCountLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetSignedCountLogQsub())
	natsSubs = append(natsSubs, logic.GetSignedCountLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetAccessKeyLogQsub())
	natsSubs = append(natsSubs, logic.GetAccessKeyLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetBucketLogQsub())
	natsSubs = append(natsSubs, logic.GetBucketLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetFolderLogQsub())
	natsSubs = append(natsSubs, logic.GetFolderLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetKeyPairLogQsub())
	natsSubs = append(natsSubs, logic.GetKeyPairLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetUserLogQsub())
	natsSubs = append(natsSubs, logic.GetUserLogQSubMsgHandler())
	stanSubs = append(stanSubs, logic.GetBandwidthQsub())
	natsSubs = append(natsSubs, logic.GetBandwidthMsgHandler())
	natsSubs = append(natsSubs, logic.GetReportQSubMsgHandler())
	println("Finish init msg handlers")

	sigs := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		println("Cleaning logger")
		for len(stanSubs) > 0 {
			_ = stanSubs[0].Unsubscribe()
			stanSubs = stanSubs[1:]
		}
		for len(natsSubs) > 0 {
			_ = natsSubs[0].Unsubscribe()
			natsSubs = natsSubs[1:]
		}
		println("Cleaning done")
		cleanupDone <- true
	}()

	fmt.Println("Listening for log events")
	//
	//go logic.TestErr()
	//go logic.TestSendFile()
	//go logic.TestFile()

	<-cleanupDone
	return nil
}
