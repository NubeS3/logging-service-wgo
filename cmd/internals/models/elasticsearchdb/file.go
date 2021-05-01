package elasticsearchdb

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
	"logging-service-wgo/cmd/internals/models/common"
	"reflect"
	"time"
)

func WriteFileLog(fileLog common.FileLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().
		Index("file-log").
		BodyJson(fileLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadFileLogInDateRange(from, to time.Time, limit, offset int) ([]common.FileLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("upload_date").From(from).To(to)
	res, err := client.Search().Index("file-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var fl common.FileLog
	logRes := []common.FileLog{}
	for _, item := range res.Each(reflect.TypeOf(fl)) {
		if f, ok := item.(common.FileLog); ok {
			logRes = append(logRes, f)
		}
	}

	return logRes, err
}

func ReadFileLog(limit, offset int) ([]common.FileLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchAllQuery()
	res, err := client.Search().Index("file-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var fl common.FileLog
	logRes := []common.FileLog{}
	for _, item := range res.Each(reflect.TypeOf(fl)) {
		if f, ok := item.(common.FileLog); ok {
			logRes = append(logRes, f)
		}
	}

	return logRes, err
}
