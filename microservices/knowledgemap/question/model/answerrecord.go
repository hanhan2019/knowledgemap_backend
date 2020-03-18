package model

import "gopkg.in/mgo.v2/bson"

const (
	ANSWER_RECORD_COLLECTION_NAME = "answer_record"
)

type AnswerRecord struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	UserID     bson.ObjectId `json:"userid" bson:"userid"`
	QuestionID bson.ObjectId `json:"question" bson:"question"`
	Answer     string        `json:"answer" bson:"answer"`
	IsTrue     bool          `json:"istrue" bson:"istrue"`
	// Knowledge  string        `json:"knowledge" bson:"knowledge"`
	Subject    string `json:"subject" bson:"subject"`
	CreateTime int64  `json:"createtime" bson:"createtime"`
}
