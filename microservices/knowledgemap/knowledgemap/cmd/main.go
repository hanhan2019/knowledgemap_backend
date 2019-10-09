package main

import (
	"fmt"

	"knowledgemap_backend/microservices/common/namespace"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/handler"
	"myProjects/collegeManage/app/common/conf"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name(namespace.GetName("knowledgemap.microservices.knowledgemap.knowledgemap")),
	)
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	conf.Init()
	handler.Init()
	dao.InitAllCourseKnowledgeMap()
	api.RegisterPassportHandler(service.Server(), new(handler.PassportService))
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
