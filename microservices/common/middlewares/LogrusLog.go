package middlewares

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const (
	CONTEXT_LOG_KEY = "context_log_key"
)

func Log(c echo.Context) *logrus.Entry {
	return c.Get(CONTEXT_LOG_KEY).(*logrus.Entry)
}

func CreateContextLogMid(log *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			uuid := ""
			cookie, err := c.Cookie("uuid")
			if err == nil {
				uuid = cookie.Value
			}

			contextLogger := log.WithFields(logrus.Fields{
				"type": "context",
				"ip":   c.RealIP(),
				"uuid": uuid,
				"id":   c.Response().Header().Get(echo.HeaderXRequestID),
			})
			c.Set(CONTEXT_LOG_KEY, contextLogger)
			return next(c)
		}
	}

}
