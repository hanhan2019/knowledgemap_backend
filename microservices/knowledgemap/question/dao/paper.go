package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/question/model"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) NewPaper(ctx context.Context, paper *model.Paper, tableName string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	return col.Insert(paper)
}

func (d *Dao) FillPaperByClassId(ctx context.Context, tableName, classid string, paper *[]*model.Paper, page, pageCount int64) (err error, allCount int) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	cont := bson.M{
		"classid": classid,
	}
	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(paper)
	if err == nil {
		allCount, err = col.Find(cont).Count()
	}
	return
}

func (d *Dao) FillPaperByClassIdLimit(ctx context.Context, tableName, classid string, paper *[]*model.Paper, limit int) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	cont := bson.M{
		"classid": classid,
	}
	err = col.Find(cont).Sort("-_id").Limit(limit).All(paper)
	return
}

func (d *Dao) FillPaperById(ctx context.Context, tableName string, id bson.ObjectId, paper *model.Paper) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	return col.FindId(id).One(paper)
}
