package handler

import (
	"context"
	"fmt"
	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"

	"knowledgemap_backend/microservices/knowledgemap/question/api"
	"knowledgemap_backend/microservices/knowledgemap/question/model"
	"reflect"

	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

const (
	INDEX_SHOW_COUNT = 6
)

// package handler

// import (
// 	"context"
// 	"fmt"
// 	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
// 	"knowledgemap_backend/microservices/knowledgemap/question/api"
// 	"knowledgemap_backend/microservices/knowledgemap/question/model"
// 	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
// 	"reflect"
// 	"time"

// 	"github.com/Sirupsen/logrus"
// 	"gopkg.in/mgo.v2/bson"
// )

// func (s *QuestionService) CreateHomeWork(ctx context.Context, req *api.CreateHomeWorkReq, rsp *api.CreateHomeWorkReply) error {
// 	logrus.Infof("CreateHomeWork req is %v ", req)
// 	id := bson.NewObjectId()
// 	question := &model.HomeWork{id, req.Name, req.Classid, req.Questions, time.Now().Unix(), req.Endtime}
// 	if err := gdao.NewHomeWork(ctx, question); err != nil {
// 		fmt.Println("create homework error", err)
// 		return fmt.Errorf("创建作业失败")
// 	} else {
// 		rsp.Homeworkid = id.Hex()
// 	}

// 	classStudents := []*capi.QueryClassUserInfoReply_StudentInfo{}
// 	if req.Classid != "" {
// 		classReq := new(capi.ClassReq)
// 		classReq.Classid = req.Classid
// 		if res, err := classSrv.QueryClassUserInfo(context.TODO(), classReq); err != nil {
// 			logrus.Error("CreateHomeWork query students in class %v error %v", req.Classid, err)
// 		} else {
// 			classStudents = res.GetStudents()
// 		}
// 	}
// 	for _, v := range classStudents {
// 		recordId := bson.NewObjectId()
// 		record := &model.HomeWorkRecord{recordId, req.Name, id.Hex(), v.Userid, model.HOMEWORK_WAITTING_DONE, time.Now().Unix(), 0}
// 		gdao.NewHomeWorkRecord(ctx, record)
// 	}
// 	for _, v := range req.GetStudents() {
// 		recordId := bson.NewObjectId()
// 		record := &model.HomeWorkRecord{recordId, req.Name, id.Hex(), v, model.HOMEWORK_WAITTING_DONE, time.Now().Unix(), 0}
// 		gdao.NewHomeWorkRecord(ctx, record)
// 	}
// 	return nil
// }

func (s *QuestionService) QueryRecommendPaper(ctx context.Context, req *api.QueryRecommendPaperReq, rsp *api.QueryRecommendPaperReply) error {
	logrus.Infof("QueryRecommendPaper req is %v\n", req)
	//homework
	records := new([]*model.PaperRecord)
	if err := gdao.FillPaperRecordByUserId(ctx, model.HOMEWORK_RECORD_COLLECTION_NAME, bson.ObjectIdHex(req.Userid), INDEX_SHOW_COUNT, records); err != nil {
		logrus.Infoln("QueryRecommendPaper query homework record error", err)
		return fmt.Errorf("查询我的作业失败")
	} else {
		for _, v := range *records {
			if v.Status == model.RECORD_WAITTING_DONE {
				homeWork := new(model.Paper)
				if err := gdao.FillPaperById(ctx, model.HOMEWORK_COLLECTION_NAME, bson.ObjectIdHex(v.PaperId), homeWork); err != nil {
					logrus.Infoln("QueryRecommendPaper query homework %v record error", v.PaperId, err)
					continue
				}
				className := ""
				if homeWork.ClassId != "" {
					classReq := new(capi.ClassReq)
					classReq.Classid = homeWork.ClassId
					if res, err := classSrv.ClassInfo(context.TODO(), classReq); err != nil {
						logrus.Error("QueryRecommendPaper classid error %v", err)
					} else {
						className = res.GetName()
					}
				}
				paper := new(api.PaperInfo)
				paper.Origin = className
				paper.Paperid = v.PaperId
				paper.Name = v.Name
				paper.Suggesttime = homeWork.SuggestTime
				paper.Score = homeWork.Totalscore
				rsp.Homework = append(rsp.Homework, paper)
				// questionsInfo := []*api.QuestionInfo{}
				// for _, questionId := range homeWork.Questions {
				// 	question := new(model.Question)
				// 	if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(questionId), question); err != nil {
				// 		logrus.Infoln("QueryMyHomeWork query question %v  error", questionId, err)
				// 		continue
				// 	}
				// 	info := &api.QuestionInfo{question.ID.Hex(), string(question.Kind), question.Content, question.QImage, question.OImage, question.Option, int64(100) / int64(len(homeWork.Questions)), question.NeedCheck}
				// 	questionsInfo = append(questionsInfo, info)
				// }
				// homeWorkInfo := &api.HomeWorkInfo{v.ID.Hex(), homeWork.Name, questionsInfo}
				// rsp.Homework = append(rsp.Homework, homeWorkInfo)
			}
		}
	}

	//exam
	classReq := new(uapi.UserReq)
	classReq.Userid = req.Userid
	if res, err := classSrv.UserClassInfo(context.TODO(), classReq); err != nil {
		logrus.Error("QueryRecommendPaper user class error %v", err)
	} else {
		limit := INDEX_SHOW_COUNT
		for _, v := range res.GetClasses() {
			papers := new([]*model.Paper)
			err := gdao.FillPaperByClassIdLimit(ctx, model.EXAM_COLLECTION_NAME, v.Classid, papers, limit)
			if err != nil {
				logrus.Error("QueryRecommendPaper query paper error", err)
				continue
			} else {
				for _, i := range *papers {
					rsp.Exam = append(rsp.Exam, &api.PaperInfo{Paperid: i.ID.Hex(), Name: i.Name, Score: i.Totalscore, Suggesttime: i.SuggestTime, Origin: v.Name})
				}
				if len(*papers) == limit {
					break
				} else {
					limit = limit - len(*papers)
				}

			}
		}
	}

	return nil
}

// func (s *QuestionService) DoHomeWork(ctx context.Context, req *api.DoHomeWorkReq, rsp *uapi.Empty) error {
// 	logrus.Infof("DoHomeWork req is %v ", req)
// 	homeWork := new(model.HomeWork)
// 	completeStudents := []string{}
// 	err := gdao.FillHomeWorkById(ctx, bson.ObjectIdHex(req.Homeworkid), homeWork)
// 	if err != nil {
// 		fmt.Println("DoHomeWork query homework error", err)
// 		return fmt.Errorf("查找作业失败")
// 	} else if len(homeWork.Questions) != len(req.Answer) {
// 		fmt.Println("DoHomeWork: len(question) and len(answer) is not equal ", err)
// 		return fmt.Errorf("作业未完成")
// 	} else {
// 		for _, v := range req.Answer {
// 			question := new(model.Question)
// 			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v.Questionid), question); err != nil {
// 				fmt.Println("DoHomeWork 查找题目 %v 失败", v.Questionid)
// 				continue
// 			}
// 			// isTrue := StringSliceReflectEqual(question.Answer, v.Answer)
// 			result := model.WRONG
// 			score := int64(10)
// 			record := &model.AnswerRecord{bson.NewObjectId(), bson.ObjectIdHex(req.Userid), req.Username, bson.ObjectIdHex(v.Questionid), v.Answer, question.AImage, result, question.Subject, req.Homeworkid, "", score, time.Now().Unix()}
// 			gdao.NewAnswerRecord(ctx, record)
// 		}
// 		// completeStudents = homeWork.CompleteStudents
// 		completeStudents = []string{}
// 		completeStudents = append(completeStudents, req.Userid)
// 	}
// 	gdao.UpdateCompeleStudent(ctx, bson.ObjectIdHex(req.Homeworkid), completeStudents)
// 	return nil
// }

func StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

// func (s *QuestionService) QueryAnswerRecord(ctx context.Context, req *api.QueryAnswerRecordReq, rsp *api.QueryAnswerRecordReply) error {
// 	logrus.Infof("QueryAnswerRecord req is %v ", req)
// 	homeWork := new(model.HomeWork)
// 	err := gdao.FillHomeWorkById(ctx, bson.ObjectIdHex(req.Homeworkid), homeWork)
// 	if err != nil {
// 		fmt.Println("DoHomeWork query homework error", err)
// 		return fmt.Errorf("查找作业失败")
// 	} else {
// 		for _, v := range homeWork.Questions {
// 			question := new(model.Question)
// 			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v), question); err != nil {
// 				fmt.Println("QueryAnswerRecord 查找题目 %v 失败", v)
// 				continue
// 			}
// 			records := new([]*model.AnswerRecord)
// 			if err := gdao.FillAnserRecordByIdAndQuestionId(ctx, "homeworkid", req.Homeworkid, bson.ObjectIdHex(v), records); err != nil {
// 				return fmt.Errorf("QueryAnswerRecord: 查找题目 %v 的答题记录失败", v)
// 				continue
// 			}
// 			userAnswerInfo := []*api.UserAnswerInfo{}
// 			for _, b := range *records {
// 				userAnswerInfo = append(userAnswerInfo, &api.UserAnswerInfo{b.UserName, b.UserID.Hex(), b.Answer})
// 			}
// 			oneQuestionAnswerRecord := &api.AllUserAnswerInfo{v, question.Content, question.Option, question.Answer, userAnswerInfo}
// 			rsp.Homeworkrecord = append(rsp.Homeworkrecord, oneQuestionAnswerRecord)
// 		}
// 	}
// 	return nil
// }

// func (s *QuestionService) QueryHomeWorkInClass(ctx context.Context, req *capi.ClassReq, rsp *api.QueryHomeWorkInClassReply) error {
// 	logrus.Infof("QueryHomeWorkInClass req is %v ", req)
// 	homeWork := new([]*model.HomeWork)
// 	err := gdao.FillHomeWorkByClassId(ctx, req.Classid, homeWork)
// 	if err != nil {
// 		fmt.Println("QueryHomeWorkInClass query homework error", err)
// 		return fmt.Errorf("查找作业失败")
// 	} else {
// 		for _, v := range *homeWork {
// 			rsp.Homework = append(rsp.Homework, &api.HomeWorkInfo{Homeworkid: v.ID.Hex(), Name: v.Name})
// 		}
// 	}
// 	return nil
// }

// func (s *QuestionService) QueryMyHomeWorkAnswerRecordList(ctx context.Context, req *api.QueryPaperAnswerRecordListReq, rsp *api.QueryHomeWorkAnswerRecordListReply) error {
// 	logrus.Infof("QueryMyHomeWorkAnswerRecordList req is %v ", req)
// 	recordList := new([]*model.HomeWorkRecord)
// 	fmt.Println("userid is", bson.ObjectIdHex(req.Userid))
// 	err, allCount := gdao.FillHomeWorkRecordList(ctx, bson.ObjectIdHex(req.Userid), recordList, req.Page, PageCount)
// 	if err != nil {
// 		logrus.Errorln("QueryMyHomeWorkAnswerRecordList query recordlist error", err)
// 		return fmt.Errorf("查找作业记录失败")
// 	} else {
// 		fmt.Println("len is", len(*recordList))
// 		for _, v := range *recordList {
// 			homeWork := new(model.HomeWork)
// 			err := gdao.FillHomeWorkById(ctx, bson.ObjectIdHex(v.HomeWorkId), homeWork)
// 			if err != nil {
// 				fmt.Println("QueryMyHomeWorkAnswerRecordList query homework %v error", v.HomeWorkId, err)
// 				continue
// 			}
// 			courseName := ""
// 			teacherName := "匿名"
// 			if homeWork.ClassId != "" {
// 				classReq := new(capi.ClassReq)
// 				classReq.Classid = homeWork.ClassId
// 				if res, err := classSrv.ClassInfo(context.TODO(), classReq); err != nil {
// 					logrus.Error("QueryMyPaperAnswerRecordList query paper classid error %v", err)
// 				} else {
// 					courseName = res.GetCourse()
// 					teacherName = res.GetTeachername()
// 				}
// 			}
// 			rsp.Recordlist = append(rsp.Recordlist, &api.QueryHomeWorkAnswerRecordListReply_OneRecord{homeWork.Name, courseName, teacherName, string(v.Status), v.CreateTime})
// 		}
// 	}
// 	rsp.Currentpage = req.Page
// 	rsp.Totalpage = int64(allCount / PageCount)
// 	fmt.Println(allCount, rsp.Currentpage, rsp.Totalpage)
// 	return nil
// }

// func (s *QuestionService) QueryHomeWorkAnswerRecord(ctx context.Context, req *api.QueryHomeWorkAnswerRecordReq, rsp *api.QueryHomeWorkAnswerRecordReply) error {
// 	logrus.Infof("QueryHomeWorkAnswerRecord req is %v ", req)
// 	homeWork := new(model.HomeWork)
// 	err := gdao.FillHomeWorkById(ctx, bson.ObjectIdHex(req.Homeworkid), homeWork)
// 	if err != nil {
// 		fmt.Println("QueryHomeWorkAnswerRecord query homework %v error", req.Homeworkid, err)
// 		return fmt.Errorf("查找作业失败")
// 	} else {
// 		needCheck := false
// 		getScore := int64(0)
// 		records := new([]*model.AnswerRecord)
// 		// if err := gdao.FillAnserRecordById(ctx, "paperid", req.Paperid, bson.ObjectIdHex(req.Userid), bson.ObjectIdHex(v.Questionid), records); err != nil {
// 		if err := gdao.FillAnserRecordById(ctx, "homeworkid", req.Homeworkid, bson.ObjectIdHex(req.Userid), records); err != nil {
// 			// return fmt.Errorf("QueryMyPaperAnswerRecord: 查找试卷%v 题目 %v 的答题记录失败", req.Paperid, v.Questionid)
// 			logrus.Errorln("QueryHomeWorkAnswerRecord: 作业%v err:%v", req.Homeworkid, err)
// 			return fmt.Errorf("查找作业内容失败")
// 		}

// 		fmt.Printf("len of records is %v\n", len(*records))
// 		for _, i := range *records {
// 			question := new(model.Question)
// 			if err := gdao.FillQuestionById(ctx, i.QuestionID, question); err != nil {
// 				logrus.Errorln("QueryHomeWorkAnswerRecord 查找题目 %v 失败 err%v", i.QuestionID.Hex(), err)
// 				continue
// 			}
// 			getScore = getScore + i.Score
// 			NEEDCHECK := false
// 			if i.Result == model.WAITTINGCHECK {
// 				NEEDCHECK = true
// 				needCheck = true
// 			}
// 			oneQuestionAnswerRecord := &api.UserPaperAnswerInfo{i.QuestionID.Hex(), string(question.Kind), question.Name, question.Content, question.QImage, question.Option, question.OImage, question.Answer, question.AImage, i.Answer, i.AImage, NEEDCHECK, question.Explain, question.Star, string(i.Result)}
// 			rsp.Homeworkrecord = append(rsp.Homeworkrecord, oneQuestionAnswerRecord)
// 		}

// 		rsp.Homeworkid = req.Homeworkid
// 		rsp.Homeworkname = homeWork.Name
// 		rsp.Totalscore = paper.Totalscore
// 		rsp.Getscore = getScore
// 		rsp.Needcheck = needCheck
// 	}
// 	return nil
// }
