package handler

import (
	"knowledgemap_backend/library/database/mongo"
	"knowledgemap_backend/microservices/common/namespace"
	"knowledgemap_backend/microservices/knowledgemap/class/dao"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"

	mclient "github.com/micro/go-micro/client"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	gdao    *dao.Dao
	userSrv uapi.UserService
)

func Init() {
	db := mongo.NewDb(viper.GetString("db.mongo.uri"))
	if db == nil {
		panic("init db error")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("db.redis.addr"),
		Password: "",                          // no password set
		DB:       viper.GetInt("db.redis.Db"), // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	gdao = dao.InitDao(db, client)
	userSrv = uapi.NewUserService(namespace.GetName("microservices.knowledgemap.user"), mclient.DefaultClient)

}
