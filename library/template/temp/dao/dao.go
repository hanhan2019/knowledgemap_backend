package dao

import (
	"knowledgemap_backend/library/database/mongo"
)

type (
	Dao struct {
		mdb *mongo.DB
	}
)

func InitDao(mongodb *mongo.DB) *Dao {
	return &Dao{mdb: mongodb}
}
