package logic

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
)

var (
	sc stan.Conn
	nc *nats.Conn

	// mailSubj              string
	reqSubj  string
	errSubj  string
	fileSubj string
	//uploadFileSuccessSubj string
	userSubj      string
	bucketSubj    string
	folderSubj    string
	accessKeySubj string
	keyPairSubj   string
	bandwidthSubj string
)

func Initialize() {
	var err error
	sc, err = stan.Connect(viper.GetString("Cluster_id"), viper.GetString("Client_id"), stan.NatsURL(viper.GetString("Nats_url")))
	if err != nil {
		panic(fmt.Errorf("fatal error connecting nats stream: %s", err))
	}

	nc, err = nats.Connect("nats://" + viper.GetString("Nats_url"))
	if err != nil {
		panic(err)
	}

	// mailSubj = env["mailSubj"]
	reqSubj = viper.GetString("reqSubj")
	errSubj = viper.GetString("errSubj")
	fileSubj = viper.GetString("fileSubj")
	//uploadFileSuccessSubj = viper.GetString("uploadFileSuccessSubj")
	userSubj = viper.GetString("userSubj")
	bucketSubj = viper.GetString("bucketSubj")
	folderSubj = viper.GetString("folderSubj")
	accessKeySubj = viper.GetString("accessKeySubj")
	keyPairSubj = viper.GetString("keyPairSubj")
	bandwidthSubj = viper.GetString("bandwidthSubj")
}
