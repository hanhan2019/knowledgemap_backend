package handler

import (
	"knowledgemap_backend/library/database/mongo"
	"knowledgemap_backend/microservices/common/namespace"
	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
	kapi "knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"
	"knowledgemap_backend/microservices/knowledgemap/question/dao"

	mclient "github.com/micro/go-micro/client"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	gdao            *dao.Dao
	knowledgeMapSrv kapi.KnowledegeMapService
	classSrv        capi.ClassService
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

	knowledgeMapSrv = kapi.NewKnowledegeMapService(namespace.GetName("microservices.knowledgemap.knowledgemap"), mclient.DefaultClient)
	classSrv = capi.NewClassService(namespace.GetName("microservices.knowledgemap.class"), mclient.DefaultClient)

}
