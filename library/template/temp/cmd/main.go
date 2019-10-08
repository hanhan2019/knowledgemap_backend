package main

import (
	"fmt"
	"knowledgemap_backend/microservices/common/conf"
	"knowledgemap_backend{{ .PackagePath }}/api"
	"knowledgemap_backend{{ .PackagePath }}/handler"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("college.{{ .PackageName }}"),
	)
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	conf.Init()
	handler.Init()

	api.Register{{.ServiceName}}Handler(service.Server(), new(handler.{{.ServiceName}}Service))
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
