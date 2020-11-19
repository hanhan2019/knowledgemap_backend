package http

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	// validator "gopkg.in/go-playground/validator.v8"

	"knowledgemap_backend/microservices/common/conf"
	"knowledgemap_backend/microservices/common/middlewares"
	"knowledgemap_backend/microservices/common/namespace"

	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
	kapi "knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/handler"
	papi "knowledgemap_backend/microservices/knowledgemap/passport/api"
	qapi "knowledgemap_backend/microservices/knowledgemap/question/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	passportSrv     papi.PassportService
	userSrv         uapi.UserService
	questionSrv     qapi.QuestionService
	knowledgeMapSrv kapi.KnowledegeMapService
	classSrv        capi.ClassService
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

	addr := "127.0.0.1:8500"
	profile := os.Getenv("profile")
	if profile != "debug" {
		addr = "172.17.9.156:8500"
	}
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{addr}
	})

	services := micro.NewService(micro.Registry(reg), micro.Name(namespace.GetName("agent.knowledgemap.client")))
	services.Init()
	micro.RegisterHandler(services.Server(), new(handler.KnowledgeMapService))
	go services.Run()
	//questionSrv = qapi.NewQuestionService(namespace.GetName("microservices.knowledgemap.question"), client.DefaultClient)
	knowledgeMapSrv = kapi.NewKnowledegeMapService(namespace.GetName("microservices.knowledgemap.knowledgemap"), client.DefaultClient)
	questionSrv = qapi.NewQuestionService(namespace.GetName("microservices.knowledgemap.question"), services.Client())
	knowledgeMapSrv = kapi.NewKnowledegeMapService(namespace.GetName("microservices.knowledgemap.knowledgemap"), services.Client())
	userSrv = uapi.NewUserService(namespace.GetName("microservices.knowledgemap.user"), services.Client())
	passportSrv = papi.NewPassportService(namespace.GetName("microservices.knowledgemap.passport"), services.Client())
	classSrv = capi.NewClassService(namespace.GetName("microservices.knowledgemap.class"), services.Client())
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

	api := e.Group("/api")
	api.Use(CorsMid)

	api.POST("/user/register", userRegister)
	api.PUT("/user/login", userLogin)

	authMid := CreateAuthMid(passportSrv)
	mustTeacherMid := CreateMustPositionMid(passportSrv, POSITION_TEACHER)
	mustStudentMid := CreateMustPositionMid(passportSrv, POSITION_STUDENT)
	api.PUT("/user/changepassword", userChangePassword, authMid)
	api.PUT("/user/changeinfo", userChangeInfo, authMid)
	api.GET("/hi", hello)
	api.GET("/user/knowledgemap/:uid/:subject/:endtime", queryUserKnowledgeMap)
	{
		api.POST("/class/create", classCreate, authMid, mustTeacherMid)
		api.PUT("/class/join", joinClass, authMid, mustStudentMid)
		api.GET("/class/query/myclasses", queryMyClasses, authMid)
		api.GET("/class/query/alluserinclass/:classid", queryAllUserInClass, authMid)
		api.GET("/class/query/classinfo/:classid", queryAllUserInClass, authMid)
		api.GET("/class/query/classes/:college/:subject/:course/:page", searchClasses, authMid)

	}

	// api.PUT("/class/invitation/create", createInvitation, authMid, mustTeacherMid)
	// api.PUT("/class/invitation/drop", dropInvitation, authMid, mustTeacherMid)
	// api.GET("/class/invitation/query/:invitationcode", queryInvitation, authMid)
	{
		api.POST("/knowledge/create", createKnowledge, authMid, mustTeacherMid)
		api.GET("/knowledge/query/:knowledgeId", queryKnowledge, authMid)
	}
	{
		api.POST("/question/create", createQuestion, authMid, mustTeacherMid)
		api.GET("/question/query/:kind/:course/:subject/:knowledge", queryQuestion, authMid)
	}
	{
		api.POST("/homework/create", createHomeWork, authMid, mustTeacherMid)
		api.GET("/homework/query/:userid/:classid", queryHomeWork, authMid)
		api.PUT("/homework/do", doHomeWork, authMid)
		api.GET("/homework/answerrecord/query/:homeworkid", queryAnswerRecord, authMid, mustTeacherMid)
		api.GET("/homework/query/info/:classid", queryHomeWorkInClass, authMid)

	}
	//api.GET("/user/allcourse/:uid/:major", getAllCourseInfo, authMid)
}
