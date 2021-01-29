package handler

import (
	"context"
	"fmt"
	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
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
	paper := &model.Paper{id, req.Name, req.Classid, questions, time.Now().Unix(), totalScore, req.Suggesttime}
	if err := gdao.NewPaper(ctx, paper, req.Paperkind); err != nil {
		fmt.Println("create paper error", err)
		return fmt.Errorf("创建失败")
	} else {
		rsp.Paperid = id.Hex()
	}
	return nil
}

func (s *QuestionService) ChangeQuestionInPaper(ctx context.Context, req *api.ChangeQuestionInPaperReq, rsp *uapi.Empty) error {
	logrus.Infof("ChangeQuestionInPaper req is %v ", req)
	questions := []model.QuestionScore{}
	totalScore := int64(0)
	for _, v := range req.Questions {
		questions = append(questions, model.QuestionScore{v.Questionid, v.Score, v.Needcheck})
		totalScore = totalScore + v.Score
	}
	if err := gdao.ChangeQuestionsInPaper(ctx, bson.ObjectIdHex(req.Paperid), req.Paperkind, questions, totalScore); err != nil {
		fmt.Println("ChangeQuestionInPaper %v error", req.Paperid, err)
		return fmt.Errorf("修改题目失败")
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
	time := time.Now().Unix()
	paperRecordId := bson.NewObjectId()
	paper := new(model.Paper)
	// completeStudents := []string{}
	err := gdao.FillPaperById(ctx, req.Paperkind, bson.ObjectIdHex(req.Paperid), paper)
	if err != nil {
		logrus.Errorf("DoPaper query paper %v error %v", req.Paperid, err)
		return fmt.Errorf("查找试卷失败")
	} else {
		paperRecordStaus := model.RECORD_DONE
		getScore := int64(0)
		for _, v := range req.Answer {
			question := new(model.Question)
			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v.Questionid), question); err != nil {
				logrus.Errorf("DoPaper 查找题目 %v err %v", v.Questionid, err)
				continue
			}
			isTrue := false
			result := model.WRONG
			score := int64(0)
			if !question.NeedCheck {
				//不需要人工判题
				rightAnswers := []string{}
				for _, i := range question.Answers {
					rightAnswers = append(rightAnswers, i.Prefix)
				}
				isTrue = StringSliceReflectEqual(rightAnswers, v.Answers)
				if isTrue {
					result = model.RIGHT
					for _, i := range paper.Questions {
						if i.Questionid == v.Questionid {
							score = i.Score
							getScore = getScore + i.Score
						}
					}
				}
			} else {
				paperRecordStaus = model.RECORD_WAITTING_CHECK
				result = model.WAITTINGCHECK
			}
			questionRecord := &model.AnswerRecord{bson.NewObjectId(), bson.ObjectIdHex(req.Userid), req.Username, bson.ObjectIdHex(v.Questionid), v.Answers, result, question.Subject, req.Paperkind, paperRecordId, req.Paperid, score, time}
			gdao.NewAnswerRecord(ctx, questionRecord)
		}
		if req.Paperkind == model.HOMEWORK_COLLECTION_NAME {
			gdao.RemovePaperRecord(ctx, bson.ObjectIdHex(req.Userid), paper.ID, getPaperRecordTableName(req.Paperkind))
		}
		// if req.Paperkind == model.EXAM_COLLECTION_NAME {
		paperRecord := &model.PaperRecord{paperRecordId, paper.Name, bson.ObjectIdHex(req.Userid), req.Username, paperRecordStaus, req.Paperid, getScore, time, time}
		gdao.NewPaperRecord(ctx, getPaperRecordTableName(req.Paperkind), paperRecord)
		// }
		// if req.Paperkind == model.HOMEWORK_COLLECTION_NAME {
		// 	record := new(model.PaperRecord)
		// 	gdao.FillPaperRecordByUserIDAndPaperId(ctx, bson.ObjectIdHex(req.Userid), paper.ID, getPaperRecordTableName(req.Paperkind), record)
		// 	gdao.UpdatePaperRecord(ctx, getPaperRecordTableName(req.Paperkind), record.ID, paperRecordStaus, time)
		// }
		// completeStudents = homeWork.CompleteStudents
		// completeStudents = append(completeStudents, req.Userid)
	}
	// gdao.UpdateCompeleStudent(ctx, bson.ObjectIdHex(req.Homeworkid), completeStudents)
	return nil
}

