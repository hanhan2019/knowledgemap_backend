package http

import (
	"context"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	papi "knowledgemap_backend/microservices/knowledgemap/passport/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"net/http"

	"github.com/labstack/echo"
)

func getUserInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	clog.Info("getUserInfo")

	uid := c.Param("uid")

	response, err := userSrv.UserInfo(context.TODO(), &uapi.UserReq{
		Userid: uid,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(200, comm.Data(response))

}

func userRegister(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(papi.RegisterReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := passportSrv.Register(context.TODO(), req); err != nil {
		clog.Errorf("call register %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(200, comm.Data("response"))
}

func userLogin(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(papi.LoginReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := passportSrv.Login(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		if res.User.Userid != "" {
			return c.JSON(http.StatusOK, comm.Data(res))
		} else {
			//如果用mcode登录,且未找到用户，尝试注册用户
			return c.JSON(http.StatusOK, comm.Err("need register", comm.STATUS_PHONE_NOT_REGISTER))
		}
	}
}
