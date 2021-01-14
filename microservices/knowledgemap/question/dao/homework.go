package dao

// func (d *Dao) NewHomeWork(ctx context.Context, homwWork *model.HomeWork) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.HOMEWORK_COLLECTION_NAME)
// 	return col.Insert(homwWork)
// }

// func (d *Dao) FillHomeWorkByStudentId(ctx context.Context, userId string, homwWork *[]*model.HomeWork) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.HOMEWORK_COLLECTION_NAME)
// 	cont := bson.M{
// 		"students": bson.M{"$in": []string{userId}},
// 	}
// 	return col.Find(cont).All(homwWork)
// }

// func (d *Dao) FillHomeWorkById(ctx context.Context, id bson.ObjectId, homwWork *model.HomeWork) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.HOMEWORK_COLLECTION_NAME)
// 	return col.FindId(id).One(homwWork)
// }

// func (d *Dao) UpdateCompeleStudent(ctx context.Context, id bson.ObjectId, completeStudents []string) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.HOMEWORK_COLLECTION_NAME)
// 	return col.UpdateId(id, bson.M{"$set": bson.M{"completestudents": completeStudents}})
// }

// func (d *Dao) FillHomeWorkByClassId(ctx context.Context, classid string, homwWork *[]*model.HomeWork) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.HOMEWORK_COLLECTION_NAME)
// 	cont := bson.M{
// 		"classid": classid,
// 	}
// 	return col.Find(cont).All(homwWork)
// }
