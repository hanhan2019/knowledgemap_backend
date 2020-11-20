package model

import "gopkg.in/mgo.v2/bson"

const (
	CLASS_COLLECTION_NAME      = "class"
	CLASS_USER_COLLECTION_NAME = "class_user"
	GEN_ID_COLLECTION_NAME     = "genid"
)

type UserStatus string

const (
	Teacher UserStatus = "teacher"
	Student UserStatus = "student"
)

type Class struct {
	ID           bson.ObjectId `bson:"_id" json:"_id"`
	Number       string        `bson:"number" json:"number"`
	Name         string        `bson:"name" json:"name"`
	College      string        `bson:"college" json:"college"`
	Sbuject      string        `bson:"subject" json:"subject"`
	Course       string        `bson:"course" json:"course"`
	TeacherName  string        `bson:"teachername" json:"teachername"`
	TeacherId    string        `bson:"teacherid" json:"teacherid"`
	CreateTime   int64         `bson:"createtime" json:"createtime"`
	Introduction string        `bson:"introduction" json:"introduction"`
}

type ClassUser struct {
	ID         bson.ObjectId `bson:"_id" json:"_id"`
	UserId     string        `bson:"userid" json:"userid"`
	UserName   string        `bson:"username" json:"username"`
	Status     UserStatus    `bson:"status" json:"status"`
	ClassId    string        `bson:"classid" json:"classid"`
	CreateTime int64         `bson:"createtime" json:"createtime"`
}
