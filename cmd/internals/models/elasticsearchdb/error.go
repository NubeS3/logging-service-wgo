package elasticsearchdb

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
	"log-service-go/cmd/internals/models/common"
	"reflect"
	"time"
)

func WriteErrLog(errLog common.ErrLog) {
	ctx, cancel := context.WithTimeout(nil, ContextDuration)
	defer cancel()
	// Index a tweet (using JSON serialization)
	_, err := client.Index().
		Index("err-log").
		BodyJson(errLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadErrLogAfterDate(from time.Time, to time.Time) ([]common.ErrLog, error) {
	ctx, cancel := context.WithTimeout(nil, ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("err-log").Query(query).Pretty(true).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var el common.ErrLog
	logRes := []common.ErrLog{}
	for _, item := range res.Each(reflect.TypeOf(el)) {
		if e, ok := item.(common.ErrLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}

func ReadErrLogByType(t string) ([]common.ErrLog, error) {
	ctx, cancel := context.WithTimeout(nil, ContextDuration)
	defer cancel()

	query := elastic.NewTermQuery("type", t)
	res, err := client.Search().Index("err-log").Query(query).Pretty(true).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var el common.ErrLog
	logRes := []common.ErrLog{}
	for _, item := range res.Each(reflect.TypeOf(el)) {
		if e, ok := item.(common.ErrLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}
