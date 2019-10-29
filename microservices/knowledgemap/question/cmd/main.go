package main

import (
	"fmt"

	"knowledgemap_backend/microservices/common/conf"
	"knowledgemap_backend/microservices/common/namespace"
	"knowledgemap_backend/microservices/knowledgemap/question/api"
	"knowledgemap_backend/microservices/knowledgemap/question/handler"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name(namespace.GetName("knowledgemap.microservices.knowledgemap.question")),
	)
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	conf.Init()
	handler.Init()

	api.RegisterQuestionHandler(service.Server(), new(handler.QuestionService))
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
