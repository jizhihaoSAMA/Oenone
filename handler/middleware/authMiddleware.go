package middleware

import (
    "Oenone/common/base"
    "Oenone/model"
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "strings"
)

func AuthMiddleware(ctx *gin.Context) {
    // 获取authorization header
    tokenString := ctx.GetHeader("Authorization")

    // 验证其是否合法
    if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
        base.UnAuth(ctx, "token不合法:1")
        ctx.Abort()
        return
    }

    tokenString = tokenString[7:]

    token, claims, err := base.ParseToken(tokenString)

    if err != nil || !token.Valid {
        base.UnAuth(ctx, "token不合法:2")
        ctx.Abort()
        return
    }

    // 验证通过后获取claim中的userId

    userId := claims.UserId
    opt := options.FindOne().SetProjection(bson.M{"_id": 1, "username": 1})
    db := base.GLOBAL_RESOURCE[base.MongoDB].(*mongo.Database)
    var user model.User
    err = db.Collection(base.Users).FindOne(context.Background(), bson.D{{
        "_id", userId,
    }}, opt).Decode(&user)
    if err != nil {
        base.ServerError(ctx, "内部错误", "校验失败 err: "+err.Error())
        ctx.Abort()
        return
    }
    // 将用户的存在写入上下文

    ctx.Set("user", user)
    ctx.Set("userID", userId)
    ctx.Next()
}
