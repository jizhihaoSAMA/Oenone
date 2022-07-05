package user

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func PostHouse(house model.House) (string, error) {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)
	houseJson := base.StructToMap(house)
	delete(houseJson, "id")
	res, err := es.Index().Index(base.HouseInfo).Routing(house.HouseLocArea).BodyJson(houseJson).Do(context.Background())
	if err != nil {
		return "", err
	}
	return res.Id, nil
}

func OfflineHouse(houseID string) error {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)
	_, err := es.Update().Index(base.HouseInfo).
		Id(houseID).
		Doc(map[string]interface{}{
			"status": model.Offline,
		}).Do(context.Background())
	return err
}

func UpdateHouse(house model.House, houseID string) error {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)

	m := base.StructToMap(house)
	for _, f := range []string{"created_time", "owner_id", "id"} {
		delete(m, f)
	}
	log.Printf("%+v", house)
	log.Println(m)
	_, err := es.Update().Index(base.HouseInfo).
		Id(houseID).
		Doc(m).Do(context.Background())
	return err
}
