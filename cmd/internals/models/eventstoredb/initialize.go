package eventstoredb

import (
	"log"

	goes "github.com/jetbasrawi/go.geteventstore"
)

var client *goes.Client

func Initialize() {
	var err error
	// client, err = goes.NewClient(nil, "http://"+viper.GetString("Esdb_url"))
	client, err = goes.NewClient(nil, "esdb://admin:changeit@127.0.0.1:2113?tls=true")
	if err != nil {
		log.Fatal(err)
	}
}
