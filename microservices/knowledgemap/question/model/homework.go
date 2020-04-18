package model

import "gopkg.in/mgo.v2/bson"

const (
	HOMEWORK_COLLECTION_NAME = "homework"
)

type HomeWork struct {
	ID               bson.ObjectId `json:"_id" bson:"_id"`
	Name             string        `json:"name" bson:"name"`
	ClassId          string        `json:"classid" bson:"classid"`
	Students         []string      `json:"students" bson:"students"`
	Questions        []string      `json:"questions" bson:"questions"`
	CompleteStudents []string      `json:"completestudents" bson:"completestudents"`
	CreateTime       int64         `json:"createtime" bson:"createtime"`
	StopTime         int64         `json:"stoptime" bson:"stoptime"`
}
