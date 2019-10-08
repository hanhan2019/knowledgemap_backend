package main

import (
	"fmt"

	"knowledgemap_backend/microservices/common/namespace"
	"knowledgemap_backend/microservices/knowledgemap/passport/api"
	"knowledgemap_backend/microservices/knowledgemap/passport/handler"
	"myProjects/collegeManage/app/common/conf"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name(namespace.GetName("knowledgemap.microservices.knowledgemap.passport")),
	)
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	conf.Init()
	handler.Init()

	api.RegisterPassportHandler(service.Server(), new(handler.PassportService))
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
