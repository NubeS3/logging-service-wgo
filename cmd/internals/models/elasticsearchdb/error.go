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
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().
		Index("err-log").
		BodyJson(errLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadErrLogInDateRange(from, to time.Time, limit, offset int) ([]common.ErrLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("err-log").Query(query).From(offset).Size(limit).Do(ctx)
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

func ReadErrLogByType(t string, limit, offset int) ([]common.ErrLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchQuery("type", t)
	res, err := client.Search().Index("err-log").Query(query).From(offset).Size(limit).Do(ctx)
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

func ReadErrLog(limit, offset int) ([]common.ErrLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchAllQuery()
	res, err := client.Search().Index("err-log").Query(query).From(offset).Size(limit).Do(ctx)
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
