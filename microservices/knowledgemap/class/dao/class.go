package dao

import (
	"context"
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
	class.Major = req.Major
	class.College = req.College
	class.TeacherName = req.Teachername
	class.CreateTime = time.Now().Unix()
	return class
}

func (d *Dao) FillClassByID(ctx context.Context, id bson.ObjectId, rsp **api.ClassReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &api.ClassReply{}
	}
	err = db.C(model.CLASS_COLLECTION_NAME).FindId(id).One(*rsp)
	return
}
