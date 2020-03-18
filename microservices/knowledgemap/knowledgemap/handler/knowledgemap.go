package handler

import (
	"context"
	"encoding/json"
	"errors"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"
	math "knowledgemap_backend/microservices/knowledgemap/knowledgemap/data/math"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/model"
	qapi "knowledgemap_backend/microservices/knowledgemap/question/api"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
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

func (s *KnowledgeMapService) GetMyKnowledegeMapBySubject(ctx context.Context, req *api.CRqQueryMyMapBySubject, rsp *api.KnowledegeMapInfo) error {
	if req.Subject == "" {
		logrus.Errorln("subject is undefine!")
		return errors.New("subject is undefine!")
	}
	mathKnowledgeMap := math.GetMathKnowledgeMap()
	questionInfo, err := questionSrv.GetMyQuestionInfo(ctx, &qapi.CRqQueryMyQuestionInfoBySubject{Uid: req.Uid, Subject: req.Subject, Endtime: req.Endtime})
	for _, v := range questionInfo.Knowledgenodes {
		node := &model.Node{}
		err := gdao.QueryNodeInfoByNodeID(ctx, bson.ObjectIdHex(v), node)
		if err != nil {
			logrus.Errorf("can't find node:%v", v)
			continue
		}
		nodeInfo := CreateKnowledgeMapFromNodes(ctx, node)
		if len(nodeInfo) != 0 {
			mathKnowledgeMap["@graph"] = append(mathKnowledgeMap["@graph"].([]map[string]interface{}), nodeInfo...)
		}
	}
	info, _ := json.Marshal(mathKnowledgeMap)
	rsp.Knowledgemap = string(info[:])
	return err
}

func CreateKnowledgeMapFromNodes(ctx context.Context, node *model.Node) []map[string]interface{} {
	info := []map[string]interface{}{}
	relations := new([]*model.Relation)
	err := gdao.QueryRelationByNodeID(ctx, node.ID, relations)
	if err != nil {
		logrus.Errorf("can't find node relations:%v", node.ID.Hex())
		return info
	}

	contains := []map[string]interface{}{}
	derives := []map[string]interface{}{}
	affects := []map[string]interface{}{}
	for _, v := range *relations {
		node := &model.Node{}
		err := gdao.QueryNodeInfoByNodeID(ctx, v.ObjectNodeID, node)
		if err != nil {
			logrus.Errorf("can't find node:%v", v)
			continue
		}
		switch v.Relation {
		case model.RelationContain:
			contains = append(contains, map[string]interface{}{"id": node.Kind})
		case model.RelationDerive:
			derives = append(derives, map[string]interface{}{"id": node.Kind})
		case model.RelationAffect:
			affects = append(affects, map[string]interface{}{"id": node.Kind})
		default:
		}
	}

	label := []map[string]interface{}{}
	for _, v := range node.Label {
		label = append(label, map[string]interface{}{"language": v.Language, "value": v.Value})
	}
	info = append(info, map[string]interface{}{
		"id":       node.Kind,
		"label":    label,
		"type":     node.Type,
		"contains": contains,
		"derives":  derives,
		"affects":  affects,
	})
	return info
}
