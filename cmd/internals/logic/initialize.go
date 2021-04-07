package logic

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	sc stan.Conn

	mailSubj              string
	errSubj               string
	uploadFileSubj        string
	downloadFileSubj      string
	stagingFileSubj       string
	uploadFileSuccessSubj string
	userSubj              string
	bucketSubj            string
	folderSubj            string
	accessKeySubj         string
	keyPairSubj           string
)

type Event struct {
	Type string    `json:"type"`
	Date time.Time `json:"event_time"`
}

func Initialize() {
	var err error
	sc, err = stan.Connect(viper.GetString("Cluster_id"), viper.GetString("Client_id"), stan.NatsURL("nats://"+viper.GetString("Nats_url")))
	if err != nil {
		panic(fmt.Errorf("fatal error connecting nats stream: %s", err))
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	var env map[string]string
	env, err = godotenv.Read()
	if err != nil {
		log.Fatal(err)
		return
	}

	mailSubj = env["mailSubj"]
	errSubj = env["errSubj"]
	uploadFileSubj = env["uploadFileSubj"]
	downloadFileSubj = env["downloadFileSubj"]
	stagingFileSubj = env["stagingFileSubj"]
	uploadFileSuccessSubj = env["uploadFileSuccessSubj"]
	userSubj = env["userSubj"]
	bucketSubj = env["bucketSubj"]
	folderSubj = env["folderSubj"]
	accessKeySubj = env["accessKeySubj"]
	keyPairSubj = env["keyPairSubj"]
}
