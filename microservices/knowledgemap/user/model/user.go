package model

import "gopkg.in/mgo.v2/bson"

const (
	STUDENT_COLLECTION_NAME   = "student"
	TEACHER_COLLECTION_NAME   = "teacher"
	SECRETARY_COLLECTION_NAME = "secretary"
)

type People struct {
	Name     string `bson:"name" json:"name"`
	Number   string `bson:"number" json:"number"`
	Major    string `bson:"major" json:"major"`
	IDCard   string `bson:"idcard" json:"idcard"`
	Account  string `bson:"account" json:"account"`
	Password string `bson:"password" json:"password"`
}

type Student struct {
	ID            bson.ObjectId `bson:"_id" json:"_id"`
	People        `bson:",inline"`
	AdmissionTime string   `bson:"admissontime" json:"admissiontime"`
	Origin        string   `bson:"origin" json:"origin"`
	Courses       []string `bson:"courses" json:"courses"`
	Class         string   `bson:"class" json:"class"`
	College       string   `bson:"college" json:"college"`
	CreateTime    int64    `bson:"createtime" json:"createtime"`
}

type Teacher struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	People  `bson:",inline"`
	Courses []string `bson:"courses" json:"courses"`
	Classes []string `bson:"classes" json:"classes"`
}

type Secretary struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	People  `bson:",inline"`
	College string `bson:"college" json:"college"`
}

type Class struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	ClassID     string        `bson:"classid" json:"classid"`
	FromCollege string        `bson:"fromcollege" json:"fromcollege"`
	Students    []*Student    `bson:"students" json:"students"`
	Major       string        `bson:"major" json:"major"`
	Header      string        `bson:"header" json:"header"`
}
