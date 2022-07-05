package utils

import "github.com/gin-gonic/gin"

func GetPostBodyJson(ctx *gin.Context) (map[string]interface{}, error) {
	var f map[string]interface{}
	err := ctx.BindJSON(&f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
