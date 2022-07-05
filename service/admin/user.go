package admin

import (
	"Oenone/common/base"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUser(userId primitive.ObjectID) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Users).DeleteOne(context.Background(), bson.M{
		"_id": userId,
	})
	return err
}
