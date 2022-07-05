package user

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(username string, password string, telephone string) error {
	db := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err := db.Collection(base.Users).InsertOne(context.Background(), model.NewUser(username, password, telephone))
	return err
}
