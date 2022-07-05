package public

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/thoas/go-funk"
	"log"
)

func GetGroupByLocation(location *utils.Location) map[string]*model.NeighborhoodCounterHit {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)

	topRight := elastic.GeoPointFromLatLon(location.TopRightLatFloat, location.TopRightLngFloat)
	bottomLeft := elastic.GeoPointFromLatLon(location.BottomLeftLatFloat, location.BottomLeftLngFloat)

	// 第一次查询，查询经纬度内所有的小区信息，
	queryRes, err := es.Search().Index(base.NeighborhoodInfo).
		Query(elastic.NewBoolQuery().Filter(
			elastic.NewGeoBoundingBoxQuery("location").
				BottomLeftFromGeoPoint(bottomLeft).
				TopRightFromGeoPoint(topRight)),
		).
		Do(context.Background())

	if err != nil {
		log.Println("[GetGroupByLocation] 查询错误, err: ", err, queryRes)
		return nil
	}

	counter := funk.Map(queryRes.Hits.Hits, func(hit *elastic.SearchHit) (string, *model.NeighborhoodCounterHit) {
		var tempNeighborhood model.Neighborhood
		json.Unmarshal(hit.Source, &tempNeighborhood)
		return tempNeighborhood.Name, &model.NeighborhoodCounterHit{
			Neighborhood: tempNeighborhood,
			Count:        0,
		}
	}).(map[string]*model.NeighborhoodCounterHit)

	// 第二次查询，查询各个小区的房屋数量的聚合情况
	queryRes, err = es.Search().Index(base.HouseInfo).
		Query(
			elastic.NewBoolQuery().Must(
				elastic.NewTermsQueryFromStrings("house_neighborhood.keyword", funk.Keys(counter).([]string)...),
				elastic.NewTermQuery("status", model.Online))).
		Aggregation("counter", elastic.NewTermsAggregation().Field("house_neighborhood.keyword")).
		Do(context.Background())

	if err != nil {
		log.Println("[GetGroupByLocation] 二次查询错误, err: ", err, queryRes)
		return nil
	}

	agg, found := queryRes.Aggregations.Terms("counter")
	if found {
		for _, bucket := range agg.Buckets {
			counter[bucket.Key.(string)].Count = bucket.DocCount
		}
	}
	return counter

}

// 正向查找，获得不为0的小区，目前代码为所有小区，即使为0也显示进去
//neighborhoods := funk.Map(queryRes.Hits.Hits, func(hit *elastic.SearchHit) (name string, neighborhood model.Neighborhood) {
//		json.Unmarshal(hit.Source, &neighborhood)
//		return neighborhood.Name, neighborhood
//	}).(map[string]model.Neighborhood)
//
//	// 第二次查询，查询各个小区的房屋数量的聚合情况
//	queryRes, err = es.Search().Index(base.HouseInfo).
//		Query(elastic.NewTermsQueryFromStrings("house_neighborhood.keyword", funk.Keys(neighborhoods).([]string)...)).
//		Aggregation("counter", elastic.NewTermsAggregation().Field("house_neighborhood.keyword")).
//		Do(context.Background())
//
//	if err != nil {
//		log.Println("[GetGroupByLocation] 二次查询错误, err: ", err, queryRes)
//		return nil
//	}
//
//	agg, found := queryRes.Aggregations.Terms("counter")
//	counter := map[string]*searchHit{}
//	if found {
//		for _, bucket := range agg.Buckets {
//			counter[bucket.Key.(string)] = &searchHit{
//				Neighborhood: neighborhoods[bucket.Key.(string)],
//				Count:        bucket.DocCount,
//			}
//		}
//	}
//	return counter
