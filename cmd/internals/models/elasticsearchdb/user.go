package elasticsearchdb

import (
	"context"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
)

func WriteUserLog(userLog common.UserLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().Index("user-log").BodyJson(userLog).Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadUserLogInDateRange(from, to time.Time, limit, offset int) ([]common.UserLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("user-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.UserLog
	logRes := []common.UserLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.UserLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}

func ReadUserLogByType(t string, limit, offset int) ([]common.UserLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchQuery("type", t)
	res, err := client.Search().Index("user-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.UserLog
	logRes := []common.UserLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.UserLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}
