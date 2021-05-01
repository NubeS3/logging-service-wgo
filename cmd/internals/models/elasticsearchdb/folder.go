package elasticsearchdb

import (
	"context"
	"log-service-go/cmd/internals/models/common"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
)

func WriteFolderLog(folderLog common.FolderLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().Index("folder-log").BodyJson(folderLog).Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadFolderLogInDateRange(from, to time.Time, limit, offset int) ([]common.FolderLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("folder-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.FolderLog
	logRes := []common.FolderLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.FolderLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}

func ReadFolderLogByType(t string, limit, offset int) ([]common.FolderLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchQuery("type", t)
	res, err := client.Search().Index("folder-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var akl common.FolderLog
	logRes := []common.FolderLog{}
	for _, item := range res.Each(reflect.TypeOf(akl)) {
		if e, ok := item.(common.FolderLog); ok {
			logRes = append(logRes, e)
		}
	}

	return logRes, err
}
