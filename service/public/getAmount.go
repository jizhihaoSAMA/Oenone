package public

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func GetAmountOfHouse() int64 {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)
	//total, err := es.Count(base.HouseInfo).Do(context.Background())
	resp, err := es.Search().Index(base.HouseInfo).Query(elastic.NewTermQuery("status", model.Online)).Do(context.Background())
	if err != nil {
		log.Println("[GetAmountOfHouse] 获取房屋数量错误，err: ", err)
		return -1
	}
	return resp.TotalHits()
}
