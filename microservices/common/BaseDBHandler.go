package common

import "myProjects/collegeManage/library/database/mongo"

type (
	BaseDBHandler struct {
		DB *mongo.DB
	}
)
