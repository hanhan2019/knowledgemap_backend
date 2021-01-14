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

const PageCount = 10

type ClassService struct{}

func (s *ClassService) CreateClass(ctx context.Context, req *api.CreateClassReq, rsp *api.ClassReply) error {
	logrus.Infof("create class req is %v ", req)
	//cheack class
	if err := gdao.CheckNewClass(ctx, req.Name, req.Subject, req.Teacherid, req.College, req.Course); err != nil {
		fmt.Println(err)
		return err
	}
	//create class
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
		// rsp.Major = class.Major
		rsp.Teachername = class.TeacherName
		rsp.College = class.College
		rsp.Course = class.Course
		rsp.Subject = class.Sbuject
		rsp.Introduction = class.Introduction
		rsp.Inclass = true
	}
}

func checkInClass(rsp *api.ClassReply, myClasses *[]*model.ClassUser) {
	rsp.Inclass = false
	for _, b := range *myClasses {
		if rsp.Classid == b.ClassId {
			rsp.Inclass = true
			break
		}
	}
	return
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
	req.Indentify = string(model.Student)
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

//查询用户加入的班级
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
		return err
	}
	allStudents := new([]*model.ClassUser)
	err, allCount := gdao.FillAllStudentsById(ctx, req.Classid, allStudents, int(req.Page), int(PageCount))
	if err != nil {
		fmt.Println("query all students by classid err", err)
		return fmt.Errorf("查无此班级学生信息")
	}
	fmt.Println(len(*allStudents))
	for _, v := range *allStudents {
		info := new(api.StudentInfo)
		info.Userid = v.UserId
		// info.Usernumber = v.Number
		if v.UserId != "" {
			fmt.Println("查学生信息")
			userInfoReq := new(uapi.UserReq)
			userInfoReq.Userid = v.UserId
			userInfoReq.Identify = uapi.Identify_STUDENT
			if res, err := userSrv.UserInfo(context.TODO(), userInfoReq); err != nil {
				logrus.Errorf("query user %v info error %v\n", v.UserId, err)
			} else {
				user := res.GetUser()
				info.Number = user.GetNumber()
				info.Username = user.GetUsername()
				info.Account = user.GetAccount()
				info.Sex = user.GetSex()
				info.College = user.GetCollege()
				info.Createtime = user.GetCreatetime()
				info.Status = 1
			}
		}
		rsp.Students = append(rsp.Students, info)
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	return nil
}

func (s *ClassService) SearchClassesInfo(ctx context.Context, req *api.SearchClassesInfoReq, rsp *api.SearchClassesInfoReply) error {
	logrus.Infof("SearchClassesInfo req is %v ", req)
	classes := new([]*model.Class)
	allCount, err := gdao.FillClassByConditions(ctx, req, classes, PageCount)
	if err != nil {
		return err
	}
	myClasses := new([]*model.ClassUser)
	gdao.FillMyAllClass(ctx, req.Userid, myClasses)
	for _, v := range *classes {
		info := new(api.ClassReply)
		convertClass(v, info)
		checkInClass(info, myClasses)
		rsp.Classes = append(rsp.Classes, info)
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	return nil
}

//TODO 补齐去重查询
func (s *ClassService) QueryFormList(ctx context.Context, req *uapi.Empty, rsp *api.QueryFormListReply) error {
	logrus.Infof("QueryFormList req is %v ", req)
	colleges := []string{"数学科学学院", "中文系", "信息工程学院", "历史学院"}
	courses := []string{"高数", "线代", "语文", "英语", "c语言"}
	subjects := []string{"数学", "历史", "语文", "计算机"}
	rsp.Colleges = colleges
	rsp.Courses = courses
	rsp.Subjects = subjects
	return nil
}

func (s *ClassService) DeleteStudent(ctx context.Context, req *api.DeleteStudentReq, rsp *uapi.Empty) error {
	logrus.Infof("DeleteStudent req is %v ", req)
	if err := gdao.DeleteStudenInClass(ctx, req.Classid, req.Userid); err != nil {
		logrus.Errorf("delete student %v in class %v error %v", req.Classid, req.Userid, err)
	}

	return nil
}

func (s *ClassService) QueryStudentInClass(ctx context.Context, req *api.QueryStudentInClassReq, rsp *api.QueryClassUserInfoReply) error {
	logrus.Infof("QueryStudentInClass req is %v ", req)
	if err := CheckClass(ctx, req.Classid); err != nil {
		return err
	}
	allStudents := new([]*model.ClassUser)
	err, allCount := gdao.FillStudentsByUserName(ctx, req.Classid, allStudents, int(req.Page), int(PageCount), req.Username)
	if err != nil {
		fmt.Println("query students by name err", err)
		return fmt.Errorf("此班级查无此学生信息")
	}
	fmt.Println(len(*allStudents))
	for _, v := range *allStudents {
		info := new(api.StudentInfo)
		info.Userid = v.UserId
		info.Username = v.UserName
		// info.Usernumber = v.Number
		if v.UserId != "" {
			userInfoReq := new(uapi.UserReq)
			userInfoReq.Userid = v.UserId
			userInfoReq.Identify = uapi.Identify_STUDENT
			if res, err := userSrv.UserInfo(context.TODO(), userInfoReq); err != nil {
				logrus.Errorf("query user %v info error %v\n", v.UserId, err)
			} else {
				user := res.GetUser()
				info.Number = user.GetNumber()
				info.Account = user.GetAccount()
				info.Sex = user.GetSex()
				info.College = user.GetCollege()
				info.Createtime = user.GetCreatetime()
				info.Status = 1
			}
		}
		rsp.Students = append(rsp.Students, info)
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	return nil
}
