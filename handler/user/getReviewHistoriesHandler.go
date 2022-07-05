package user

import (
	"Oenone/common/base"
	"Oenone/service/user"
	"github.com/gin-gonic/gin"
)

func GetReviewHistoriesHandler(ctx *gin.Context) {
	houseID := ctx.Query("houseID")
	val, err := user.GetReviewHistories(houseID)
	if err != nil {
		base.ServerError(ctx, "", "[GetReviewsHistories] err: "+err.Error())
		return
	}

	base.Success(ctx, "ok", gin.H{"review_histories": val.Histories})
}
