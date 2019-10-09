package handler

import (
	"knowledgemap_backend/library/database/mongo"
	"knowledgemap_backend/microservices/knowledgemap/passport/dao"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	gdao *dao.Dao
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
}
