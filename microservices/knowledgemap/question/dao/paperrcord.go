package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/question/model"

	"gopkg.in/mgo.v2/bson"
)

// func (d *Dao) QueryUserAnswerRecords(ctx context.Context, userId bson.ObjectId, subject string, endTime int64, records *[]*model.AnswerRecord) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	err = db.C(model.ANSWER_RECORD_COLLECTION_NAME).Find(bson.M{"userid": userId, "subject": subject, "createtime": bson.M{"$lte": endTime}}).All(records)
// 	return
// }

func (d *Dao) NewPaperRecord(ctx context.Context, tableName string, paperRecord *model.PaperRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	return col.Insert(paperRecord)
}

func (d *Dao) FillPaperRecordList(ctx context.Context, userid bson.ObjectId, tableName string, records *[]*model.PaperRecord, page, pageCount int64) (err error, allCount int) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	cont := bson.M{
		// "question": questionId,
		"userid": userid,
	}
	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(records)
	if err == nil {
		allCount, err = col.Find(cont).Count()
	}
	return
}

func (d *Dao) FillPaperRecordByID(ctx context.Context, id bson.ObjectId, tableName string, record *model.PaperRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	return col.FindId(id).One(record)
}

func (d *Dao) FillPaperRecordByUserId(ctx context.Context, tableName string, userId bson.ObjectId, count int, record *[]*model.PaperRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	cont := bson.M{
		"userid": userId,
	}
	return col.Find(cont).Sort("-createtime").Limit(count).All(record)
}

func (d *Dao) UpdatePaperRecord(ctx context.Context, tableName string, id bson.ObjectId, status model.RecordStatus, doneTime int64) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	return col.UpdateId(id, bson.M{
		"$set": bson.M{
			"status":   status,
			"donetime": doneTime,
		},
	})
}

func (d *Dao) FillPaperRecordByUserIDAndPaperId(ctx context.Context, userId, paperId bson.ObjectId, tableName string, record *model.PaperRecord) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	return col.Find(bson.M{"userid": userId, "paperid": paperId}).Sort("-creattime").One(record)
}

func (d *Dao) RemovePaperRecord(ctx context.Context, userId, paperId bson.ObjectId, tableName string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(tableName)
	return col.Remove(bson.M{"userid": userId, "paperid": paperId})
}
