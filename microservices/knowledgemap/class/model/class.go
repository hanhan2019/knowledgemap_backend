package model

import "gopkg.in/mgo.v2/bson"

const (
	CLASS_COLLECTION_NAME      = "class"
	CLASS_USER_COLLECTION_NAME = "class_user"
	GEN_ID_COLLECTION_NAME     = "genid"
)

type Class struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Number      string        `bson:"number" json:"number"`
	Name        string        `bson:"name" json:"name"`
	College     string        `bson:"college" json:"college"`
	Major       string        `bson:"major" json:"major"`
	TeacherName string        `bson:"teachername" json:"teachername"`
	CreateTime  int64         `bson:"createtime" json:"createtime"`
}

type ClassUser struct {
	ID         bson.ObjectId `bson:"_id" json:"_id"`
	UserId     string        `bson:"userid" json:"userid"`
	ClassId    string        `bson:"classid" json:"classid"`
	CreateTime int64         `bson:"createtime" json:"createtime"`
}
