package user

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUnreadAmount(userID primitive.ObjectID) (int, error) {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	var resultUser model.User
	err := mongoDB.Collection(base.Users).FindOne(context.Background(), bson.M{
		"_id": userID,
	}, &options.FindOneOptions{Projection: bson.M{"unread_amount": 1}}).Decode(&resultUser)

	if err != nil {
		return 0, err
	}
	return resultUser.UnreadAmount, nil
}

func NoticeAllRead(userID primitive.ObjectID) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Users).UpdateByID(context.Background(), userID, bson.M{
		"$set": bson.M{
			"unread_amount": 0,
		},
	})
	return err
}
