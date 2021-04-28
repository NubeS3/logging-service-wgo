package elasticsearchdb

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

const ContextDuration = time.Second * 30

var (
	client *elastic.Client
)

const errLogMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"type":{
				"type":"keyword"
			},
			"content":{
				"type":"text"
			},
			"at":{
				"type":"date"
			}
		}
	}
}`

func Initialize() {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	dbUrl := viper.GetString("ELASTIC_URL")

	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(dbUrl),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ERR ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "ELASTIC INFO ", log.LstdFlags)),
	)

	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(dbUrl).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	exists, err := client.IndexExists("error-log").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("error-log").BodyString(errLogMapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}
