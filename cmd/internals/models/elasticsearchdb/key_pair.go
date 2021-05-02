package elasticsearchdb

import (
	"context"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
)

func WriteKeyPairLog(keyPairLog common.KeyPairLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().Index("key-pair-log").BodyJson(keyPairLog).Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadKeyPairLogInDateRange(from, to time.Time, limit, offset int) ([]common.KeyPairLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("key-pair-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.KeyPairLog
	logRes := []common.KeyPairLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.KeyPairLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}

func ReadKeyPairLogByType(t string, limit, offset int) ([]common.KeyPairLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchQuery("type", t)
	res, err := client.Search().Index("key-pair-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.KeyPairLog
	logRes := []common.KeyPairLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.KeyPairLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}
