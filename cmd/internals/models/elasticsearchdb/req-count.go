package elasticsearchdb

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/prometheus/common/log"
	"logging-service-wgo/cmd/internals/models/common"
	"reflect"
	"time"
)

func WriteUnauthReqCountLog(countLog common.UnauthReqLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().
		Index("unauth-req-log").
		BodyJson(countLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func WriteAuthReqCountLog(countLog common.AuthReqLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().
		Index("auth-req-log").
		BodyJson(countLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func WriteAccessKeyReqCountLog(countLog common.AccessKeyReqLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().
		Index("access-key-req-log").
		BodyJson(countLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func WriteSignedReqCountLog(countLog common.SignedReqLog) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()
	_, err := client.Index().
		Index("signed-req-log").
		BodyJson(countLog).
		Do(ctx)
	if err != nil {
		log.Error(err)
	}
}

func ReadUnauthReqCountInDateRange(from, to time.Time, limit, offset int) ([]common.UnauthReqLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("unauth-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var l common.UnauthReqLog
	logRes := []common.UnauthReqLog{}
	for _, item := range res.Each(reflect.TypeOf(l)) {
		if rl, ok := item.(common.UnauthReqLog); ok {
			logRes = append(logRes, rl)
		}
	}

	return logRes, err
}

func CountUnauthReqCountInDateRange(from, to time.Time, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewRangeQuery("at").From(from).To(to)
	res, err := client.Search().Index("unauth-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

func ReadUnauthReqCountLog(limit, offset int) ([]common.UnauthReqLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchAllQuery()
	res, err := client.Search().Index("unauth-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var l common.UnauthReqLog
	logRes := []common.UnauthReqLog{}
	for _, item := range res.Each(reflect.TypeOf(l)) {
		if rl, ok := item.(common.UnauthReqLog); ok {
			logRes = append(logRes, rl)
		}
	}

	return logRes, err
}

func CountUnauthReqCountLog(limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	query := elastic.NewMatchAllQuery()
	res, err := client.Search().Index("unauth-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

func ReadAuthReqCountInDateRange(uid string, from, to time.Time, limit, offset int) ([]common.AuthReqLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	timeQuery := elastic.NewRangeQuery("at").From(from).To(to)
	uidQuery := elastic.NewMatchQuery("user_id", uid)
	query := elastic.NewBoolQuery().Must(timeQuery).Must(uidQuery)
	res, err := client.Search().Index("auth-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var l common.AuthReqLog
	logRes := []common.AuthReqLog{}
	for _, item := range res.Each(reflect.TypeOf(l)) {
		if rl, ok := item.(common.AuthReqLog); ok {
			logRes = append(logRes, rl)
		}
	}

	return logRes, err
}

func CountAuthReqCountInDateRange(uid string, from, to time.Time, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	timeQuery := elastic.NewRangeQuery("at").From(from).To(to)
	uidQuery := elastic.NewMatchQuery("user_id", uid)
	query := elastic.NewBoolQuery().Must(timeQuery).Must(uidQuery)
	res, err := client.Search().Index("auth-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

//

func ReadAccessKeyReqCountInDateRange(accessKey string, from, to time.Time, limit, offset int) ([]common.AccessKeyReqLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	timeQuery := elastic.NewRangeQuery("at").From(from).To(to)
	keyQuery := elastic.NewMatchQuery("key", accessKey)
	query := elastic.NewBoolQuery().Must(timeQuery).Must(keyQuery)
	res, err := client.Search().Index("access-key-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var l common.AccessKeyReqLog
	logRes := []common.AccessKeyReqLog{}
	for _, item := range res.Each(reflect.TypeOf(l)) {
		if rl, ok := item.(common.AccessKeyReqLog); ok {
			logRes = append(logRes, rl)
		}
	}

	return logRes, err
}

func CountAccessKeyReqCountInDateRange(accessKey string, from, to time.Time, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	timeQuery := elastic.NewRangeQuery("at").From(from).To(to)
	keyQuery := elastic.NewMatchQuery("key", accessKey)
	query := elastic.NewBoolQuery().Must(timeQuery).Must(keyQuery)
	res, err := client.Search().Index("access-key-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

func ReadAccessKeyReqCount(accessKey string, limit, offset int) ([]common.AccessKeyReqLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	keyQuery := elastic.NewMatchQuery("key", accessKey)
	res, err := client.Search().Index("access-key-req-log").Query(keyQuery).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var l common.AccessKeyReqLog
	logRes := []common.AccessKeyReqLog{}
	for _, item := range res.Each(reflect.TypeOf(l)) {
		if rl, ok := item.(common.AccessKeyReqLog); ok {
			logRes = append(logRes, rl)
		}
	}

	return logRes, err
}

func CountAccessKeyReqCount(accessKey string, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	keyQuery := elastic.NewMatchQuery("key", accessKey)
	res, err := client.Search().Index("access-key-req-log").Query(keyQuery).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

//

func ReadSignedReqCountInDateRange(accessKey string, from, to time.Time, limit, offset int) ([]common.SignedReqLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	timeQuery := elastic.NewRangeQuery("at").From(from).To(to)
	keyQuery := elastic.NewMatchQuery("public", accessKey)
	query := elastic.NewBoolQuery().Must(timeQuery).Must(keyQuery)
	res, err := client.Search().Index("signed-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var l common.SignedReqLog
	logRes := []common.SignedReqLog{}
	for _, item := range res.Each(reflect.TypeOf(l)) {
		if rl, ok := item.(common.SignedReqLog); ok {
			logRes = append(logRes, rl)
		}
	}

	return logRes, err
}

func CountSignedReqCountInDateRange(accessKey string, from, to time.Time, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	timeQuery := elastic.NewRangeQuery("at").From(from).To(to)
	keyQuery := elastic.NewMatchQuery("public", accessKey)
	query := elastic.NewBoolQuery().Must(timeQuery).Must(keyQuery)
	res, err := client.Search().Index("signed-req-log").Query(query).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}

func ReadSignedReqCount(accessKey string, limit, offset int) ([]common.SignedReqLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	keyQuery := elastic.NewMatchQuery("public", accessKey)
	res, err := client.Search().Index("signed-req-log").Query(keyQuery).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	var l common.SignedReqLog
	logRes := []common.SignedReqLog{}
	for _, item := range res.Each(reflect.TypeOf(l)) {
		if rl, ok := item.(common.SignedReqLog); ok {
			logRes = append(logRes, rl)
		}
	}

	return logRes, err
}

func CountSignedReqCount(accessKey string, limit, offset int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextDuration)
	defer cancel()

	keyQuery := elastic.NewMatchQuery("public", accessKey)
	res, err := client.Search().Index("signed-req-log").Query(keyQuery).From(offset).Size(limit).Do(ctx)
	if err != nil {
		log.Error(err)
	}

	return res.TotalHits(), err
}
