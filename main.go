package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gprc-template/internal/config"
	gRPC "gprc-template/internal/grpc"
	"gprc-template/internal/service"
	"gprc-template/pkg/pb/generated/api"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println(os.Getenv(config.Env))
	if os.Getenv(config.Env) == "" {
		configError := config.ReadConfig(viper.GetString(config.Env))
		if configError != nil {
			log.Panicf("error in building configuration: %v\n", configError)
		}
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt(config.GRPCPort)))
	if err != nil {
		log.Panicln(err)
	}
	var opts []grpc.ServerOption
	server := gRPC.CreateServerWithLogV1(opts)

	api.RegisterSearchServer(server, service.NewSearchImpl())

	go func() {
		if err := server.Serve(listen); err != http.ErrServerClosed {
			log.Fatalf("failed to setup gRPC connection: %v", err)
		}
		log.Println("Server started")
	}()

	shutdown(server, listen)
}

func shutdown(server *grpc.Server,  l net.Listener) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-stop

	log.Println("Shutting down server")
	defer func() {
		_ = l.Close()
	}()
	server.GracefulStop()
}

