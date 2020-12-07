package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	PRACTICE_COLLECTION_NAME          = "practice_summary"
	PRACTICE_QUESTION_COLLECTION_NAME = "practice_question"
)

type PSTYPE int

const (
	DEFAULT    PSTYPE = 1
	COLLECTION PSTYPE = 2
	MISTAKES   PSTYPE = 3
)

type QuestionInfo struct {
	ID           bson.ObjectId `json:"_id" bson:"_id"`
	KnowlegeName string        `json:"knowlegename" bson:"knowlegename"`
}

type PracticeSummary struct {
	ID           bson.ObjectId `json:"_id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Pstype       PSTYPE        `json:"pstype" bson:"pstype"`
	OwnerId      string        `json:"owmerid" bson:"ownerid"`
	OwnerName    string        `json:"ownername" bson:"ownername"`
	ClassId      string        `json:"classid" json:"classid"`
	CourseName   string        `json:"coursename" json:"coursename"`
	Introduction string        `json:"introduction" json:"introdution"`
	//Questions    []QuestionInfo `json:"questions" bson:"questions"`
	CreateTime int64 `json:"createtime" bson:"createtime"`
}

type PracticeQuestion struct {
	ID           bson.ObjectId `json:"_id" bson:"_id"`
	PSId         bson.ObjectId `json:"psid" bson:"psid"`
	QuestionId   bson.ObjectId `json:"questionid" bson:"questionid"`
	QuestionKind Qusetion_Kind `json:"questionkind" bson:"questionkind"`
	KnowlegeId   bson.ObjectId `json:"knowlegeid" bson:"knowlegeid"`
	KnowlegeName string        `json:"knowlegename" bson:"knowlegename"`
	CreateTime   int64         `json:"createtime" bson:"createtime"`
}
