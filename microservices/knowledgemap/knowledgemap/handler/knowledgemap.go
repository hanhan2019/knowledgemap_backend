package handler

import (
	"context"
	"errors"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"

	"github.com/Sirupsen/logrus"
)

type KnowledgeMapService struct{}

// func (s *PassportService) Register(ctx context.Context, req *api.RegisterReq, rsp *api.PassportUserReply) error {
// 	fmt.Println("here")
// 	if err := CheckIDCard(ctx, req.Idcard); err != nil {
// 		fmt.Printf("CheckIDCard error %v ", err)
// 		return errors.New("CheckIDCard error")
// 	}

// 	if err := CheckAccount(ctx, req.Account); err != nil {
// 		fmt.Printf("CheckAccount error %v ", err)
// 		return errors.New("CheckAccount error")
// 	}
// 	// if api.RegisterReq_STUDENT == req.Rtype {
// 	// 	gdao.NewUser(ctx, req)
// 	// } else if api.RegisterReq_WITH_FACEBOOK_TOKEN == req.Rtype {
// 	// 	gdao.NewUserWithFacebookToken(ctx, req)
// 	// }
// 	switch req.Rtype {
// 	case api.RegisterReq_STUDENT:
// 		gdao.NewStudent(ctx, req)
// 		if err := gdao.FillUserByIDCardInStudent(ctx, req.Idcard, &rsp.User); err != nil {
// 			return err
// 		}
// 		return gdao.GenerateLoginToken(ctx, &rsp)
// 	case api.RegisterReq_TEACHER:
// 		gdao.NewTeacher(ctx, req)
// 		if err := gdao.FillUserByIDCardInTeacher(ctx, req.Idcard, &rsp.User); err != nil {
// 			return err
// 		}
// 		return gdao.GenerateLoginToken(ctx, &rsp)
// 	default:
// 		fmt.Println("there_2")
// 		return nil
// 	}
// }

// func CheckIDCard(ctx context.Context, idCard string) error {
// 	if err := gdao.CheckIDCardInStudent(ctx, idCard); err != nil {
// 		return err
// 	}
// 	if err := gdao.CheckIDCardInTeacher(ctx, idCard); err != nil {
// 		return err
// 	}
// 	if err := gdao.CheckIDCardInSecretary(ctx, idCard); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func CheckAccount(ctx context.Context, account string) error {
// 	if err := gdao.CheckIDCardInStudent(ctx, account); err != nil {
// 		return err
// 	}
// 	if err := gdao.CheckIDCardInTeacher(ctx, account); err != nil {
// 		return err
// 	}
// 	if err := gdao.CheckIDCardInSecretary(ctx, account); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *PassportService) Login(ctx context.Context, req *api.LoginReq, rsp *api.PassportUserReply) error {
// 	logrus.Infof("Login req is %v ", req)
// 	loginType := FindUserByAccount(ctx, req.Account, rsp)
// 	//check error
// 	logrus.Infof("loginType is %v ", loginType)
// 	logrus.Infof("login rsp is %v ", rsp)
// 	rsp.User.Usertype = int64(loginType)
// 	if loginType == api.LoginReq_NOTFOUND {
// 		return errors.New("登陆失败")
// 	}
// 	if !checkPassWord(req.Passward, rsp.User.Password) {
// 		logrus.Infof("password err!")
// 		return errors.New("密码错误!")
// 	}
// 	return gdao.GenerateLoginToken(ctx, &rsp)
// }

// func checkPassWord(in, orign string) bool {
// 	if in == orign {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func FindUserByAccount(ctx context.Context, account string, rsp *api.PassportUserReply) api.LoginReq_LoginType {
// 	if gdao.FillStudentByAccount(ctx, account, &rsp.User) == nil {
// 		return api.LoginReq_STUDENT
// 	} else if gdao.FillTeacherByAccount(ctx, account, &rsp.User) == nil {
// 		return api.LoginReq_TEACHER
// 	} else if gdao.FillSecretaryByAccount(ctx, account, &rsp.User) == nil {
// 		return api.LoginReq_SECRETARY
// 	}
// 	return api.LoginReq_NOTFOUND
// }

// func (s *PassportService) CheckSToken(ctx context.Context, req *api.SessionTokenReq, rsp *uapi.Empty) error {
// 	return gdao.CheckSessionToken(ctx, req.Uid, req.Stoken)
// }

func (s *KnowledgeMapService) GetKnowledegeMapBySubject(ctx context.Context, req *api.CRqQueryMapBySubject, rsp *api.KnowledegeMapInfo) error {
	if req.Subject == "" {
		logrus.Errorln("subject is undefine!")
		return errors.New("subject is undefine!")
	}
	knowledgeMap, err := gdao.QueryKnowledgeMapByCourse(req.Subject)
	rsp.Knowledgemap = knowledgeMap
	return err
}
