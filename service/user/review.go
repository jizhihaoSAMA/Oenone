package user

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetReviewHistories(houseID string) (*model.Review, error) {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	var review model.Review
	err := mongoDB.Collection(base.Reviews).FindOne(context.Background(), bson.M{
		"_id": houseID,
	}).Decode(&review)
	if err != nil {
		return nil, err
	}
	return &review, nil
}
