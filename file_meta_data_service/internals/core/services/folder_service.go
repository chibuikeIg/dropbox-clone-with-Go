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

type FolderService struct {
	repo, cacheDB ports.FileMetaDataDBRepository
}

func NewFolderService(repo, cacheDB ports.FileMetaDataDBRepository) *FolderService {

	return &FolderService{repo: repo, cacheDB: cacheDB}
}

func (fs *FolderService) Folders(userId string) []domain.Folder {

	var folders []domain.Folder

	// load from cache and if
	// not available load from DB
	// then save in cache
	data, err := fs.cacheDB.Table("folders:" + userId).Get()
	if err != nil {
		log.Println(err)
	}

	if data != nil {

		err = json.Unmarshal([]byte(data.(string)), &folders)
		if err != nil {
			log.Println(err)
		}
	}

	if len(folders) == 0 {

		fs.repo.Table("folders").Find([]string{
			"userId",
			userId,
		}, &folders)

		data, err := json.Marshal(folders)
		if err != nil {
			log.Println(err)
		}

		_, err = fs.cacheDB.Table("folders:" + userId).Create(string(data))
		if err != nil {
			log.Println(err)
		}
	}

	return folders
}

func (fs *FolderService) Store(request fr.FolderRequest, userId string) (domain.Folder, error) {

	folder := domain.Folder{
		ID:        uuid.NewString(),
		Name:      request.Name,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := fs.repo.Table("folders").Create(folder)

	fs.cacheDB.Table("folders:" + userId).Delete()

	return folder, err
}

func (fs *FolderService) GetFolderFiles(folder_id string, user_id string) ([]domain.File, error) {

	var files []domain.File

	data, err := fs.cacheDB.Table("files:" + folder_id).Get()
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

		fs.repo.Table("files").Find([]string{
			"folderId",
			folder_id,
		}, &files)

		data, err := json.Marshal(files)
		if err != nil {
			log.Println(err)
		}

		_, err = fs.cacheDB.Table("files:" + folder_id).Create(string(data))
		if err != nil {
			log.Println(err)
		}
	}

	return files, nil
}

func (fs *FolderService) GetFolder(userId, folderId string) (*domain.Folder, error) {
	var folder domain.Folder

	data, err := fs.cacheDB.Table("folders:" + folderId + "," + userId).Get()

	if err != nil {
		log.Println(err)
	}

	if data != nil {
		err = json.Unmarshal([]byte(data.(string)), &folder)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	if folder.ID == "" {

		// query db for file
		data, err := fs.repo.Table("folders").Where([][]any{
			{"id", folderId},
			{"userId", userId},
		}).First()

		if err != nil {
			log.Println(err)
			return nil, err
		}

		/// convert fetched record to file
		err = app.Decode(data.(map[string]interface{}), &folder)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if folder.ID == "" {
			return nil, errors.New("not found")
		}

		// marshal file data and store in cache
		folder_data, err := json.Marshal(folder)
		if err != nil {
			log.Println(err)
		}

		_, err = fs.cacheDB.Table("folders:" + folderId + "," + userId).Create(string(folder_data))
		if err != nil {
			log.Println(err)
		}
	}

	return &folder, nil
}

func (fs *FolderService) Update(request fr.FolderRequest, userId, folderId string) (domain.Folder, error) {

	folder := domain.Folder{
		Name:   request.Name,
		UserId: userId,
		ID:     folderId,
	}

	_, err := fs.repo.Table("folders").Where([][]any{
		{"id", folder.ID},
		{"userId", folder.UserId},
	}).Update(map[string]string{
		"name":      folder.Name,
		"updatedAt": time.Now().String(),
	})

	fs.cacheDB.Table("folders:" + folder.UserId).Delete()

	return folder, err
}

func (fs *FolderService) Delete(userId string, folder_id string) error {

	folder, _ := fs.repo.Table("folders").Where([][]any{
		{"id", folder_id},
		{"userId", userId},
	}).First()

	if len(folder.(map[string]interface{})) == 0 {
		return errors.New("record not found")
	}

	err := fs.repo.Table("folders").Where([][]any{
		{"id", folder_id},
		{"userId", userId},
	}).Delete()

	fs.cacheDB.Table("folders:" + userId).Delete()
	fs.cacheDB.Table("folders:" + folder_id + "," + userId).Delete()

	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}
