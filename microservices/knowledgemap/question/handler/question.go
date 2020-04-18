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
	if err := gdao.FillQuestionBySubject(ctx, req.Kind, req.Subject, req.Course, bson.ObjectIdHex(req.Knowledge), questions); err != nil {
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
