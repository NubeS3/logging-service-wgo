package internals

import (
	"fmt"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/logic"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/elasticsearchdb"
	"github.com/nats-io/nats.go"
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
	var natsSubs []*nats.Subscription
	natsSubs = append(natsSubs, logic.GetErrLogQsub())
	natsSubs = append(natsSubs, logic.GetErrLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetFileLogQsub())
	natsSubs = append(natsSubs, logic.GetFileLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetUnauthCountLogQsub())
	natsSubs = append(natsSubs, logic.GetUnauthCountLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetAuthCountLogQsub())
	natsSubs = append(natsSubs, logic.GetAuthCountLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetAccessKeyCountLogQsub())
	natsSubs = append(natsSubs, logic.GetAccessKeyCountLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetSignedCountLogQsub())
	natsSubs = append(natsSubs, logic.GetSignedCountLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetAccessKeyLogQsub())
	natsSubs = append(natsSubs, logic.GetAccessKeyLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetBucketLogQsub())
	natsSubs = append(natsSubs, logic.GetBucketLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetFolderLogQsub())
	natsSubs = append(natsSubs, logic.GetFolderLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetKeyPairLogQsub())
	natsSubs = append(natsSubs, logic.GetKeyPairLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetUserLogQsub())
	natsSubs = append(natsSubs, logic.GetUserLogQSubMsgHandler())
	natsSubs = append(natsSubs, logic.GetBandwidthQsub())
	natsSubs = append(natsSubs, logic.GetBandwidthMsgHandler())
	natsSubs = append(natsSubs, logic.GetReportQSubMsgHandler())
	println("Finish init msg handlers")

	sigs := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		println("Cleaning logger")
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
