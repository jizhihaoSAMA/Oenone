package admin

import (
	"Oenone/common/base"
	"Oenone/common/utils"
	"Oenone/service/admin"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	if !base.IsSuperuser(ctx) {
		base.Response(ctx, 403, "权限不足", nil)
		return
	}

	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "", "[DeleteUserHandler] err: "+err.Error())
		return
	}

	userId_, ok := body["userId"]
	if !ok {
		base.Fail(ctx)
		return
	}

	userId := utils.UserID2ObjectID(userId_.(string))

	err = admin.DeleteUser(userId)
	if err != nil {
		base.ServerError(ctx, "", "[DeleteUserHandler]操作DB失败err: "+err.Error())
		return
	}
	base.Success(ctx, "ok", nil)

}
