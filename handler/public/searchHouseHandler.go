package public

import (
	"Oenone/common/base"
	"Oenone/service/public"
	"github.com/gin-gonic/gin"
)

func SearchHouseHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		base.Fail(ctx, "name字段不能为空")
		return
	}

	res, err := public.SearchHouse(name, ctx.Query("area"), ctx.Query("bc"), ctx.Query("neighborhood"))
	if err != nil {
		base.ServerError(ctx, "系统服务异常，查询错误")
		return
	}
	base.Success(ctx, "ok", gin.H{"suggestions": res})
}