func getPaperRecordTableName(paperKind string) string {
	return paperKind + "_record"
}

func (s *QuestionService) QueryMyPaperAnswerRecord(ctx context.Context, req *api.QueryPaperAnswerRecordReq, rsp *api.QueryPaperAnswerRecordReply) error {
	logrus.Infof("QueryMyPaperAnswerRecord req is %v ", req)
	if req.Paperkind == "question" {
		record := new(model.AnswerRecord)
		if err := gdao.FillAnserRecord(ctx, bson.ObjectIdHex(req.Paperrecordid), record); err != nil {
			return fmt.Errorf("查找答题%v的记录失败%v", req.Paperrecordid, err)
		}
		questionTitleItems := []*api.TitleItems{&api.TitleItems{Name: "单选题"}, &api.TitleItems{Name: "多选题"}, &api.TitleItems{Name: "判断题"}, &api.TitleItems{Name: "填空题"}, &api.TitleItems{Name: "简答题"}, &api.TitleItems{Name: "图片题"}, &api.TitleItems{Name: "文件题"}}
		question := new(model.Question)
		if err := gdao.FillQuestionById(ctx, record.QuestionID, question); err != nil {
			logrus.Errorln("QueryMyPaperAnswerRecord 查找题目 %v err %v", record.QuestionID, err)
			return fmt.Errorf("查找答题%v的题目%v失败%v", req.Paperrecordid, record.QuestionID, err)
		}
		shiftQuestionsWithAnswerInfo(question, &questionTitleItems, record.Answer, question.NeedCheck, question.Explain, int64(record.Result))
		rsp.Paperrecord = questionTitleItems
		return nil
	}

	paperRecord := new(model.PaperRecord)
	err := gdao.FillPaperRecordByID(ctx, bson.ObjectIdHex(req.Paperrecordid), getPaperRecordTableName(req.Paperkind), paperRecord)
	if err != nil {
		logrus.Errorf("QueryMyPaperAnswerRecord query paperrecord %v error %v \n", req.Paperrecordid, err)
		return fmt.Errorf("查找失败")
	} else {
		paper := new(model.Paper)
		err := gdao.FillPaperById(ctx, req.Paperkind, bson.ObjectIdHex(paperRecord.PaperId), paper)
		if err != nil {
			logrus.Errorf("QueryMyPaperAnswerRecord query paper %v error %v\n", paperRecord.PaperId, err)
			return fmt.Errorf("查找失败")
		} else {
			paperNeedCheck := false
			paperGetScore := int64(0)
			records := new([]*model.AnswerRecord)
			// if err := gdao.FillAnserRecordById(ctx, "paperid", req.Paperid, bson.ObjectIdHex(req.Userid), bson.ObjectIdHex(v.Questionid), records); err != nil {
			if err := gdao.FillAnserRecordByPaperRecordId(ctx, req.Paperkind, bson.ObjectIdHex(req.Paperrecordid), bson.ObjectIdHex(req.Userid), records); err != nil {
				// return fmt.Errorf("QueryMyPaperAnswerRecord: 查找试卷%v 题目 %v 的答题记录失败", req.Paperid, v.Questionid)
				return fmt.Errorf("查找试卷类型%v的答卷记录%v的题目失败%v", req.Paperkind, req.Paperrecordid, err)
			}
			questionTitleItems := []*api.TitleItems{&api.TitleItems{Name: "单选题"}, &api.TitleItems{Name: "多选题"}, &api.TitleItems{Name: "判断题"}, &api.TitleItems{Name: "填空题"}, &api.TitleItems{Name: "简答题"}, &api.TitleItems{Name: "图片题"}, &api.TitleItems{Name: "文件题"}}
			for _, v := range paper.Questions {
				question := new(model.Question)
				if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v.Questionid), question); err != nil {
					logrus.Errorln("QueryMyPaperAnswerRecord 查找题目 %v err %v", v.Questionid, err)
					continue
				}
				record := new(model.AnswerRecord)
				err := gdao.FillAnserRecordBySomeIds(ctx, req.Paperkind, bson.ObjectIdHex(req.Paperrecordid), bson.ObjectIdHex(req.Userid), bson.ObjectIdHex(v.Questionid), record)
				if err == nil {
					paperGetScore = paperGetScore + record.Score
					if record.Result == model.WAITTINGCHECK {
						paperNeedCheck = true
					}
				}
				shiftQuestionsWithAnswerInfo(question, &questionTitleItems, record.Answer, question.NeedCheck, question.Explain, int64(record.Result))
			}
			rsp.Paperrecord = questionTitleItems
			// for _, i := range *records {
			// 	question := new(model.Question)
			// 	if err := gdao.FillQuestionById(ctx, i.QuestionID, question); err != nil {
			// 		logrus.Errorf("QueryMyPaperAnswerRecord 查找题目 %v err %v", i.QuestionID.Hex(), err)
			// 		continue
			// 	}
			// 	paperGetScore = paperGetScore + i.Score
			// 	if i.Result == model.WAITTINGCHECK {
			// 		paperNeedCheck = true
			// 	}
			// 	shiftQuestionsWithAnswerInfo(question, &questionTitleItems, i.Answer, question.NeedCheck, question.Explain, int64(i.Result))
			// 	rsp.Paperrecord = questionTitleItems
			// }

			rsp.Paperid = paper.ID.Hex()
			rsp.Papername = paper.Name
			rsp.Score = paper.Totalscore
			rsp.Getscore = paperGetScore
			rsp.Needcheck = paperNeedCheck
		}
	}
	// paper := new(model.Paper)
	// err := gdao.FillPaperById(ctx, req.Paperkind, bson.ObjectIdHex(req.Paperid), paper)
	// if err != nil {
	// 	fmt.Println("QueryMyPaperAnswerRecord query paper error", err)
	// 	return fmt.Errorf("查找失败")
	// } else {
	// 	paperNeedCheck := false
	// 	paperGetScore := int64(0)
	// 	records := new([]*model.AnswerRecord)
	// 	// if err := gdao.FillAnserRecordById(ctx, "paperid", req.Paperid, bson.ObjectIdHex(req.Userid), bson.ObjectIdHex(v.Questionid), records); err != nil {
	// 	if err := gdao.FillAnserRecordById(ctx, req.Paperkind, req.Paperid, bson.ObjectIdHex(req.Userid), records); err != nil {
	// 		// return fmt.Errorf("QueryMyPaperAnswerRecord: 查找试卷%v 题目 %v 的答题记录失败", req.Paperid, v.Questionid)
	// 		return fmt.Errorf("QueryMyPaperAnswerRecord: 查找试卷类型%v的%v 题目err", req.Paperkind, req.Paperid, err)
	// 	}
	// 	for _, i := range *records {
	// 		question := new(model.Question)
	// 		if err := gdao.FillQuestionById(ctx, i.QuestionID, question); err != nil {
	// 			fmt.Println("QueryMyPaperAnswerRecord 查找题目 %v err %v", i.QuestionID.Hex(), err)
	// 			continue
	// 		}
	// 		paperGetScore = paperGetScore + i.Score
	// 		NEEDCHECK := false
	// 		if i.Result == model.WAITTINGCHECK {
	// 			NEEDCHECK = true
	// 			paperNeedCheck = true
	// 		}
	// 		oneQuestionAnswerRecord := &api.UserPaperAnswerInfo{i.QuestionID.Hex(), string(question.Kind), question.Name, question.Content, question.QImage, question.Option, question.OImage, question.Answer, question.AImage, i.Answer, i.AImage, NEEDCHECK, question.Explain, question.Star, string(i.Result)}
	// 		rsp.Paperrecord = append(rsp.Paperrecord, oneQuestionAnswerRecord)
	// 	}

	return nil
}

