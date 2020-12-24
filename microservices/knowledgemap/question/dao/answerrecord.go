package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/question/model"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) QueryUserAnswerRecords(ctx context.Context, userId bson.ObjectId, subject string, endTime int64, records *[]*model.AnswerRecord) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	err = db.C(model.ANSWER_RECORD_COLLECTION_NAME).Find(bson.M{"userid": userId, "subject": subject, "createtime": bson.M{"$lte": endTime}}).All(records)
	return
}

func (d *Dao) NewAnswerRecord(ctx context.Context, answerRecord *model.AnswerRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.ANSWER_RECORD_COLLECTION_NAME)
	return col.Insert(answerRecord)
}

func (d *Dao) FillAnserRecordByIdAndQuestionId(ctx context.Context, Idname, Id string, questionId bson.ObjectId, records *[]*model.AnswerRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.ANSWER_RECORD_COLLECTION_NAME)
	cont := bson.M{
		Idname:     Id,
		"question": questionId,
	}
	return col.Find(cont).All(records)
}
