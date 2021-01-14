package model

import "gopkg.in/mgo.v2/bson"

const (
	HOMEWORK_TASK_COLLECTION_NAME = "homeworktask"
)

type MyHomeWork struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	UserId      bson.ObjectId `json:"userid" bson:"userid"`
	HomeworkId  bson.ObjectId `json:"homeworkid" bson:"homeworkid"`
	CreateTime  int64         `json:"createtime" bson:"createtime"`
	SuggestTime int64         `json:"suggesttime" bson:"suggesttime"`
	DoneTime    int64         `json:"donetime" bson:"donetime"`
}
