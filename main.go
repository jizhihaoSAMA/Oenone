package main

import (
	"Oenone/common/base"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

func main() {
	base.InitService()
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()
	BindRoutes(r)
	r.Static("api/house", "./static/house")

	//db := base.GLOBAL_RESOURCE[base.MySQLClient].(*gorm.DB)
	//
	//err := db.AutoMigrate(
	//	&model.user{},
	//)

	//if err != nil {
	//    log.Println("[Init] Auto migrate error, err is", err)
	//}

	port := viper.GetString("port.server")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}
