package main

import (
	"context"

	"go-micro.dev/v5"
	logger "go-micro.dev/v5/logger"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/registry/consul"
	"go-micro.dev/v5/server"
	"go-micro.dev/v5/server/grpc"

	pb "github.com/robert-pkg/go-micro/my-examples/hello-world/GreeterServer/api"
	log "github.com/robert-pkg/go-micro/my-examples/hello-world/common/log"
)

type Greeter struct{}

func (g *Greeter) SayHello(ctx context.Context, req *pb.SayHelloRequest, rsp *pb.SayHelloResponse) error {

	logger.Info("recv", " name=", req.Name)
	rsp.Greeting = "Hello " + req.Name

	return nil
}

func main() {
	if err := log.Init("", true); err != nil {
		panic(err)
	}

	consulRegistry := consul.NewConsulRegistry(
		registry.Addrs("127.0.0.1:8500"), // consul 地址
	)

	logger.Info("start")

	grpcServer := grpc.NewServer(
		server.Name("GreeterServer"),
		server.Version("1.0.1"),
		server.Registry(consulRegistry),
	)
	srv := micro.NewService(
		micro.Server(grpcServer),
	)

	srv.Init()

	pb.RegisterGreeterServerHandler(srv.Server(), &Greeter{})

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
