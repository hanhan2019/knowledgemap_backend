package handler

import (
	"context"

	"gopkg.in/mgo.v2/bson"

	"knowledgemap_backend/microservices/knowledgemap/user/api"

	"github.com/sirupsen/logrus"
)

type UserService struct{}

func (u *UserService) UserInfo(ctx context.Context, req *api.UserReq, rsp *api.UserReply) error {
	logrus.Infof("Userinfo req is %v ", req)
	if err := gdao.FillUserById(ctx, bson.ObjectIdHex(req.Userid), &rsp); err != nil {
		logrus.Errorf("GetUserById error %v ", err)
		return err
	} else {
		rsp.Userid = bson.ObjectId(rsp.Userid).Hex()
		return nil
	}
}

func (u *UserService) UserClassInfo(ctx context.Context, req *api.ClassReq, rsp *api.ClassReply) error {
	logrus.Infof("UserClassInfo req is %v ", req)
	if err := gdao.GetAllStudentInClass(ctx, req.Classid, &rsp); err != nil {
		logrus.Errorf("UserClassInfo error %v ", err)
		return err
	} else {
		return nil
	}
}
