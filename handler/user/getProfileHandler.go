package user

import (
	"Oenone/common/base"
	"Oenone/model"
	"Oenone/service/user"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

func GetProfileHandler(ctx *gin.Context) {
	userID := ctx.MustGet("user").(model.User).ID
	field := ctx.Query("field")
	page := ctx.Query("page")
	if field == "" {
		base.Fail(ctx, "field字段错误")
		return
	}

	pageConverted, err := strconv.Atoi(page)
	if field == "house" && (page == "" || err != nil || pageConverted <= 0) {
		base.Fail(ctx, "page字段错误")
		return
	}
	switch field {
	case "house":
		data, total, err := user.GetOwnHouses(userID, pageConverted)
		if err != nil {
			base.ServerError(ctx, "", "[GetProfileHandler] err: "+err.Error())
			return
		}
		base.Success(ctx, "", gin.H{"houses": data, "total": total, "total_page": math.Ceil(float64(total) / 12)})
		return

	case "notice":
		notices, err := user.GetNotices(userID)
		if err != nil {
			base.Fail(ctx, "用户不存在")
			return
		}
		base.Success(ctx, "", gin.H{"notices": notices})
		return
	case "star":
		stars, err := user.GetStars(userID)
		if err != nil {
			base.Fail(ctx, "用户不存在")
			return
		}
		base.Success(ctx, "", gin.H{"stars": stars})
		return
	case "appointment":
		appointments, err := user.GetAppointmentsByUserID(userID)
		if err != nil {
			base.Fail(ctx, "用户不存在")
		}
		base.Success(ctx, "", gin.H{"appointments": appointments})
		return
	default:
		base.Fail(ctx, "请求field有误")
		return
	}
}
