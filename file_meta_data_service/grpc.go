package main

import (
	"filemetadata-service/internals/adapters/handlers/grpc_handler"
	"filemetadata-service/internals/adapters/repository"
	"filemetadata-service/internals/core/services"
	uploaddataservice "grpc-codes/upload_data"

	"google.golang.org/grpc"
)

func registerGRPCServices(server *grpc.Server) {

	dynamoDB := repository.NewDynamoDbRepository()

	// part upload data handler & service
	uds := services.NewUploadDataService(dynamoDB)
	udh := grpc_handler.NewUploadDataHandler(uds)

	// register services
	uploaddataservice.RegisterUploadDataServer(server, udh)
}
