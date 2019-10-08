package http

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"

	"knowledgemap_backend/microservices/common/conf"

	"knowledgemap_backend/microservices/common/middlewares"

	"knowledgemap_backend/microservices/knowledgemap/passport/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/micro/go-micro/client"
	"github.com/sirupsen/logrus"
)

var (
	passSrv api.PassportService
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func Init() *echo.Echo {
	//init middlewares

	e := echo.New()
	e.Use(middleware.RequestID(), middleware.CORS() /* cusmid.CreateAccessLogMid() */)
	e.Validator = &CustomValidator{validator.New()}
	//context log
	{
		contextlog := logrus.New()
		contextlog.SetFormatter(&logrus.JSONFormatter{})
		e.Use(middlewares.CreateContextLogMid(contextlog))
	}
	if conf.IsDebugEnv() {
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			fmt.Printf("%s\n", resBody)
		}))
	}

	passSrv = api.NewPassportService("collegemanage.app.college.passport", client.DefaultClient)
	// userSrv = uapi.NewUserService("collegemanage.app.college.user", client.DefaultClient)
	// courseSrv = capi.NewCourseService("collegemanage.app.college.course", client.DefaultClient)
	InitRouter(e)
	/*
		//using gomcro web
			service := web.NewService(
				web.Name("go.micro.api.api.user"),
			)
			service.Init()
			service.Handle("/", e)
			go service.Run()
	*/

	go e.Start(viper.GetString("web.listenaddr"))
	return e
}

func InitRouter(e *echo.Echo) {

	// api := e.Group("/api")

	// api.POST("/user/register", userRegister)
	// api.PUT("/user/login", userLogin)

	// authMid := CreateAuthMid(passSrv)
	// api.GET("/user/classmate/:cid", getClassMates, authMid)
	// api.GET("/user/allcourse/:uid/:major", getAllCourseInfo, authMid)
	// api.GET("/user/teacherallcourse/:uid", getTeacherAllCourseInfo, authMid)
	// api.GET("/user/teacherallstudents/:uid", getTeacherAllStudent, authMid)
	// api.GET("/user/allelectivedcourse/:uid", getElectivedCourseInfo, authMid)
	// api.PUT("/user/givegrade", giveGrade, authMid)
	// api.POST("/user/addcourse", addCourse)
	// api.PUT("/user/choosecourse", chooseCourse, authMid)
	// api.PUT("/user/deletecourse", deleteCourse, authMid)
}
