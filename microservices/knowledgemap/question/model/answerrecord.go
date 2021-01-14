package model

import "gopkg.in/mgo.v2/bson"

const (
	ANSWER_RECORD_COLLECTION_NAME = "answer_record"
)

type Result int64

const (
	RIGHT         Result = 1 // "正确"
	WRONG         Result = 2 //"错误"
	NOTBEST       Result = 3 //"非最佳答案"
	WAITTINGCHECK Result = 4 //"等待批改"
)

type AnswerRecord struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	UserID     bson.ObjectId `json:"userid" bson:"userid"`
	UserName   string        `json:"username" bson:"username"`
	QuestionID bson.ObjectId `json:"question" bson:"question"`
	Answer     []string      `json:"answer" bson:"answer"`
	Result     Result        `json:"result" bson:"result"`
	// Knowledge  string        `json:"knowledge" bson:"knowledge"`
	Subject   string `json:"subject" bson:"subject"`
	PaperKind string `json:"paperkind" bson:"paperkind"` //哪个类型的试卷
	// HomeWorkId string `json:"homeworkid" bson:"homeworkid"`
	PaperRecordID bson.ObjectId `json:"paperrecordid" bson:"paperrecordid"` //答卷记录id
	PaperId       string        `json:"paperid" bson:"paperid"`             //试卷id是多少
	Score         int64         `json:"score" bson:"socre"`
	CreateTime    int64         `json:"createtime" bson:"createtime"`
}
