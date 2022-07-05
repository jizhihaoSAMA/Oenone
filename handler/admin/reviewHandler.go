package admin

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/service/admin"
	"github.com/gin-gonic/gin"
)

func ReviewHandler(ctx *gin.Context) {
	if !base.IsAdmin(ctx) {
		base.UnAuth(ctx)
		return
	}

	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "")
		return
	}

	houseID := body["houseID"].(string)
	msg := body["msg"].(string)
	pass := body["pass"].(bool)

	err = admin.AddReview(houseID, msg, pass)
	if err != nil {
		base.ServerError(ctx, "", "[ReviewHandler] review出错，错误为: "+err.Error())
		return
	}
}
