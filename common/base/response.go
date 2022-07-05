package base

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"log"
)

func Response(ctx *gin.Context, respCode int, msg string, data gin.H) {
	ctx.JSON(respCode, gin.H{"code": respCode, "msg": msg, "data": data})
}

func Success(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(200, gin.H{"code": 200, "msg": funk.GetOrElse(msg, "ok"), "data": data})
}

func UnAuth(ctx *gin.Context, msg ...interface{}) {
	if len(msg) == 0 {
		ctx.JSON(401, gin.H{
			"code": 401,
			"msg":  "权限不足",
			"data": nil,
		})
		return
	}
	if len(msg) == 2 && msg[1] != "" {
		log.Println(ctx, "err msg: ", msg[1])
	}
	ctx.JSON(401, gin.H{
		"code": 401,
		"msg":  msg[0],
		"data": nil,
	})
}

// Fail
// @ctx 函数上下文
// @Msg 传达信息，第一个字符串为传递给前端的信息，第二个字符串为内部日志输出的信息
func Fail(ctx *gin.Context, msg ...interface{}) {
	if len(msg) == 0 {
		msg = append(msg, "请求参数有误")
	}

	if len(msg) == 2 && msg[1] != "" {
		log.Println(ctx, "err msg: ", msg[1])
	}
	ctx.JSON(400, gin.H{
		"code": 400,
		"msg":  funk.GetOrElse(msg[0], "请求参数有误"),
		"data": nil,
	})
}

func ServerError(ctx *gin.Context, msg ...interface{}) {
	if len(msg) == 0 {
		ctx.JSON(500, gin.H{
			"code": 500,
			"msg":  "内部服务错误",
			"data": nil,
		})
		return
	}

	if len(msg) == 2 && msg[1] != "" {
		log.Println(ctx, "err msg: ", msg[1])
	}

	ctx.JSON(500, gin.H{
		"code": 500,
		"msg":  funk.GetOrElse(msg[0], "内部服务错误"),
		"data": nil,
	})
}
