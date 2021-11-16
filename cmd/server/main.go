package main

import (
	v1 "github.com/jxlwqq/todo/api/proto/v1"
	"github.com/jxlwqq/todo/internal/config"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
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

	todoService, _ := InitTodoService(conf.DSN)

	server := grpc.NewServer()
	reflection.Register(server)

	v1.RegisterTodoServiceServer(server, todoService)
	lis, err := net.Listen("tcp", conf.GRPC.Port)

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
