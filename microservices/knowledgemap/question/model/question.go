package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	QUESTION_COLLECTION_NAME = "questions"
)

type QUESTION_KIND int

const (
	SIMPLE_CHOICE_QUESTION   QUESTION_KIND = 1 //"单选题"
	MULTIPLE_CHOICE_QUESTION QUESTION_KIND = 2 //"多选题"
	JUDGMENT_QUESTION        QUESTION_KIND = 3 //"判断题"
	BLACK_FILL_QUESTION      QUESTION_KIND = 4 //填空题"
	ESSAY_QUESTION           QUESTION_KIND = 5 //"简答题"
	PICTURE_QUESTION         QUESTION_KIND = 6 //"图片题"
	DOCUMENT_QUESTION        QUESTION_KIND = 7 //"文件题"

)

type KnowledgeInfo struct {
	ID   bson.ObjectId `json:"_id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

type Question struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Kind    QUESTION_KIND `json:"kind" bson:"kind"`
	Content string        `json:"content" bson:"content"`
	IsQImg  bool          `json:"isqimg" bson:"isqimg"`
	Options []Option      `json:"options" json:"options"`
	// OImage     bool          `json:"oimage" bson:"oimage"`
	Answers []Option `json:"answers" bson:"answers"` //选择题的话存数组下标
	// AImage     bool          `json:"aimage" bson:"aimage"`
	Subject     string        `json:"subject" bson:"subject"`
	Course      string        `json:"course" bson:"course"`
	Knowledge   bson.ObjectId `json:"knowledge" bson:"knowledge"`
	CreaterName string        `json:"createname" bson:"createname"`
	CreateTime  int64         `json:"createtime" bson:"createtime"`
	NeedCheck   bool          `json:"needcheck" bson:"needcheck"`
	Explain     string        `json:"explain" bson:"explain"`
	Star        int64         `json:"star" bson:"star"`
}

type Option struct {
	Prefix  string `json:"prefix" bson:"prefix"`
	Content string `json:"content" bson:"content"`
	IsImg   bool   `json:"isimg" bson:"isime"`
}
