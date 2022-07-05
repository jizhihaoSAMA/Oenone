package user

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"time"
)

func RegisterCaptchaHandler(ctx *gin.Context) {
	body, err := utils.GetPostBodyJson(ctx)

	telephone, ok := (body["telephone"]).(string)
	if !ok {
		base.Fail(ctx, "请求参数有误或参数为空")
		return
	}

	captcha, err := utils.SendCaptcha(telephone)
	if err != nil {
		base.Fail(ctx, "发送消息失败", err)
		return
	}

	rdb := base.GLOBAL_RESOURCE[base.RedisClient].(*redis.Client)
	err = rdb.Set("[reg]"+telephone, captcha, 5*time.Minute).Err()
	if err != nil {
		base.ServerError(ctx, "系统内部错误，发送消息失败", err)
		return
	}

	base.Success(ctx, "发送成功", nil)
}
