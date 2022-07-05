package user

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/model"
	"github.com/gin-gonic/gin"
)

func GetInfoHandler(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	base.Success(ctx, "", gin.H{"user": utils.ToUserDto(user.(model.User))})
}
