package main

import (
	"filemetadata-service/internals/app"
	consul_api "filemetadata-service/internals/app/consul"
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("error loading .env file. error:" + err.Error())
	}

	// consul service registration
	consul_api.ServiceRegistration()

	//grpc server
	go runGRPCServer()

	// REST endpoints registration
	r := gin.Default()
	LoadRoutes(r)

	if len(os.Args) >= 2 {
		os.Exit(0)
	}

	r.Run(app.Port())
}

func runGRPCServer() {

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	registerGRPCServices(server)

	address := app.HostName() + ":" + app.Port()

	listener, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}

	if err := server.Serve(listener); err != nil {
		fmt.Println(err)
	}
}
