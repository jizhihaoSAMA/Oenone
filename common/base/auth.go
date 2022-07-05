package base

import (
	"Oenone/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const (
	Superuser = 4
	Agent     = 3
	User      = 2
	Visitor   = 1
)

func GetRole(ctx *gin.Context) int {
	// 获取authorization header
	tokenString := ctx.GetHeader("Authorization")

	// 验证其是否合法
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return Visitor
	}

	tokenString = tokenString[7:]

	token, claims, err := ParseToken(tokenString)

	if err != nil || !token.Valid {
		return Visitor
	}

	// 验证通过后获取claim中的userId

	userId := claims.UserId
	db := GLOBAL_RESOURCE[MongoDB].(*mongo.Database)
	var tempUser model.User
	err = db.Collection(Users).FindOne(context.Background(), bson.D{{
		"_id", userId,
	}}).Decode(&tempUser)
	if err != nil {
		return Visitor
	}

	// 用户查找 但不存在
	if tempUser.Password == "" {
		return Visitor
	}

	ctx.Set("userID", userId)
	ctx.Set("user", tempUser)
	return User
}

func GetAdminRole(ctx *gin.Context) int {
	// 获取authorization header
	tokenString := ctx.GetHeader("Authorization")

	// 验证其是否合法
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return Visitor
	}

	tokenString = tokenString[7:]

	token, claims, err := ParseToken(tokenString)

	if err != nil || !token.Valid {
		return Visitor
	}

	// 验证通过后获取claim中的userId
	userId := claims.UserId
	db := GLOBAL_RESOURCE[MongoDB].(*mongo.Database)
	var tempUser model.Admin
	err = db.Collection(Admins).FindOne(context.Background(), bson.D{{
		"_id", userId,
	}}).Decode(&tempUser)
	if err != nil {
		return Visitor
	}

	ctx.Set("adminID", userId)
	ctx.Set("admin", tempUser)

	return tempUser.Role
}

func IsSuperuser(ctx *gin.Context) bool {
	if val, exists := ctx.Get("admin"); !exists {
		return GetAdminRole(ctx) == Superuser
	} else {
		return val.(model.Admin).Role == Superuser
	}
}

func IsAdmin(ctx *gin.Context) bool {
	if val, exists := ctx.Get("admin"); !exists {
		return funk.ContainsInt([]int{Agent, Superuser}, GetAdminRole(ctx))
	} else {
		return funk.ContainsInt([]int{Agent, Superuser}, val.(model.Admin).Role)
	}
}

func IsVisitor(ctx *gin.Context) bool {
	return GetRole(ctx) == Visitor
}

func IsUser(ctx *gin.Context) bool {
	return GetRole(ctx) == User
}
