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
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// func GetLoginTokenRedisKey(uid string) string {
// 	return fmt.Sprintf("passport:logintoken:%v", uid)
// }

// func GetLoginTokenRedisKey(uid string) string {
// 	return fmt.Sprintf("passport:logintoken:%v", uid)
// }

func GetLoginTokenRedisKey(uid string, identify int) string {
	return fmt.Sprintf("%v_%v", uid, identify)
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
	teacherId, _ := genIdCol.Inc(pmodel.GEN_TEACHER_ID_KEY_NAME)
	teacher.Number = fmt.Sprintf("%v", teacherId)
	return db.C(model.TEACHER_COLLECTION_NAME).Insert(teacher)
}

func (d *Dao) newSecretary(ctx context.Context, secretary *model.Secretary) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	genIdCol := mongo.NewNumberCollection(db.C(pmodel.GEN_ID_COLLECTION_NAME))
	//fill username
	secretaryId, _ := genIdCol.Inc(pmodel.GEN_SECRETARY_ID_KEY_NAME)
	secretary.Number = fmt.Sprintf("%v", secretaryId)
	return db.C(model.SECRETARY_COLLECTION_NAME).Insert(secretary)
}

func createDefaultStudent(req *api.RegisterReq) *model.Student {
	user := new(model.Student)
	user.ID = bson.NewObjectId()
	user.Name = req.Username
	user.Major = req.Major
	// user.IDCard = req.Idcard
	user.Account = req.Account
	user.Password = req.Password
	// user.Origin = req.Origin
	// user.Class = req.Class
	user.College = req.College
	user.Sex = req.Sex
	//user.AdmissionTime = req.Admissontime
	user.CreateTime = time.Now().Unix()
	return user
}

func createDefaultTeacher(req *api.RegisterReq) *model.Teacher {
	user := new(model.Teacher)
	user.ID = bson.NewObjectId()
	user.Name = req.Username
	user.Major = req.Major
	//user.IDCard = req.Idcard
	user.Account = req.Account
	user.Password = req.Password
	//user.Courses = []string{req.Course}
	user.College = req.College
	user.Sex = req.Sex
	user.CreateTime = time.Now().Unix()
	return user
}

func createDefaultSecretary(req *api.RegisterReq) *model.Secretary {
	user := new(model.Secretary)
	user.ID = bson.NewObjectId()
	user.Name = req.Username
	user.Major = req.Major
	//user.IDCard = req.Idcard
	user.Account = req.Account
	user.Password = req.Password
	user.College = req.College
	user.Sex = req.Sex
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

func (d *Dao) NewSecretary(ctx context.Context, req *api.RegisterReq) error {
	secreatary := createDefaultSecretary(req)
	return d.newSecretary(ctx, secreatary)
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

// func (d *Dao) GenerateLoginToken(ctx context.Context, rsp **api.PassportUserReply) error {
// 	userid := (*rsp).User.Userid
// 	token := utils.GenSession(userid)
// 	(*rsp).Token = token
// 	expires := 24 * time.Hour
// 	(*rsp).Expires = time.Now().Add(expires).Unix()

// 	set := d.redis.Set(GetLoginTokenRedisKey(userid), token, expires)
// 	return set.Err()

// }

func (d *Dao) GenerateLoginToken(ctx context.Context, identify int, rsp **api.PassportUserReply) error {
	userid := (*rsp).User.Userid
	token := utils.GenSession(userid)
	(*rsp).Token = token
	expires := 24 * 365 * time.Hour
	(*rsp).Expires = time.Now().Add(expires).Unix()

	set := d.redis.Set(token, GetLoginTokenRedisKey(userid, identify), expires)
	return set.Err()
}

// func (d *Dao) CheckSessionToken(ctx context.Context, uid, token string) error {
// 	res := d.redis.Get(GetLoginTokenRedisKey(uid))
// 	if savedToken, err := res.Result(); err != nil {
// 		return err
// 	} else {
// 		if savedToken != token {
// 			fmt.Println("token 不正确", savedToken, token)
// 			return pmodel.ErrorSessionTokenNotValidate
// 		}
// 	}
// 	return nil
// }

func (d *Dao) CheckSessionToken(ctx context.Context, token string, user *uapi.UserReply) error {
	res := d.redis.Get(token)
	if saved, err := res.Result(); err != nil {
		return err
	} else {
		data := strings.Split(saved, "_")
		uid := data[0]
		identify, _ := strconv.Atoi(data[1])

		// if savedToken != token {
		// 	fmt.Println("token 不正确", savedToken, token)
		// 	return pmodel.ErrorSessionTokenNotValidate
		// }
		switch uapi.Identify(identify) {
		case uapi.Identify_STUDENT:
			d.FillStudentById(ctx, bson.ObjectIdHex(uid), &user)
			user.Usertype = int64(uapi.Identify_STUDENT)
		case uapi.Identify_TEACHER:
			d.FillTeacherById(ctx, bson.ObjectIdHex(uid), &user)
			user.Usertype = int64(uapi.Identify_TEACHER)
		case uapi.Identify_SECRETARY:
			d.FillSecretaryById(ctx, bson.ObjectIdHex(uid), &user)
			user.Usertype = int64(uapi.Identify_SECRETARY)
		}
		if user.Userid == "" {
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

func (d *Dao) ChangePassword(ctx context.Context, userid, password string, identify uapi.Identify, rsp **uapi.Empty) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if *rsp == nil {
		*rsp = &uapi.Empty{}
	}
	dbName := ""
	switch uapi.Identify(identify) {
	case uapi.Identify_STUDENT:
		dbName = model.STUDENT_COLLECTION_NAME
	case uapi.Identify_TEACHER:
		dbName = model.TEACHER_COLLECTION_NAME
	case uapi.Identify_SECRETARY:
		dbName = model.SECRETARY_COLLECTION_NAME
	default:
		return errors.New("errors.wrong_type")
	}
	err = db.C(dbName).UpdateId(bson.ObjectIdHex(userid), bson.M{
		"$set": bson.M{
			"password": password,
		},
	})
	return
}

func (d *Dao) ChangeUserInfo(ctx context.Context, req *api.ChangeUserInfoReq) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	dbName := ""
	switch req.Usertype {
	case uapi.Identify_STUDENT:
		dbName = model.STUDENT_COLLECTION_NAME
	case uapi.Identify_TEACHER:
		dbName = model.TEACHER_COLLECTION_NAME
	case uapi.Identify_SECRETARY:
		dbName = model.SECRETARY_COLLECTION_NAME
	default:
		return errors.New("errors.wrong_type")
	}
	err = db.C(dbName).UpdateId(bson.ObjectIdHex(req.Userid), bson.M{
		"$set": bson.M{
			// "password": req.Password,
			"major":   req.Major,
			"college": req.College,
			"sex":     req.Sex,
			"name":    req.Username,
		},
	})
	return
}
