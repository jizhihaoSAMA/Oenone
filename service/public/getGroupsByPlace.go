package public

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func GetGroupByPlaces(field string, places []string) map[string]interface{} {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)
	aggs := elastic.NewTermsAggregation().Field(field + ".keyword")

	var (
		res *elastic.SearchResult
		err error
	)

	filter := elastic.NewTermQuery("status", model.Online)

	if len(places) == 0 {
		res, err = es.Search().
			Index(base.HouseInfo).
			Aggregation("counter", aggs).
			Query(filter).
			Do(context.Background())
	} else {
		res, err = es.Search().
			Index(base.HouseInfo).
			Query(elastic.NewBoolQuery().Must(elastic.NewTermsQueryFromStrings(field+".keyword", places...), filter)).
			Aggregation("counter", aggs).
			Do(context.Background())
	}

	if err != nil {
		log.Println("[GetGroupByPlaces] 获取数量错误，err: ", err)
		return nil
	}
	counter := make(map[string]interface{})
	agg, found := res.Aggregations.Terms("counter")
	if found {
		for _, bucket := range agg.Buckets {
			counter[bucket.Key.(string)] = bucket.DocCount
		}

	}
	return counter
}