func (s *QuestionService) QueryPaperInClass(ctx context.Context, req *api.QueryPaperInClassReq, rsp *api.QueryPaperInClassReply) error {
	logrus.Infof("QueryPaperInClass req is %v ", req)
	papers := new([]*model.Paper)
	err, allCount := gdao.FillPaperByClassId(ctx, req.Paperkind, req.Classid, papers, req.Page, PageCount)
	if err != nil {
		fmt.Println("query paper error", err)
		return fmt.Errorf("查找试卷失败")
	} else {
		for _, v := range *papers {
			rsp.Paper = append(rsp.Paper, &api.PaperInfo{Paperid: v.ID.Hex(), Name: v.Name, Score: v.Totalscore, Suggesttime: v.SuggestTime})
		}
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	fmt.Println(allCount, rsp.Currentpage, rsp.Totalpage)
	return nil
}

func convertQuestionKind(kind model.QUESTION_KIND) string {
	switch kind {
	case model.SIMPLE_CHOICE_QUESTION:
		return "单选题"
	case model.MULTIPLE_CHOICE_QUESTION:
		return "多选题"
	case model.JUDGMENT_QUESTION:
		return "判断题"
	case model.BLACK_FILL_QUESTION:
		return "填空题"
	case model.ESSAY_QUESTION:
		return "简答题"
	case model.PICTURE_QUESTION:
		return "图片题"
	case model.DOCUMENT_QUESTION:
		return "文件题"
	default:
		return ""
	}
}

func shiftModelOptions2Proto(modelOptions *[]model.Option) []*api.Option {
	options := []*api.Option{}
	for _, i := range *modelOptions {
		options = append(options, &api.Option{i.Prefix, i.Content, i.IsImg})
	}
	return options
}

// func shiftModelAnswers2Proto(modelAnswers *[]model.Option) []*api.Option {
// 	answers := []*api.Option{}
// 	for _, i := range *modelAnswers {
// 		answers = append(answers, &api.Option{i.Prefix, i.Content, i.IsImg})
// 	}
// 	return answers
// }

func shiftProto2ModelOptions(protoOptions []*api.Option) []model.Option {
	options := []model.Option{}
	for _, i := range protoOptions {
		options = append(options, model.Option{i.Prefix, i.Content, i.Iscontentimg})
	}
	return options
}

// func shiftProto2ModelAnswers(protoAnswers []*api.Option) []model.Option {
// 	answers := []model.Option{}
// 	for _, i := range protoAnswers {
// 		answers = append(answers, model.Option{i.Prefix, i.Content, i.Iscontentimg})
// 	}
// 	return answers
// }
func shiftQuestionsWithoutAnswerInfo(question *model.Question, items *[]*api.TitleItems, needCheck bool) {
	options := shiftModelOptions2Proto(&question.Options)
	// &api.TitleItems{Name: "单选题"},&api.TitleItems{Name: "多选题"},&api.TitleItems{Name: "判断题"},&api.TitleItems{Name: "填空题",&api.TitleItems{Name: "简答题"},&api.TitleItems{Name: "图片题"},&api.TitleItems{Name: "文件题"}
	switch question.Kind {
	case model.SIMPLE_CHOICE_QUESTION:
		(*items)[0].Name = "单选题"
		(*items)[0].Questionitems = append((*items)[0].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, nil, []string{}, needCheck, "", 4})
	case model.MULTIPLE_CHOICE_QUESTION:
		(*items)[1].Name = "多选题"
		(*items)[1].Questionitems = append((*items)[1].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, nil, []string{}, needCheck, "", 4})
	case model.JUDGMENT_QUESTION:
		(*items)[2].Name = "判断题"
		(*items)[2].Questionitems = append((*items)[2].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, nil, []string{}, needCheck, "", 4})
	case model.BLACK_FILL_QUESTION:
		(*items)[3].Name = "填空题"
		(*items)[3].Questionitems = append((*items)[3].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, nil, []string{}, needCheck, "", 4})
	case model.ESSAY_QUESTION:
		(*items)[4].Name = "简答题"
		(*items)[4].Questionitems = append((*items)[4].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, nil, []string{}, needCheck, "", 4})
	case model.PICTURE_QUESTION:
		(*items)[5].Name = "图片题"
		(*items)[5].Questionitems = append((*items)[5].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, nil, []string{}, needCheck, "", 4})
	case model.DOCUMENT_QUESTION:
		(*items)[6].Name = "文件题"
		(*items)[6].Questionitems = append((*items)[6].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, nil, []string{}, needCheck, "", 4})
	default:
		return
	}
}

