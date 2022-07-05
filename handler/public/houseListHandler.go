package public

import (
	"Oenone/common/base"
	"Oenone/common/crud"
	"Oenone/model"
	"Oenone/service/public"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"strconv"
)

func HouseListHandler(ctx *gin.Context) {
	if trend := ctx.Query("trend"); trend == "1" {
		res, err := public.GetHotHouse()
		if err != nil {
			base.ServerError(ctx, "", "[HouseListHandler] 获取热门房屋ID出错, err: "+err.Error())
			return
		}
		base.Success(ctx, "ok", gin.H{"houses": res})
		return
	}

	area := ctx.Query("area")
	bc := ctx.Query("bc")
	rentType := ctx.Query("rentType")
	priceStart := ctx.Query("priceStart")
	priceEnd := ctx.Query("priceEnd")

	supportShortTermRent := ctx.Query("supportShortTermRent")
	hasLift := ctx.Query("hasLift")
	hasSingleToilet := ctx.Query("hasSingleToilet")
	hasSingleBalcony := ctx.Query("hasSingleBalcony")

	neighborhood := ctx.Query("neighborhood")

	var (
		priceStartConverted int
		priceEndConverted   int
		pageConverted       = 1
		status              int

		supportShortTermRentConverted bool
		hasLiftConverted              bool
		hasSingleToiletConverted      bool
		hasSingleBalconyConverted     bool

		err error
	)

	if !base.IsAdmin(ctx) {
		status = model.Online
	}

	if supportShortTermRent != "" {
		supportShortTermRentConverted, err = strconv.ParseBool(supportShortTermRent)
		if err != nil {
			base.Fail(ctx, "请求参数有误: supportShortTermRent")
		}
	}
	if hasLift != "" {
		hasLiftConverted, err = strconv.ParseBool(hasLift)
		if err != nil {
			base.Fail(ctx, "请求参数有误: hasLift")
		}
	}
	if hasSingleToilet != "" {
		hasSingleToiletConverted, err = strconv.ParseBool(hasSingleToilet)
		if err != nil {
			base.Fail(ctx, "请求参数有误: hasSingleToilet")
		}
	}
	if hasSingleBalcony != "" {
		hasSingleBalconyConverted, err = strconv.ParseBool(hasSingleBalcony)
		if err != nil {
			base.Fail(ctx, "请求参数有误: hasSingleBalcony")
		}
	}

	page := ctx.Query("page")
	m := ctx.Query("fieldMapper")

	if area == "" && bc != "" {
		base.Fail(ctx, "请求参数有误: area, bc")
		return
	}

	if priceStart != "" {
		priceStartConverted, err = strconv.Atoi(priceStart)
		if err != nil {
			base.Fail(ctx, "请求参数有误: priceStart")
			return
		}
	}

	if priceStart != "" {
		priceEndConverted, err = strconv.Atoi(priceEnd)
		if err != nil {
			base.Fail(ctx, "请求参数有误: priceEnd")
			return
		}
	}

	if priceEndConverted != 0 && priceStartConverted != 0 && priceEndConverted <= priceStartConverted {
		base.Fail(ctx, "请求参数有误: priceStart, priceEnd")
		return
	}

	if page != "" {
		pageConverted, err = strconv.Atoi(page)
		if err != nil {
			base.Fail(ctx, "请求参数有误: page")
			return
		}
	}

	v, total, err := public.HouseList(primitive.ObjectID{}, area, bc, neighborhood, rentType, priceStartConverted, priceEndConverted, pageConverted, supportShortTermRentConverted, hasLiftConverted, hasSingleToiletConverted, hasSingleBalconyConverted, status, crud.FieldMapper[m])
	if err != nil {
		base.ServerError(ctx)
		return
	}

	base.Success(ctx, "ok", gin.H{"houses": v, "total_page": math.Ceil(float64(total) / 12), "current_page": page, "total_count": total})
	return
}
