package main

import (
	"fmt"

	"knowledgemap_backend/microservices/common/conf"
	"knowledgemap_backend/microservices/common/namespace"
	"knowledgemap_backend/microservices/knowledgemap/user/api"
	"knowledgemap_backend/microservices/knowledgemap/user/handler"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(micro.Registry(reg), micro.Name(namespace.GetName("microservices.knowledgemap.user")))
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	conf.Init()
	handler.Init()

	api.RegisterUserHandler(service.Server(), new(handler.UserService))
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
