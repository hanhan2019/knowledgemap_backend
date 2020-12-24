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

func (s *QuestionService) CreatePaper(ctx context.Context, req *api.CreatePaperReq, rsp *api.CreatePaperReply) error {
	logrus.Infof("CreatePaper req is %v ", req)
	id := bson.NewObjectId()
	questions := []model.QuestionScore{}
	totalScore := int64(0)
	for _, v := range req.Questions {
		questions = append(questions, model.QuestionScore{v.Questionid, v.Score, v.Needcheck})
		totalScore = totalScore + v.Score
	}
	paper := &model.Paper{id, req.Name, req.Classid, questions, time.Now().Unix(), totalScore, req.Continuingtime}
	if err := gdao.NewPaper(ctx, paper); err != nil {
		fmt.Println("create paper error", err)
		return fmt.Errorf("创建试卷失败")
	} else {
		rsp.Paperid = id.Hex()
	}
	return nil
}

// func (s *QuestionService) QueryMyHomeWork(ctx context.Context, req *api.QueryMyHomeWorkReq, rsp *api.QueryMyHomeWorkReply) error {
// 	logrus.Infof("QueryMyHomeWork req is %v ", req)
// 	homeWork := new([]*model.HomeWork)
// 	if err := gdao.FillHomeWorkByStudentId(ctx, req.Userid, homeWork); err != nil {
// 		fmt.Println("query homework error", err)
// 		return fmt.Errorf("查询作业失败")
// 	} else {
// 		for _, v := range *homeWork {
// 			isOver := false
// 			for _, studentId := range v.CompleteStudents {
// 				if req.Userid == studentId {
// 					isOver = true
// 					break
// 				}
// 			}
// 			if !isOver {
// 				questionsInfo := []*api.QuestionInfo{}
// 				for _, questionId := range v.Questions {
// 					question := new(model.Qusetion)
// 					gdao.FillQuestionById(ctx, bson.ObjectIdHex(questionId), question)
// 					info := &api.QuestionInfo{question.ID.Hex(), int64(question.Kind), question.Content, question.Option}
// 					questionsInfo = append(questionsInfo, info)
// 				}
// 				homeWorkInfo := &api.HomeWorkInfo{v.ID.Hex(), v.Name, questionsInfo}
// 				rsp.Homework = append(rsp.Homework, homeWorkInfo)
// 			}
// 		}
// 	}
// 	return nil
// }

func (s *QuestionService) DoPaper(ctx context.Context, req *api.DoPaperReq, rsp *uapi.Empty) error {
	logrus.Infof("DoPaper req is %v ", req)
	paper := new(model.Paper)
	// completeStudents := []string{}
	err := gdao.FillPaperById(ctx, bson.ObjectIdHex(req.Paperid), paper)
	if err != nil {
		fmt.Println("DoPaper query paper error", err)
		return fmt.Errorf("查找试卷失败")
	} else {
		for _, v := range req.Answer {
			question := new(model.Qusetion)
			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v.Questionid), question); err != nil {
				fmt.Println("DoPaper 查找题目 %v 失败", v.Questionid)
				continue
			}
			isTrue := false
			if !question.NeedCheck {
				isTrue = StringSliceReflectEqual(question.Answer, v.Answer)
			}
			result := model.WRONG
			score := int64(0)
			if isTrue {
				result = model.RIGHT
				for _, i := range paper.Questions {
					if i.Questionid == v.Questionid {
						score = i.Score
					}
				}
			}
			record := &model.AnswerRecord{bson.NewObjectId(), bson.ObjectIdHex(req.Userid), req.Username, bson.ObjectIdHex(v.Questionid), v.Answer, v.Aimage, result, question.Subject, "", req.Paperid, score, time.Now().Unix()}
			gdao.NewAnswerRecord(ctx, record)
		}
		// completeStudents = homeWork.CompleteStudents
		// completeStudents = append(completeStudents, req.Userid)
	}
	// gdao.UpdateCompeleStudent(ctx, bson.ObjectIdHex(req.Homeworkid), completeStudents)
	return nil
}

