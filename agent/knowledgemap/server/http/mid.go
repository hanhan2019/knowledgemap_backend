package http

import (
	"context"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/conf"
	"knowledgemap_backend/microservices/knowledgemap/passport/api"
	"net/http"

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
