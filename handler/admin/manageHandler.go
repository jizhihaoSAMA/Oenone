package admin

import (
	"Oenone/common/base"
	"Oenone/common/crud"
	"Oenone/common/utils"
	"Oenone/model"
	"Oenone/service/admin"
	"Oenone/service/public"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"strconv"
)

func ManageHandler(ctx *gin.Context) {
	pass := base.IsAdmin(ctx)
	if !pass {
		base.UnAuth(ctx, "权限有误")
		return
	}

	field := ctx.Query("field")
	if field == "" {
		base.Fail(ctx)
		return
	}
	page := ctx.Query("page")
	if page == "" {
		base.Fail(ctx)
		return
	}

	pageConverted, _ := strconv.Atoi(page)

	switch field {
	case "review":
		res, total, err := public.HouseList(primitive.ObjectID{}, "", "", "", "", 0, 0, pageConverted, false, false, false, false, model.PendingReView, crud.FieldMapper["admin"])
		if err != nil {
			base.ServerError(ctx, "", "[ManageHandler] 获取审核信息出错, err: "+err.Error())
			return
		}
		base.Success(ctx, "ok", gin.H{"reviews": res, "total_page": math.Ceil(float64(total) / 12), "current_page": page, "total_count": total})
		return

	case "appointment":
		res, total, err := admin.GetAppointments(pageConverted)
		if err != nil {
			base.ServerError(ctx, "", "[ManageHandler] 获取预约信息异常, err: "+err.Error())
			return
		}
		base.Success(ctx, "ok", gin.H{"appointments": res, "total_page": math.Ceil(float64(total) / 12), "current_page": page, "total_count": total})
		return

	case "house":
		res, total, err := public.HouseList(primitive.ObjectID{}, "", "", "", "", 0, 0, pageConverted, false, false, false, false, 0, crud.FieldMapper["admin"])
		if err != nil {
			base.ServerError(ctx, "", "[ManageHandler] 获取房屋管理信息出错, err: "+err.Error())
			return
		}
		base.Success(ctx, "ok", gin.H{"houses": res, "total_page": math.Ceil(float64(total) / 12), "current_page": page, "total_count": total})
		return
	case "admin":
		if !base.IsSuperuser(ctx) {
			base.Response(ctx, 403, "权限不足, 您不是超级管理员", nil)
			return
		}

		res, total, err := admin.GetAdmins(pageConverted)
		if err != nil {
			base.ServerError(ctx, "", "[ManageHandler] 获取管理员信息出错, err: "+err.Error())
		}
		base.Success(ctx, "ok", gin.H{"admins": res, "total_page": math.Ceil(float64(total) / 12), "current_page": page, "total_count": total})
		return
	}

}

func AddAdminHandler(ctx *gin.Context) {
	if !base.IsSuperuser(ctx) {
		base.Response(ctx, 403, "权限不够", nil)
		return
	}

	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx)
		return
	}

	err = admin.AddAdmins(body["username"].(string), body["password"].(string), int(body["role"].(float64)))
	if err != nil {
		base.Fail(ctx, "添加管理错误", "[AddAdminHandler] err: "+err.Error())
		return
	}

}

func DeleteAdminHandler(ctx *gin.Context) {
	if !base.IsSuperuser(ctx) {
		base.Response(ctx, 403, "权限不够", nil)
		return
	}

	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx)
		return
	}

	val, ok := body["adminId"]
	if !ok {
		base.Fail(ctx)
		return
	}

	err = admin.DeleteAdmin(utils.UserID2ObjectID(val.(string)))
}
