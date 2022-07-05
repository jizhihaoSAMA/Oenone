package user

import (
    "Oenone/common/base"
    "Oenone/service/user"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUnreadNoticeAmountHandler(ctx *gin.Context) {
    userID := ctx.MustGet("userID").(primitive.ObjectID)

    amount, err := user.GetUnreadAmount(userID)
    if err != nil {
        base.Fail(ctx, "", "[GetUnreadAmountHandler]获取数量错误, err: "+err.Error())
        return
    }
    base.Success(ctx, "ok", gin.H{"unread_amount": amount})
}

func NoticeAllReadHandler(ctx *gin.Context) {
    userID := ctx.MustGet("userID").(primitive.ObjectID)
    err := user.NoticeAllRead(userID)
    if err != nil {
        base.Fail(ctx, "", "[NoticeAllReadHandler] 全部已读错误, err:"+err.Error())
        return
    }
    base.Success(ctx, "ok", nil)
}