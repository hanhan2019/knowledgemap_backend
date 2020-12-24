package model

import "gopkg.in/mgo.v2/bson"

const (
	PAPER_COLLECTION_NAME = "paper"
)

type QuestionScore struct {
	Questionid string `json:"questionid" bson:"questionid"`
	Score      int64  `json:"score" bson:"score"`
	NeedCheck  bool   `json:"needcheck" bson:"needcheck"`
}

type Paper struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	ClassId string        `json:"classid" bson:"classid"`
	// Students         []string      `json:"students" bson:"students"`
	Questions []QuestionScore `json:"questions" bson:"questions"`
	// CompleteStudents []string      `json:"completestudents" bson:"completestudents"`
	CreateTime     int64 `json:"createtime" bson:"createtime"`
	Totalscore     int64 `json:"totalscore" bson:"totalscore"`
	ContinuingTime int64 `json:"stoptime" bson:"stoptime"`
}
