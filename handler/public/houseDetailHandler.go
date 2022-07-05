package public

import (
	"Oenone/common/base"
	"Oenone/common/crud"
	"Oenone/model"
	"Oenone/service/public"
	"Oenone/service/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func HouseDetailHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		base.Fail(ctx, "字段有误，id")
		return
	}

	house, err := public.HouseDetail(id, 0)
	if err != nil {
		base.ServerError(ctx, "服务异常，查询错误", "[HouseDetailHandler] 查询出错，错误: "+err.Error())
		return
	}

	// 权限校验
	if house.Status != model.Online { // 不在线
		if !base.IsAdmin(ctx) && (base.GetRole(ctx) == base.Visitor || ctx.MustGet("userID").(primitive.ObjectID) != house.OwnerID) { // 非管理( 游客、用户ID 不是本人不可见)
			base.Response(ctx, 403, "权限不足，当前房屋不可见", nil)
			return
		}
	}

	res := base.StructToMap(house)

	res["starred"] = user.CheckStarred(ctx, house.ID)
	res["appointed"] = user.CheckAppointed(ctx, house.ID)

	// ZSet score + 0.05
	err = crud.ZSetIncrBy(base.GetHouseQualifyZSetKey(), id, 0.05)
	if err != nil {
		log.Println("[HouseDetailHandler] ZSet add score err: ", err)
	}
	base.Success(ctx, "ok", gin.H{"house": res})
}
