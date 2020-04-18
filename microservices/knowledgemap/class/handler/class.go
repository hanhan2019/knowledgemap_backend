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
	logrus.Infof("UserClassInfo req is %v ", req)
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
