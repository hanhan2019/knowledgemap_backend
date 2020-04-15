package handler

import (
	"context"
	"fmt"
	"knowledgemap_backend/microservices/knowledgemap/class/api"
	"knowledgemap_backend/microservices/knowledgemap/class/model"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type ClassService struct{}

// func (s *PassportService) Register(ctx context.Context, req *api.RegisterReq, rsp *api.PassportUserReply) error {
// 	if err := CheckAccount(ctx, req.Account); err != nil {
// 		fmt.Printf("CheckAccount error %v ", err)
// 		return errors.New("CheckAccount error")
// 	}
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

// func CheckAccount(ctx context.Context, account string) error {
// 	if err := gdao.CheckIDCardInStudent(ctx, account); err != nil {
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
// 	if !checkPassWord(req.Password, rsp.User.Password) {
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

// func (s *PassportService) ChangePassword(ctx context.Context, req *api.ChangePasswordReq, rsp *uapi.Empty) error {
// 	return gdao.ChangePassword(ctx, req.Account, req.Password, &rsp)
// }

func (s *ClassService) CreateClass(ctx context.Context, req *api.CreateClassReq, rsp *api.ClassReply) error {
	logrus.Infof("create class req is %v ", req)
	if class, err := gdao.NewClass(ctx, req); err != nil {
		fmt.Println("create class error", err)
		return fmt.Errorf("创建班级失败")
	} else {
		joinClass := &api.JoinClassReq{req.Teacherid, class.ID.Hex(), req.Teachername, string(model.Teacher)}
		if err := gdao.NewClassUser(ctx, joinClass); err != nil {
			fmt.Println("create class teacher info error", err)
			return fmt.Errorf("初始化班级信息失败")
		}
		convertClass(class, rsp)
	}
	return nil
}

func convertClass(class *model.Class, rsp *api.ClassReply) {
	if class != nil {
		rsp.Classid = class.ID.Hex()
		rsp.Name = class.Name
		rsp.Major = class.Major
		rsp.Teachername = class.TeacherName
		rsp.College = class.College
	}
}

func (s *ClassService) ClassInfo(ctx context.Context, req *api.ClassReq, rsp *api.ClassReply) error {
	logrus.Infof("ClassInfo req is %v ", req)
	if err := gdao.FillClassByID(ctx, bson.ObjectIdHex(req.Classid), &rsp); err != nil {
		fmt.Println("query class info error", err)
		return fmt.Errorf("查询班级信息失败")
	}
	return nil
}

func (s *ClassService) JoinClass(ctx context.Context, req *api.JoinClassReq, rsp *api.UserClassReply) error {
	if err := gdao.CheckInClass(ctx, req.Userid, req.Classid); err != nil {
		fmt.Printf("Check in class error %v ", err)
		return err
	}
	if err := gdao.NewClassUser(ctx, req); err != nil {
		return err
	} else {
		myClasses := new([]*model.ClassUser)
		if err = gdao.FillMyAllClass(ctx, req.Userid, myClasses); err != nil {
			return err
		}
		for _, v := range *myClasses {
			info := new(api.ClassReply)
			gdao.FillClassByID(ctx, bson.ObjectIdHex(v.ClassId), &info)
			rsp.Classes = append(rsp.Classes, info)
		}
	}
	return nil
}

func (s *ClassService) UserClassInfo(ctx context.Context, req *uapi.UserReq, rsp *api.UserClassReply) error {
	myClasses := new([]*model.ClassUser)
	if err := gdao.FillMyAllClass(ctx, req.Userid, myClasses); err != nil {
		return err
	}
	for _, v := range *myClasses {
		info := new(api.ClassReply)
		gdao.FillClassByID(ctx, bson.ObjectIdHex(v.ClassId), &info)
		rsp.Classes = append(rsp.Classes, info)
	}
	return nil
}

func (s *ClassService) QueryClassUserInfo(ctx context.Context, req *api.ClassReq, rsp *api.QueryClassUserInfoReply) error {
	logrus.Infof("QueryClassUserInfo req is %v ", req)
	if err := CheckClass(ctx, req.Classid); err != nil {
		return nil
	}
	allStudents := new([]*model.ClassUser)
	if err := gdao.FillAllStudentsById(ctx, req.Classid, allStudents); err != nil {
		fmt.Println("query all students by classid err", err)
		return fmt.Errorf("查无此班级学生信息")
	}
	for _, v := range *allStudents {
		info := new(api.QueryClassUserInfoReply_StudentInfo)
		info.Userid = v.UserId
		info.Username = v.UserName
		rsp.Students = append(rsp.Students, info)
	}
	return nil
}
