package public

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
)

func HouseDetail(id string, status int) (*model.House, error) {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)

	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("_id", id))
	if status != 0 {
		query.Must(elastic.NewTermQuery("status", status))
	}

	queryResult, err := es.Search().Index(base.HouseInfo).Query(query).Do(context.Background())
	if err != nil {
		log.Println("[HouseDetail] 查询详情出错：" + err.Error())
		return nil, err
	}

	if queryResult.TotalHits() == 0 {
		return nil, nil
	}
	var result model.House
	json.Unmarshal(queryResult.Hits.Hits[0].Source, &result)
	result.ID = queryResult.Hits.Hits[0].Id
	return &result, nil
}
