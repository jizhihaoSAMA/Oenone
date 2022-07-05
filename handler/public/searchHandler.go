package public

import (
    "Oenone/common/base"
    "github.com/gin-gonic/gin"
)

var handlerMapper = map[string]gin.HandlerFunc{
    "house":        SearchHouseHandler,
    "neighborhood": SearchNeighborhoodHandler,
}

func SearchHandler(ctx *gin.Context) {
    field := ctx.Query("field")
    if field == "" {
        base.Fail(ctx, "field字段不能为空")
        return
    }
    f, ok := handlerMapper[field]
    if !ok {
        base.Fail(ctx, "field字段有误")
        return
    }
    f(ctx)
}
