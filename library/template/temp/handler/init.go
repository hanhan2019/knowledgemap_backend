package handler

import (
	"knowledgemap_backend/{{ .PackagePath }}/dao"
	"knowledgemap_backend/library/database/mongo"
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
	gdao = dao.InitDao(db)
}
