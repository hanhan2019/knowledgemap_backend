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
	switch req.Usertype {
	case uapi.Identify_STUDENT:
		gdao.NewStudent(ctx, req)
		if err := gdao.FillStudentByAccount(ctx, req.Account, &rsp.User); err != nil {
			return err
		}
		rsp.User.Usertype = int64(req.Usertype)
		return gdao.GenerateLoginToken(ctx, int(req.Usertype), &rsp)
	case uapi.Identify_TEACHER:
		gdao.NewTeacher(ctx, req)
		if err := gdao.FillTeacherByAccount(ctx, req.Account, &rsp.User); err != nil {
			return err
		}
		rsp.User.Usertype = int64(req.Usertype)
		return gdao.GenerateLoginToken(ctx, int(req.Usertype), &rsp)
	case uapi.Identify_SECRETARY:
		gdao.NewTeacher(ctx, req)
		if err := gdao.FillTeacherByAccount(ctx, req.Account, &rsp.User); err != nil {
			return err
		}
		rsp.User.Usertype = int64(req.Usertype)
		return gdao.GenerateLoginToken(ctx, int(req.Usertype), &rsp)
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
	logrus.Infof("loginType is %v ", req.Usertype)
	logrus.Infof("login rsp is %v ", rsp)
	if err := FindUserByLtypeAndAccount(ctx, req.Account, req.Usertype, rsp); err != nil {
		return err
	}
	rsp.User.Usertype = int64(req.Usertype)
	if !checkPassWord(req.Password, rsp.User.Password) {
		logrus.Infof("password err!")
		return errors.New("password weong")
	}
	rsp.User.Password = ""
	return gdao.GenerateLoginToken(ctx, int(req.Usertype), &rsp)
}

func checkPassWord(in, orign string) bool {
	if in == orign {
		return true
	} else {
		return false
	}
}
func FindUserByLtypeAndAccount(ctx context.Context, account string, ltype uapi.Identify, rsp *api.PassportUserReply) error {
	var err error
	switch ltype {
	case uapi.Identify_STUDENT:
		err = gdao.FillStudentByAccount(ctx, account, &rsp.User)
	case uapi.Identify_TEACHER:
		err = gdao.FillTeacherByAccount(ctx, account, &rsp.User)
	case uapi.Identify_SECRETARY:
		err = gdao.FillSecretaryByAccount(ctx, account, &rsp.User)
	default:
		return errors.New("unknow ltype")
	}
	if err != nil {
		return errors.New("need register")
	}
	return nil
}

func (s *PassportService) CheckSToken(ctx context.Context, req *api.SessionTokenReq, rsp *uapi.UserReply) error {
	logrus.Infof("check token req is %v ", req)
	//FindUserById(ctx, req.Type, req.Uid, rsp)
	return gdao.CheckSessionToken(ctx, req.Cookie, rsp)
}

func (s *PassportService) ChangePassword(ctx context.Context, req *api.ChangePasswordReq, rsp *uapi.Empty) error {
	logrus.Infof("change password req is %v ", req)
	return gdao.ChangePassword(ctx, req.Userid, req.Password, req.Identify, &rsp)
}

func (s *PassportService) CheckIndentify(ctx context.Context, req *uapi.UserReq, rsp *api.IdentifyReply) error {
	logrus.Infof("check indentify req is %v ", req)
	identify := FindUserIndentifyById(ctx, bson.ObjectIdHex(req.Userid))
	fmt.Println("indentify is", identify)
	rsp.Ltype = uapi.Identify(identify)
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

func FindUserById(ctx context.Context, indentify uapi.Identify, uid string, rsp **uapi.UserReply) {
	switch indentify {
	case uapi.Identify_STUDENT:
		gdao.FillStudentById(ctx, bson.ObjectIdHex(uid), rsp)
	case uapi.Identify_TEACHER:
		gdao.FillTeacherById(ctx, bson.ObjectIdHex(uid), rsp)
	case uapi.Identify_SECRETARY:
		gdao.FillSecretaryById(ctx, bson.ObjectIdHex(uid), rsp)
	}
	return
}

func (s *PassportService) ChangeUserInfo(ctx context.Context, req *api.ChangeUserInfoReq, rsp *uapi.UserInfoReply) error {
	logrus.Infof("change userInfo req is %v ", req)
	// rsp = FindUserById(ctx, bson.ObjectIdHex(req.Userid))
	gdao.ChangeUserInfo(ctx, req)
	FindUserById(ctx, req.Usertype, req.Userid, &rsp.User)
	rsp.User.Password = ""
	rsp.User.Usertype = int64(req.Usertype)
	return nil
}
