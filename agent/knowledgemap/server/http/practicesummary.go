package http

import (
	"context"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	qapi "knowledgemap_backend/microservices/knowledgemap/question/api"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func psCreate(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.CreatePracticeSummaryReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	req.Ownerid = c.Get("userId").(string)
	req.Ownername = c.Get("userName").(string)

	if res, err := questionSrv.CreatePracticeSummary(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryPSInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryPracticeSummaryReq)
	req.Practicesummaryid = c.Param("psid")
	if res, err := questionSrv.QueryPracticeSummaryInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func queryPSDetailInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryPracticeSummaryReq)
	req.Practicesummaryid = c.Param("psid")
	req.Page, _ = strconv.ParseInt(c.Param("page"), 10, 64)
	if res, err := questionSrv.GetPracticeSummary(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func queryMyPSInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.QueryMyPracticeSummaryReq)
	req.Userid = c.Get("userId").(string)
	if res, err := questionSrv.QueryMyPracticeSummary(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func addQuestionInPS(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.ControllQuestionInPSReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := questionSrv.AddQuestionInPS(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func deleteQuestionInPS(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(qapi.ControllQuestionInPSReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := questionSrv.DelteQuestionInPS(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}
