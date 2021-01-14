package http

import (
	"context"
	"fmt"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	qapi "knowledgemap_backend/microservices/knowledgemap/question/api"
	"strconv"

	"net/http"

	"github.com/labstack/echo"
)

func createQuestion(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.CreateQuestionReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := questionSrv.CreateQuestion(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryQuestion(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryQuestionReq)
	req.Questiontype, _ = strconv.ParseInt(c.QueryParam("questiontype"), 10, 64)
	req.Course = c.QueryParam("course")
	req.Subject = c.QueryParam("subject")
	// req.Knowledge = c.QueryParam("knowledge")
	req.Page, _ = strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if res, err := questionSrv.QueryQuestion(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func doQuestion(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.DoQuestionInfo)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	req.Userid = c.Get("userId").(string)
	req.Username = c.Get("userName").(string)
	if res, err := questionSrv.DoQuestion(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

// func createHomeWork(c echo.Context) error {
// 	clog := middlewares.Log(c)
// 	req := new(qapi.CreateHomeWorkReq)
// 	if err := comm.VBind(c, req); err != nil {
// 		clog.Errorf("参数错误:%v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
// 	}
// 	if res, err := questionSrv.CreateHomeWork(context.TODO(), req); err != nil {
// 		clog.Error("error %v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
// 	} else {
// 		return c.JSON(http.StatusOK, comm.Data(res))
// 	}
// }

// func queryHomeWork(c echo.Context) error {
// 	clog := middlewares.Log(c)
// 	req := new(qapi.QueryMyHomeWorkReq)
// 	req.Userid = c.QueryParam("userid")
// 	req.Classid = c.QueryParam("classid")
// 	if res, err := questionSrv.QueryMyHomeWork(context.TODO(), req); err != nil {
// 		clog.Error("error %v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
// 	} else {
// 		return c.JSON(http.StatusOK, comm.Data(res))
// 	}
// 	return c.JSON(http.StatusOK, comm.Data(nil))
// }

// func doHomeWork(c echo.Context) error {
// 	clog := middlewares.Log(c)
// 	req := new(qapi.DoHomeWorkReq)
// 	if err := comm.VBind(c, req); err != nil {
// 		clog.Errorf("参数错误:%v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
// 	}
// 	if res, err := questionSrv.DoHomeWork(context.TODO(), req); err != nil {
// 		clog.Error("error %v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
// 	} else {
// 		return c.JSON(http.StatusOK, comm.Data(res))
// 	}
// }

// func queryAnswerRecord(c echo.Context) error {
// 	clog := middlewares.Log(c)
// 	req := new(qapi.QueryAnswerRecordReq)
// 	req.Homeworkid = c.QueryParam("homeworkid")
// 	if res, err := questionSrv.QueryAnswerRecord(context.TODO(), req); err != nil {
// 		clog.Error("error %v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
// 	} else {
// 		return c.JSON(http.StatusOK, comm.Data(res))
// 	}
// }

// func queryHomeWorkInClass(c echo.Context) error {
// 	clog := middlewares.Log(c)
// 	req := new(capi.ClassReq)
// 	req.Classid = c.QueryParam("classid")
// 	if res, err := questionSrv.QueryHomeWorkInClass(context.TODO(), req); err != nil {
// 		clog.Error("error %v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
// 	} else {
// 		return c.JSON(http.StatusOK, comm.Data(res))
// 	}
// }

func createPaper(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.CreatePaperReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := questionSrv.CreatePaper(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryPaper(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryPaperInClassReq)
	req.Userid = c.Get("userId").(string)
	req.Classid = c.QueryParam("classid")
	req.Paperkind = c.QueryParam("paperkind")
	req.Page, _ = strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if res, err := questionSrv.QueryPaperInClass(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func queryRecommendPaper(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryRecommendPaperReq)
	req.Userid = c.Get("userId").(string)
	if res, err := questionSrv.QueryRecommendPaper(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func queryPaperQuestion(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryPaperQuestionsReq)
	req.Paperid = c.QueryParam("paperid")
	req.Paperkind = c.QueryParam("paperkind")
	fmt.Println(req.Paperid, req.Paperkind)
	if res, err := questionSrv.QueryPaperQuestions(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func doPaper(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.DoPaperReq)
	a := c.Request()
	body := []byte{}
	a.Body.Read(body)
	fmt.Println("request body is", string(body))
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	req.Userid = c.Get("userId").(string)
	req.Username = c.Get("userName").(string)
	fmt.Println(req)
	fmt.Println(req.Paperid)
	fmt.Println(req.Answer)
	fmt.Println(req.Paperkind)
	if res, err := questionSrv.DoPaper(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryPaperAnswerRecord(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryPaperAnswerRecordReq)
	req.Paperrecordid = c.QueryParam("paperrecordid")
	req.Paperkind = c.QueryParam("paperkind")
	req.Userid = c.Get("userId").(string)
	if res, err := questionSrv.QueryMyPaperAnswerRecord(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryPaperAnswerRecordList(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryPaperAnswerRecordListReq)
	req.Userid = c.Get("userId").(string)
	req.Page, _ = strconv.ParseInt(c.QueryParam("page"), 10, 64)
	req.Paperkind = c.QueryParam("paperkind")
	if res, err := questionSrv.QueryMyPaperAnswerRecordList(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}
