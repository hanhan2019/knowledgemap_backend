package model

import "gopkg.in/mgo.v2/bson"

const ()

type HomeWorkRecordStatus int

const (
	HOMEWORK_DONE           HomeWorkRecordStatus = 1 //"完成"
	HOMEWORK_WAITTING_DONE  HomeWorkRecordStatus = 2 //"待完成"
	HOMEWORK_WAITTING_CHECK HomeWorkRecordStatus = 3 //"待批改"
)

type HomeWorkRecord struct {
	ID           bson.ObjectId        `json:"_id" bson:"_id"`
	HomeWorkName string               `json:"homeworkname" bson:"homeworkname"`
	HomeWorkId   string               `json:"homeworkid" bson:"homeworkid"`
	StudentId    string               `json:"studentid" bson:"studentid"`
	Status       HomeWorkRecordStatus `json:"status" bson:"status"`
	CreateTime   int64                `json:"createtime" bson:"createtime"`
	DoneTime     int64                `json:"donetime" bson:"donetime"`
}
