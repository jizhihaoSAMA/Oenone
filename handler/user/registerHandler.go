package user

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/model"
	"Oenone/service/user"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func validateRegInfo(body map[string]interface{}) (bool, func(ctx *gin.Context, Msg ...interface{}), string) {
	requiredKeys := map[string]string{
		"username":        "用户名",
		"telephone":       "电话",
		"password":        "密码",
		"repeat_password": "确认密码",
		"captcha":         "验证码",
	}
	for k, v := range requiredKeys {
		if content, ok := body[k]; !ok || content.(string) == "" {
			return false, base.Fail, v + "不能为空"
		}
	}

	if strings.HasPrefix(body["username"].(string), "test_") { //测试用户直接过
		return true, nil, ""
	}

	if body["password"] != body["repeat_password"] {
		return false, base.Fail, "两次密码不一致"
	}

	if len(body["telephone"].(string)) != 11 {
		return false, base.Fail, "电话格式不对"
	}

	if len(body["password"].(string)) < 6 {
		return false, base.Fail, "密码不得小于6位数"
	}

	rdb := base.GLOBAL_RESOURCE[base.RedisClient].(*redis.Client)
	authCaptcha, err := rdb.Get("[reg]" + (body["telephone"].(string))).Result()
	if err != nil {
		return false, base.ServerError, "验证码失效"
	}

	if authCaptcha != body["captcha"].(string) {
		return false, base.Fail, "验证码错误"
	}

	var temp model.User
	db := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	err = db.Collection(base.Users).FindOne(context.Background(), bson.D{{
		"telephone", body["telephone"].(string),
	}}).Decode(&temp)
	if err == nil {
		return false, base.Fail, "电话号码已存在"
	}

	return true, nil, ""
}

func RegisterHandler(ctx *gin.Context) {
	bodyJson, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "请求参数有误")
		return
	}

	if pass, f, msg := validateRegInfo(bodyJson); !pass {
		f(ctx, msg)
		return
	}

	username, telephone, password := bodyJson["username"].(string), bodyJson["telephone"].(string), bodyJson["password"].(string)

	err = user.Register(username, password, telephone)
	if err != nil {
		base.ServerError(ctx, "注册失败，服务器内部异常", "注册失败")
		return
	}
	base.Success(ctx, "注册成功", nil)
}
