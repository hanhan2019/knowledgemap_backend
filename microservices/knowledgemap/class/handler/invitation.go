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

func (s *ClassService) CreateInvitaion(ctx context.Context, req *api.InvitationReq, rsp *uapi.Empty) error {
	logrus.Infof("create invitaion req is %v ", req)
	if err := CheckClass(ctx, req.Classid); err != nil {
		return err
	}
	if invitaion, err := gdao.NewInvitation(ctx, req); err != nil {
		fmt.Println("create invitaion err", err)
		return fmt.Errorf("邀请码已被使用")
	} else if invitaion.DropTime != 0 {
		return fmt.Errorf("邀请码已被废弃")
	}
	return nil
}

func (s *ClassService) StopInvitaion(ctx context.Context, req *api.InvitationReq, rsp *uapi.Empty) error {
	logrus.Infof("stop invitaion req is %v ", req)
	if err := gdao.StopInvitaion(ctx, req.Invitaioncode); err != nil {
		fmt.Println("stop invitaion err", err)
		return fmt.Errorf("停用邀请码失败")
	}
	return nil
}

func (s *ClassService) InvitaionInfo(ctx context.Context, req *api.InvitationReq, rsp *api.ClassReply) error {
	logrus.Infof("query invitaion req is %v ", req)
	invitaion := &model.Invitation{}
	if err := gdao.FillInvitaion(ctx, req.Invitaioncode, invitaion); err != nil {
		fmt.Println("query invitaion err", err)
		return fmt.Errorf("查找邀请码失败")
	} else if invitaion.ClassId != "" {
		err = gdao.FillClassByID(ctx, bson.ObjectIdHex(invitaion.ClassId), &rsp)
		if err != nil {
			fmt.Println("query invitaion of class err", err)
			return fmt.Errorf("查找邀请码对应班级信息失败")
		}
	}
	return nil
}

func CheckClass(ctx context.Context, classId string) error {
	class := &api.ClassReply{}
	if err := gdao.FillClassByID(ctx, bson.ObjectIdHex(classId), &class); err != nil {
		fmt.Println("query invitaion of class err", err)
		return fmt.Errorf("查无此班级")
	}
	return nil
}
