package dao

import (
	"myProjects/collegeManage/library/database/mongo"

	"github.com/spf13/viper"
)

type Dao struct {
	db *mongo.DB
}

func NewDAO() (dao *Dao) {
	dao = &Dao{
		db: mongo.NewDb(viper.GetString("db.mongo.uri")),
	}

	if dao.db == nil {
		panic("mongo db connect failed")
	}
	return dao
}
