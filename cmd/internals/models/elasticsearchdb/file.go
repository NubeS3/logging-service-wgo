package elasticsearchdb

import (
	"context"
	"errors"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
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

func CountUploadedObjectByUidInDateRange(uid string, from, to time.Time, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	dateQuery := elastic.NewRangeQuery("at").From(from).To(to)
	uidQuery := elastic.NewMatchQuery("uid", uid)
	typeQuery := elastic.NewMatchQuery("type", "Upload")
	query := elastic.NewBoolQuery().Must(uidQuery).Must(dateQuery).Must(typeQuery)
	res, err := client.Search().Index("file-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

func CountDeletedObjectByUidInDateRange(uid string, from, to time.Time, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	dateQuery := elastic.NewRangeQuery("at").From(from).To(to)
	uidQuery := elastic.NewMatchQuery("uid", uid)
	typeQuery := elastic.NewMatchQuery("type", "Delete")
	query := elastic.NewBoolQuery().Must(uidQuery).Must(dateQuery).Must(typeQuery)
	res, err := client.Search().Index("file-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

func CountAvgObjectByUidInDateRange(uid string, from, to time.Time, limit, offset int) (int64, error) {
	uploaded, err := CountUploadedObjectByUidInDateRange(uid, from, to, limit, offset)
	if err != nil {
		return 0, err
	}
	deleted, err := CountDeletedObjectByUidInDateRange(uid, from, to, limit, offset)
	if err != nil {
		return 0, err
	}

	avg := uploaded - deleted
	if avg < 0 {
		avg = 0
	}

	return avg, nil
}

func AvgUploadedFileSizeByUidInDateRange(uid string, from, to time.Time, limit, offset int) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	dateQuery := elastic.NewRangeQuery("at").From(from).To(to)
	uidQuery := elastic.NewMatchQuery("uid", uid)
	typeQuery := elastic.NewMatchQuery("type", "Upload")
	query := elastic.NewBoolQuery().Must(typeQuery).Must(uidQuery).Must(dateQuery)

	agg := elastic.NewSumAggregation().Field("size")

	res, err := client.Search().Index("file-log").Query(query).From(offset).Size(limit).Aggregation("sum-size-agg", agg).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	aggr := res.Aggregations
	sumAggRes, _ := aggr.Sum("sum-size-agg")
	if sumAggRes.Value == nil {
		return 0, errors.New("sum agg of bandwidth: value not found")
	}

	return *sumAggRes.Value, nil
}

func AvgDeletedFileSizeByUidInDateRange(uid string, from, to time.Time, limit, offset int) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	dateQuery := elastic.NewRangeQuery("at").From(from).To(to)
	uidQuery := elastic.NewMatchQuery("uid", uid)
	typeQuery := elastic.NewMatchQuery("type", "Delete")
	query := elastic.NewBoolQuery().Must(typeQuery).Must(uidQuery).Must(dateQuery)

	agg := elastic.NewSumAggregation().Field("size")

	res, err := client.Search().Index("file-log").Query(query).From(offset).Size(limit).Aggregation("sum-size-agg", agg).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	aggr := res.Aggregations
	sumAggRes, _ := aggr.Sum("sum-size-agg")
	if sumAggRes.Value == nil {
		return 0, errors.New("sum agg of bandwidth: value not found")
	}

	return *sumAggRes.Value, nil
}

func AvgSizeStoredByUidInDateRange(uid string, from, to time.Time, limit, offset int) (float64, error) {
	uploaded, err := AvgUploadedFileSizeByUidInDateRange(uid, from, to, limit, offset)
	if err != nil {
		return 0, err
	}
	deleted, err := AvgDeletedFileSizeByUidInDateRange(uid, from, to, limit, offset)
	if err != nil {
		return 0, err
	}

	avg := uploaded - deleted
	if avg < 0 {
		avg = 0.0
	}

	return avg, nil
}
