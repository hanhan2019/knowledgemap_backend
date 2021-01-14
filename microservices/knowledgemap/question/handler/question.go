package handler

import (
	"context"
	"fmt"
	"knowledgemap_backend/microservices/knowledgemap/question/api"
	"knowledgemap_backend/microservices/knowledgemap/question/model"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"time"

	"github.com/Sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

type QuestionService struct{}

const SIMPLE_PAPER_RECORD_ID string = "000000000000000000000000"

func (s *QuestionService) GetMyQuestionInfo(ctx context.Context, req *api.CRqQueryMyQuestionInfoBySubject, rsp *api.CRpMyQuestionInfoBySubject) error {
	logrus.Printf("GetMyQuestionInfo, uid:%v, subject:%v, endtime:%v", req.Uid, req.Subject, req.Endtime)
	records := new([]*model.AnswerRecord)
	if err := gdao.QueryUserAnswerRecords(ctx, bson.ObjectIdHex(req.Uid), req.Subject, req.Endtime, records); err != nil {
		logrus.Errorf("GetMyQuestionInfo wrong:%v ,uid:%v, subject:%v", err, req.Uid, req.Subject)
		return err
	}
	for _, v := range *records {
		questionInfo := new(model.Question)
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
	question := &model.Question{id, req.Name, model.QUESTION_KIND(req.Questiontype), req.Title, req.Istitleimg, shiftProto2ModelOptions(req.Options), shiftProto2ModelOptions(req.Answers), req.Subject, req.Course, bson.ObjectIdHex(req.Knowledge), req.Creatername, time.Now().Unix(), req.Needcheck, req.Explain, req.Star}
	if err := gdao.NewQuestion(ctx, question); err != nil {
		fmt.Println("create question error", err)
		return fmt.Errorf("创建题目失败")
	} else {
		rsp.Id = id.Hex()
		rsp.Questiontype = req.Questiontype
		rsp.Name = req.Name
		rsp.Knowledge = req.Knowledge
		rsp.Options = req.Options
		rsp.Subject = req.Subject
		rsp.Answers = req.Answers
		rsp.Title = req.Title
		rsp.Course = req.Course
		rsp.Explain = req.Explain
		rsp.Star = req.Star
	}
	return nil
}

func (s *QuestionService) QueryQuestion(ctx context.Context, req *api.QueryQuestionReq, rsp *api.QueryQuestionReply) error {
	logrus.Infof("QueryQuestion req is %v ", req)
	questions := new([]*model.Question)
	err, allCount := gdao.FillQuestionByConditions(ctx, req.Subject, req.Course, req.Questiontype, questions, req.Page, PageCount)
	if err != nil {
		fmt.Println("query questions info error", err)
		return fmt.Errorf("查询题目失败")
	} else {
		for _, v := range *questions {
			info := new(api.QuestionInfoReply)
			info.Id = v.ID.Hex()
			info.Name = v.Name
			info.Questiontype = int64(v.Kind)
			info.Subject = v.Subject
			info.Course = v.Course
			info.Title = v.Content
			info.Istitleimg = v.IsQImg
			info.Options = shiftModelOptions2Proto(&v.Options)
			info.Answers = shiftModelOptions2Proto(&v.Answers)
			info.Knowledge = v.Knowledge.Hex()
			info.Explain = v.Explain
			info.Star = v.Star
			info.Needcheck = v.NeedCheck
			rsp.Questions = append(rsp.Questions, info)
		}
		rsp.Currentpage = req.Page
		rsp.Totalpage = int64(allCount / PageCount)
	}
	return nil
}

func (s *QuestionService) QueryQuestionKind(ctx context.Context, req *uapi.Empty, rsp *api.QuestionKindInfoReply) error {
	logrus.Infof("QueryQuestionKind req is %v ", req)
	rsp.Kindlist = []string{string(model.SIMPLE_CHOICE_QUESTION), string(model.MULTIPLE_CHOICE_QUESTION), string(model.ESSAY_QUESTION)}
	return nil
}

func (s *QuestionService) DoQuestion(ctx context.Context, req *api.DoQuestionInfo, rsp *api.QuestionItems) error {
	logrus.Infof("DoQuestion req is %v ", req)
	fmt.Println("更新了")
	question := new(model.Question)
	if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(req.Questionid), question); err != nil {
		logrus.Errorf("DoPaper 查找题目 %v 失败\n", req.Questionid)
		return fmt.Errorf("查询题目失败")
	}
	time := time.Now().Unix()
	result := model.WRONG
	isTrue := false
	if !question.NeedCheck {
		// questionAnswers := shiftProto2ModelAnswers(req.Answers)
		rightAnswers := []string{}
		for _, i := range question.Answers {
			rightAnswers = append(rightAnswers, i.Prefix)
		}
		isTrue = StringSliceReflectEqual(rightAnswers, req.Answers)
	} else {
		return fmt.Errorf("该题目请自行比对答案")
	}
	if isTrue {
		result = model.RIGHT
	}
	score := int64(100)
	//&api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result}
	questionRecord := &model.AnswerRecord{bson.NewObjectId(), bson.ObjectIdHex(req.Userid), req.Username, bson.ObjectIdHex(req.Questionid), req.Answers, result, question.Subject, "", bson.ObjectIdHex(SIMPLE_PAPER_RECORD_ID), "", score, time}
	err := gdao.NewAnswerRecord(ctx, questionRecord)
	if err != nil {
		logrus.Errorf("DoPaper 插入答题记录 %v 失败\n", err)
	}

	rsp.Id = req.Questionid
	rsp.Questiontype = int64(question.Kind)
	rsp.Title = question.Content
	rsp.Istitleimg = question.IsQImg
	rsp.Options = shiftModelOptions2Proto(&question.Options)
	rsp.Rightoptions = shiftModelOptions2Proto(&question.Answers)
	rsp.Getoptions = req.Answers
	rsp.Needcheck = question.NeedCheck
	rsp.Explain = question.Explain
	rsp.Result = int64(result)
	return nil
}
