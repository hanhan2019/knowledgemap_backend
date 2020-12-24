package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/question/model"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) QueryQuestionInfo(ctx context.Context, questionId bson.ObjectId, info *model.Qusetion) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	err = db.C(model.QUESTION_COLLECTION_NAME).FindId(questionId).One(info)
	return
}

func (d *Dao) NewQuestion(ctx context.Context, question *model.Qusetion) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.QUESTION_COLLECTION_NAME)
	return col.Insert(question)
}

func (d *Dao) FillQuestionBySubject(ctx context.Context, kind int64, subject, course string, knowledge bson.ObjectId, questions *[]*model.Qusetion, page, pageCount int64) (err error, allCount int) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.QUESTION_COLLECTION_NAME)
	cont := bson.M{
		"kind":      kind,
		"subject":   subject,
		"course":    course,
		"knowledge": knowledge,
	}
	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(questions)
	if err == nil {
		allCount, err = col.Find(cont).Count()
	}
	return
}

func (d *Dao) FillQuestionById(ctx context.Context, id bson.ObjectId, question *model.Qusetion) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.QUESTION_COLLECTION_NAME)
	return col.FindId(id).One(question)
}

func (d *Dao) QueryMulQuestionInfo(ctx context.Context, questionId []bson.ObjectId, info *[]*model.Qusetion) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	cont := bson.M{
		"_id": bson.M{"$in": questionId},
	}
	err = db.C(model.QUESTION_COLLECTION_NAME).Find(cont).All(info)
	return
}
