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

func (d *Dao) FillQuestionBySubject(ctx context.Context, kind int64, subject, course string, knowledge bson.ObjectId, questions *[]*model.Qusetion) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.QUESTION_COLLECTION_NAME)
	cont := bson.M{
		"kind":      kind,
		"subject":   subject,
		"course":    course,
		"knowledge": knowledge,
	}
	return col.Find(cont).All(questions)
}

func (d *Dao) FillQuestionById(ctx context.Context, id bson.ObjectId, question *model.Qusetion) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.QUESTION_COLLECTION_NAME)
	return col.FindId(id).One(question)
}
