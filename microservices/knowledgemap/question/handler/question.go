package handler

import (
	"context"
	"fmt"
	"knowledgemap_backend/microservices/knowledgemap/question/api"
	"knowledgemap_backend/microservices/knowledgemap/question/model"
	"time"

	"github.com/Sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

type QuestionService struct{}

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

func (s *QuestionService) GetMyQuestionInfo(ctx context.Context, req *api.CRqQueryMyQuestionInfoBySubject, rsp *api.CRpMyQuestionInfoBySubject) error {
	logrus.Printf("GetMyQuestionInfo, uid:%v, subject:%v, endtime:%v", req.Uid, req.Subject, req.Endtime)
	records := new([]*model.AnswerRecord)
	if err := gdao.QueryUserAnswerRecords(ctx, bson.ObjectIdHex(req.Uid), req.Subject, req.Endtime, records); err != nil {
		logrus.Errorf("GetMyQuestionInfo wrong:%v ,uid:%v, subject:%v", err, req.Uid, req.Subject)
		return err
	}
	for _, v := range *records {
		questionInfo := new(model.Qusetion)
		if err := gdao.QueryQuestionInfo(ctx, v.QuestionID, questionInfo); err != nil {
			logrus.Errorf("QueryQuestionInfo wrong:%v ,questionId:%v", err, v.QuestionID)
			continue
		}
		rsp.Knowledgenodes = append(rsp.Knowledgenodes, questionInfo.Knowledge.Hex())
	}
	return nil
}

func (s *QuestionService) CreateQuestion(ctx context.Context, req *api.CreateQuestionReq, rsp *api.QuestionInfoReply) error {
	logrus.Infof("CreateQuestion req is %v ", req)
	id := bson.NewObjectId()
	question := &model.Qusetion{id, model.Qusetion_Kind(req.Kind), req.Content, req.Option, req.Answer, req.Subject, req.Course, bson.ObjectIdHex(req.Knowledge), time.Now().Unix()}
	if err := gdao.NewQuestion(ctx, question); err != nil {
		fmt.Println("create question error", err)
		return fmt.Errorf("创建题目失败")
	} else {
		rsp.Id = id.Hex()
		rsp.Kind = req.Kind
		rsp.Knowledge = req.Knowledge
		rsp.Option = req.Option
		rsp.Subject = req.Subject
		rsp.Answer = req.Answer
		rsp.Content = req.Content
		rsp.Course = req.Course
	}
	return nil
}

func (s *QuestionService) QueryQuestion(ctx context.Context, req *api.QueryQuestionReq, rsp *api.QueryQuestionReply) error {
	logrus.Infof("QueryQuestion req is %v ", req)
	questions := new([]*model.Qusetion)
	if err := gdao.FillQuestionBySubject(ctx, req.Kind, req.Subject, req.Course, req.Knowledge, questions); err != nil {
		fmt.Println("query questions info error", err)
		return fmt.Errorf("查询题目失败")
	} else {
		for _, v := range *questions {
			info := new(api.QuestionInfoReply)
			info.Id = v.ID.Hex()
			info.Kind = int64(v.Kind)
			info.Subject = v.Subject
			info.Content = v.Content
			info.Course = v.Course
			info.Answer = v.Answer
			info.Option = v.Option
			info.Knowledge = v.Knowledge.Hex()
			rsp.Questions = append(rsp.Questions, info)
		}
	}
	return nil
}
