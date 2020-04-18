package http

import (
	"context"
	"fmt"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	kapi "knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	// clog := middlewares.Log(c)
	return c.JSON(http.StatusOK, comm.Data("hello"))
}

func queryUserKnowledgeMap(c echo.Context) error {
	clog := middlewares.Log(c)
	uid := c.Param("uid")
	subject := c.Param("subject")
	endTimeStr := c.Param("endtime")
	endTime, _ := strconv.ParseInt(endTimeStr, 10, 64)
	req := &kapi.CRqQueryMyMapBySubject{
		Uid:     uid,
		Subject: subject,
		Endtime: endTime,
	}
	fmt.Println(req.Uid, req.Subject, req.Endtime)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := knowledgeMapSrv.GetMyKnowledegeMapBySubject(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))

	}
}

func createKnowledge(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(kapi.CreateKnowledegeReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := knowledgeMapSrv.CreateKnowledege(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryKnowledge(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(kapi.QueryKnowledegeInfoReq)
	req.Id = c.Param("knowledgeId")
	if res, err := knowledgeMapSrv.QueryKnowledegeInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}
