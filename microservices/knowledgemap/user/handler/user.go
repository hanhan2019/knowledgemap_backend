package handler

import (
	"context"
	"errors"

	"knowledgemap_backend/microservices/knowledgemap/user/api"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct{}

func (u *UserService) UserInfo(ctx context.Context, req *api.UserReq, rsp *api.UserInfoReply) error {
	logrus.Infof("Userinfo req is %v ", req)
	FindUserById(ctx, req.Identify, req.Userid, &rsp.User)
	if rsp.User.Userid == "" {
		logrus.Errorf("GetUserById error")
		errors.New("修改用户信息失败")
	}
	rsp.User.Password = ""
	rsp.User.Usertype = int64(req.Identify)
	rsp.User.Imagepath = "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=1028479771,2944343576&fm=26&gp=0.jpg"
	return nil
	// if err := gdao.FillUserById(ctx, bson.ObjectIdHex(req.Userid), &rsp); err != nil {
	// 	logrus.Errorf("GetUserById error %v ", err)
	// 	return err
	// } else {
	// 	rsp.Userid = bson.ObjectId(rsp.Userid).Hex()
	// 	return nil
	// }
}
func (u *UserService) QueryUserInfo(ctx context.Context, req *api.QueryUserInfoReq, rsp *api.QueryUserInfoReply) error {
	logrus.Infof("Userinfo req is %v ", req)
	FindUserByName(ctx, req.Identify, req.Username, &rsp.Users)
	for _, v := range rsp.Users {
		if v.Userid == "" {
			logrus.Errorf("GetUserById error")
		}
		v.Usertype = int64(req.Identify)
		v.Imagepath = "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=1028479771,2944343576&fm=26&gp=0.jpg"
	}
	return nil
	// if err := gdao.FillUserById(ctx, bson.ObjectIdHex(req.Userid), &rsp); err != nil {
	// 	logrus.Errorf("GetUserById error %v ", err)
	// 	return err
	// } else {
	// 	rsp.Userid = bson.ObjectId(rsp.Userid).Hex()
	// 	return nil
	// }
}

// func (u *UserService) UserClassInfo(ctx context.Context, req *capi.ClassReq, rsp *api.ClassReply) error {
// 	logrus.Infof("UserClassInfo req is %v ", req)
// 	if err := gdao.GetAllStudentInClass(ctx, req.Classid, &rsp); err != nil {
// 		logrus.Errorf("UserClassInfo error %v ", err)
// 		return err
// 	} else {
// 		return nil
// 	}
// }

func FindUserById(ctx context.Context, indentify api.Identify, uid string, rsp **api.UserReply) {
	switch indentify {
	case api.Identify_STUDENT:
		gdao.FillStudentById(ctx, bson.ObjectIdHex(uid), rsp)
	case api.Identify_TEACHER:
		gdao.FillTeacherById(ctx, bson.ObjectIdHex(uid), rsp)
	case api.Identify_SECRETARY:
		gdao.FillSecretaryById(ctx, bson.ObjectIdHex(uid), rsp)
	}
	return
}

func FindUserByName(ctx context.Context, indentify api.Identify, uName string, rsp *[]*api.UserReply) {
	switch indentify {
	case api.Identify_STUDENT:
		gdao.FillStudentByName(ctx, uName, rsp)
	case api.Identify_TEACHER:
		gdao.FillTeacherByName(ctx, uName, rsp)
	case api.Identify_SECRETARY:
		gdao.FillSecretaryByName(ctx, uName, rsp)
	}
	return
}
