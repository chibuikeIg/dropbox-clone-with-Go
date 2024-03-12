package grpc_handler

import (
	"context"
	"encoding/json"
	"filemetadata-service/internals/core/services"
	uploaddataservice "grpc-codes/upload_data"
	"net/http"
)

type UploadDataHandler struct {
	uploaddataservice.UnimplementedUploadDataServer
	uds *services.UploadDataService
}

func NewUploadDataHandler(uds *services.UploadDataService) *UploadDataHandler {
	return &UploadDataHandler{uds: uds}
}

func (udh UploadDataHandler) FetchUploadData(ctx context.Context, req *uploaddataservice.UploadDataRequest) (*uploaddataservice.UploadDataResponse, error) {

	data := udh.uds.FetchUploadData(req)

	upload_data, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return &uploaddataservice.UploadDataResponse{Data: string(upload_data), StatusCode: http.StatusOK}, nil
}

func (udh UploadDataHandler) SaveUploadData(ctx context.Context, req *uploaddataservice.SaveUploadDataRequest) (*uploaddataservice.SaveUploadDataResponse, error) {

	_, err := udh.uds.CreateUploadData(req)

	if err != nil {
		return &uploaddataservice.SaveUploadDataResponse{Message: "unable to save upload data", StatusCode: http.StatusInternalServerError}, err
	}

	return &uploaddataservice.SaveUploadDataResponse{Message: "saved upload data", StatusCode: http.StatusOK}, nil
}

func (udh UploadDataHandler) DeleteUploadData(ctx context.Context, req *uploaddataservice.DeleteUploadDataRequest) (*uploaddataservice.DeleteUploadDataResponse, error) {

	err := udh.uds.DeleteUploadData(req)

	if err != nil {
		return nil, err
	}

	return &uploaddataservice.DeleteUploadDataResponse{Message: "upload data removed", StatusCode: http.StatusOK}, nil
}
