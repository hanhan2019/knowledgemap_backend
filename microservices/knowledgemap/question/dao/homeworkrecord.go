package dao

// func (d *Dao) NewHomeWorkRecord(ctx context.Context, record *model.HomeWorkRecord) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.HOMEWORK_RECORD_COLLECTION_NAME)
// 	return col.Insert(record)
// }

// func (d *Dao) FillHomeWorkRecordList(ctx context.Context, userid bson.ObjectId, records *[]*model.HomeWorkRecord, page, pageCount int64) (err error, allCount int) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.HOMEWORK_RECORD_COLLECTION_NAME)
// 	cont := bson.M{
// 		// "question": questionId,
// 		"studentid": userid,
// 	}
// 	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(records)
// 	if err == nil {
// 		allCount, err = col.Find(cont).Count()
// 	}
// 	return
// }

// func (d *Dao) UpdateHomeWorkRecord(ctx context.Context, id bson.ObjectId, paper *model.Paper) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.PAPER_COLLECTION_NAME)
// 	return col.FindId(id).One(paper)
// }
