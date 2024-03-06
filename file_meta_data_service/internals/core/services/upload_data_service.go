package services

import (
	"filemetadata-service/internals/core/domain"
	"filemetadata-service/internals/core/ports"
	uploaddataservice "grpc-codes/upload_data"
	"time"
)

type UploadDataService struct {
	Repo ports.FileMetaDataDBRepository
}

func NewUploadDataService(repo ports.FileMetaDataDBRepository) *UploadDataService {
	return &UploadDataService{repo}
}

func (uds UploadDataService) FetchUploadData(req *uploaddataservice.UploadDataRequest) []domain.UploadData {

	var uploadData []domain.UploadData

	uds.Repo.Table("upload_data").Find([]string{
		"uploadId",
		req.UploadId,
	}, &uploadData)

	return uploadData
}

func (uds UploadDataService) CreateUploadData(req *uploaddataservice.SaveUploadDataRequest) (any, error) {

	data, err := uds.Repo.Table("upload_data").Create(domain.UploadData{
		UserId:    req.UserId,
		UploadId:  req.UploadId,
		PartNum:   req.PartNum,
		Etag:      req.Etag,
		CreatedAt: time.Now(),
	})

	return data, err
}

func (uds UploadDataService) DeleteUploadData(req *uploaddataservice.DeleteUploadDataRequest) error {

	err := uds.Repo.Table("upload_data").Where([][]any{
		{"userId", req.UserId},
		{"uploadId", req.UploadId},
	}).Delete()

	return err
}
