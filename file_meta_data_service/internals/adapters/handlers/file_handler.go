package handlers

import (
	fr "filemetadata-service/internals/app/form-request"
	"filemetadata-service/internals/app/generator"
	rv "filemetadata-service/internals/app/request-validator"
	"filemetadata-service/internals/core/domain"
	"filemetadata-service/internals/core/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FileHandler struct {
	fs *services.FileService
}

func NewFileHandler(fs *services.FileService) *FileHandler {
	return &FileHandler{fs}
}

func (fh FileHandler) Index(c *gin.Context) {

	files := fh.fs.GetFiles(c.GetHeader("UserId"))

	if len(files) == 0 {
		files = []domain.File{}
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func (fh FileHandler) Store(c *gin.Context) {

	var request fr.FileRequest

	// validate request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": rv.Validator{Err: err.(validator.ValidationErrors)}.Errors(),
		})
		return
	}
	userId := c.GetHeader("UserId")
	file, err := fh.fs.SaveFile(request, userId)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurred",
		})

		log.Printf("unable to create new folder %v", err)
		return
	}

	//generate download url
	downloadUrl := fh.fileDownloadUrl(userId, request.FolderId, file.Name)

	c.JSON(http.StatusOK, gin.H{
		"fileId":      file.ID,
		"downloadUrl": downloadUrl,
	})
}

func (fh FileHandler) Show(c *gin.Context) {

	file, err := fh.fs.GetFile(c.GetHeader("UserId"), c.Param("file"))

	if err != nil && err.Error() == "not found" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//generate download url
	downloadUrl := fh.fileDownloadUrl(file.UserId, file.FolderId, file.Name)

	c.JSON(http.StatusOK, gin.H{
		"fileId":      file.ID,
		"downloadUrl": downloadUrl,
	})
}

func (fh FileHandler) Delete(c *gin.Context) {

	err := fh.fs.Delete(c.GetHeader("UserId"), c.Param("file"))

	if err != nil && err.Error() != "record not found" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "file not found",
		})
		return
	}

	if err != nil && err.Error() != "internal server error" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "file deleted",
	})
}

func (fh FileHandler) fileDownloadUrl(userId, folderId, fileName string) string {

	bucket := "myawsbucket899"
	objectKey := userId + "/"

	if folderId != "" {
		folderService := services.NewFolderService(fh.fs.Repo, fh.fs.CacheDB)
		folder, err := folderService.GetFolder(userId, folderId)
		if err != nil {
			log.Printf("unable to get folder %v", err)
		}

		objectKey += folder.Name + "/"
	}

	objectKey += fileName

	awsPresignUrl := generator.NewAwsPresignUrl(&objectKey, &bucket)
	downloadUrl, err := awsPresignUrl.Generate()

	if err != nil {
		log.Printf("unable to generate aws presigned url %v", err)
	}

	return downloadUrl
}
