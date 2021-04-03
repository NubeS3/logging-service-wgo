package eventstoredb

import (
	goes "github.com/jetbasrawi/go.geteventstore"
	"github.com/spf13/viper"
	"log"
)

var client *goes.Client

func Initialize() {
	var err error
	client, err = goes.NewClient(nil, "http://"+viper.GetString("Esdb_url"))
	if err != nil {
		log.Fatal(err)
	}
}
