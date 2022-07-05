package admin

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func AddReview(houseID string, msg string, pass bool) error {
	statusCode := map[bool]int{false: model.RejectReview, true: model.Online}[pass]
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Reviews).UpdateByID(context.Background(), houseID, bson.M{
		"$push": bson.M{
			"histories": model.NewReviewHistory(pass, msg),
		},
	})
	if err != nil {
		return err
	}

	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)
	queryResult, err := es.Update().Index(base.HouseInfo).
		Id(houseID).
		Doc(map[string]interface{}{
			"status": statusCode,
		}).FetchSource(true).Do(context.Background())
	if err != nil {
		log.Println(base.GetErrorDetailFromES(err))
		return err
	}

	var house model.House
	_ = json.Unmarshal(queryResult.GetResult.Source, &house)

	var notice *model.Notice
	if pass {
		notice = model.NewHouseApprovedNotice(houseID)
	} else {
		notice = model.NewHouseRejectedNotice(houseID)
	}

	log.Println(house)
	_, err = mongoDB.Collection(base.Users).UpdateByID(context.Background(), house.OwnerID, bson.M{
		"$push": bson.M{
			"notices": bson.M{
				"$each":     bson.A{notice},
				"$position": 0,
			},
		},
		"$inc": bson.M{
			"unread_amount": 1,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
