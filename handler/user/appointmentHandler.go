package user

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/model"
	"Oenone/service/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func MakeAppointmentHandler(ctx *gin.Context) {
	userID := ctx.MustGet("user").(model.User).ID
	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx)
		return
	}
	appointedTime, ok := body["appointedTime"].(string)
	if !ok {
		base.Fail(ctx)
		return
	}
	houseID, ok := body["houseID"].(string)
	if !ok {
		base.Fail(ctx)
		return
	}

	t, err := time.Parse("2006-01-02", appointedTime)
	if err != nil {
		base.Fail(ctx)
		return
	}

	err = user.MakeAppointment(userID, houseID, primitive.NewDateTimeFromTime(t))
	if err != nil {
		base.ServerError(ctx)
		return
	}
	base.Success(ctx, "", nil)

}

func DelAppointmentHandler(ctx *gin.Context) {
	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "请求有误", "[DelAppointmentHandler][GetBodyJson] err: "+err.Error())
		return
	}
	userID := ctx.MustGet("user").(model.User).ID
	houseID := body["houseID"].(string)
	if houseID == "" {
		base.Fail(ctx, "houseID有误")
		return
	}

	err = user.DelAppointment(userID, houseID)
	if err != nil {
		base.ServerError(ctx, "", "[DelAppointmentHandler] 取消预约失败： "+err.Error())
		return
	}
	base.Success(ctx, "", nil)
}
