package elasticsearchdb

import (
	"context"
	"errors"
	"github.com/NubeS3/logging-service-wgo/cmd/internals/models/common"
	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
	"time"
)

func WriteBandwidthLog(bandwidthLog common.BandwidthLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().
		Index("bandwidth-log").
		BodyJson(bandwidthLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func GetTotalBandwidthInDateByUid(from, to time.Time, uid string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	aggregate := elastic.NewSumAggregation().Field("size")

	uidQuery := elastic.NewMatchQuery("uid", uid)
	timeQeury := elastic.NewRangeQuery("at").From(from).To(to)
	query := elastic.NewBoolQuery().Must(uidQuery).Must(timeQeury)
	builder := client.Search().Index("bandwidth-log").Query(query).Pretty(true).Aggregation("sumBandwidth", aggregate)

	searchResult, err := builder.Do(ctx)
	if err != nil {
		return 0, err
	}

	agg := searchResult.Aggregations
	sumAggRes, _ := agg.Sum("sumBandwidth")
	if sumAggRes.Value == nil {
		return 0, errors.New("sum agg of bandwidth: value not found")
	}

	return *sumAggRes.Value, nil
}

func GetTotalBandwidthInDateByBucketId(from, to time.Time, bucketId string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	aggregate := elastic.NewSumAggregation().Field("size")

	uidQuery := elastic.NewMatchQuery("bucket_id", bucketId)
	timeQeury := elastic.NewRangeQuery("at").From(from).To(to)
	query := elastic.NewBoolQuery().Must(uidQuery).Must(timeQeury)
	builder := client.Search().Index("bandwidth-log").Query(query).Pretty(true).Aggregation("sumBandwidth", aggregate)

	searchResult, err := builder.Do(ctx)
	if err != nil {
		return 0, err
	}

	agg := searchResult.Aggregations
	sumAggRes, _ := agg.Sum("sumBandwidth")
	if sumAggRes.Value == nil {
		return 0, errors.New("sum agg of bandwidth: value not found")
	}

	return *sumAggRes.Value, nil
}

func GetTotalBandwidthInDateByFrom(from, to time.Time, fromSrc string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	aggregate := elastic.NewSumAggregation().Field("size")
	query := elastic.NewMatchQuery("from", fromSrc)
	builder := client.Search().Index("bandwidth-log").Query(query).Pretty(true).Aggregation("sumBandwidth", aggregate)

	searchResult, err := builder.Do(ctx)
	if err != nil {
		return 0, err
	}

	agg := searchResult.Aggregations
	sumAggRes, _ := agg.Sum("sumBandwidth")
	if sumAggRes.Value == nil {
		return 0, errors.New("sum agg of bandwidth: value not found")
	}

	return *sumAggRes.Value, nil
}
