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

func (d *Dao) FillAnserRecordById(ctx context.Context, paperKind, Id string, userid bson.ObjectId, records *[]*model.AnswerRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.ANSWER_RECORD_COLLECTION_NAME)
	cont := bson.M{
		"paperid":   Id,
		"paperkind": paperKind,
		// "question": questionId,
		"userid": userid,
	}
	return col.Find(cont).All(records)
}

func (d *Dao) FillAnserRecordByPaperRecordId(ctx context.Context, paperKind string, Id, userid bson.ObjectId, records *[]*model.AnswerRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.ANSWER_RECORD_COLLECTION_NAME)
	cont := bson.M{
		"paperrecordid": Id,
		"paperkind":     paperKind,
		"userid":        userid,
	}
	return col.Find(cont).All(records)
}

func (d *Dao) FillSimpleAnserRecordInPage(ctx context.Context, paperKind string, Id, userid bson.ObjectId, pageCount, page int, records *[]*model.AnswerRecord) (err error, allCount int) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.ANSWER_RECORD_COLLECTION_NAME)
	cont := bson.M{
		"paperrecordid": Id,
		"paperkind":     paperKind,
		"userid":        userid,
	}
	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(records)
	if err == nil {
		allCount, err = col.Find(cont).Count()
	}
	return
}

func (d *Dao) FillAnserRecordBySomeIds(ctx context.Context, paperKind string, Id, userid, questionId bson.ObjectId, record *model.AnswerRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.ANSWER_RECORD_COLLECTION_NAME)
	cont := bson.M{
		"paperrecordid": Id,
		"paperkind":     paperKind,
		"userid":        userid,
		"question":      questionId,
	}
	return col.Find(cont).One(record)
}

func (d *Dao) FillAnserRecord(ctx context.Context, Id bson.ObjectId, record *model.AnswerRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.ANSWER_RECORD_COLLECTION_NAME)
	return col.FindId(Id).One(record)
}
