package admin

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAppointments 分页(管理员管理)
func GetAppointments(page int) ([]*model.User, int64, error) {
	// 每页15个用户
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)

	opt := options.Find().SetProjection(bson.M{"_id": 1, "username": 1, "appointments": 1, "telephone": 1})
	if page != 0 {
		opt.SetSkip(int64(15 * (page - 1))).SetLimit(15)
	}

	cur, err := mongoDB.Collection(base.Users).Find(context.Background(), bson.M{}, opt)
	if err != nil {
		return nil, 0, err
	}

	var results []*model.User
	for cur.Next(context.Background()) {
		var user model.User
		cur.Decode(&user)
		results = append(results, &user)
	}
	// 查全量
	total, err := mongoDB.Collection(base.Users).CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}
