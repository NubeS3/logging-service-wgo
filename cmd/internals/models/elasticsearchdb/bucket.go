package elasticsearchdb

import (
	"context"
	"logging-service-wgo/cmd/internals/models/common"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
)

func WriteBucketLog(bucketLog common.BucketLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().Index("bucket-log").BodyJson(bucketLog).Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadBucketLogInDateRange(from, to time.Time, limit, offset int) ([]common.BucketLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("bucket-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.BucketLog
	logRes := []common.BucketLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.BucketLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}

func ReadBucketLogByType(t string, limit, offset int) ([]common.BucketLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchQuery("type", t)
	res, err := client.Search().Index("bucket-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.BucketLog
	logRes := []common.BucketLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.BucketLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}
