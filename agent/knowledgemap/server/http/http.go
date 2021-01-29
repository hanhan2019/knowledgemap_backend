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

	authMid := CreateAuthMid(passportSrv)
	mustTeacherMid := CreateMustPositionMid(passportSrv, POSITION_TEACHER)
	mustStudentMid := CreateMustPositionMid(passportSrv, POSITION_STUDENT)
	api.GET("/hi", hello)
	api.POST("/file/upload", transport)
	{
		api.POST("/user/register", userRegister)
		api.PUT("/user/login", userLogin)
		api.PUT("/user/changepassword", userChangePassword, authMid)
		api.PUT("/user/changeinfo", userChangeInfo, authMid)
		api.GET("/user/query/info", userInfo, authMid)
	}
	{
		api.POST("/class/create", classCreate, authMid, mustTeacherMid)
		api.PUT("/class/join", joinClass, authMid, mustStudentMid)
		api.GET("/class/query/myclasses", queryMyClasses, authMid)
		api.GET("/class/query/alluserinclass", queryAllUserInClass, authMid)
		api.GET("/class/query/classinfo", queryClassInfo, authMid)
		api.GET("/class/query/classes", searchClasses, authMid)
		api.GET("/class/query/formlist", queryFormList)
		api.POST("/class/deletestudent", deleteStudentInClass, authMid)
		api.GET("/class/query/student", queryStudentInClass, authMid)
		api.POST("/class/delete", deleteClass, authMid)
		api.GET("/class/query/myinfo", queryMyCollegeInfo, authMid)
	}
	// api.PUT("/class/invitation/create", createInvitation, authMid, mustTeacherMid)
	// api.PUT("/class/invitation/drop", dropInvitation, authMid, mustTeacherMid)
	// api.GET("/class/invitation/query/:invitationcode", queryInvitation, authMid)
	{
		api.POST("/practice/create", psCreate, authMid)
		api.GET("/practice/query/psinfo", queryPSInfo, authMid)
		api.GET("/practice/query/psdetailinfo", queryPSDetailInfo, authMid)
		api.GET("/practice/query/mypsinfo", queryMyPSInfo, authMid)
		api.POST("/practice/addquestion", addQuestionInPS, authMid)
		api.POST("/practice/deletequestion", deleteQuestionInPS, authMid)

	}
	{
		api.POST("/knowledge/create", createKnowledge, authMid)
		api.GET("/knowledge/query", queryKnowledge, authMid)
	}
	{
		api.POST("/question/create", createQuestion, authMid)
		api.GET("/question/query", queryQuestion, authMid)
		api.PUT("/question/do", doQuestion, authMid)
		api.GET("/user/knowledgemap", queryUserKnowledgeMap)
	}
	{
		// api.POST("/homework/create", createHomeWork, authMid, mustTeacherMid)
		// api.GET("/homework/query", queryHomeWork, authMid)
		// api.PUT("/homework/do", doHomeWork, authMid)
		// api.GET("/homework/answerrecord/query", queryAnswerRecord, authMid, mustTeacherMid)
		// api.GET("/homework/query/info", queryHomeWorkInClass, authMid)
	}
	{
		api.POST("/paper/create", createPaper, authMid, mustTeacherMid)
		api.POST("/paper/changequestions", changeQuestionsInPaper, authMid, mustTeacherMid)
		api.GET("/paper/query", queryPaper, authMid)                    //根据班级号查试卷
		api.GET("/paper/query/recommend", queryRecommendPaper, authMid) //查待做作业和推荐试卷
		api.GET("/paper/query/questions", queryPaperQuestion, authMid)  //根据试卷号查试题目
		api.PUT("/paper/do", doPaper, authMid)
		api.GET("/paper/answerrecord/query", queryPaperAnswerRecord, authMid)          // 查询具体某次试卷的答题记录
		api.GET("/paper/answerrecord/query/list", queryPaperAnswerRecordList, authMid) //查询我答过哪些试卷

	}
	//api.GET("/user/allcourse/:uid/:major", getAllCourseInfo, authMid)
}
