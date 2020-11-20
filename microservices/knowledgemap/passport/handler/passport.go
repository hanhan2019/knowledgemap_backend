package handler

import (
	"context"
	"errors"
	"fmt"
	"knowledgemap_backend/microservices/knowledgemap/passport/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"

	"gopkg.in/mgo.v2/bson"

	"github.com/sirupsen/logrus"
)

type PassportService struct{}

func (s *PassportService) Register(ctx context.Context, req *api.RegisterReq, rsp *api.PassportUserReply) error {
	logrus.Infof("Register req is %v ", req)
	if err := CheckAccount(ctx, req.Account); err != nil {
		fmt.Printf("CheckAccount error %v ", err)
		return errors.New("CheckAccount error")
	}
	switch req.Rtype {
	case api.Indentify_STUDENT:
		gdao.NewStudent(ctx, req)
		if err := gdao.FillStudentByAccount(ctx, req.Account, &rsp.User); err != nil {
			return err
		}
		return gdao.GenerateLoginToken(ctx, &rsp)
	case api.Indentify_TEACHER:
		gdao.NewTeacher(ctx, req)
		if err := gdao.FillTeacherByAccount(ctx, req.Account, &rsp.User); err != nil {
			return err
		}
		return gdao.GenerateLoginToken(ctx, &rsp)
	case api.Indentify_SECRETARY:
		gdao.NewTeacher(ctx, req)
		if err := gdao.FillTeacherByAccount(ctx, req.Account, &rsp.User); err != nil {
			return err
		}
		return gdao.GenerateLoginToken(ctx, &rsp)
	default:
		fmt.Println("账号类型错误")
		return errors.New("rtype error")
	}
}

func CheckAccount(ctx context.Context, account string) error {
	// if err := gdao.CheckIDCardInStudent(ctx, account); err != nil {
	// 	return err
	// }
	if err := gdao.CheckAccountInStudent(ctx, account); err != nil {
		return err
	}
	// if err := gdao.CheckIDCardInTeacher(ctx, account); err != nil {
	// 	return err
	// }
	if err := gdao.CheckAccountInTeacher(ctx, account); err != nil {
		return err
	}
	// if err := gdao.CheckIDCardInSecretary(ctx, account); err != nil {
	// 	return err
	// }
	if err := gdao.CheckAccountInSecretary(ctx, account); err != nil {
		return err
	}
	return nil
}

func (s *PassportService) Login(ctx context.Context, req *api.LoginReq, rsp *api.PassportUserReply) error {
	logrus.Infof("Login req is %v ", req)
	//loginType := FindUserIndentifyByAccount(ctx, req.Account, rsp)
	logrus.Infof("loginType is %v ", req.Ltype)
	logrus.Infof("login rsp is %v ", rsp)
	if err := FindUserByLtypeAndAccount(ctx, req.Account, req.Ltype, rsp); err != nil {
		return err
	}
	rsp.User.Usertype = int64(req.Ltype)
	if !checkPassWord(req.Password, rsp.User.Password) {
		logrus.Infof("password err!")
		return errors.New("password weong")
	}
	return gdao.GenerateLoginToken(ctx, &rsp)
}

func checkPassWord(in, orign string) bool {
	if in == orign {
		return true
	} else {
		return false
	}
}
func FindUserByLtypeAndAccount(ctx context.Context, account string, ltype api.Indentify, rsp *api.PassportUserReply) error {
	var err error
	switch ltype {
	case api.Indentify_STUDENT:
		err = gdao.FillStudentByAccount(ctx, account, &rsp.User)
	case api.Indentify_TEACHER:
		err = gdao.FillTeacherByAccount(ctx, account, &rsp.User)
	case api.Indentify_SECRETARY:
		err = gdao.FillSecretaryByAccount(ctx, account, &rsp.User)
	default:
		return errors.New("unknow ltype")
	}
	if err != nil {
		return errors.New("need register")
	}
	return nil
}

func (s *PassportService) CheckSToken(ctx context.Context, req *api.SessionTokenReq, rsp *uapi.Empty) error {
	logrus.Infof("check token req is %v ", req)
	//FindUserById(ctx, req.Type, req.Uid, rsp)
	return gdao.CheckSessionToken(ctx, req.Uid, req.Stoken)
}

func (s *PassportService) ChangePassword(ctx context.Context, req *api.ChangePasswordReq, rsp *uapi.Empty) error {
	logrus.Infof("change password req is %v ", req)
	return gdao.ChangePassword(ctx, req.Account, req.Password, &rsp)
}

func (s *PassportService) CheckIndentify(ctx context.Context, req *uapi.UserReq, rsp *api.IndentifyReply) error {
	logrus.Infof("check indentify req is %v ", req)
	indentify := FindUserIndentifyById(ctx, bson.ObjectIdHex(req.Userid))
	fmt.Println("indentify is", indentify)
	rsp.Ltype = api.Indentify(indentify)
	return nil
}

func FindUserIndentifyById(ctx context.Context, id bson.ObjectId) int64 {
	rsp := new(api.PassportUserReply)
	if gdao.FillStudentById(ctx, id, &rsp.User) == nil {
		return 0
	} else if gdao.FillTeacherById(ctx, id, &rsp.User) == nil {
		return 1
	} else if gdao.FillSecretaryById(ctx, id, &rsp.User) == nil {
		return 2
	}
	return -1
}
func FindUserById(ctx context.Context, indentify api.Indentify, uid string, rsp *api.PassportUserReply) {
	switch indentify {
	case api.Indentify_STUDENT:
		gdao.FillStudentById(ctx, bson.ObjectIdHex(uid), &rsp.User)
	case api.Indentify_TEACHER:
		gdao.FillTeacherById(ctx, bson.ObjectIdHex(uid), &rsp.User)
	case api.Indentify_SECRETARY:
		gdao.FillSecretaryById(ctx, bson.ObjectIdHex(uid), &rsp.User)
	}
	return
}
func (s *PassportService) ChangeUserInfo(ctx context.Context, req *api.ChangeUserInfoReq, rsp *api.PassportUserReply) error {
	logrus.Infof("change userInfo req is %v ", req)
	// rsp = FindUserById(ctx, bson.ObjectIdHex(req.Userid))
	gdao.ChangeUserInfo(ctx, req)
	FindUserById(ctx, req.Usertype, req.Userid, rsp)
	return nil
}
