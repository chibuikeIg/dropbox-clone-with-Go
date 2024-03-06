package grpc_handler

import (
	"context"
	"filemetadata-service/internals/core/services"
	uploaddataservice "grpc-codes/upload_data"
)

type UploadDataHandler struct {
	uploaddataservice.UnimplementedUploadDataServer
	uds *services.UploadDataService
}

func NewUploadDataHandler(uds *services.UploadDataService) *UploadDataHandler {
	return &UploadDataHandler{uds: uds}
}

func (udh UploadDataHandler) FetchUploadData(ctx context.Context, req uploaddataservice.UploadDataRequest) (*uploaddataservice.UploadDataResponse, error) {
	return nil, nil
}

func (udh UploadDataHandler) SaveUploadData(ctx context.Context, req uploaddataservice.SaveUploadDataRequest) (*uploaddataservice.SaveUploadDataResponse, error) {
	return nil, nil
}

func (udh UploadDataHandler) DeleteUploadData(ctx context.Context, req uploaddataservice.DeleteUploadDataRequest) (*uploaddataservice.DeleteUploadDataResponse, error) {
	return nil, nil
}
