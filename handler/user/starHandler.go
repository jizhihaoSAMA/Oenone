package user

import (
	"Oenone/common/base"
	"Oenone/common/crud"
	"Oenone/common/utils"
	"Oenone/model"
	"Oenone/service/user"
	"github.com/gin-gonic/gin"
)

func StarHandler(ctx *gin.Context) {
	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "请求有误", "[StarHandler][GetBodyJson] err: "+err.Error())
		return
	}
	userID := ctx.MustGet("user").(model.User).ID
	houseID := body["houseID"].(string)
	if houseID == "" {
		base.Fail(ctx, "houseID有误")
		return
	}

	err = user.Star(userID, houseID)
	if err != nil {
		base.ServerError(ctx, "", "[StarHandler] 收藏错误： "+err.Error())
		return
	}

	// ZSet score + 1
	err = crud.ZSetIncrBy(base.GetHouseQualifyZSetKey(), houseID, 1)
	if err != nil {
		base.ServerError(ctx, "", "[StarHandler] ZSet添加出错： "+err.Error())
		return
	}
	base.Success(ctx, "", nil)
}

func UnstarHandler(ctx *gin.Context) {
	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "请求有误", "[UnstarHandler][GetBodyJson] err: "+err.Error())
		return
	}
	userID := ctx.MustGet("user").(model.User).ID
	houseID := body["houseID"].(string)
	if houseID == "" {
		base.Fail(ctx, "houseID有误")
		return
	}

	err = user.Unstar(userID, houseID)
	if err != nil {
		base.ServerError(ctx, "", "[UnStarHandler] 收藏错误： "+err.Error())
		return
	}

	// ZSet score - 1
	err = crud.ZSetIncrBy(base.GetHouseQualifyZSetKey(), houseID, -1)
	if err != nil {
		base.ServerError(ctx, "", "[UnStarHandler] ZSet添加出错： "+err.Error())
		return
	}

	base.Success(ctx, "", nil)
}
