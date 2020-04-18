package http

import (
	"context"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
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
	kind, _ := strconv.ParseInt(c.Param("kind"), 10, 64)
	req.Kind = kind
	req.Course = c.Param("course")
	req.Subject = c.Param("subject")
	req.Knowledge = c.Param("knowledge")
	if res, err := questionSrv.QueryQuestion(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func createHomeWork(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.CreateHomeWorkReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := questionSrv.CreateHomeWork(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryHomeWork(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryMyHomeWorkReq)
	req.Userid = c.Param("userid")
	req.Classid = c.Param("classid")
	if res, err := questionSrv.QueryMyHomeWork(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func doHomeWork(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.DoHomeWorkReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := questionSrv.DoHomeWork(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryAnswerRecord(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryAnswerRecordReq)
	req.Homeworkid = c.Param("homeworkid")
	if res, err := questionSrv.QueryAnswerRecord(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryHomeWorkInClass(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.ClassReq)
	req.Classid = c.Param("classid")
	if res, err := questionSrv.QueryHomeWorkInClass(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}
