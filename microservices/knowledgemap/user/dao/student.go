package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/user/api"
	"knowledgemap_backend/microservices/knowledgemap/user/model"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) FillUserById(ctx context.Context, id bson.ObjectId, rsp **api.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	logrus.Infof("Id is %s", id.Hex())
	err = db.C(model.STUDENT_COLLECTION_NAME).FindId(id).One(*rsp)
	return
}

// func (d *Dao) FillUserById(ctx context.Context, id bson.ObjectId, user *model.Student) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	logrus.Infof("Id is %s", id.Hex())
// 	err = db.C(model.STUDENT_COLLECTION_NAME).FindId(id).One(user)
// 	return
// }

func (d *Dao) GetAllStudentInClass(ctx context.Context, cid string, rsp **api.ClassReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	logrus.Infof("Id is %s", cid)
	err = db.C(model.STUDENT_COLLECTION_NAME).Find(bson.M{"class": cid}).All(&(*rsp).Students)
	return
}
