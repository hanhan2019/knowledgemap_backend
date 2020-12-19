package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	QUESTION_COLLECTION_NAME = "questions"
)

type Qusetion_Kind int

const (
	Choice_Question Qusetion_Kind = 1
	Essay_Question  Qusetion_Kind = 2
)

type KnowledgeInfo struct {
	ID   bson.ObjectId `json:"_id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

type Qusetion struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Kind       Qusetion_Kind `json:"kind" bson:"kind"`
	Content    string        `json:"content" bson:"content"`
	Option     []string      `json:"option" json:"option"`
	Answer     []string      `json:"answer" bson:"answer"`
	Subject    string        `json:"subject" bson:"subject"`
	Course     string        `json:"course" bson:"course"`
	Knowledge  bson.ObjectId `json:"knowledge" bson:"knowledge"`
	CreateTime int64         `json:"createtime" bson:"createtime"`
}
