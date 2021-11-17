package main

import (
	v1 "github.com/jxlwqq/todo/api/todo/v1"
	"github.com/jxlwqq/todo/internal/config"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var flagConfig = flag.String("config", "./configs/config.yaml", "path to the config file")

func main() {
	flag.Parse()
	conf, err := config.Load(*flagConfig)

	if err != nil {
		log.Fatal(err)
	}

	todoServer, _ := InitTodoServer(conf.DSN)
	healthServer := health.NewServer()

	server := grpc.NewServer()
	reflection.Register(server)

	v1.RegisterTodoServer(server, todoServer)
	grpc_health_v1.RegisterHealthServer(server, healthServer)

	lis, err := net.Listen("tcp", conf.GRPC.Port)
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
