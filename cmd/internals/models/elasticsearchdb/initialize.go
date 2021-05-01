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
const fileLogMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"type":{
				"type":"keyword"
			},
			"at":{
				"type":"date"
			},
			"file_id": {
				"type":"keyword"
			},
			"file_name": {
				"type": "text"
			},
			"size": {
				"type": "long"
			},
			"bucket_id": {
				"type": "keyword"
			},
			"content_type": {
				"type": "keyword"
			},
			"upload_date": {
				"type": "date"	
			},
			"path": {
				"type": "text"
			},
			"is_hidden": {
				"type": "boolean"
			}
		}
	}
}`
const unauthReqCountMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"type":{
				"type":"keyword"
			},
			"at":{
				"type":"date"
			},
			"method":{
				"type":"keyword"
			},
			"req":{
				"type":"text"
			},
			"source_ip":{
				"type":"keyword"
			}
		}
	}
}`
const authReqCountMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"type":{
				"type":"keyword"
			},
			"at":{
				"type":"date"
			},
			"method":{
				"type":"keyword"
			},
			"req":{
				"type":"text"
			},
			"source_ip":{
				"type":"keyword"
			},
			"user_id":{
				"type":"keyword"
			}
		}
	}
}`
const accessKeyReqCountMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"type":{
				"type":"keyword"
			},
			"at":{
				"type":"date"
			},
			"method":{
				"type":"keyword"
			},
			"req":{
				"type":"text"
			},
			"source_ip":{
				"type":"keyword"
			},
			"key":{
				"type":"keyword"
			}
		}
	}
}`
const signedCountMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"type":{
				"type":"keyword"
			},
			"at":{
				"type":"date"
			},
			"method":{
				"type":"keyword"
			},
			"req":{
				"type":"text"
			},
			"source_ip":{
				"type":"keyword"
			},
			"public":{
				"type":"keyword"
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

	exists, err = client.IndexExists("file-log").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("file-log").BodyString(fileLogMapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	exists, err = client.IndexExists("unauth-req-log").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("unauth-req-log").BodyString(unauthReqCountMapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	exists, err = client.IndexExists("auth-req-log").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("auth-req-log").BodyString(authReqCountMapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	exists, err = client.IndexExists("access-key-req-log").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("access-key-req-log").BodyString(accessKeyReqCountMapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	exists, err = client.IndexExists("signed-req-log").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("signed-req-log").BodyString(signedCountMapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}
