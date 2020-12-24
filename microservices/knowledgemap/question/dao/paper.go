package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/question/model"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) NewPaper(ctx context.Context, paper *model.Paper) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.PAPER_COLLECTION_NAME)
	return col.Insert(paper)
}

func (d *Dao) FillPaperByClassId(ctx context.Context, classid string, paper *[]*model.Paper, page, pageCount int64) (err error, allCount int) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.PAPER_COLLECTION_NAME)
	cont := bson.M{
		"classid": classid,
	}
	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(paper)
	if err == nil {
		allCount, err = col.Find(cont).Count()
	}
	return
}

func (d *Dao) FillPaperById(ctx context.Context, id bson.ObjectId, paper *model.Paper) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.PAPER_COLLECTION_NAME)
	return col.FindId(id).One(paper)
}
