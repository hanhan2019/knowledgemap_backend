package handler

import (
	"context"
	"fmt"
	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
	"knowledgemap_backend/microservices/knowledgemap/question/api"
	"knowledgemap_backend/microservices/knowledgemap/question/model"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"reflect"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

func (s *QuestionService) CreateHomeWork(ctx context.Context, req *api.CreateHomeWorkReq, rsp *api.CreateHomeWorkReply) error {
	logrus.Infof("CreateHomeWork req is %v ", req)
	id := bson.NewObjectId()
	question := &model.HomeWork{id, req.Name, req.Classid, req.Students, req.Questions, []string{}, time.Now().Unix(), time.Now().Unix()}
	if err := gdao.NewHomeWork(ctx, question); err != nil {
		fmt.Println("create homework error", err)
		return fmt.Errorf("创建作业失败")
	} else {
		rsp.Homeworkid = id.Hex()
	}
	return nil
}

func (s *QuestionService) QueryMyHomeWork(ctx context.Context, req *api.QueryMyHomeWorkReq, rsp *api.QueryMyHomeWorkReply) error {
	logrus.Infof("QueryMyHomeWork req is %v ", req)
	homeWork := new([]*model.HomeWork)
	if err := gdao.FillHomeWorkByStudentId(ctx, req.Userid, homeWork); err != nil {
		fmt.Println("query homework error", err)
		return fmt.Errorf("查询作业失败")
	} else {
		for _, v := range *homeWork {
			isOver := false
			for _, studentId := range v.CompleteStudents {
				if req.Userid == studentId {
					isOver = true
					break
				}
			}
			if !isOver {
				questionsInfo := []*api.QuestionInfo{}
				for _, questionId := range v.Questions {
					question := new(model.Qusetion)
					gdao.FillQuestionById(ctx, bson.ObjectIdHex(questionId), question)
					info := &api.QuestionInfo{question.ID.Hex(), int64(question.Kind), question.Content, question.Option}
					questionsInfo = append(questionsInfo, info)
				}
				homeWorkInfo := &api.HomeWorkInfo{v.ID.Hex(), v.Name, questionsInfo}
				rsp.Homework = append(rsp.Homework, homeWorkInfo)
			}
		}
	}
	return nil
}

func (s *QuestionService) DoHomeWork(ctx context.Context, req *api.DoHomeWorkReq, rsp *uapi.Empty) error {
	logrus.Infof("DoHomeWork req is %v ", req)
	homeWork := new(model.HomeWork)
	completeStudents := []string{}
	err := gdao.FillHomeWorkById(ctx, bson.ObjectIdHex(req.Homeworkid), homeWork)
	if err != nil {
		fmt.Println("DoHomeWork query homework error", err)
		return fmt.Errorf("查找作业失败")
	} else if len(homeWork.Questions) != len(req.Answer) {
		fmt.Println("DoHomeWork: len(question) and len(answer) is not equal ", err)
		return fmt.Errorf("作业未完成")
	} else {
		for _, v := range req.Answer {
			question := new(model.Qusetion)
			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v.Questionid), question); err != nil {
				fmt.Println("DoHomeWork 查找题目 %v 失败", v.Questionid)
				continue
			}
			isTrue := StringSliceReflectEqual(question.Answer, v.Answer)
			record := &model.AnswerRecord{bson.NewObjectId(), bson.ObjectIdHex(req.Userid), req.Username, bson.ObjectIdHex(v.Questionid), v.Answer, isTrue, question.Subject, req.Homeworkid, time.Now().Unix()}
			gdao.NewAnswerRecord(ctx, record)
		}
		completeStudents = homeWork.CompleteStudents
		completeStudents = append(completeStudents, req.Userid)
	}
	gdao.UpdateCompeleStudent(ctx, bson.ObjectIdHex(req.Homeworkid), completeStudents)
	return nil
}

func StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func (s *QuestionService) QueryAnswerRecord(ctx context.Context, req *api.QueryAnswerRecordReq, rsp *api.QueryAnswerRecordReply) error {
	logrus.Infof("QueryAnswerRecord req is %v ", req)
	homeWork := new(model.HomeWork)
	err := gdao.FillHomeWorkById(ctx, bson.ObjectIdHex(req.Homeworkid), homeWork)
	if err != nil {
		fmt.Println("DoHomeWork query homework error", err)
		return fmt.Errorf("查找作业失败")
	} else {
		for _, v := range homeWork.Questions {
			question := new(model.Qusetion)
			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v), question); err != nil {
				fmt.Println("QueryAnswerRecord 查找题目 %v 失败", v)
				continue
			}
			records := new([]*model.AnswerRecord)
			if err := gdao.FillAnserRecordByHomeWorkIdAndQuestionId(ctx, req.Homeworkid, bson.ObjectIdHex(v), records); err != nil {
				return fmt.Errorf("QueryAnswerRecord: 查找题目 %v 的答题记录失败", v)
				continue
			}
			userAnswerInfo := []*api.UserAnswerInfo{}
			for _, b := range *records {
				userAnswerInfo = append(userAnswerInfo, &api.UserAnswerInfo{b.UserName, b.UserID.Hex(), b.Answer})
			}
			oneQuestionAnswerRecord := &api.AllUserAnswerInfo{v, question.Content, question.Option, question.Answer, userAnswerInfo}
			rsp.Homeworkrecord = append(rsp.Homeworkrecord, oneQuestionAnswerRecord)
		}
	}
	return nil
}

func (s *QuestionService) QueryHomeWorkInClass(ctx context.Context, req *capi.ClassReq, rsp *api.QueryHomeWorkInClassReply) error {
	logrus.Infof("QueryHomeWorkInClass req is %v ", req)
	homeWork := new([]*model.HomeWork)
	err := gdao.FillHomeWorkByClassId(ctx, req.Classid, homeWork)
	if err != nil {
		fmt.Println("QueryHomeWorkInClass query homework error", err)
		return fmt.Errorf("查找作业失败")
	} else {
		for _, v := range *homeWork {
			rsp.Homework = append(rsp.Homework, &api.HomeWorkInfo{Homeworkid: v.ID.Hex(), Name: v.Name})
		}
	}
	return nil
}
