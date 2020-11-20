package http

import (
	"context"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/knowledgemap/passport/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

const (
	POSITION_STUDENT = 0
	POSITION_TEACHER = 1
)

func CreateAuthMid(passSrv api.PassportService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// if conf.IsDebugEnv() {
			// 	return next(c)
			// }
			uid := c.Request().Header.Get("auth-uid")
			token := c.Request().Header.Get("auth-session")
			userType, _ := strconv.ParseInt(c.Request().Header.Get("auth-type"), 10, 64)
			_, err := passSrv.CheckSToken(context.TODO(), &api.SessionTokenReq{Uid: uid, Stoken: token, Type: api.Indentify(userType)})
			if err != nil {
				return c.JSON(http.StatusBadRequest, comm.Err("请先登陆", comm.STATUS_NEED_LOGIN))
			}
			// c.Set("userName", rsp.User.Username)
			return next(c)
		}
	}
}

func CreateMustPositionMid(passSrv api.PassportService, position int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// if conf.IsDebugEnv() {
			// 	return next(c)
			// }
			uid := c.Request().Header.Get("auth-uid")
			if resp, err := passSrv.CheckIndentify(context.TODO(), &uapi.UserReq{Userid: uid}); err != nil {
				return c.JSON(http.StatusBadRequest, comm.Err("请先登陆", comm.STATUS_NEED_LOGIN))
			} else if resp != nil {
				if int(resp.Ltype) == position {
					return next(c)
				}
			}
			return c.JSON(http.StatusBadRequest, comm.Err("身份认证不通过，您无此操作权限", comm.STATUS_NEED_LOGIN))
		}
	}
}

func CorsMid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))                                 // 设置跨域请求
		c.Response().Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization") // 设置跨域请求
		// c.Response().Header().Set("Content-Type", "application/json")                                     // 设置跨域请求
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,HEAD") // 设置跨域请求
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("MaxAge", "86400")
		return next(c)
	}
}
