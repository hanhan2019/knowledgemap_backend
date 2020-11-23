package http

import (
	"context"
	"fmt"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	papi "knowledgemap_backend/microservices/knowledgemap/passport/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"net/http"

	"github.com/labstack/echo"
)

func userInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	clog.Info("getUserInfo")

	// userId := c.Request().Header.Get("auth-uid")
	userId := c.Get("userId").(string)
	// fmt.Println(userId, c.Get("userType"))
	response, err := userSrv.UserInfo(context.TODO(), &uapi.UserReq{
		Userid:   userId,
		Identify: uapi.Identify(c.Get("userType").(int64)),
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
	//fmt.Println("req is", req.Rtype, req.Name, req.Major, req.College, req.Sex, req.Account, req.Password)
	if res, err := passportSrv.Register(context.TODO(), req); err != nil {
		clog.Errorf("call register %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	// return c.JSON(http.StatusOK, comm.Data("response"))
}

func userLogin(c echo.Context) error {
	fmt.Println("进login")
	clog := middlewares.Log(c)
	req := new(papi.LoginReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	fmt.Println(req.Usertype, req.Password, req.Account)
	if res, err := passportSrv.Login(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		if res.User.Userid != "" {
			return c.JSON(http.StatusOK, comm.Data(res))
		}
		return c.JSON(http.StatusOK, comm.Err("need register", comm.STATUS_PHONE_NOT_REGISTER))
	}
}

func userChangePassword(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(papi.ChangePasswordReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	req.Userid = c.Get("userId").(string)
	req.Identify = uapi.Identify(c.Get("userType").(int64))
	if res, err := passportSrv.ChangePassword(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func userChangeInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(papi.ChangeUserInfoReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	// req.Userid = c.Request().Header.Get("auth-uid")
	// userType, _ := strconv.ParseInt(c.Request().Header.Get("auth-type"), 10, 64)
	req.Userid = c.Get("userId").(string)
	req.Usertype = uapi.Identify(c.Get("userType").(int64))
	if res, err := passportSrv.ChangeUserInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}