func shiftQuestionsWithAnswerInfo(question *model.Question, items *[]*api.TitleItems, getAnswers []string, needCheck bool, explain string, result int64) {
	options := shiftModelOptions2Proto(&question.Options)
	rightOptions := shiftModelOptions2Proto(&question.Answers)
	switch question.Kind {
	case model.SIMPLE_CHOICE_QUESTION:
		(*items)[0].Name = "单选题"
		(*items)[0].Questionitems = append((*items)[0].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result})
	case model.MULTIPLE_CHOICE_QUESTION:
		(*items)[1].Name = "多选题"
		(*items)[1].Questionitems = append((*items)[1].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result})
	case model.JUDGMENT_QUESTION:
		(*items)[2].Name = "判断题"
		(*items)[2].Questionitems = append((*items)[2].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result})
	case model.BLACK_FILL_QUESTION:
		(*items)[3].Name = "填空题"
		(*items)[3].Questionitems = append((*items)[3].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result})
	case model.ESSAY_QUESTION:
		(*items)[4].Name = "简答题"
		(*items)[4].Questionitems = append((*items)[4].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result})
	case model.PICTURE_QUESTION:
		(*items)[5].Name = "图片题"
		(*items)[5].Questionitems = append((*items)[5].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result})
	case model.DOCUMENT_QUESTION:
		(*items)[6].Name = "文件题"
		(*items)[6].Questionitems = append((*items)[6].Questionitems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.IsQImg, options, rightOptions, getAnswers, needCheck, explain, result})
	default:
		return
	}
}

