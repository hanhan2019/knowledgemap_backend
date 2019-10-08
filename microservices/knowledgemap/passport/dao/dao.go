package dao

import (
	"knowledgemap_backend/library/database/mongo"

	"github.com/go-redis/redis"
)

type (
	Dao struct {
		mdb   *mongo.DB
		redis *redis.Client
	}
)

func InitDao(mongodb *mongo.DB, r *redis.Client) *Dao {
	return &Dao{mdb: mongodb, redis: r}
}
