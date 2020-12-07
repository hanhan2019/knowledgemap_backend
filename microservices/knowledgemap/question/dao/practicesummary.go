package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/question/model"

	"gopkg.in/mgo.v2/bson"
)

// func (d *Dao) QueryQuestionInfo(ctx context.Context, questionId bson.ObjectId, info *model.Qusetion) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	err = db.C(model.QUESTION_COLLECTION_NAME).FindId(questionId).One(info)
// 	return
// }

func (d *Dao) NewPracticeSummary(ctx context.Context, ps *model.PracticeSummary) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.PRACTICE_COLLECTION_NAME)
	return col.Insert(ps)
}

func (d *Dao) FillPSByID(ctx context.Context, id bson.ObjectId, ps **model.PracticeSummary) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	err = db.C(model.PRACTICE_COLLECTION_NAME).FindId(id).One(*ps)
	return
}

func (d *Dao) FillPSByUserID(ctx context.Context, id string, page int64, ps *[]*model.PracticeSummary, pageCount int64) (allCount int, err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	err = db.C(model.PRACTICE_COLLECTION_NAME).Find(bson.M{"ownerid": id}).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(ps)
	allCount, err = db.C(model.PRACTICE_COLLECTION_NAME).Find(bson.M{"ownerid": id}).Count()
	return
}

func (d *Dao) QueryPSQuestionCount(ctx context.Context, psId bson.ObjectId) (error, int64) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.PRACTICE_QUESTION_COLLECTION_NAME)
	cnt, err := col.Find(bson.M{"psid": psId}).Count()
	if err != nil {
		return err, 0
	}
	return nil, int64(cnt)
}

func (d *Dao) FillQuestionsByPSId(ctx context.Context, psId bson.ObjectId, page int64, ps *[]*model.PracticeQuestion, pageCount int64) (allCount int, err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	cont := bson.M{"psid": psId}
	err = db.C(model.PRACTICE_QUESTION_COLLECTION_NAME).Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(ps)
	allCount, err = db.C(model.PRACTICE_QUESTION_COLLECTION_NAME).Find(cont).Count()
	return
}

func (d *Dao) NewPracticeQuestion(ctx context.Context, ps *model.PracticeQuestion) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.PRACTICE_QUESTION_COLLECTION_NAME)
	return col.Insert(ps)
}

func (d *Dao) DeletePracticeQuestion(ctx context.Context, psId, questionId bson.ObjectId) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.PRACTICE_QUESTION_COLLECTION_NAME)
	cont := bson.M{"psid": psId, "questionid": questionId}
	return col.Remove(cont)
}
