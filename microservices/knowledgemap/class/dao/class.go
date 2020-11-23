package dao

import (
	"context"
	"errors"
	"fmt"
	"knowledgemap_backend/library/database/mongo"
	"knowledgemap_backend/microservices/knowledgemap/class/api"
	"knowledgemap_backend/microservices/knowledgemap/class/model"
	pmodel "knowledgemap_backend/microservices/knowledgemap/passport/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) newClass(ctx context.Context, class *model.Class) (*model.Class, error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	genIdCol := mongo.NewNumberCollection(db.C(pmodel.GEN_ID_COLLECTION_NAME))
	classId, _ := genIdCol.Inc(pmodel.GEN_CLASS_ID_KEY_NAME)
	class.Number = fmt.Sprintf("%v", classId)
	col := db.C(model.CLASS_COLLECTION_NAME)
	return class, col.Insert(class)
}

func (d *Dao) NewClass(ctx context.Context, req *api.CreateClassReq) (*model.Class, error) {
	class := createDefaultClass(req)
	return d.newClass(ctx, class)
}

func createDefaultClass(req *api.CreateClassReq) *model.Class {
	class := new(model.Class)
	class.ID = bson.NewObjectId()
	class.Name = req.Name
	class.Course = req.Course
	class.College = req.College
	class.Sbuject = req.Subject
	class.TeacherName = req.Teachername
	class.TeacherId = req.Teacherid
	class.Introduction = req.Introduction
	class.CreateTime = time.Now().Unix()
	return class
}

func (d *Dao) CheckNewClass(ctx context.Context, name, subject, college, teacherid, course string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.CLASS_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"name": name, "subject": subject, "teacherid": teacherid, "college": college, "course": course}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.class-duplicated")
	}
	return nil
}

func (d *Dao) FillClassByID(ctx context.Context, id bson.ObjectId, rsp **api.ClassReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &api.ClassReply{}
	}
	err = db.C(model.CLASS_COLLECTION_NAME).FindId(id).One(*rsp)
	if err == nil {
		(*rsp).Classid = bson.ObjectId((*rsp).Classid).Hex()
	}
	return
}

func (d *Dao) FillClassByConditions(ctx context.Context, req *api.SearchClassesInfoReq, classes *[]*model.Class, pageCount int64) (allCount int, err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	cont := bson.M{}
	if req.College != "" {
		cont["college"] = req.College
	}
	if req.Subject != "" {
		cont["subject"] = req.Subject
	}
	if req.Course != "" {
		cont["course"] = req.Course
	}
	if req.Subject != "" {
		cont["teachername"] = req.Teachername
	}
	// fmt.Println(cont)
	// fmt.Println(req.Subject = )

	err = db.C(model.CLASS_COLLECTION_NAME).Find(cont).Sort("-_id").Limit(int(pageCount)).Skip(int(req.Page * pageCount)).All(classes)
	allCount, err = db.C(model.CLASS_COLLECTION_NAME).Find(cont).Count()
	return
}
