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
