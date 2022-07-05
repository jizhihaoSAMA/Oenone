package user

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Star(userID primitive.ObjectID, houseID string) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Users).UpdateByID(context.Background(), userID, bson.M{
		"$push": bson.M{
			"stars": model.NewStar(houseID),
		},
	})
	return err
}

func Unstar(userID primitive.ObjectID, houseID string) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Users).UpdateByID(context.Background(), userID, bson.M{
		"$pull": bson.M{
			"stars": bson.M{
				"_id": houseID,
			},
		},
	})
	return err
}
