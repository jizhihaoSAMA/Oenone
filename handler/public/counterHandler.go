package public

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/service/public"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"
)

//var QueryMap = map[int]func(){
//
//}

// HouseCounterHandler 四种查询方式，
// 1. 传递经纬度，返回当前经纬度内所有小区的聚合结果。 zoomLevel = 4
// 2. 传递商圈，返回聚合结果。 zoomLevel = 3
// 3. 返回区县聚合结果。 zoomLevel = 2
// 4. 什么都不传，返回总数。zoomLevel = 1
func HouseCounterHandler(ctx *gin.Context) {
	zl := ctx.Query("zoomLevel")
	if zoomLevel, err := strconv.Atoi(zl); err != nil {
		base.Fail(ctx, "请求参数有误: zoomLevel")
		return

	} else if zoomLevel == 1 {
		if total := public.GetAmountOfHouse(); total == -1 {
			base.ServerError(ctx, "内部查询错误")
		} else {
			base.Success(ctx, "ok", gin.H{"total": total})
		}
		return

	} else if zoomLevel == 2 {
		// 获取所有区县结果
		if res := public.GetGroupByPlaces("house_loc_area", nil); res == nil {
			base.ServerError(ctx, "查询有误，内部服务异常")
		} else {
			base.Success(ctx, "ok", gin.H{"counter": res})
		}
		return

	} else if zoomLevel == 3 {
		// 查询商圈，返回所有商圈商圈结果
		// 校验请求信息places字段
		placeArr := strings.Split(ctx.Query("places"), ",")
		if len(placeArr) == 0 {
			base.Fail(ctx, "请求参数有误: places")
			return
		}

		if res := public.GetGroupByPlaces("house_loc_bc", placeArr); res == nil {
			base.ServerError(ctx, "查询有误，内部服务异常")
		} else {
			base.Success(ctx, "ok", gin.H{"counter": res})
		}
		return

	} else if zoomLevel == 4 {
		// 经纬度查矩阵，传入方式lat1,lng1-lat2,lng2
		loc := ctx.Query("location")
		location := utils.GetLocationByParams(loc)
		if location == nil {
			base.Fail(ctx, "参数有误")
			return
		}
		if math.Abs(location.BottomLeftLatFloat-location.TopRightLatFloat) > 1 || math.Abs(location.BottomLeftLngFloat-location.TopRightLngFloat) > 1 {
			base.Fail(ctx, "矩阵过大")
			return
		}

		counter := public.GetGroupByLocation(location)
		base.Success(ctx, "ok", gin.H{"counter": counter})
		return

	} else {
		base.Fail(ctx, "请求参数有误: zoomLevel")
		return
	}
}