func (s *QuestionService) QueryPaperQuestions(ctx context.Context, req *api.QueryPaperQuestionsReq, rsp *api.QueryPaperQuestionsReply) error {
	logrus.Infof("QueryPaperQuestions req is %v ", req)
	paper := new(model.Paper)
	err := gdao.FillPaperById(ctx, req.Paperkind, bson.ObjectIdHex(req.Paperid), paper)
	if err != nil {
		fmt.Println("QueryPaperQuestions query paper %v error", req.Paperid, err)
		return fmt.Errorf("查找试卷失败")
	} else {
		questionTitleItems := []*api.TitleItems{&api.TitleItems{Name: "单选题"}, &api.TitleItems{Name: "多选题"}, &api.TitleItems{Name: "判断题"}, &api.TitleItems{Name: "填空题"}, &api.TitleItems{Name: "简答题"}, &api.TitleItems{Name: "图片题"}, &api.TitleItems{Name: "文件题"}}
		for _, v := range paper.Questions {
			question := new(model.Question)
			if err := gdao.FillQuestionById(ctx, bson.ObjectIdHex(v.Questionid), question); err != nil {
				logrus.Errorln("QueryMyPaperAnswerRecord 查找题目 %v err %v", v.Questionid, err)
				continue
			}
			shiftQuestionsWithoutAnswerInfo(question, &questionTitleItems, question.NeedCheck)
			// options := []*api.Option{}
			// for _, i := range question.Options {
			// 	options = append(options, &api.Option{i.Prefix, i.Content, i.IsImage})
			// }
			// if question.Kind == model.SIMPLE_CHOICE_QUESTION {
			// 	simpleChoice.QuestionItems = append(simpleChoice.QuestionItems, &api.QuestionItems{question.ID.Hex(), int64(question.Kind), question.Content, question.QImage, options})
			// }

			// rsp.TitleItems = append(rsp.TitleItems, &api.TitleItems{convertQuestionKind(question.Kind), string(question.Kind), question.Content, question.QImage, question.OImage, question.Option, v.Score, question.NeedCheck})
		}
		rsp.Titleitems = questionTitleItems
	}
	rsp.Name = paper.Name
	rsp.Score = paper.Totalscore
	rsp.Suggesttime = paper.SuggestTime
	return nil
}

