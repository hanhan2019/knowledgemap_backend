package dao

import (
	"context"
	"errors"
	"fmt"
	"knowledgemap_backend/library/database/mongo"
	"knowledgemap_backend/microservices/knowledgemap/passport/api"
	pmodel "knowledgemap_backend/microservices/knowledgemap/passport/model"
	"knowledgemap_backend/microservices/knowledgemap/passport/utils"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"knowledgemap_backend/microservices/knowledgemap/user/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func GetLoginTokenRedisKey(uid string) string {
	return fmt.Sprintf("passport:logintoken:%v", uid)
}

func (d *Dao) CheckIDCardInStudent(ctx context.Context, idCard string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.STUDENT_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"idcard": idCard}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.idcard-duplicated")
	}
	return nil
}

func (d *Dao) CheckAccountInStudent(ctx context.Context, account string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.STUDENT_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"account": account}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.idcard-duplicated")
	}
	return nil
}

func (d *Dao) CheckIDCardInTeacher(ctx context.Context, idCard string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.TEACHER_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"idcard": idCard}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.idcard-duplicated")
	}
	return nil
}
func (d *Dao) CheckAccountInTeacher(ctx context.Context, account string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.TEACHER_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"account": account}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.idcard-duplicated")
	}
	return nil
}

func (d *Dao) CheckIDCardInSecretary(ctx context.Context, idCard string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.SECRETARY_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"idcard": idCard}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.idcard-duplicated")
	}
	return nil
}

func (d *Dao) CheckAccountInSecretary(ctx context.Context, account string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.SECRETARY_COLLECTION_NAME)
	if cnt, err := col.Find(bson.M{"account": account}).Count(); err != nil {
		return err
	} else if cnt > 0 {
		return errors.New("errors.idcard-duplicated")
	}
	return nil
}

func (d *Dao) newStudent(ctx context.Context, student *model.Student) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	genIdCol := mongo.NewNumberCollection(db.C(pmodel.GEN_ID_COLLECTION_NAME))
	//fill username
	studentId, _ := genIdCol.Inc(pmodel.GEN_STUDENT_ID_KEY_NAME)
	student.Number = fmt.Sprintf("%v", studentId)
	col := db.C(model.STUDENT_COLLECTION_NAME)
	return col.Insert(student)
}

func (d *Dao) newTeacher(ctx context.Context, teacher *model.Teacher) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	genIdCol := mongo.NewNumberCollection(db.C(pmodel.GEN_ID_COLLECTION_NAME))
	//fill username
	teacherId, _ := genIdCol.Inc(pmodel.GEN_STUDENT_ID_KEY_NAME)
	teacher.Number = fmt.Sprintf("%v", teacherId)
	return db.C(model.TEACHER_COLLECTION_NAME).Insert(teacher)
}

func createDefaultStudent(req *api.RegisterReq) *model.Student {
	user := new(model.Student)
	user.ID = bson.NewObjectId()
	user.Name = req.Name
	user.Major = req.Major
	user.IDCard = req.Idcard
	user.Account = req.Account
	user.Password = req.Password
	user.Origin = req.Origin
	user.Class = req.Class
	user.College = req.College
	user.AdmissionTime = req.Admissontime
	user.CreateTime = time.Now().Unix()
	return user
}

func createDefaultTeacher(req *api.RegisterReq) *model.Teacher {
	user := new(model.Teacher)
	user.ID = bson.NewObjectId()
	user.Name = req.Name
	user.Major = req.Major
	user.IDCard = req.Idcard
	user.Account = req.Account
	user.Password = req.Password
	user.Courses = []string{req.Course}
	user.CreateTime = time.Now().Unix()
	return user
}

func (d *Dao) NewStudent(ctx context.Context, req *api.RegisterReq) error {
	student := createDefaultStudent(req)
	return d.newStudent(ctx, student)
}

func (d *Dao) NewTeacher(ctx context.Context, req *api.RegisterReq) error {
	teacher := createDefaultTeacher(req)
	return d.newTeacher(ctx, teacher)
}
func (d *Dao) FillUserByIDCardInStudent(ctx context.Context, idCard string, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.STUDENT_COLLECTION_NAME).Find(bson.M{"idcard": idCard}).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}

func (d *Dao) FillUserByIDCardInTeacher(ctx context.Context, idCard string, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.TEACHER_COLLECTION_NAME).Find(bson.M{"idcard": idCard}).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}

func (d *Dao) GenerateLoginToken(ctx context.Context, rsp **api.PassportUserReply) error {
	userid := (*rsp).User.Userid
	token := utils.GenSession(userid)
	(*rsp).Token = token
	expires := 24 * time.Hour
	(*rsp).Expires = time.Now().Add(expires).Unix()

	set := d.redis.Set(GetLoginTokenRedisKey(userid), token, expires)
	return set.Err()

}

func (d *Dao) CheckSessionToken(ctx context.Context, uid, token string) error {
	res := d.redis.Get(GetLoginTokenRedisKey(uid))
	if savedToken, err := res.Result(); err != nil {
		return err
	} else {
		if savedToken != token {
			fmt.Println("token 不正确")
			return pmodel.ErrorSessionTokenNotValidate
		}
	}
	return nil
}

func (d *Dao) FillStudentByAccount(ctx context.Context, account string, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.STUDENT_COLLECTION_NAME).Find(bson.M{"account": account}).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}

func (d *Dao) FillTeacherByAccount(ctx context.Context, account string, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.TEACHER_COLLECTION_NAME).Find(bson.M{"account": account}).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}
func (d *Dao) FillSecretaryByAccount(ctx context.Context, account string, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.SECRETARY_COLLECTION_NAME).Find(bson.M{"account": account}).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}

func (d *Dao) FillStudentById(ctx context.Context, id bson.ObjectId, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.STUDENT_COLLECTION_NAME).FindId(id).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}

func (d *Dao) FillTeacherById(ctx context.Context, id bson.ObjectId, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.TEACHER_COLLECTION_NAME).FindId(id).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}
func (d *Dao) FillSecretaryById(ctx context.Context, id bson.ObjectId, rsp **uapi.UserReply) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.UserReply{}
	}
	err = db.C(model.SECRETARY_COLLECTION_NAME).FindId(id).One(*rsp)
	if err == nil {
		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
	}
	return
}

func (d *Dao) ChangePassword(ctx context.Context, account, password string, rsp **uapi.Empty) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.Empty{}
	}
	err = db.C(model.STUDENT_COLLECTION_NAME).Update(bson.M{"account": account}, bson.M{
		"$set": bson.M{
			"password": password,
		},
	})
	return
}
