package services

import "filemetadata-service/internals/core/ports"

type UploadDataService struct {
	Repo ports.FileMetaDataDBRepository
}

func NewUploadDataService(repo ports.FileMetaDataDBRepository) *UploadDataService {
	return &UploadDataService{repo}
}
