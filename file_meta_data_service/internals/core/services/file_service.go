package services

import (
	"encoding/json"
	"errors"
	"filemetadata-service/internals/app"
	fr "filemetadata-service/internals/app/form-request"
	"filemetadata-service/internals/core/domain"
	"filemetadata-service/internals/core/ports"
	"log"
	"time"

	"github.com/google/uuid"
)

type FileService struct {
	Repo, CacheDB ports.FileMetaDataDBRepository
}

func NewFileService(repo, cacheDB ports.FileMetaDataDBRepository) *FileService {
	return &FileService{Repo: repo, CacheDB: cacheDB}
}

func (fs FileService) GetFiles(userId string) []domain.File {

	var files []domain.File

	// load from cache and if
	// not available load from DB
	// then save in cache
	data, err := fs.CacheDB.Table("files:" + userId).Get()
	if err != nil {
		log.Println(err)
	}

	if data != nil {

		err = json.Unmarshal([]byte(data.(string)), &files)
		if err != nil {
			log.Println(err)
		}
	}

	if len(files) == 0 {

		fs.Repo.Table("files").Find([]string{
			"userId",
			userId,
		}, &files)

		data, err := json.Marshal(files)
		if err != nil {
			log.Println(err)
		}

		_, err = fs.CacheDB.Table("files:" + userId).Create(string(data))
		if err != nil {
			log.Println(err)
		}
	}

	return files
}

func (fs FileService) SaveFile(request fr.FileRequest, userId string) (*domain.File, error) {

	file := domain.File{
		ID:        uuid.NewString(),
		UserId:    userId,
		Name:      request.Name,
		FolderId:  "nil",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if request.FolderId != "" {

		file.FolderId = request.FolderId
	}

	_, err := fs.Repo.Table("files").Create(file)

	if err != nil {
		return nil, err
	}

	fs.CacheDB.Table("files:" + userId).Delete()

	return &file, nil
}

func (fs FileService) GetFile(userId, fileId string) (*domain.File, error) {

	var file domain.File

	data, err := fs.CacheDB.Table("files:" + fileId + "," + userId).Get()

	if err != nil {
		log.Println(err)
	}

	if data != nil {
		err = json.Unmarshal([]byte(data.(string)), &file)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	if file.ID == "" {

		// query db for file
		data, err := fs.Repo.Table("files").Where([][]any{
			{"id", fileId},
			{"userId", userId},
		}).First()

		if err != nil {
			log.Println(err)
			return nil, err
		}

		/// convert fetched record to file
		err = app.Decode(data.(map[string]interface{}), &file)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if file.ID == "" {
			return nil, errors.New("not found")
		}

		// marshal file data and store in cache
		file_data, err := json.Marshal(file)
		if err != nil {
			log.Println(err)
		}

		_, err = fs.CacheDB.Table("files:" + fileId + "," + userId).Create(string(file_data))
		if err != nil {
			log.Println(err)
		}
	}

	return &file, nil
}

func (fs FileService) Delete(userId, fileId string) error {

	var file domain.File

	data, _ := fs.Repo.Table("files").Where([][]any{
		{"id", fileId},
		{"userId", userId},
	}).First()

	err := app.Decode(data.(map[string]interface{}), &file)

	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	if file.ID == "" {
		return errors.New("record not found")
	}

	err = fs.Repo.Table("files").Where([][]any{
		{"id", fileId},
		{"userId", userId},
	}).Delete()

	fs.CacheDB.Table("files:" + userId).Delete()
	fs.CacheDB.Table("files:" + fileId + "," + userId).Delete()
	fs.CacheDB.Table("files:" + file.FolderId).Delete()

	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}
