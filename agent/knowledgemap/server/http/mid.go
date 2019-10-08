package http

import (
	"context"
	"net/http"

	"myProjects/collegeManage/app/common/conf"

	"myProjects/collegeManage/agent/college/server/http/comm"

	"myProjects/collegeManage/app/college/passport/api"

	"github.com/labstack/echo"
)

func CreateAuthMid(passSrv api.PassportService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if conf.IsDebugEnv() {
				return next(c)
			}
			uid := c.Request().Header.Get("auth-uid")
			token := c.Request().Header.Get("auth-session")
			if _, err := passSrv.CheckSToken(context.TODO(), &api.SessionTokenReq{Uid: uid, Stoken: token}); err != nil {
				return c.JSON(http.StatusBadRequest, comm.Err("", comm.STATUS_NEED_LOGIN))
			}
			return next(c)
		}
	}
}
