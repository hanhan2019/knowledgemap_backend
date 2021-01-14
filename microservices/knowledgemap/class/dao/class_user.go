package dao

import (
	"context"
	"errors"
	"knowledgemap_backend/microservices/knowledgemap/class/api"
	"knowledgemap_backend/microservices/knowledgemap/class/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) newClassUser(ctx context.Context, record *model.ClassUser) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.CLASS_USER_COLLECTION_NAME)
	return col.Insert(record)
}

func (d *Dao) NewClassUser(ctx context.Context, req *api.JoinClassReq) error {
	record := createClassUserRecord(req)
	return d.newClassUser(ctx, record)
}

func createClassUserRecord(req *api.JoinClassReq) *model.ClassUser {
	record := new(model.ClassUser)
	record.ID = bson.NewObjectId()
	record.ClassId = req.Classid
	record.UserId = req.Userid
	record.UserName = req.Username
	record.Status = model.UserStatus(req.Indentify)
	record.CreateTime = time.Now().Unix()
	return record
}

func (d *Dao) FillMyAllClass(ctx context.Context, userid string, classes *[]*model.ClassUser) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.CLASS_USER_COLLECTION_NAME)
	cont := bson.M{
		"userid": userid,
	}
	return col.Find(cont).All(classes)
}

func (d *Dao) CheckInClass(ctx context.Context, userId, classId string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.CLASS_USER_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"userid": userId, "classid": classId}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.inclass-duplicated")
	}
	return nil
}

func (d *Dao) FillAllStudentsById(ctx context.Context, classid string, classes *[]*model.ClassUser, page, pageCount int) (err error, allCount int) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.CLASS_USER_COLLECTION_NAME)
	cont := bson.M{
		"classid": classid,
		"status":  model.Student,
	}
	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(classes)
	if err == nil {
		allCount, err = col.Find(cont).Count()
	}
	return
}
func (d *Dao) FillStudentsByUserName(ctx context.Context, classid string, classes *[]*model.ClassUser, page, pageCount int, userName string) (err error, allCount int) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.CLASS_USER_COLLECTION_NAME)
	cont := bson.M{
		"classid":  classid,
		"status":   model.Student,
		"username": userName,
	}
	err = col.Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(page * pageCount)).All(classes)
	if err == nil {
		allCount, err = col.Find(cont).Count()
	}
	return
}

func (d *Dao) DeleteStudenInClass(ctx context.Context, classId, userId string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.CLASS_USER_COLLECTION_NAME)
	cont := bson.M{"classid": classId, "userid": userId}
	return col.Remove(cont)
}
