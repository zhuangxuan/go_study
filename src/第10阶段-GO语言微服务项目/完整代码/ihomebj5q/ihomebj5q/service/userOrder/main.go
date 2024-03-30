package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"ihomebj5q/service/userOrder/handler"

	"github.com/micro/go-micro/registry/consul"
	"ihomebj5q/service/userOrder/model"
	userOrder "ihomebj5q/service/userOrder/proto/userOrder"
)

func main() {
	// New Service
	consulReg := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("go.micro.srv.userOrder"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9986"),
	)

	// Initialise service
	service.Init()
	model.InitDb()

	// Register Handler
	userOrder.RegisterUserOrderHandler(service.Server(), new(handler.UserOrder))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
