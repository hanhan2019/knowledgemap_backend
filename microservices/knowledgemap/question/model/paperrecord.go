package model

import "gopkg.in/mgo.v2/bson"

const (
	EXAM_RECORD_COLLECTION_NAME     = "exam_record"
	HOMEWORK_RECORD_COLLECTION_NAME = "homework_record"
)

type RecordStatus int64

const (
	RECORD_DONE           RecordStatus = 1 //"完成"
	RECORD_WAITTING_CHECK RecordStatus = 2 //"待批改"
	RECORD_WAITTING_DONE  RecordStatus = 3 //"待完成"
)

type PaperRecord struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	UserID   bson.ObjectId `json:"userid" bson:"userid"`
	UserName string        `json:"username" bson:"username"`
	Status   RecordStatus  `json:"status" bson:"status"`
	// TypeId   string        `json:"typeid" bson:"typeid"`
	PaperId    string `json:"paperid" bson:"paperid"`
	Score      int64  `json:"score" bson:"socre"`
	CreateTime int64  `json:"createtime" bson:"createtime"`
	DoneTime   int64  `json:"donetime" bson:"donetime"`
}
