package main

import (
	"context"
	"fmt"
	"github.com/jxlwqq/todo/api/protobuf"
	"github.com/jxlwqq/todo/internal/pkg/config"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	protobuf.RegisterTodoServer(grpcServer, todoServer)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	lis, err := net.Listen("tcp", conf.GRPC.Port)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("grpc server start at: ", conf.GRPC.Port)
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Shutting down the grpcServer...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	grpcServer.GracefulStop()
	<-ctx.Done()
	close(ch)
	fmt.Println("Graceful Shutdown end ")
}
