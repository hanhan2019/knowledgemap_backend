package http

import (
	"context"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/knowledgemap/passport/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"net/http"

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
			if _, err := passSrv.CheckSToken(context.TODO(), &api.SessionTokenReq{Uid: uid, Stoken: token}); err != nil {
				return c.JSON(http.StatusBadRequest, comm.Err("请先登陆", comm.STATUS_NEED_LOGIN))
			}
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
