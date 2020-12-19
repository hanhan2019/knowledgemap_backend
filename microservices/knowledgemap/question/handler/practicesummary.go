package handler

import (
	"context"
	"fmt"
	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
	kapi "knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"
	"knowledgemap_backend/microservices/knowledgemap/question/api"
	"knowledgemap_backend/microservices/knowledgemap/question/model"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"

	"time"

	"github.com/Sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

const PageCount = 10

func (s *QuestionService) CreatePracticeSummary(ctx context.Context, req *api.CreatePracticeSummaryReq, rsp *api.CreatePracticeSummaryReply) error {
	logrus.Infof("CreatePracticeSummary req is %v ", req)
	id := bson.NewObjectId()
	// questionids := []bson.ObjectId{}
	// for _, v := range req.Questions {
	// 	questionids = append(questionids, bson.ObjectIdHex(v))
	// }
	// questionsInfo := new([]*model.Qusetion)
	// err := gdao.QueryMulQuestionInfo(ctx, questionids, questionsInfo)
	courseName := ""
	if req.Classid != "" {
		classReq := new(capi.ClassReq)
		classReq.Classid = req.Classid
		if res, err := classSrv.ClassInfo(context.TODO(), classReq); err != nil {
			logrus.Error("query practice summary classid error %v", err)
		} else {
			courseName = res.GetCourse()
		}
	}
	ps := &model.PracticeSummary{id, req.Name, model.PSTYPE(int(req.Pstype)), req.Ownerid, req.Ownername, req.Classid, courseName, req.Introduction, time.Now().Unix()}
	if err := gdao.NewPracticeSummary(ctx, ps); err != nil {
		fmt.Println("create practice summary error", err)
		return fmt.Errorf("创建题库失败")
	} else {
		rsp.Practicesummaryid = id.Hex()
	}
	return nil
}
func GetQuestionInfoById(questions []string) {

}
func (s *QuestionService) QueryPracticeSummaryInfo(ctx context.Context, req *api.QueryPracticeSummaryReq, rsp *api.PracticeSummaryInfo) error {
	logrus.Printf("QueryPracticeSummaryInfo req is", req)
	ps := new(model.PracticeSummary)
	if err := gdao.FillPSByID(ctx, bson.ObjectIdHex(req.Practicesummaryid), &ps); err != nil {
		logrus.Println("query practice summary info error", err)
		return fmt.Errorf("查询题库信息失败")
	}
	rsp.Practicesummaryid = ps.ID.Hex()
	rsp.Name = ps.Name
	rsp.Introduction = ps.Introduction
	_, number := gdao.QueryPSQuestionCount(ctx, bson.ObjectIdHex(req.Practicesummaryid))
	rsp.Questionnumbers = number
	rsp.Pstype = api.PracticeSummaryType(ps.Pstype)
	rsp.Ownername = ps.OwnerName
	rsp.Coursename = ps.CourseName
	// if ps.ClassId != "" {
	// 	classReq := new(capi.ClassReq)
	// 	classReq.Classid = ps.ClassId
	// 	if res, err := classSrv.ClassInfo(context.TODO(), classReq); err != nil {
	// 		logrus.Error("query practice summary classid error %v", err)
	// 	} else {
	// 		rsp.Coursename = res.GetCourse()
	// 	}
	// }
	return nil
}

func (s *QuestionService) GetPracticeSummary(ctx context.Context, req *api.QueryPracticeSummaryReq, rsp *api.PracticeSummaryDetailInfo) error {
	logrus.Infof("GetPracticeSummary req is %v ", req)
	ps := new([]*model.PracticeQuestion)
	allCount, err := gdao.FillQuestionsByPSId(ctx, bson.ObjectIdHex(req.Practicesummaryid), req.Page, ps, PageCount)
	if err != nil {
		fmt.Println("query questions info error", err)
		return fmt.Errorf("查询题目失败")
	}
	for _, v := range *ps {
		question := new(api.QuestionPre)
		question.Name = v.QuestionName
		question.Questionid = v.QuestionId.Hex()
		question.Kind = int64(v.QuestionKind)
		question.Knowledgename = v.KnowlegeName
		rsp.Questions = append(rsp.Questions, question)
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	return nil
}

func (s *QuestionService) QueryMyPracticeSummary(ctx context.Context, req *api.QueryMyPracticeSummaryReq, rsp *api.QueryMyPracticeSummaryReply) error {
	logrus.Infof("QueryMyPracticeSummary req is %v ", req)
	ps := new([]*model.PracticeSummary)
	allCount, err := gdao.FillPSByUserID(ctx, req.Userid, req.Page, ps, PageCount)
	if err != nil {
		fmt.Println("query my ps error", err)
		return fmt.Errorf("查询用户题目失败")
	}
	fmt.Println("count is ", allCount)
	for _, v := range *ps {
		onePs := new(api.PracticeSummaryInfo)
		onePs.Practicesummaryid = v.ID.Hex()
		onePs.Name = v.Name
		onePs.Pstype = api.PracticeSummaryType(int(v.Pstype))
		onePs.Introduction = v.Introduction
		_, number := gdao.QueryPSQuestionCount(ctx, v.ID)
		onePs.Questionnumbers = number
		onePs.Ownername = v.OwnerName
		onePs.Coursename = v.CourseName
		rsp.Practicesummary = append(rsp.Practicesummary, onePs)
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	return nil
}

func (s *QuestionService) AddQuestionInPS(ctx context.Context, req *api.ControllQuestionInPSReq, rsp *uapi.Empty) error {
	logrus.Infof("AddQuestionInPS req is %v ", req)
	for _, v := range req.Questions {
		oneQuestion := new(model.Qusetion)
		err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v), oneQuestion)
		if err != nil {
			logrus.Errorf("AddQuestionInPS Error,question id is %v ,err :%v", v, err)
			continue
		}
		knowledgeName := ""
		if oneQuestion.Knowledge.Hex() != "" {
			knowledgeReq := new(kapi.QueryKnowledegeInfoReq)
			knowledgeReq.Id = oneQuestion.Knowledge.Hex()
			if res, err := knowledgeMapSrv.QueryKnowledegeInfo(context.TODO(), knowledgeReq); err != nil {
				logrus.Error("query knowledge info %v error %v", oneQuestion.Knowledge.Hex(), err)
			} else {
				knowledgeName = res.GetName()
			}
		}
		psQuestion := &model.PracticeQuestion{bson.NewObjectId(), bson.ObjectIdHex(req.Practicesummaryid), oneQuestion.Name, oneQuestion.ID, oneQuestion.Kind, oneQuestion.Knowledge, knowledgeName, time.Now().Unix()}
		if err := gdao.NewPracticeQuestion(ctx, psQuestion); err != nil {
			logrus.Error("create practice question error", err)
			continue
		}
	}
	return nil
}

func (s *QuestionService) DelteQuestionInPS(ctx context.Context, req *api.ControllQuestionInPSReq, rsp *uapi.Empty) error {
	logrus.Infof("DelteQuestionInPS req is %v ", req)
	for _, v := range req.Questions {
		if err := gdao.DeletePracticeQuestion(ctx, bson.ObjectIdHex(req.Practicesummaryid), bson.ObjectIdHex(v)); err != nil {
			logrus.Error("delete question in ps error", err)
			continue
		}
	}
	return nil
}