func (s *QuestionService) QueryMyPaperAnswerRecord(ctx context.Context, req *api.QueryPaperAnswerRecordReq, rsp *api.QueryPaperAnswerRecordReply) error {
	logrus.Infof("QueryMyPaperAnswerRecord req is %v ", req)
	paper := new(model.Paper)
	err := gdao.FillPaperById(ctx, bson.ObjectIdHex(req.Paperid), paper)
	if err != nil {
		fmt.Println("QueryMyPaperAnswerRecord query paper error", err)
		return fmt.Errorf("查找试卷失败")
	} else {
		paperNeedCheck := false
		paperGetScore := int64(0)
		records := new([]*model.AnswerRecord)
		// if err := gdao.FillAnserRecordById(ctx, "paperid", req.Paperid, bson.ObjectIdHex(req.Userid), bson.ObjectIdHex(v.Questionid), records); err != nil {
		if err := gdao.FillAnserRecordById(ctx, "paperid", req.Paperid, bson.ObjectIdHex(req.Userid), records); err != nil {
			// return fmt.Errorf("QueryMyPaperAnswerRecord: 查找试卷%v 题目 %v 的答题记录失败", req.Paperid, v.Questionid)
			return fmt.Errorf("QueryMyPaperAnswerRecord: 查找试卷%v 题目", req.Paperid)
		}

		fmt.Printf("len of records is %v\n", len(*records))
		for _, i := range *records {
			question := new(model.Qusetion)
			if err := gdao.FillQuestionById(ctx, i.QuestionID, question); err != nil {
				fmt.Println("QueryMyPaperAnswerRecord 查找题目 %v 失败", i.QuestionID.Hex())
				continue
			}
			paperGetScore = paperGetScore + i.Score
			NEEDCHECK := false
			if i.IsTrue == model.WAITTINGCHECK {
				NEEDCHECK = true
				paperNeedCheck = true
			}
			oneQuestionAnswerRecord := &api.UserPaperAnswerInfo{i.QuestionID.Hex(), int64(question.Kind), question.Name, question.Content, question.QImage, question.Option, question.OImage, question.Answer, question.AImage, i.Answer, i.AImage, NEEDCHECK, question.Explain, question.Star}
			rsp.Paperrecord = append(rsp.Paperrecord, oneQuestionAnswerRecord)
		}

		rsp.Paperid = req.Paperid
		rsp.Papername = paper.Name
		rsp.Totalscore = paper.Totalscore
		rsp.Getscore = paperGetScore
		rsp.Needcheck = paperNeedCheck
	}
	return nil
}

func (s *QuestionService) QueryPaperInClass(ctx context.Context, req *api.QueryPaperInClassReq, rsp *api.QueryPaperInClassReply) error {
	logrus.Infof("QueryPaperInClass req is %v ", req)
	papers := new([]*model.Paper)
	err, allCount := gdao.FillPaperByClassId(ctx, req.Classid, papers, req.Page, PageCount)
	if err != nil {
		fmt.Println("query paper error", err)
		return fmt.Errorf("查找试卷失败")
	} else {
		for _, v := range *papers {
			rsp.Paper = append(rsp.Paper, &api.PaperInfo{Paperid: v.ID.Hex(), Name: v.Name, Totalscore: v.Totalscore, Continuingtime: v.ContinuingTime})
		}
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	fmt.Println(allCount, rsp.Currentpage, rsp.Totalpage)
	return nil
}

func (s *QuestionService) QueryPaperQuestions(ctx context.Context, req *api.QueryPaperQuestionsReq, rsp *api.QueryPaperQuestionsReply) error {
	logrus.Infof("QueryPaperQuestions req is %v ", req)
	paper := new(model.Paper)
	err := gdao.FillPaperById(ctx, bson.ObjectIdHex(req.Paperid), paper)
	if err != nil {
		fmt.Println("query paper error", err)
		return fmt.Errorf("查找试卷失败")
	} else {
		for _, v := range paper.Questions {
			question := new(model.Qusetion)
			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v.Questionid), question); err != nil {
				fmt.Println("QueryMyPaperAnswerRecord 查找题目 %v 失败", v.Questionid)
				continue
			}
			rsp.Paper = append(rsp.Paper, &api.QuestionInfo{v.Questionid, int64(question.Kind), question.Content, question.QImage, question.OImage, question.Option, v.Score})
		}
	}
	rsp.Name = paper.Name
	rsp.Totalscore = paper.Totalscore
	rsp.Continuingtime = paper.ContinuingTime
	return nil
}
