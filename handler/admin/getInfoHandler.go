package admin

import (
	"Oenone/common/base"
	"Oenone/model"
	"github.com/gin-gonic/gin"
)

func GetAdminInfoHandler(ctx *gin.Context) {
	role := base.GetAdminRole(ctx)
	if role == base.Visitor {
		base.UnAuth(ctx, "校验错误")
		return
	}
	user := ctx.MustGet("admin").(model.Admin)
	base.Success(ctx, "", gin.H{"admin": gin.H{"id": ctx.MustGet("adminID"), "role": role, "admin": user}})
}
