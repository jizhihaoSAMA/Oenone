package user

import (
	"Oenone/common/base"
	"Oenone/common/crud"
	"Oenone/common/utils"
	"Oenone/model"
	"Oenone/service/user"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func PostHouseHandler(ctx *gin.Context) {
	var house model.House
	if ctx.ShouldBind(&house) != nil {
		base.Fail(ctx, "请求参数有误", "BindJson错误")
		return
	}

	form, err := ctx.MultipartForm()

	house.OwnerID = ctx.MustGet("user").(model.User).ID
	house.Status = model.PendingReView

	if err != nil {
		house.ImageAmount = 0
	} else {
		house.ImageAmount = len(form.File["imageList"])
	}

	house.CreatedTime = time.Now()

	id, err := user.PostHouse(house)
	if err != nil {
		log.Println("Post house err: ", err.Error())
	}
	// 下载图片
	if house.ImageAmount != 0 {
		if !utils.SaveUploadImages(ctx, form.File["imageList"], "house\\"+id, ".png") {
			base.ServerError(ctx, "图片上传失败，内部错误")
			return
		}
	}
	// 审核列表
	mongoDB := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
	_, err = mongoDB.Collection(base.Reviews).InsertOne(context.Background(), model.NewReview(id))
	if err != nil {
		base.ServerError(ctx, "审核信息失败", "审核信息失败: err: "+err.Error())
		return
	}
	base.Success(ctx, "上传成功，请等待内部管理员审核", nil)
}

func OfflineHouseHandler(ctx *gin.Context) {
	body, err := utils.GetPostBodyJson(ctx)
	if err != nil {
		base.Fail(ctx, "请求有误")
		return
	}

	houseID := body["houseID"].(string)
	err = user.OfflineHouse(houseID)
	if err != nil {
		base.ServerError(ctx, "删除失败", "[OfflineHouseHandler] err :"+err.Error())
		return
	}

	err = crud.DelZSetMember(base.GetHouseQualifyZSetKey(), houseID)
	if err != nil {
		log.Println("[OfflineHouseHandler] 删除zset失败: " + err.Error())
	}

	base.Success(ctx, "删除成功", nil)
}

func UpdateHouseHandler(ctx *gin.Context) {
	var house model.House
	if ctx.ShouldBind(&house) != nil {
		base.Fail(ctx, "请求参数有误", "BindJson错误")
		return
	}
	houseID := ctx.PostForm("id")

	form, _ := ctx.MultipartForm()
	house.Status = model.PendingReView
	house.ImageAmount = len(form.File["imageList"])
	err := user.UpdateHouse(house, houseID)
	if err != nil {
		base.ServerError(ctx, "更新房屋信息失败", "[UpdateHouseHandler] err: "+err.Error())
		return
	}
	// 下载图片
	if len(form.File["imageList"]) != 0 {
		if !utils.SaveUploadImages(ctx, form.File["imageList"], "house\\"+houseID, ".png") {
			base.ServerError(ctx, "图片上传失败，内部错误")
			return
		}
	}
	base.Success(ctx, "ok", nil)
}
