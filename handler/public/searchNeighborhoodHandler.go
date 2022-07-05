package public

import (
	"Oenone/common/base"
	"Oenone/service/public"
	"github.com/gin-gonic/gin"
	"log"
)

func SearchNeighborhoodHandler(ctx *gin.Context) {
	neighborhood, ok := ctx.GetQuery("neighborhood")
	if !ok {
		base.Fail(ctx, "小区不能为空")
		return
	}

	area, ok := ctx.GetQuery("area")
	if !ok {
		base.Fail(ctx, "地区不能为空")
		return
	}

	log.Println(neighborhood, area)
	result, err := public.SearchNeighborhood(neighborhood, area, 3)
	if err != nil {
		base.ServerError(ctx, "内部查询错误", "[SearchNeighborhoodHandler] 查询ES出错, err: "+err.Error())
		return
	}
	log.Println(result)

	base.Success(ctx, "ok", gin.H{
		"neighborhood": result,
	})
	return
}
