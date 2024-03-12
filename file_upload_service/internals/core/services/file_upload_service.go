package services

import (
	"context"
	"encoding/json"
	"errors"
	"file-upload-service/internals/adapters/repository"
	consulapi "file-upload-service/internals/app/consul"
	fr "file-upload-service/internals/app/form-request"
	"file-upload-service/internals/core/domain"
	"file-upload-service/internals/core/ports"
	"fmt"
	uploaddataservice "grpc-codes/upload_data"
	"log"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FileUploadService struct {
	objectStorage   ports.ObjectStorageRepository
	dbRepo, redisDB ports.FileUploadDBRepository
	// grpc upload data client
	grpcUpDClient uploaddataservice.UploadDataClient
}

func NewFileUploadService(objectStorage ports.ObjectStorageRepository, dbRepo ports.FileUploadDBRepository) *FileUploadService {
	redis := repository.NewRedisDBRepository()

	//service discovery
	address, port, err := consulapi.ServiceDiscovery("filemetadata")

	if err != nil {
		log.Fatalf("unable to retreive service, error:%s", err.Error())
	}

	grpcServerAddr := fmt.Sprintf("%s:%v", address, port)
	// dial grpc server and get conn
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to grpc upload data service: %v", err)
	}
	defer conn.Close()

	// create new upload data client
	grpcUpDClient := uploaddataservice.NewUploadDataClient(conn)

	return &FileUploadService{objectStorage: objectStorage, dbRepo: dbRepo, redisDB: redis, grpcUpDClient: grpcUpDClient}
}

func (fus FileUploadService) InitiateMultipartUpload(req fr.FileUploadRequest, userId string) (string, string, error) {

	objectKey := userId + "/"

	if req.FolderName != "" {
		objectKey += req.FolderName + "/"
	}

	objectKey += req.FileName

	uploadID, err := fus.objectStorage.CreateMultipartUpload(map[string]*string{
		"content_type": &req.ContentType,
		"object_key":   &objectKey,
	})

	return uploadID, objectKey, err

}

// improve code this code
func (fus FileUploadService) SavePartUpload(req fr.FilePartUploadRequest, userId string) (string, error) {

	// check if partNum exists on redis using the uploadId
	// as key. randomly generate new part num if previously exists
	// clear the cache db once uploads are completed or aborted

	data, _ := fus.redisDB.Table("uploadId:" + req.UploadId).Get()
	var partNums map[int32]struct{}

	if data != nil {
		err := json.Unmarshal([]byte(data.(string)), &partNums)
		if err != nil {
			log.Println(err)
		}
	}

	partNum := fus.generatePartNum(partNums)

	fileUpload, err := req.File.Open()

	if err != nil {
		return "", err
	}

	defer fileUpload.Close()

	etag, err := fus.objectStorage.UploadPart(map[string]any{
		"part_number": &partNum,
		"object_key":  &req.ObjectKey,
		"upload_id":   &req.UploadId,
		"requestBody": fileUpload,
	})

	if err != nil {
		return "", err
	}

	// create and store partNums in cacheDB
	partNums = map[int32]struct{}{
		partNum: {},
	}

	partNumData, err := json.Marshal(partNums)
	if err != nil {
		log.Println(err)
	}

	_, err = fus.redisDB.Table("uploadId:" + req.UploadId).Create(string(partNumData))
	if err != nil {
		log.Println(err)
	}

	// store part number and etag in db
	_, err = fus.grpcUpDClient.SaveUploadData(context.Background(), &uploaddataservice.SaveUploadDataRequest{
		UserId:   userId,
		UploadId: req.UploadId,
		PartNum:  partNum,
		Etag:     string(etag),
	})

	if err != nil {
		// abort part uploads & return error
		fus.objectStorage.AbortMultipartUpload(map[string]any{
			"object_key": &req.ObjectKey,
			"upload_id":  &req.UploadId,
		})

		// delete from cache db
		fus.redisDB.Table("uploadId:" + req.UploadId).Delete()

		return "", errors.New("error occured while saving upload parts, error: " + err.Error())
	}

	return etag, nil
}

func (fus FileUploadService) CompleteMultiUpload(req fr.CompleteMultiUploadRequest, userId string) (string, error) {

	var uploadParts []domain.UploadPart

	// retreive stored uploaded parts
	fus.dbRepo.Table("upload_parts").Find([]string{
		"uploadId",
		req.UploadId,
	}, &uploadParts)

	if len(uploadParts) == 0 {
		return "", errors.New("records not found")
	}

	completed_parts := make([]map[string]any, len(uploadParts))

	for _, v := range uploadParts {
		completed_parts = append(completed_parts, map[string]any{
			"ETag":       &v.Etag,
			"PartNumber": &v.PartNum,
		})
	}

	response, err := fus.objectStorage.CompleteMultipartUpload(map[string]any{
		"object_key":      &req.ObjectKey,
		"upload_id":       &req.UploadId,
		"completed_parts": completed_parts,
	})

	if err != nil {
		return "", err
	}

	// delete stored part uploads from DB
	fus.dbRepo.Table("upload_parts").Where([][]any{
		{"userId", userId},
		{"uploadId", req.UploadId},
	}).Delete()

	// delete from cache db
	fus.redisDB.Table("uploadId:" + req.UploadId).Delete()

	return response, nil
}

func (fus FileUploadService) AbortMultiUpload(uploadID, objectKey, userId string) error {

	_, err := fus.objectStorage.AbortMultipartUpload(map[string]any{
		"object_key": &objectKey,
		"upload_id":  &uploadID,
	})

	if err != nil {
		return err
	}

	// delete stored part uploads from DB
	fus.dbRepo.Table("upload_parts").Where([][]any{
		{"userId", userId},
		{"uploadId", uploadID},
	}).Delete()

	// delete from cache db
	fus.redisDB.Table("uploadId:" + uploadID).Delete()

	return nil
}

func (fus FileUploadService) generatePartNum(partNums map[int32]struct{}) int32 {

	randNum := int32(rand.Intn(10000-1) + 1)

	if len(partNums) == 0 {
		return randNum
	}

	if _, ok := partNums[randNum]; !ok {
		return randNum
	}

	return fus.generatePartNum(partNums)
}
