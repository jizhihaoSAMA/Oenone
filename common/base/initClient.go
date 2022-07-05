package base

import (
	"context"
	"flag"
	"fmt"
	"github.com/cloopen/go-sms-sdk/cloopen"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var GLOBAL_RESOURCE = make(map[int]interface{})

// client 常量
const (
	CaptchaClientConfig = 1
	RedisClient         = 2
	MySQLClient         = 3
	ESClient            = 4
	MongoDB             = 5
	WorkDir             = 6
)

// 数据库表、ES索引、mongo-collection
const (
	// ES索引常量
	NeighborhoodInfo = "neighborhood_info"
	HouseInfo        = "house_info"

	// mongo-collection
	Users   = "users"
	Reviews = "reviews"
	Admins  = "admins"

	// redis Zset
	HouseQualify = "house_qualify"
)

func InitService() {
	initConfig()
	initCaptchaConfig()
	// 初始话数据库
	initDB()
}

func initDB() {
	initRedis()
	initMySQL()
	initES()
	initMongo()
}

func initConfig() {
	workDir, _ := os.Getwd()
	GLOBAL_RESOURCE[WorkDir] = workDir
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	if flag.Lookup("test.v") != nil { //测试
		viper.AddConfigPath("E:/work/Compile/go/Oenone" + "/config")
	} else {
		viper.AddConfigPath(workDir + "/config")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initMongo() {
	uri := fmt.Sprintf("%s://%s:%s",
		viper.GetString("dataSource.MongoDB.driverName"),
		viper.GetString("dataSource.MongoDB.host"),
		viper.GetString("dataSource.MongoDB.port"),
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("连接mongo失败err:" + err.Error())
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic("无法ping通mongo, err:" + err.Error())
	}

	GLOBAL_RESOURCE[MongoDB] = client.Database(viper.GetString("dataSource.MongoDB.database"))
}

func initCaptchaConfig() {
	cfg := cloopen.DefaultConfig().
		WithAPIAccount(viper.GetString("messageInfo.APIAccount")).
		WithAPIToken(viper.GetString("messageInfo.APIToken"))

	GLOBAL_RESOURCE[CaptchaClientConfig] = cfg
}

func initRedis() {
	port := viper.GetString("dataSource.Redis.port")
	password := viper.GetString("dataSource.Redis.password")
	db := viper.GetInt("dataSource.Redis.db")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic("链接Redis失败：" + err.Error())
	}

	GLOBAL_RESOURCE[RedisClient] = rdb
}

func initMySQL() {
	host := viper.GetString("dataSource.MySQL.host")
	port := viper.GetString("dataSource.MySQL.port")
	database := viper.GetString("dataSource.MySQL.database")
	username := viper.GetString("dataSource.MySQL.username")
	password := viper.GetString("dataSource.MySQL.password")
	charset := viper.GetString("dataSource.MySQL.charset")
	loc := viper.GetString("dataSource.MySQL.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("链接MYSQL失败，错误:" + err.Error())
	}
	GLOBAL_RESOURCE[MySQLClient] = db
}

func initES() {
	port := viper.GetString("dataSource.ES.port")
	es, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:"+port), elastic.SetSniff(false))
	if err != nil {
		panic("链接ES失败，错误:" + err.Error())
	}

	GLOBAL_RESOURCE[ESClient] = es
}

func GetHouseQualifyZSetKey() string {
	return fmt.Sprintf("[%s]%s", time.Now().Format("2006-01-02"), HouseQualify)
}