func (s *QuestionService) QueryMyPaperAnswerRecordList(ctx context.Context, req *api.QueryPaperAnswerRecordListReq, rsp *api.QueryPaperAnswerRecordListReply) error {
	logrus.Infof("QueryMyPaperAnswerRecordList req is %v ", req)
	if req.Paperkind == "question" {
		records := new([]*model.AnswerRecord)
		err, allCount := gdao.FillSimpleAnserRecordInPage(ctx, "", bson.ObjectIdHex(SIMPLE_PAPER_RECORD_ID), bson.ObjectIdHex(req.Userid), PageCount, int(req.Page), records)
		if err != nil {
			return fmt.Errorf("查找单答卷记录失败%v", err)
		}
		fmt.Println(allCount)
		for _, v := range *records {
			fmt.Println("进来了")
			question := new(model.Question)
			err := gdao.FillQuestionById(ctx, v.QuestionID, question)
			if err != nil {
				fmt.Println("查不到题目")
				logrus.Errorf("QueryMyPaperAnswerRecordList query question %v error %v", v.QuestionID, err)
				continue
			}
			rsp.Recordlist = append(rsp.Recordlist, &api.QueryPaperAnswerRecordListReply_OneRecord{v.ID.Hex(), "", question.Course, question.CreaterName, int64(model.RECORD_DONE), v.CreateTime})
		}
		rsp.Currentpage = req.Page
		rsp.Totalpage = int64(allCount / PageCount)
		return nil
	}

	recordList := new([]*model.PaperRecord)
	err, allCount := gdao.FillPaperRecordList(ctx, bson.ObjectIdHex(req.Userid), getPaperRecordTableName(req.Paperkind), recordList, req.Page, PageCount)
	if err != nil {
		logrus.Errorln("query recordlist error", err)
		return fmt.Errorf("查找答卷记录失败")
	} else {
		for _, v := range *recordList {
			paper := new(model.Paper)
			err := gdao.FillPaperById(ctx, req.Paperkind, bson.ObjectIdHex(v.PaperId), paper)
			if err != nil {
				fmt.Println("QueryMyPaperAnswerRecordList query paper error", err)
				continue
			}
			courseName := ""
			teacherName := "匿名"
			if paper.ClassId != "" {
				classReq := new(capi.ClassReq)
				classReq.Classid = paper.ClassId
				if res, err := classSrv.ClassInfo(context.TODO(), classReq); err != nil {
					logrus.Error("QueryMyPaperAnswerRecordList query paper classid error %v", err)
				} else {
					courseName = res.GetCourse()
					teacherName = res.GetTeachername()
				}
			}
			rsp.Recordlist = append(rsp.Recordlist, &api.QueryPaperAnswerRecordListReply_OneRecord{v.ID.Hex(), paper.Name, courseName, teacherName, int64(v.Status), v.CreateTime})
		}
	}
	rsp.Currentpage = req.Page
	rsp.Totalpage = int64(allCount / PageCount)
	fmt.Println(allCount, rsp.Currentpage, rsp.Totalpage)
	return nil
}
