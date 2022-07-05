package admin

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/model"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var resultUser model.Admin

func validateLoginInfo(body map[string]interface{}) (bool, func(ctx *gin.Context, Msg ...interface{}), string) {
	requiredKeys := map[string]string{
		"username": "用户名",
		"password": "密码",
	}

	for k, v := range requiredKeys {
		if content, ok := body[k]; !ok || content.(string) == "" {
			return false, base.Fail, v + "不能为空"
		}
	}

	db := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	err := db.Collection(base.Admins).FindOne(context.Background(), bson.D{{"username", body["username"].(string)}}).Decode(&resultUser)
	if err != nil {
		return false, base.Fail, "用户不存在"
	}
	log.Println(resultUser.Password, body["password"])
	if resultUser.Password != body["password"].(string) {
		return false, base.Fail, "密码错误"
	}
	return true, nil, ""
}

func LoginHandler(ctx *gin.Context) {
	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "请求错误")
		return
	}

	if pass, f, msg := validateLoginInfo(body); !pass {
		f(ctx, msg)
		return
	}

	log.Println(body)

	token, err := base.GetToken(resultUser.ID)
	if err != nil {
		base.ServerError(ctx, "内部服务错误", "登录无法生成token")
		return
	}

	base.Success(ctx, "", gin.H{"token": token})
}
