package admin

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAdmins(page int) ([]*model.Admin, int64, error) {
	// 15个
	opt := options.Find().SetSkip(int64(15 * (page - 1))).SetLimit(15)
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	cur, err := mongoDB.Collection(base.Admins).Find(context.TODO(), bson.M{}, opt)
	if err != nil {
		return nil, 0, err
	}

	var results []*model.Admin
	for cur.Next(context.Background()) {
		var admin model.Admin
		cur.Decode(&admin)
		results = append(results, &admin)
	}
	// 查全量
	total, err := mongoDB.Collection(base.Admins).CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}

func AddAdmins(username string, password string, role int) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)

	_, err := mongoDB.Collection(base.Admins).InsertOne(context.Background(), model.NewAdmin(username, password, role))
	return err
}

func DeleteAdmin(id primitive.ObjectID) error {
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := mongoDB.Collection(base.Admins).DeleteOne(context.Background(), bson.M{
		"_id": id,
	})
	return err
}
