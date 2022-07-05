package public

import (
	"Oenone/common/base"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func HouseList(userID primitive.ObjectID, area string, bc string, neighborhood string, rentType string, priceStart int, priceEnd int, page int,
	supportShortTermRent bool, hasLift bool, hasSingleToilet bool, hasSingleBalcony bool, status int, fields []string) (interface{}, int64, error) {
	// 每页12个
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)

	query := elastic.NewBoolQuery()

	if !userID.IsZero() {
		query.Must(elastic.NewTermQuery("owner_id", userID))
	}

	if area != "" {
		query.Must(elastic.NewTermQuery("house_loc_area.keyword", area))
	}

	if bc != "" {
		query.Must(elastic.NewTermQuery("house_loc_bc.keyword", bc))
	}

	if neighborhood != "" {
		query.Must(elastic.NewTermQuery("house_neighborhood.keyword", neighborhood))
	}

	if priceStart != 0 || priceEnd != 0 {
		rangeQuery := elastic.NewRangeQuery("house_price")
		if priceStart != 0 {
			rangeQuery.Gte(priceStart)
		}
		if priceEnd != 0 {
			rangeQuery.Lte(priceEnd)
		}
		query.Must(rangeQuery)
	}

	if rentType != "" {
		if rentType == "full" {
			query.Must(elastic.NewTermQuery("is_full_rent", true))
		}
		if rentType == "part" {
			query.Must(elastic.NewTermQuery("is_full_rent", false))
		}
	}

	if supportShortTermRent {
		query.Must(elastic.NewTermQuery("support_short_term_rent", true))
	}
	if hasLift {
		query.Must(elastic.NewTermQuery("has_lift", true))
	}
	if hasSingleToilet {
		query.Must(elastic.NewTermQuery("has_single_toilet", true))
	}
	if hasSingleBalcony {
		query.Must(elastic.NewTermQuery("has_single_balcony", true))
	}

	if status != 0 {
		query.Must(elastic.NewTermQuery("status", status))
	}

	builder := es.Search().
		Index(base.HouseInfo).
		Query(query)

	if len(fields) != 0 {
		builder.FetchSourceContext(elastic.NewFetchSourceContext(true).Include(fields...))
	}

	if page != 0 { //不等于0就翻页，否则全量
		builder.Size(12).
			From(12 * (page - 1))
	}

	queryResult, err := builder.Do(context.Background())
	if err != nil {
		log.Printf("[HouseList] 查询ES出错, err: %s\n", base.GetErrorDetailFromES(err))
		return nil, 0, err
	}

	res := funk.Map(queryResult.Hits.Hits, func(hit *elastic.SearchHit) map[string]interface{} {
		result := make(map[string]interface{})
		json.Unmarshal(hit.Source, &result)
		result["id"] = hit.Id
		return result
	}).([]map[string]interface{})
	log.Println(len(res))
	return res, queryResult.TotalHits(), nil
}
