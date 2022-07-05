package user

import (
	"Oenone/common/base"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckAppointed(ctx *gin.Context, houseID string) bool {
	if base.IsVisitor(ctx) {
		return false
	}
	userID := ctx.MustGet("userID").(primitive.ObjectID)

	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	err := mongoDB.Collection(base.Users).FindOne(context.Background(), bson.M{"_id": userID, "appointments._id": houseID}).Err()
	if err != nil {
		return false
	}
	return true

}
