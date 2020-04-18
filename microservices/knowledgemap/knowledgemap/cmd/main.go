package main

import (
	"fmt"
	"os"

	"knowledgemap_backend/microservices/common/conf"
	"knowledgemap_backend/microservices/common/namespace"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/api"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/handler"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	addr := "127.0.0.1:8500"
	profile := os.Getenv("profile")
	if profile != "debug" {
		addr = "172.17.9.156:8500"
	}
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{addr}
	})
	service := micro.NewService(micro.Registry(reg), micro.Name(namespace.GetName("microservices.knowledgemap.knowledgemap")))
	// service := micro.NewService(micro.Name(namespace.GetName("microservices.knowledgemap.knowledgemap")))
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	conf.Init()
	handler.Init()

	api.RegisterKnowledegeMapHandler(service.Server(), new(handler.KnowledgeMapService))
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
