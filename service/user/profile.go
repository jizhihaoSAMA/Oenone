package user

import (
	"Oenone/common/base"
	"Oenone/common/crud"
	"Oenone/model"
	"Oenone/service/public"
	"context"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOwnHouses(userID primitive.ObjectID, page int) (interface{}, int64, error) {
	return public.HouseList(userID, "", "", "", "", 0, 0, page, false, false, false, false, 0, crud.FieldMapper["profile"])
}

// GetNotices 不分页
func GetNotices(userID primitive.ObjectID) ([]*model.Notice, error) {
	opt := options.FindOne().SetProjection(bson.M{"notices": 1})
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)

	var resultUser model.User
	err := mongoDB.Collection(base.Users).FindOne(context.Background(), bson.M{"_id": userID}, opt).Decode(&resultUser)
	if err != nil {
		return nil, err
	}
	return resultUser.Notices, nil
}

// GetStars 不分页
func GetStars(userID primitive.ObjectID) ([]*model.StarDto, error) {
	opt := options.FindOne().SetProjection(bson.M{"stars": 1})
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)

	var resultUser model.User
	err := mongoDB.Collection(base.Users).FindOne(context.Background(), bson.M{"_id": userID}, opt).Decode(&resultUser)
	if err != nil {
		return nil, err
	}

	res := funk.Map(resultUser.Stars, func(s *model.Star) *model.StarDto {
		house, _ := public.HouseDetail(s.HouseID, model.Online)
		if house != nil {
			return model.NewStarDto(house, s)
		}
		return nil
	}).([]*model.StarDto)

	return res, nil
}

// GetAppointmentsByUserID 不分页
func GetAppointmentsByUserID(userID primitive.ObjectID) ([]*model.AppointmentDto, error) {
	opt := options.FindOne().SetProjection(bson.M{"appointments": 1})
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)

	var resultUser model.User
	err := mongoDB.Collection(base.Users).FindOne(context.Background(), bson.M{"_id": userID}, opt).Decode(&resultUser)
	if err != nil {
		return nil, err
	}

	res := funk.Map(resultUser.Appointments, func(s *model.Appointment) *model.AppointmentDto {
		house, _ := public.HouseDetail(s.HouseID, model.Online)
		if house != nil {
			return model.NewAppointmentDto(s, house)
		}
		return nil
	}).([]*model.AppointmentDto)

	return res, nil
}

// 一次查询、开销可能太高昂
//aggregate([
// {"$project": {"stars": 1}},
//  {
//      "$facet": {
//          "data": [
//              { "$match": { "_id":ObjectId("627d1b95bc1c4fe8944528f8")  }},
//              { "$skip": 0 },
//              { "$limit": 10 },
//
//          ],
//          "totalCount": [
//              { "$count": "count" }
//          ]
//      }
//  }
//
//])
