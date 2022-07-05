package user

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeAppointment(userID primitive.ObjectID, houseID string, time primitive.DateTime) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Users).UpdateByID(context.Background(), userID, bson.M{
		"$push": bson.M{
			"appointments": model.NewAppointment(houseID, time),
		},
	})
	return err
}

func DelAppointment(userID primitive.ObjectID, houseID string) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Users).UpdateByID(context.Background(), userID, bson.M{
		"$pull": bson.M{
			"appointments": bson.M{
				"_id": houseID,
			},
		},
	})
	return err
}
