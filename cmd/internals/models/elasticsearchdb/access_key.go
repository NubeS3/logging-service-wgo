package elasticsearchdb

import (
	"context"
	"log-service-go/cmd/internals/models/common"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
)

func WriteAccessKeyLog(accessKeyLog common.AccessKeyLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().Index("access-key-log").BodyJson(accessKeyLog).Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadAccessKeyLogInDateRange(from, to time.Time, limit, offset int) ([]common.AccessKeyLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("access-key-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.AccessKeyLog
	logRes := []common.AccessKeyLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.AccessKeyLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}

func ReadAccessKeyLogByType(t string, limit, offset int) ([]common.AccessKeyLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchQuery("type", t)
	res, err := client.Search().Index("access-key-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.AccessKeyLog
	logRes := []common.AccessKeyLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.AccessKeyLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}
